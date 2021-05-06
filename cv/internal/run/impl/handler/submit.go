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

package handler

import (
	"context"
	"fmt"
	"sort"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/gerrit"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/state"
	"go.chromium.org/luci/cv/internal/run/impl/submit"
	"go.chromium.org/luci/cv/internal/tree"
)

// OnReadyForSubmission implements Handler interface.
func (impl *Impl) OnReadyForSubmission(ctx context.Context, rs *state.RunState) (*Result, error) {
	switch status := rs.Run.Status; {
	case run.IsEnded(status):
		// It is safe to discard this event because this event either
		//  * arrives after Run gets cancelled while waiting for submission.
		//  * is sent by OnCQDVerificationCompleted handler as a fail-safe and Run
		//    submission has already completed.
		logging.Debugf(ctx, "received ReadyForSubmission event when Run is %s", status)
		// Under certain race condition, this Run may still occupy the submit
		// queue. So, check first without a transaction and then initiate a
		// transaction to release if this Run currently occupies the submit queue.
		switch current, err := submit.CurrentRun(ctx, rs.Run.ID.LUCIProject()); {
		case err != nil:
			return nil, err
		case current == rs.Run.ID:
			var innerErr error
			err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				innerErr = submit.Release(ctx, impl.RM.NotifyReadyForSubmission, rs.Run.ID)
				return innerErr
			}, nil)
			switch {
			case innerErr != nil:
				return nil, innerErr
			case err != nil:
				return nil, errors.Annotate(err, "failed to release submit queue").Tag(transient.Tag).Err()
			}
		}
		return &Result{State: rs}, nil
	case status == run.Status_SUBMITTING:
		return continueSubmissionIfPossible(ctx, rs, impl.RM)
	case status == run.Status_RUNNING || status == run.Status_WAITING_FOR_SUBMISSION:
		// TODO(yiwzhang): fail if partially submitted.
		rs, err := markSubmitting(ctx, rs)
		if err != nil {
			return nil, err
		}
		s, err := constructSubmitter(ctx, rs, impl.RM)
		if err != nil {
			return nil, err
		}
		return &Result{
			State:         rs,
			PostProcessFn: s.submit,
		}, nil
	default:
		panic(fmt.Errorf("impossible status %s", status))
	}
}

func continueSubmissionIfPossible(ctx context.Context, rs *state.RunState, rm RM) (*Result, error) {
	deadline := rs.Run.Submission.GetDeadline()
	taskID := rs.Run.Submission.GetTaskId()
	switch {
	case deadline == nil:
		panic(fmt.Errorf("impossible: run %q is submitting but Run.Submission.Deadline is not set", rs.Run.ID))
	case taskID == "":
		panic(fmt.Errorf("impossible: run %q is submitting but Run.Submission.TaskId is not set", rs.Run.ID))
	}

	switch expired := clock.Now(ctx).After(deadline.AsTime()); {
	case expired:
		// TODO(yiwzhang): fail if partially submitted.
		// Deadline has already expired. Try to acquire submit queue again
		// and attempt another submission if not waitlisted. Otherwise,
		// falls back to WAITING_FOR_SUBMISSION status.
		switch waitlist, err := acquireSubmitQueue(ctx, rs, rm); {
		case err != nil:
			return nil, err
		case waitlist:
			rs = rs.ShallowCopy()
			rs.Run.Status = run.Status_WAITING_FOR_SUBMISSION
			rs.Run.Submission.Deadline = nil
			rs.Run.Submission.TaskId = ""
			return &Result{State: rs}, nil
		default:
			rs, err := markSubmitting(ctx, rs)
			if err != nil {
				return nil, err
			}
			s, err := constructSubmitter(ctx, rs, rm)
			if err != nil {
				return nil, err
			}
			return &Result{
				State:         rs,
				PostProcessFn: s.submit,
			}, nil
		}
	case taskID == mustTaskIDFromContext(ctx):
		// Matching taskID indicates current task is the retry of a previous
		// submitting task that has failed transiently. Continue the submission.
		s, err := constructSubmitter(ctx, rs, rm)
		if err != nil {
			return nil, err
		}
		return &Result{
			State:         rs,
			PostProcessFn: s.submit,
		}, nil
	default:
		// Presumably another task is working on the submission at this time. So
		// poke as soon as the deadline expires.
		if err := rm.PokeAt(ctx, rs.Run.ID, deadline.AsTime()); err != nil {
			return nil, err
		}
		return &Result{State: rs}, nil
	}
}

const submissionDuration = 20 * time.Minute

