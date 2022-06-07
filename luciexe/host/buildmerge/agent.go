// Copyright 2019 The LUCI Authors.
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

// Package buildmerge implements the build.proto tracking and merging logic for
// luciexe host applications.
//
// You probably want to use `go.chromium.org/luci/luciexe/host` instead.
//
// This package is separate from luciexe/host to avoid unnecessary entaglement
// with butler/logdog; All the logic here is implemented to avoid:
//
//   * interacting with the environment
//   * interacting with butler/logdog (except by implementing callbacks for
//     those, but only acting on simple datastructures/proto messages)
//   * handling errors in any 'brutal' ways (all errors in this package are
//     handled by reporting them directly in the data structures that this
//     package manipulates).
//
// This is done to simplify testing (as much as it can be) by concentrating all
// the environment stuff into luciexe/host, and all the 'pure' functional stuff
// here (search "imperative shell, functional core").
package buildmerge

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/proto/reflectutil"
	"go.chromium.org/luci/common/sync/dispatcher"
	"go.chromium.org/luci/common/sync/dispatcher/buffer"
	"go.chromium.org/luci/logdog/api/logpb"
	"go.chromium.org/luci/logdog/client/butler"
	"go.chromium.org/luci/logdog/common/types"
	"go.chromium.org/luci/luciexe"
)

// CalcURLFn is a stateless function which can calculate the absolute url and
// viewUrl from a given logdog namespace (with trailing slash) and streamName.
type CalcURLFn func(namespaceSlash, streamName types.StreamName) (url, viewUrl string)

// Agent holds all the logic around merging build.proto streams.
type Agent struct {
	// MergedBuildC is the channel of all the merged builds generated by this
	// Agent.
	//
	// The rate at which Agent merges Builds is governed by the consumption of
	// this channel; Consuming it slowly will have Agent merge less frequently,
	// and consuming it rapidly will have Agent merge more frequently.
	//
	// The last build before the channel closes will always be the final state of
	// all builds at the time this Agent was Close()'d.
	MergedBuildC <-chan *bbpb.Build

	// Wait on this channel for the Agent to drain. Will only drain after calling
	// Close() at least once.
	DrainC <-chan struct{}

	// used to cancel in-progress sendMerge calls.
	ctx context.Context

	// mergedBuildC is the send side of MergedBuildC
	mergedBuildC chan<- *bbpb.Build

	// userNamespace is the logdog namespace (with a trailing slash) which we'll
	// use to determine if a new stream is potentially monitored, or not.
	userNamespace types.StreamName

	// userRootURL is the full url ('logdog://.../stream/build.proto') of the
	// user's "root" build.proto stream (i.e. the one emitted by the top level
	// luciexe implementation.
	//
	// This is used as a key to start the merge process.
	userRootURL string
	baseBuild   *bbpb.Build

	// statesMu covers `states`. It must be held when reading or writing to
	// `states`, but doesn't need to be held while interacting with an individual
	// *buildState obtained from the map.
	statesMu sync.RWMutex

	// states maps a stream URL (i.e. `logdog://.../stream/build.proto`) to the
	// state tracker for that stream.
	states map[string]*buildStateTracker

	// mergeCh is used in production mode to send pings via informNewData
	mergeCh dispatcher.Channel

	// informNewData is used to 'ping' mergeCh; it's overwritten in tests.
	informNewData func()

	// done is an atomically-accessed boolean
	done int32

	// calculateURLs is a function which can convert a logdog namespace and
	// streamname into both the full 'Url' and 'ViewUrl' values for a Log message.
	// This is used by the buildMerger itself when deriving keys for the `states`
	// map, as well as for individual buildState objects to adjust their build's
	// logs' URLs.
	calculateURLs CalcURLFn
}

