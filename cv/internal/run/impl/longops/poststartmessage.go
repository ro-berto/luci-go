// Copyright 2021 The LUCI Authors.
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

package longops

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/common/lease"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/gerrit"
	"go.chromium.org/luci/cv/internal/gerrit/botdata"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/usertext"
)

// staleCLAgeThreshold controls when CV should fetch latest Gerrit info before
// trying to post the starting message.
//
// Large values increase the chance of duplicate starting messages,
// while low values increase load on Gerrit.
const staleCLAgeThreshold = 10 * time.Second

// notPosted means start message wasn't posted yet.
//
// exists for readability only.
var notPosted time.Time

// PostStartMessageOp posts a start message on each of the Run CLs.
//
// PostStartMessageOp is a single-use object.
type PostStartMessageOp struct {
	// All public fields must be set.

	*Base
	GFactory gerrit.Factory
	Env      *common.Env

	// These private fields are set internally as implementation details.

	rcls       []*run.RunCL
	cfg        *prjcfg.ConfigGroup
	botdataCLs []botdata.ChangeID

	lock           sync.Mutex
	latestPostedAt time.Time
}

// Do actually posts the start messages.
func (op *PostStartMessageOp) Do(ctx context.Context) (*eventpb.LongOpCompleted, error) {
	op.assertCalledOnce()

	if op.IsCancelRequested() {
		return &eventpb.LongOpCompleted{Status: eventpb.LongOpCompleted_CANCELLED}, nil
	}

	if err := op.prepare(ctx); err != nil {
		return nil, err
	}

	errs := make(errors.MultiError, len(op.rcls))
	poolError := parallel.WorkPool(min(len(op.rcls), 8), func(work chan<- func() error) {
		for i, rcl := range op.rcls {
			i, rcl := i, rcl
			work <- func() error {
				switch posted, err := op.doCL(ctx, rcl); {
				case err != nil:
					// all errors will be aggregated across all CLs below.
					errs[i] = err
				default:
					op.lock.Lock()
					if op.latestPostedAt.IsZero() || op.latestPostedAt.Before(posted) {
						op.latestPostedAt = posted
					}
					op.lock.Unlock()
				}
				return nil
			}
		}
	})
	if poolError != nil {
		panic(fmt.Errorf("unexpected WorkPool error %s", poolError))
	}

	// Aggregate all errors into the final result.
	psm := &eventpb.LongOpCompleted_PostStartMessage{}
	result := &eventpb.LongOpCompleted{
		Status: eventpb.LongOpCompleted_SUCCEEDED, // be hopeful by default
		Result: &eventpb.LongOpCompleted_PostStartMessage_{
			PostStartMessage: psm,
		},
	}
	var firstError error
	for i, rcl := range op.rcls {
		switch err := errs[i]; {
		case err == nil:
			psm.Posted = append(psm.Posted, int64(rcl.ID))
		case errors.Unwrap(err) == errCancelHonored:
			result.Status = eventpb.LongOpCompleted_CANCELLED
			errs[i] = nil
		case !transient.Tag.In(err):
			if psm.GetPermanentErrors() == nil {
				psm.PermanentErrors = make(map[int64]string, 1)
				firstError = err
			}
			psm.GetPermanentErrors()[int64(rcl.ID)] = err.Error()
			logging.Errorf(ctx, "failed to post start message on CL %d %q: %s", rcl.ID, rcl.ExternalID, err)
		case firstError == nil:
			// First transient error.
			firstError = err
		}
	}

	switch {
	case len(psm.GetPermanentErrors()) > 0:
		result.Status = eventpb.LongOpCompleted_FAILED
		return result, firstError
	case firstError != nil:
		// Must be a transient error, expect a retry.
		return nil, firstError
	default:
		psm.Time = timestamppb.New(op.latestPostedAt)
		return result, nil
	}
}