func markSubmitting(ctx context.Context, rs *state.RunState) (*state.RunState, error) {
	ret := rs.ShallowCopy()
	ret.Run.Status = run.Status_SUBMITTING
	if ret.Run.Submission == nil {
		ret.Run.Submission = &run.Submission{}
		var err error
		if ret.Run.Submission.Cls, err = orderCLIDsInSubmissionOrder(ctx, ret.Run.CLs, ret.Run.ID, ret.Run.Submission); err != nil {
			return nil, err
		}
	}
	ret.Run.Submission.Deadline = timestamppb.New(clock.Now(ctx).UTC().Add(submissionDuration))
	ret.Run.Submission.AttemptCount += 1
	ret.Run.Submission.TaskId = mustTaskIDFromContext(ctx)
	return ret, nil
}

var fakeTaskIDKey = "used in handler tests only for setting the mock taskID"

func mustTaskIDFromContext(ctx context.Context) string {
	if taskID, ok := ctx.Value(&fakeTaskIDKey).(string); ok {
		return taskID
	}
	switch executionInfo := tq.TaskExecutionInfo(ctx); {
	case executionInfo == nil:
		panic("must be called within a task handler")
	case executionInfo.TaskID == "":
		panic("taskID in task executionInfo is empty")
	default:
		return executionInfo.TaskID
	}
}

func constructSubmitter(ctx context.Context, rs *state.RunState, rm RM) (*submitter, error) {
	cg, err := rs.LoadConfigGroup(ctx)
	if err != nil {
		return nil, err
	}
	submission := rs.Run.Submission
	unsubmittedCLs := make(common.CLIDs, 0, len(submission.GetCls())-len(submission.GetSubmittedCls()))
	submitted := common.MakeCLIDs(submission.GetSubmittedCls()...).Set()
	for _, cl := range submission.GetCls() {
		clid := common.CLID(cl)
		if _, ok := submitted[clid]; !ok {
			unsubmittedCLs = append(unsubmittedCLs, clid)
		}
	}
	return &submitter{
		runID:    rs.Run.ID,
		deadline: submission.GetDeadline().AsTime(),
		attempt:  submission.GetAttemptCount(),
		treeURL:  cg.Content.GetVerifiers().GetTreeStatus().GetUrl(),
		clids:    unsubmittedCLs,
		rm:       rm,
	}, nil
}

// OnCLSubmitted implements Handler interface.
func (*Impl) OnCLSubmitted(ctx context.Context, rs *state.RunState, clids common.CLIDs) (*Result, error) {
	rs = rs.ShallowCopy()
	sub := rs.Run.Submission
	submitted := clids.Set()
	for _, clid := range sub.GetSubmittedCls() {
		submitted[common.CLID(clid)] = struct{}{}
	}
	if sub.GetSubmittedCls() != nil {
		sub.SubmittedCls = sub.SubmittedCls[:0]
	}
	for _, cl := range sub.GetCls() {
		clid := common.CLID(cl)
		if _, ok := submitted[clid]; ok {
			sub.SubmittedCls = append(sub.SubmittedCls, cl)
			delete(submitted, clid)
		}
	}
	if len(submitted) > 0 {
		unexpected := make(sort.IntSlice, 0, len(submitted))
		for clid := range submitted {
			unexpected = append(unexpected, int(clid))
		}
		unexpected.Sort()
		return nil, errors.Reason("received CLSubmitted event for cls not belonging to this Run: %v", unexpected).Err()
	}
	return &Result{State: rs}, nil
}

// OnSubmissionCompleted implements Handler interface.
func (*Impl) OnSubmissionCompleted(ctx context.Context, rs *state.RunState, sr eventpb.SubmissionResult, attempt int32) (*Result, error) {
	panic("implement")
}

func acquireSubmitQueue(ctx context.Context, rs *state.RunState, rm RM) (waitlisted bool, err error) {
	cg, err := rs.LoadConfigGroup(ctx)
	if err != nil {
		return false, err
	}
	now := clock.Now(ctx).UTC()
	rid := rs.Run.ID
	var innerErr error
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		waitlisted, innerErr = submit.TryAcquire(ctx, rm.NotifyReadyForSubmission, rid, cg.SubmitOptions)
		switch {
		case innerErr != nil:
			return innerErr
		case !waitlisted:
			// If not waitlisted, RM will proceed as if ReadyForSubmission event is
			// received. Sends a ReadyForSubmission event 10 seconds later in case
			// the event processing has failed in the middle.
			return rm.NotifyReadyForSubmission(ctx, rid, now.Add(10*time.Second))
		default:
			return nil
		}
	}, nil)
	switch {
	case innerErr != nil:
		return false, innerErr
	case err != nil:
		return false, errors.Annotate(err, "failed to run the transaction to acquire submit queue").Tag(transient.Tag).Err()
	default:
		return waitlisted, nil
	}
}