// New returns a new Agent.
//
// Args:
//   * ctx - used for logging, clock and cancelation. When canceled, the Agent
//     will cease sending updates on MergedBuildC, but you must still invoke
//     Agent.Close() in order to clean up all resources associated with the
//     Agent.
//   * userNamespace - The logdog namespace (with a trailing slash) under which
//     we should monitor streams.
//   * base - The "model" Build message that all generated builds should start
//     with. All build proto streams will be merged onto a copy of this message.
//     Any Output.Log's which have non-absolute URLs will have their Url and
//     ViewUrl absolutized relative to userNamespace using calculateURLs.
//   * calculateURLs - A function to calculate Log.Url and Log.ViewUrl values.
//     Should be a pure function.
//
// The following fields will be merged into `base` from the user controlled
// build.proto stream(s):
//
//   Steps
//   SummaryMarkdown
//   Status
//   StatusDetails
//   UpdateTime
//   Tags
//   EndTime
//   Output
//
// The frequency of updates from this Agent is governed by how quickly the
// caller consumes from Agent.MergedBuildC.
func New(ctx context.Context, userNamespace types.StreamName, base *bbpb.Build, calculateURLs CalcURLFn) (*Agent, error) {
	userNamespace = userNamespace.AsNamespace()

	ch := make(chan *bbpb.Build)
	userRootURL, _ := calculateURLs(userNamespace, luciexe.BuildProtoStreamSuffix)

	ret := &Agent{
		ctx: ctx,

		MergedBuildC: ch,

		mergedBuildC:  ch,
		states:        map[string]*buildStateTracker{},
		calculateURLs: calculateURLs,
		userNamespace: userNamespace,
		userRootURL:   userRootURL,
		baseBuild:     proto.Clone(base).(*bbpb.Build),
	}
	for _, log := range ret.baseBuild.GetOutput().GetLogs() {
		var err error
		log.Url, log.ViewUrl, err = absolutizeURLs(log.Url, log.ViewUrl, userNamespace, calculateURLs)
		if err != nil {
			return nil, errors.Annotate(err, "build.output.logs[%q]", log.Name).Err()
		}
	}

	var err error
	ret.mergeCh, err = dispatcher.NewChannel(ctx, &dispatcher.Options{
		Buffer: buffer.Options{
			MaxLeases:     1,
			BatchItemsMax: 1,
			FullBehavior:  &buffer.DropOldestBatch{},
		},
		DropFn:    dispatcher.DropFnQuiet,
		DrainedFn: ret.finalize,
	}, ret.sendMerge)
	if err != nil {
		return nil, err // creating dispatcher with static config should never fail
	}
	ret.informNewData = func() {
		ret.mergeCh.C <- nil // content doesn't matter
	}
	ret.DrainC = ret.mergeCh.DrainC

	return ret, nil
}

// Attach should be called once to attach this to a Butler.
//
// This must be done before the butler receives any build.proto streams.
func (a *Agent) Attach(b *butler.Butler) {
	b.AddStreamRegistrationCallback(a.onNewStream, true)
}

var validContentTypes = stringset.NewFromSlice(
	luciexe.BuildProtoContentType,
	luciexe.BuildProtoZlibContentType,
)

func (a *Agent) onNewStream(desc *logpb.LogStreamDescriptor) butler.StreamChunkCallback {
	if !a.collectingData() {
		return nil
	}

	namespace, base := types.StreamName(desc.Name).Split()

	var err error
	zlib := false
	switch validStreamT, validContentT := desc.StreamType == logpb.StreamType_DATAGRAM, validContentTypes.Has(desc.ContentType); {
	case validStreamT && validContentT:
		zlib = desc.ContentType == luciexe.BuildProtoZlibContentType
	case validStreamT && !validContentT:
		err = errors.Reason("stream %q has content type %q, expected one of %v", desc.Name, desc.ContentType, validContentTypes.ToSortedSlice()).Err()
	case !validStreamT && validContentT:
		err = errors.Reason("build proto stream %q has type %q, expected %q", desc.Name, desc.StreamType, logpb.StreamType_DATAGRAM).Err()
	case strings.HasPrefix(desc.Name, string(a.userNamespace)) && base == luciexe.BuildProtoStreamSuffix:
		err = errors.Reason("build.proto stream %q has stream type %q and content type %q, expected %q and one of %v", desc.Name, desc.StreamType, desc.ContentType, logpb.StreamType_DATAGRAM, validContentTypes.ToSortedSlice()).Err()
	default:
		// neither a ".../build.proto" stream nor a stream with valid stream type
		// or content type.
		return nil
	}

	url, _ := a.calculateURLs("", types.StreamName(desc.Name))
	bState := newBuildStateTracker(a.ctx, a, namespace, zlib, err)

	a.statesMu.Lock()
	defer a.statesMu.Unlock()
	a.states[url] = bState
	if err == nil {
		return bState.handleNewData
	}
	return nil // no need to handle invalid stream.
}

