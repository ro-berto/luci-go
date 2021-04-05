// Copyright 2020 The LUCI Authors.
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

package artifactcontent

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"golang.org/x/sync/errgroup"
	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/trace"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/router"
)

// RBEConn creates a gRPC connection to RBE authenticated as self.
func RBEConn(ctx context.Context) (*grpc.ClientConn, error) {
	creds, err := auth.GetPerRPCCredentials(ctx, auth.AsSelf, auth.WithScopes(auth.CloudOAuthScopes...))
	if err != nil {
		return nil, err
	}

	return grpc.Dial(
		"remotebuildexecution.googleapis.com:443",
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
		grpc.WithPerRPCCredentials(creds),
	)
}

// RegisterRBEInstanceFlag registers -artifact-rbe-instance flag.
func RegisterRBEInstanceFlag(fs *flag.FlagSet, target *string) {
	fs.StringVar(
		target,
		"artifact-rbe-instance",
		"",
		"Name of the RBE instance to use for artifact storage",
	)
}

// handleRBECASContent serves artifact content stored in RBE-CAS.
func (r *contentRequest) handleRBECASContent(c *router.Context, hash string) {
	// Protocol:
	// https://github.com/bazelbuild/remote-apis/blob/7802003e00901b4e740fe0ebec1243c221e02ae2/build/bazel/remote/execution/v2/remote_execution.proto#L229-L233
	// https://github.com/googleapis/googleapis/blob/c8e291e6a4d60771219205b653715d5aeec3e96b/google/bytestream/bytestream.proto#L50-L53

	// Start a reading stream.
	stream, err := r.ReadCASBlob(c.Context, &bytestream.ReadRequest{
		ResourceName: resourceName(r.RBECASInstanceName, hash, r.size.Int64),
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// Do not lose the original error message.
			logging.Warningf(c.Context, "RBE-CAS responded: %s", err)
			err = appstatus.Errorf(codes.NotFound, "artifact content no longer exists")
		}
		r.sendError(c.Context, err)
		return
	}

	// Forward the blob to the client.
	wroteHeader := false
	for {
		_, readSpan := trace.StartSpan(c.Context, "resultdb.readChunk")
		chunk, err := stream.Recv()
		if err == nil {
			readSpan.Attribute("size", len(chunk.Data))
		}
		readSpan.End(err)

		switch {
		case err == io.EOF:
			// We are done.
			return

		case err != nil:
			if wroteHeader {
				// The response was already partially written, so it is too late to
				// write headers. Write at least something indicating the incomplete
				// response.
				fmt.Fprintf(c.Writer, "\nResultDB: internal error while writing the response!\n")
				logging.Errorf(c.Context, "Failed to read from RBE-CAS in the middle of response: %s", err)
			} else {
				r.sendError(c.Context, err)
			}
			return

		default:
			// Forward the chunk.
			if !wroteHeader {
				r.writeContentHeaders()
				wroteHeader = true
			}

			_, writeSpan := trace.StartSpan(c.Context, "resultdb.writeChunk")
			writeSpan.Attribute("size", len(chunk.Data))
			_, err := c.Writer.Write(chunk.Data)
			writeSpan.End(err)
			if err != nil {
				logging.Warningf(c.Context, "Failed to write a response chunk: %s", err)
				return
			}
		}
	}
}

func resourceName(instance, hash string, size int64) string {
	return fmt.Sprintf("%s/blobs/%s/%d", instance, strings.TrimPrefix(hash, "sha256:"), size)
}

// Reader reads the artifact content from RBE-CAS.
type Reader struct {
	// RBEInstance is the name of the RBE instance where the artifact is stored.
	// Example: "projects/luci-resultdb/instances/artifacts".
	RBEInstance string
	// Hash is the hash of the artifact content stored in RBE-CAS.
	Hash string
	// Size is the content size in bytes.
	Size int64
}

// DownloadRBECASContent calls f for the downloaded artifact content.
func (r *Reader) DownloadRBECASContent(ctx context.Context, bs bytestream.ByteStreamClient, f func(context.Context, io.Reader) error) error {
	stream, err := bs.Read(ctx, &bytestream.ReadRequest{
		ResourceName: resourceName(r.RBEInstance, r.Hash, r.Size),
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			logging.Warningf(ctx, "RBE-CAS responded: %s", err)
		}
		return err
	}

	pr, pw := io.Pipe()
	eg, ctx := errgroup.WithContext(ctx)
	defer eg.Wait()
	eg.Go(func() error {
		defer pr.Close()
		return f(ctx, pr)
	})

	eg.Go(func() error {
		defer pw.Close()
		for {
			_, readSpan := trace.StartSpan(ctx, "resultdb.readChunk")
			chunk, err := stream.Recv()
			if err == nil {
				readSpan.Attribute("size", len(chunk.Data))
			}
			readSpan.End(err)

			switch {
			case err == io.EOF:
				// We are done.
				return nil

			case err != nil:
				return err

			default:
				if _, err := pw.Write(chunk.Data); err != nil {
					if err == io.ErrClosedPipe {
						// If f() exits early, return nil here to see what error that f() returns.
						return nil
					}
					return errors.Annotate(err, "write to pipe").Err()
				}
			}
		}
	})

	return eg.Wait()
}