func orderCLIDsInSubmissionOrder(ctx context.Context, clids common.CLIDs, runID common.RunID, sub *run.Submission) ([]int64, error) {
	cls, err := run.LoadRunCLs(ctx, runID, clids)
	if err != nil {
		return nil, err
	}
	cls, err = submit.ComputeOrder(cls)
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(cls))
	for i, cl := range cls {
		ret[i] = int64(cl.ID)
	}
	return ret, nil
}

type submitter struct {
	// All fields are immutable.

	// runID is the ID of the Run to be submitted.
	runID common.RunID
	// deadline is when this submission should be stopped.
	deadline time.Time
	// attempt is the current submission attempt count.
	attempt int32
	// treeURL is used to check if tree is closed at the beginning
	// of submission.
	treeURL string
	// clids contains ids of cls to be submitted in submission order.
	clids common.CLIDs
	// rm is used to interact with Run Manager.
	rm RM
}

const defaultFatalMsg = "CV failed to submit your change because of " +
	"unexpected internal error. Please contact LUCI team: https://bit.ly/3sMReYs"

// ErrTransientSubmissionFailure indicates that the submission has failed
// transiently and the same task should be retried.
var ErrTransientSubmissionFailure = errors.New("submission failed transiently", transient.Tag)

func (s submitter) submit(ctx context.Context) error {
	sc := &eventpb.SubmissionCompleted{
		Result:  eventpb.SubmissionResult_SUCCEEDED,
		Attempt: s.attempt,
	}
	switch passed, err := s.checkPrecondition(ctx); {
	case err != nil:
		sc = s.computeResultEvent(ctx, err, defaultFatalMsg)
	case !passed:
		sc = &eventpb.SubmissionCompleted{
			Result:  eventpb.SubmissionResult_FAILED_PRECONDITION,
			Attempt: s.attempt,
		}
	default: // precondition check passed
		if cls, err := run.LoadRunCLs(ctx, s.runID, s.clids); err != nil {
			sc = s.computeResultEvent(ctx, err, defaultFatalMsg)
		} else {
			dctx, cancel := clock.WithDeadline(ctx, s.deadline)
			defer cancel()
			if fatalMsg, err := s.submitCLs(dctx, cls); err != nil {
				sc = s.computeResultEvent(ctx, err, fatalMsg)
			}
		}
	}

	if sc.Result == eventpb.SubmissionResult_FAILED_TRANSIENT {
		// Do not release queue for transient failure.
		if err := s.rm.NotifySubmissionCompleted(ctx, s.runID, sc, true); err != nil {
			return err
		}
		return ErrTransientSubmissionFailure
	}
	var innerErr error
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		if innerErr = submit.Release(ctx, s.rm.NotifyReadyForSubmission, s.runID); innerErr != nil {
			return innerErr
		}
		if innerErr = s.rm.NotifySubmissionCompleted(ctx, s.runID, sc, false); innerErr != nil {
			return innerErr
		}
		return nil
	}, nil)
	switch {
	case innerErr != nil:
		return innerErr
	case err != nil:
		return errors.Annotate(err, "failed to release submit queue and notify RM").Tag(transient.Tag).Err()
	}
	// TODO(yiwzhang): optimization for happy path: for successful submission,
	// invoke the RM within the same task to reduce latency.
	return s.rm.Invoke(ctx, s.runID, time.Time{})
}

var perCLRetryFactory retry.Factory = transient.Only(func() retry.Iterator {
	return &retry.ExponentialBackoff{
		Limited: retry.Limited{
			Delay:   200 * time.Millisecond,
			Retries: 10,
		},
		Multiplier: 2,
	}
})

// submitCLs sequentially submits the provided slice of CLs and retries on
// transient failure of submitting individual CL based on `perCLRetryFactory`.
//
// Returns the first fatal error encountered or the first transient error if
// the retry quota has been exhausted. The entire submission will be retried
// later by RM for transient failure. For fatal error, RM will fail the
// submission and post `fatalMsg` on all not-yet-submitted CLs and notify the
// users. Therefore, please be aware of what's included in the `fatalMsg` to
// avoid accidental leak of information.
func (s submitter) submitCLs(ctx context.Context, cls []*run.RunCL) (fatalMsg string, err error) {
	for _, cl := range cls {
		var submitted bool
		err := retry.Retry(ctx, perCLRetryFactory, func() error {
			if !submitted {
				switch err := s.submitCL(ctx, cl); {
				case err == nil:
					submitted = true
				default:
					if fatalMsg = fatalGerritErrMsg(err); fatalMsg != "" {
						// Ensure err is not tagged with transient.
						return transient.Tag.Off().Apply(err)
					}
					return transient.Tag.Apply(err)
				}
			}
			return s.rm.NotifyCLSubmitted(ctx, s.runID, cl.ID)
		}, retry.LogCallback(ctx, fmt.Sprintf("submit cl [id=%d, external_id=%q]", cl.ID, cl.ExternalID)))
		switch {
		case err == nil:
		case fatalMsg != "":
			return fatalMsg, err
		case transient.Tag.In(err):
			return "", err
		default:
			return defaultFatalMsg, err
		}
	}
	return "", nil
}