// Close causes the Agent to stop collecting data, emit a final merged build,
// and then shut down all internal routines.
func (a *Agent) Close() {
	// stops accepting new trackers
	if atomic.SwapInt32(&a.done, 1) == 1 {
		return
	}

	// close all states' and process their final work items. Closure should be
	// very quick and will activate all final processing in parallel. GetFinal
	// ensures that the state is completely settled.
	states := a.snapStates()
	for _, t := range states {
		t.Close()
	}
	for _, t := range states {
		t.GetFinal()
	}

	// tells our merge Channel to process all the current (now-final) states one
	// last time.
	a.informNewData()

	// shut down the mergeCh so it will no longer accept new informNewData calls.
	a.mergeCh.Close()
}

func (a *Agent) snapStates() map[string]*buildStateTracker {
	a.statesMu.RLock()
	trackers := make(map[string]*buildStateTracker, len(a.states))
	for k, v := range a.states {
		trackers[k] = v
	}
	a.statesMu.RUnlock()
	return trackers
}

func (a *Agent) sendMerge(_ *buffer.Batch) error {
	trackers := a.snapStates()

	builds := make(map[string]*bbpb.Build, len(trackers))
	stepCount := 0
	for k, v := range trackers {
		build := v.getLatest().build
		stepCount += len(build.GetSteps())
		builds[k] = build
	}

	base := reflectutil.ShallowCopy(a.baseBuild).(*bbpb.Build)
	base.Steps = nil
	if stepCount > 0 {
		base.Steps = make([]*bbpb.Step, 0, stepCount)
	}

	var insertSteps func(stepNS []string, streamURL string) *bbpb.Build
	insertSteps = func(stepNS []string, streamURL string) *bbpb.Build {
		build, ok := builds[streamURL]
		if !ok {
			return nil
		}
		for _, step := range build.GetSteps() {
			isMergeStep := luciexe.IsMergeStep(step)
			if isMergeStep || len(stepNS) > 0 {
				step = reflectutil.ShallowCopy(step).(*bbpb.Step)
			}
			baseName := step.Name
			if len(stepNS) > 0 {
				step.Name = strings.Join(append(stepNS, step.Name), "|")
			}

			base.Steps = append(base.Steps, step)

			if isMergeStep {
				subBuildStreamURL := step.Logs[0].Url
				subBuild := insertSteps(append(stepNS, baseName), subBuildStreamURL)
				if subBuild == nil {
					var sb strings.Builder
					if step.SummaryMarkdown != "" {
						sb.WriteString(step.SummaryMarkdown)
						sb.WriteString("\n\n")
					}
					if _, ok := builds[subBuildStreamURL]; ok {
						sb.WriteString(fmt.Sprintf("build.proto stream: %q is empty", subBuildStreamURL))
					} else {
						sb.WriteString(fmt.Sprintf("build.proto stream: %q is not registered", subBuildStreamURL))
					}
					step.SummaryMarkdown = sb.String()
				} else {
					updateStepFromBuild(step, subBuild)
				}
			}
		}
		return build
	}
	updateBaseFromUserBuild(base, insertSteps(nil, a.userRootURL))

	select {
	case a.mergedBuildC <- base:
	case <-a.ctx.Done():
		a.Close()
	}

	return nil
}

func (a *Agent) finalize() {
	close(a.mergedBuildC)
}

func (a *Agent) collectingData() bool {
	return atomic.LoadInt32(&a.done) == 0
}

// Used for minting protobuf timestamps for buildStateTrackers
func (a *Agent) clockNow() *timestamppb.Timestamp {
	ret, err := ptypes.TimestampProto(clock.Now(a.ctx))
	if err != nil {
		panic(err)
	}
	return ret
}