func (op *PostStartMessageOp) prepare(ctx context.Context) error {
	eg, eCtx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		op.rcls, err = run.LoadRunCLs(eCtx, op.Run.ID, op.Run.CLs)
		return
	})
	eg.Go(func() (err error) {
		op.cfg, err = prjcfg.GetConfigGroup(eCtx, op.Run.ID.LUCIProject(), op.Run.ConfigGroupID)
		return
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	op.botdataCLs = make([]botdata.ChangeID, len(op.rcls))
	for i, rcl := range op.rcls {
		op.botdataCLs[i].Host = rcl.Detail.GetGerrit().GetHost()
		op.botdataCLs[i].Number = rcl.Detail.GetGerrit().GetInfo().GetNumber()
	}
	return nil
}

func (op *PostStartMessageOp) doCL(ctx context.Context, rcl *run.RunCL) (time.Time, error) {
	if rcl.Detail.GetGerrit() == nil {
		panic(fmt.Errorf("CL %d is not a Gerrit CL", rcl.ID))
	}

	if op.IsCancelRequested() {
		return notPosted, errCancelHonored
	}
	switch postedAt, err := op.isAlreadyPosted(ctx, rcl); {
	case err != nil:
		return notPosted, errors.Annotate(err, "failed to check if message was already posted").Err()
	case postedAt != notPosted:
		logging.Debugf(ctx, "CL %d %s already has the starting message at %s", rcl.ID, rcl.ExternalID, postedAt)
		return postedAt, nil
	}

	if op.IsCancelRequested() {
		return notPosted, errCancelHonored
	}
	return op.post(ctx, rcl)
}

func (op *PostStartMessageOp) isAlreadyPosted(ctx context.Context, rcl *run.RunCL) (time.Time, error) {
	// Check if Gerrit CL already has the starting message by first checking CV's
	// own cache (CL entity in Datastore) and if the cache is old, fetching latest
	// info from Gerrit.
	cl := changelist.CL{ID: rcl.ID}
	switch err := datastore.Get(ctx, &cl); {
	case err == datastore.ErrNoSuchEntity:
		return notPosted, errors.Annotate(err, "CL no longer exists").Err()
	case err != nil:
		return notPosted, errors.Annotate(err, "failed to load CL").Tag(transient.Tag).Err()
	}

	switch posted := op.findMessageTS(rcl, cl.Snapshot.GetGerrit().GetInfo()); {
	case posted != notPosted:
		return posted, nil
	case clock.Since(ctx, cl.Snapshot.GetExternalUpdateTime().AsTime()) < staleCLAgeThreshold:
		// Accept possibility of duplicate messages within the staleCLAgeThreshold.
		return notPosted, nil
	}

	// Fetch the latest CL details from Gerrit.
	gc, err := op.GFactory.MakeClient(ctx, rcl.Detail.GetGerrit().GetHost(), op.Run.ID.LUCIProject())
	if err != nil {
		return notPosted, err
	}

	req := &gerritpb.GetChangeRequest{
		Project: rcl.Detail.GetGerrit().GetInfo().GetProject(),
		Number:  rcl.Detail.GetGerrit().GetInfo().GetNumber(),
		Options: []gerritpb.QueryOption{gerritpb.QueryOption_MESSAGES},
	}
	var ci *gerritpb.ChangeInfo
	outerErr := op.GFactory.MakeMirrorIterator(ctx).RetryIfStale(func(opt grpc.CallOption) error {
		ci, err = gc.GetChange(ctx, req, opt)
		switch grpcutil.Code(err) {
		case codes.OK:
			return nil
		case codes.PermissionDenied:
			// This is permanent error which shouldn't be retried.
			return err
		case codes.NotFound:
			return gerrit.ErrStaleData
		default:
			err = gerrit.UnhandledError(ctx, err, "Gerrit.GetChange")
			return err
		}
	})
	switch {
	case err != nil:
		return notPosted, errors.Annotate(err, "failed to get the latest Gerrit ChangeInfo").Err()
	case outerErr != nil:
		// Shouldn't happen, unless Mirror iterate itself errors out for some
		// reason.
		return notPosted, outerErr
	default:
		return op.findMessageTS(rcl, ci), nil
	}
}

// findMessageTS returns when the start message was posted on a CL or `notPosted`.
func (op *PostStartMessageOp) findMessageTS(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
	// Look at latest messages first for efficiency,
	// and skip all messages which are too old.
	clTriggeredAt := rcl.Trigger.Time.AsTime()
	for i := len(ci.GetMessages()) - 1; i >= 0; i-- {
		m := ci.GetMessages()[i]
		t := m.GetDate().AsTime()
		if t.Before(clTriggeredAt) {
			// i-th message is too old, no need to check even older ones.
			return notPosted
		}
		switch data, ok := botdata.Parse(m); {
		case !ok:
		case data.Action != botdata.Start:
		case data.TriggeredAt != clTriggeredAt:
		case data.Revision != rcl.Detail.GetGerrit().GetInfo().GetCurrentRevision():
		case len(data.CLs) != len(op.Run.CLs):
		default:
			return t
		}
	}
	return notPosted
}

func (op *PostStartMessageOp) post(ctx context.Context, rcl *run.RunCL) (time.Time, error) {
	msg, err := op.makeMessage(rcl)
	if err != nil {
		return notPosted, errors.Annotate(err, "failed to generate the starting message").Err()
	}
	gc, err := op.GFactory.MakeClient(ctx, rcl.Detail.GetGerrit().GetHost(), op.Run.ID.LUCIProject())
	if err != nil {
		return notPosted, err
	}

	ctx, cancelLease, err := lease.ApplyOnCL(ctx, rcl.ID, 2*time.Minute, "post-start-message")
	if err != nil {
		return notPosted, err
	}
	defer cancelLease()

	req := &gerritpb.SetReviewRequest{
		Project:    rcl.Detail.GetGerrit().GetInfo().GetProject(),
		Number:     rcl.Detail.GetGerrit().GetInfo().GetNumber(),
		RevisionId: rcl.Detail.GetGerrit().GetInfo().GetCurrentRevision(),
		// Kept compatible with CQDaemon to avoid breaking user filters.
		Tag:     "autogenerated:cq",
		Notify:  gerritpb.Notify_NOTIFY_NONE,
		Message: msg,
	}
	outerErr := op.GFactory.MakeMirrorIterator(ctx).RetryIfStale(func(opt grpc.CallOption) error {
		_, err = gc.SetReview(ctx, req, opt)
		switch grpcutil.Code(err) {
		case codes.OK:
			return nil
		case codes.PermissionDenied:
			// This is a permanent error which shouldn't be retried.
			return err
		case codes.NotFound:
			// This is known to happen on new CLs or on recently created revisions.
			return gerrit.ErrStaleData
		default:
			err = gerrit.UnhandledError(ctx, err, "Gerrit.SetReview")
			return err
		}
	})
	switch {
	case err != nil:
		return notPosted, errors.Annotate(err, "failed to Gerrit.SetReview").Err()
	case outerErr != nil:
		// Shouldn't happen, unless MirrorIterator itself errors out for some
		// reason.
		return notPosted, outerErr
	default:
		// NOTE: to avoid another round-trip to Gerrit, use the CV time here even
		// though it isn't the same as what Gerrit recorded.
		return clock.Now(ctx).Truncate(time.Second), nil
	}
}

func (op *PostStartMessageOp) makeMessage(rcl *run.RunCL) (string, error) {
	humanMsg := usertext.OnRunStartedGerritMessage(op.Run, op.cfg, op.Env)
	bd := botdata.BotData{
		Action:      botdata.Start,
		CLs:         op.botdataCLs,
		Revision:    rcl.Detail.GetGerrit().GetInfo().GetCurrentRevision(),
		TriggeredAt: rcl.Trigger.Time.AsTime(),
	}
	return botdata.Append(humanMsg, bd)
}