func (s submitter) checkPrecondition(ctx context.Context) (passed bool, err error) {
	switch cur, err := submit.CurrentRun(ctx, s.runID.LUCIProject()); {
	case err != nil:
		return false, err
	case cur != s.runID:
		logging.Warningf(ctx, "run no longer holds submit queue, currently held by %q", cur)
		return false, nil
	}

	if s.treeURL != "" {
		switch status, err := tree.FetchLatest(ctx, s.treeURL); {
		case err != nil:
			return false, err
		case status.State != tree.Open && status.State != tree.Throttled:
			logging.Warningf(ctx, "tree %q is closed when submission starts", s.treeURL)
			return false, nil
		}
	}

	if clock.Now(ctx).After(s.deadline) {
		logging.Warningf(ctx, "submit deadline has already expired at %s", s.deadline)
		return false, nil
	}
	return true, nil
}

func (s submitter) submitCL(ctx context.Context, cl *run.RunCL) error {
	gc, err := gerrit.CurrentClient(ctx, cl.Detail.GetGerrit().GetHost(), s.runID.LUCIProject())
	if err != nil {
		return err
	}
	ci := cl.Detail.GetGerrit().GetInfo()
	_, submitErr := gc.SubmitRevision(ctx, &gerritpb.SubmitRevisionRequest{
		Number:     ci.GetNumber(),
		RevisionId: ci.GetCurrentRevision(),
		Project:    ci.GetProject(),
	})
	if submitErr == nil {
		return nil
	}
	// Sometimes, Gerrit may return error but change is actually merged.
	// Load the change again to check whether it is actually merged.
	latest, getErr := gc.GetChange(ctx, &gerritpb.GetChangeRequest{
		Number:  ci.GetNumber(),
		Project: ci.GetProject(),
	})
	if getErr == nil && latest.Status == gerritpb.ChangeStatus_MERGED {
		// It is possible that somebody else submitted the change, but this is
		// unlikely enough that we presume CV did it. If necessary, it's possible
		// to examine Change messages to see who actually did it.
		return nil
	}
	return submitErr
}

// TODO(yiwzhang/tandrii): normalize message with the template function
// used in clpurger/user_text.go.
const (
	permDeniedMsg = "CV couldn't submit your CL because CV is not " +
		"allowed to do so in your Gerrit project config. Contact your " +
		"project admin or Chrome Operations team https://goo.gl/f3mzjN"
	failedPreconditionMsgFmt = "Gerrit rejected submission with error: " +
		"%s\nHint: rebasing CL in Gerrit UI and re-submitting through CV " +
		"usually works"
	unexpectedMsgFmt = "CV failed to submit your change because of unexpected error from Gerrit: %s"
)

// fatalGerritErrMsg returns non-empty message if the provided error is fatal.
func fatalGerritErrMsg(err error) string {
	switch grpcutil.Code(err) {
	case codes.PermissionDenied:
		return permDeniedMsg
	case codes.FailedPrecondition:
		// Gerrit returns 409. Either because change can't be merged, or
		// this revision isn't the latest.
		return fmt.Sprintf(failedPreconditionMsgFmt, err)
	case codes.ResourceExhausted, codes.Internal:
		return ""
	default:
		return fmt.Sprintf(unexpectedMsgFmt, err)
	}
}

func (s submitter) computeResultEvent(ctx context.Context, err error, fatalMsg string) *eventpb.SubmissionCompleted {
	switch {
	case err == nil:
		return &eventpb.SubmissionCompleted{
			Result:  eventpb.SubmissionResult_SUCCEEDED,
			Attempt: s.attempt,
		}
	case transient.Tag.In(err):
		errors.Log(ctx, err)
		return &eventpb.SubmissionCompleted{
			Result:  eventpb.SubmissionResult_FAILED_TRANSIENT,
			Attempt: s.attempt,
		}
	default:
		errors.Log(ctx, err)
		return &eventpb.SubmissionCompleted{
			Result:       eventpb.SubmissionResult_FAILED_PERMANENT,
			FatalMessage: fatalMsg,
			Attempt:      s.attempt,
		}
	}
}
