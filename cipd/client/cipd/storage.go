// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cipd

import (
	"context"
	"fmt"
	"hash"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/context/ctxhttp"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"

	"go.chromium.org/luci/cipd/client/cipd/ui"
)

const (
	// uploadChunkSize is the max size of a single PUT HTTP request during upload.
	uploadChunkSize int64 = 4 * 1024 * 1024
	// uploadMaxErrors is the number of transient errors after which upload is aborted.
	uploadMaxErrors = 20
	// downloadMaxAttempts is how many times to retry a download on errors.
	downloadMaxAttempts = 10
)

// errTransientError is returned by getNextOffset in case of retryable error.
var errTransientError = errors.New("transient error in getUploadedOffset", transient.Tag)

func isTemporaryNetError(err error) bool {
	// net/http.Client seems to be wrapping errors into *url.Error. Unwrap if so.
	if uerr, ok := err.(*url.Error); ok {
		err = uerr.Err
	}
	// TODO(vadimsh): Figure out how to recognize dial timeouts, read timeouts,
	// etc. For now all network related errors that end up here are considered
	// temporary.
	switch err {
	case context.Canceled, context.DeadlineExceeded:
		return false
	default:
		return true
	}
}

// isTemporaryHTTPError returns true for HTTP status codes that indicate
// a temporary error that may go away if request is retried.
func isTemporaryHTTPError(statusCode int) bool {
	return statusCode >= http.StatusInternalServerError || statusCode == http.StatusRequestTimeout || statusCode == http.StatusBadRequest
}

type storage interface {
	upload(ctx context.Context, url string, data io.ReadSeeker) error
	download(ctx context.Context, url string, output io.WriteSeeker, h hash.Hash) error
}

// storageImpl implements 'storage' via Google Storage signed URLs.
type storageImpl struct {
	chunkSize int64
	userAgent string
	client    *http.Client
}

// Google Storage resumable upload protocol.
// See https://cloud.google.com/storage/docs/concepts-techniques#resumable

func (s *storageImpl) upload(ctx context.Context, url string, data io.ReadSeeker) error {
	activity := ui.CurrentActivity(ctx)

	// Grab the total length of the file.
	length, err := data.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}
	_, err = data.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	// Offset of data known to be persisted in the storage or -1 if needed to ask
	// Google Storage for it.
	offset := int64(0)
	// Number of transient errors reported so far.
	errors := 0

	// Called when some new data is uploaded.
	reportProgress := func(title string, offset int64) {
		activity.Progress(ctx, title, ui.UnitBytes, offset, length)
	}

	// Called when transient error happens.
	reportTransientError := func() {
		// Need to query for latest uploaded offset to resume.
		errors++
		offset = -1
		if err := ctx.Err(); err == nil {
			if errors < uploadMaxErrors {
				logging.Warningf(ctx, "Transient error, retrying...")
				clock.Sleep(ctx, 2*time.Second)
			}
		}
	}

	for errors < uploadMaxErrors {
		// Context canceled?
		if err := ctx.Err(); err != nil {
			return err
		}

		// Grab latest uploaded offset if not known.
		if offset == -1 {
			logging.Infof(ctx, "Resuming...")
			offset, err = s.getNextOffset(ctx, url, length)
			if err == errTransientError {
				reportTransientError()
				continue
			}
			if err != nil {
				return err
			}
			if offset == length {
				return nil
			}
		}

		// Length of a chunk to upload.
		chunk := s.chunkSize
		if chunk > length-offset {
			chunk = length - offset
		}

		// Upload the chunk.
		reportProgress("Uploading", offset)
		if _, err := data.Seek(offset, os.SEEK_SET); err != nil {
			return err
		}
		r, err := http.NewRequest("PUT", url, io.LimitReader(data, chunk))
		if err != nil {
			return err
		}
		rangeHeader := fmt.Sprintf("bytes %d-%d/%d", offset, offset+chunk-1, length)
		r.Header.Set("Content-Range", rangeHeader)
		r.Header.Set("Content-Length", fmt.Sprintf("%d", chunk))
		r.Header.Set("User-Agent", s.userAgent)
		resp, err := ctxhttp.Do(ctx, s.client, r)
		if err != nil {
			if isTemporaryNetError(err) {
				reportTransientError()
				continue
			}
			return err
		}
		resp.Body.Close()

		// Partially or fully uploaded.
		if resp.StatusCode == http.StatusPermanentRedirect || resp.StatusCode == http.StatusOK {
			offset += chunk
			if offset == length {
				reportProgress("Uploading", offset)
				return nil
			}
			continue
		}

		// Transient error.
		if isTemporaryHTTPError(resp.StatusCode) {
			reportTransientError()
			continue
		}

		// Fatal error.
		return fmt.Errorf("unexpected response during file upload: HTTP %d", resp.StatusCode)
	}

	return ErrUploadError
}

