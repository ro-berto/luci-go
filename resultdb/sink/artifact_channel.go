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

package sink

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/sync/dispatcher"
	"go.chromium.org/luci/common/sync/dispatcher/buffer"

	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
	sinkpb "go.chromium.org/luci/resultdb/sink/proto/v1"
)

// MaxBatchableArtifactSize is the maximum size of an artifact that can be added to
// batchChannel.
const MaxBatchableArtifactSize = 2 * 1024 * 1024

type uploadTask struct {
	art     *sinkpb.Artifact
	artName string
	size    int64 // content size
}

// newUploadTask constructs an uploadTask for the artifact.
//
// If FilePath is set on the artifact, this calls os.Stat to obtain the file information,
// and may return an error if the Stat call fails. e.g., permission denied, not found.
// It also returns an error if the artifact file path is a directory.
func newUploadTask(name string, art *sinkpb.Artifact) (*uploadTask, error) {
	ret := &uploadTask{
		art:     art,
		artName: name,
		size:    int64(len(art.GetContents())),
	}

	// Find and save the content size on uploadTask creation, so that the task scheduling
	// and processing logic can use the size information w/o issuing system calls.
	if fp := art.GetFilePath(); fp != "" {
		st, err := os.Stat(fp)
		switch {
		case err != nil:
			return nil, errors.Annotate(err, "querying file info").Err()
		case st.Mode().IsRegular():
			// break

		// Return a more human friendly error than 1000....0.
		case st.IsDir():
			return nil, errors.Reason("%q is a directory", fp).Err()
		default:
			return nil, errors.Reason("%q is not a regular file: %s", fp, strconv.FormatInt(int64(st.Mode()), 2)).Err()
		}
		ret.size = st.Size()
	}
	return ret, nil
}

// CreateRequest returns a CreateArtifactRequest for the upload task.
//
// Note that this will open and read content from the file, the artifact is set with
// Artifact_FilePath. Save the returned request to avoid unnecessary I/Os,
// if necessary.
func (t *uploadTask) CreateRequest() (*pb.CreateArtifactRequest, error) {
	invID, tID, rID, aID, err := pbutil.ParseArtifactName(t.artName)
	req := &pb.CreateArtifactRequest{
		Artifact: &pb.Artifact{
			ArtifactId:  aID,
			ContentType: t.art.GetContentType(),
			SizeBytes:   t.size,
			Contents:    t.art.GetContents(),
		},
	}

	// parent
	switch {
	case err != nil:
		// This should not happend.
		// uploadTask should be created with validated artifacts only.
		panic(fmt.Sprintf("invalid uploadTask.artName %q: %s", t.artName, err))
	case tID == "":
		// Invocation-level artifact
		req.Parent = pbutil.InvocationName(invID)
	default:
		req.Parent = pbutil.TestResultName(invID, tID, rID)
	}

	// contents
	if fp := t.art.GetFilePath(); fp != "" {
		if req.Artifact.Contents, err = ioutil.ReadFile(fp); err != nil {
			return nil, err
		}
	}
	// If the size of the read content is different to what stat claimed initially, then
	// return an error, so that the batching logic can be kept simple. Test frameworks
	// should send finalized artifacts only.
	if int64(len(req.Artifact.Contents)) != t.size {
		return nil, errors.Reason(
			"the size of the artifact contents changed from %d to %d",
			t.size, len(req.Artifact.Contents)).Err()
	}
	return req, nil
}

type artifactChannel struct {
	// batchChannel uploads artifacts via pb.BatchCreateArtifacts().
	//
	// This batches input artifacts and uploads them all at once.
	// This is suitable for uploading a large number of small artifacts.
	//
	// The downside of this channel is that there is a limit on the maximum size of
	// an artifact that can be included in a batch. Use streamChannel for artifacts
	// bigger than MaxBatchableArtifactSize.
	batchChannel dispatcher.Channel

	// streamChannel uploads artifacts in a streaming manner via HTTP.
	//
	// This is suitable for uploading large files, but with limited parallelism.
	// Use batchChannel, if possible.
	streamChannel dispatcher.Channel

	// wgActive indicates if there are active goroutines invoking reportTestResults.
	//
	// reportTestResults can be invoked by multiple goroutines in parallel. wgActive is used
	// to ensure that all active goroutines finish enqueuing messages to the channel before
	// closeAndDrain closes and drains the channel.
	wgActive sync.WaitGroup

	// 1 indicates that artifactChannel started the process of closing and draining
	// the channel. 0, otherwise.
	closed int32
}

func newArtifactChannel(ctx context.Context, cfg *ServerConfig) *artifactChannel {
	var err error
	c := &artifactChannel{}
	au := artifactUploader{
		Recorder:     cfg.Recorder,
		StreamClient: cfg.ArtifactStreamClient,
		StreamHost:   cfg.ArtifactStreamHost,
	}

	// batchChannel
	bcOpts := &dispatcher.Options{
		Buffer: buffer.Options{
			// BatchRequest can include up to 500 requests. KEEP BatchSize <= 500
			// For more details, visit
			// https://godoc.org/go.chromium.org/luci/resultdb/proto/v1#BatchCreateArtifactsRequest
			BatchSize:    500,
			MaxLeases:    int(cfg.ArtChannelMaxLeases),
			FullBehavior: &buffer.BlockNewItems{MaxItems: 8000},
		},
	}
	c.batchChannel, err = dispatcher.NewChannel(ctx, bcOpts, func(b *buffer.Batch) error {
		// TODO(ddoman): implement me
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create batch channel for artifacts: %s", err))
	}

	// streamChannel
	stOpts := &dispatcher.Options{
		Buffer: buffer.Options{
			// BatchSize MUST be 1.
			BatchSize:    1,
			MaxLeases:    int(cfg.ArtChannelMaxLeases),
			FullBehavior: &buffer.BlockNewItems{MaxItems: 4000},
		},
	}
	c.streamChannel, err = dispatcher.NewChannel(ctx, stOpts, func(b *buffer.Batch) error {
		return errors.Annotate(
			au.StreamUpload(ctx, b.Data[0].(*uploadTask), cfg.UpdateToken),
			"StreamUpload").Err()
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create stream channel for artifacts: %s", err))
	}
	return c
}

func (c *artifactChannel) closeAndDrain(ctx context.Context) {
	// mark the channel as closed, so that schedule() won't accept new tasks.
	if !atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		return
	}
	// wait for all the active sessions to finish enquing tests results to the channel
	c.wgActive.Wait()

	var draining sync.WaitGroup
	draining.Add(2)
	go func() {
		defer draining.Done()
		c.batchChannel.CloseAndDrain(ctx)
	}()
	go func() {
		defer draining.Done()
		c.streamChannel.CloseAndDrain(ctx)
	}()
	draining.Wait()
}

func (c *artifactChannel) schedule(tasks ...*uploadTask) {
	c.wgActive.Add(1)
	defer c.wgActive.Done()
	// if the channel already has been closed, drop the test results.
	if atomic.LoadInt32(&c.closed) == 1 {
		return
	}

	for _, task := range tasks {
		// TODO(ddoman): send small artifacts to batchChannel
		c.streamChannel.C <- task
	}
}