// getNextOffset queries the storage for size of persisted data.
func (s *storageImpl) getNextOffset(ctx context.Context, url string, length int64) (offset int64, err error) {
	r, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return offset, err
	}
	r.Header.Set("Content-Range", fmt.Sprintf("bytes */%d", length))
	r.Header.Set("Content-Length", "0")
	r.Header.Set("User-Agent", s.userAgent)
	resp, err := ctxhttp.Do(ctx, s.client, r)
	if err != nil {
		if isTemporaryNetError(err) {
			err = errTransientError
		}
		return offset, err
	}
	resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Fully uploaded.
		offset = length
	} else if resp.StatusCode == http.StatusPermanentRedirect {
		// Partially uploaded, extract last uploaded offset from Range header.
		rangeHeader := resp.Header.Get("Range")
		if rangeHeader != "" {
			_, err = fmt.Sscanf(rangeHeader, "bytes=0-%d", &offset)
			if err == nil {
				// |offset| is an *offset* of a last uploaded byte, not the data length.
				offset++
			}
		}
	} else if isTemporaryHTTPError(resp.StatusCode) {
		err = errTransientError
	} else {
		err = fmt.Errorf("unexpected response (HTTP %d) when querying for uploaded offset", resp.StatusCode)
	}
	return offset, err
}

// TODO(vadimsh): Use resumable download protocol.

func (s *storageImpl) download(ctx context.Context, url string, output io.WriteSeeker, h hash.Hash) error {
	activity := ui.CurrentActivity(ctx)

	// download is a separate function to be able to use deferred close.
	download := func(out io.Writer, src io.ReadCloser, totalLen int64) error {
		defer src.Close()

		progress := func(read int64) {
			activity.Progress(ctx, "Fetching", ui.UnitBytes, read, totalLen)
		}

		progress(0)
		_, err := io.Copy(out, &readerWithProgress{
			reader:   src,
			callback: progress,
		})
		return err
	}

	// reportTransientError logs the error and sleep few seconds.
	reportTransientError := func(msg string, args ...interface{}) {
		if err := ctx.Err(); err != nil {
			return
		}
		logging.Warningf(ctx, msg, args...)
		clock.Sleep(ctx, 2*time.Second)
	}

	// Download the actual data (several attempts).
	for attempt := 0; attempt < downloadMaxAttempts; attempt++ {
		// Context canceled?
		if err := ctx.Err(); err != nil {
			return err
		}

		// Rewind output to zero offset.
		h.Reset()
		_, err := output.Seek(0, os.SEEK_SET)
		if err != nil {
			return err
		}

		// Send the request.
		logging.Infof(ctx, "Connecting...")
		var req *http.Request
		var resp *http.Response
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Set("User-Agent", s.userAgent)
		resp, err = ctxhttp.Do(ctx, s.client, req)
		if err != nil {
			if isTemporaryNetError(err) {
				reportTransientError("Failed to connect: %s", err)
				continue
			}
			return err
		}

		// Transient error, retry.
		if isTemporaryHTTPError(resp.StatusCode) {
			resp.Body.Close()
			reportTransientError("Transient HTTP error %d", resp.StatusCode)
			continue
		}

		// Fatal error, abort.
		if resp.StatusCode >= 400 {
			resp.Body.Close()
			return fmt.Errorf("server replied with HTTP code %d", resp.StatusCode)
		}

		// Try to fetch (will close resp.Body when done).
		err = download(io.MultiWriter(output, h), resp.Body, resp.ContentLength)
		if err != nil {
			reportTransientError("Transient error: %s", err)
			continue
		}

		// Success.
		return nil
	}

	return ErrDownloadError
}

// readerWithProgress is io.Reader that calls callback whenever something is
// read from it.
type readerWithProgress struct {
	reader   io.Reader
	total    int64
	callback func(total int64)
}

func (r *readerWithProgress) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.total += int64(n)
	r.callback(r.total)
	return n, err
}
