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
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/gerrit"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/state"
	"go.chromium.org/luci/cv/internal/run/impl/submit"
	"go.chromium.org/luci/cv/internal/run/impl/util"
)

// OnReadyForSubmission implements Handler interface.
func (impl *Impl) OnReadyForSubmission(ctx context.Context, rs *state.RunState) (*Result, error) {
	switch status := rs.Status; {
	case run.IsEnded(status):
		// It is safe to discard this event because this event arrives after Run
		// gets cancelled while waiting for submission.
		logging.Debugf(ctx, "received ReadyForSubmission event when Run is %s", status)
		rs = rs.ShallowCopy()
		// Under certain race conditions, this Run may still occupy the submit
		// queue. So, check first without a transaction and then initiate a
		// transaction to release if this Run currently occupies the submit queue.
		if err := releaseSubmitQueueIfTaken(ctx, rs, impl.RM); err != nil {
			return nil, err
		}
		return &Result{State: rs}, nil
	case status == run.Status_SUBMITTING:
		// Discard this event if this Run is currently submitting. If submission
		// is stopped and should be resumed (e.g. transient failure, app crashing),
		// it should be handled in `OnSubmissionCompleted` or `TryResumeSubmission`.
		logging.Debugf(ctx, "received ReadyForSubmission event when Run is submitting")
		return &Result{State: rs}, nil
	case status == run.Status_RUNNING:
		// This may happen when this Run transitioned from RUNNING status to
		// WAITING_FOR_SUBMISSION, prepared for submission but failed to
		// save the state transition. This Run is receiving this event because
		// of the fail-safe task sent while acquiring the Submit Queue. CV should
		// treat this Run as WAITING_FOR_SUBMISSION status.
		rs = rs.ShallowCopy()
		rs.Status = run.Status_WAITING_FOR_SUBMISSION
		fallthrough
	case status == run.Status_WAITING_FOR_SUBMISSION:
		if len(rs.Submission.GetSubmittedCls()) > 0 {
			panic(fmt.Errorf("impossible; Run %q is in Status_WAITING_FOR_SUBMISSION status but has submitted CLs ", rs.ID))
		}
		rs = rs.ShallowCopy()
		switch waitlisted, err := acquireSubmitQueue(ctx, rs, impl.RM); {
		case err != nil:
			return nil, err
		case waitlisted:
			// This Run will be notified by Submit Queue once its turn comes.
			return &Result{State: rs}, nil
		}
		rs.CloneSubmission()
		switch treeOpen, treeErr := rs.CheckTree(ctx, impl.TreeClient); {
		case treeErr != nil && clock.Since(ctx, rs.Submission.TreeErrorSince.AsTime()) > treeStatusFailureTimeLimit:
			// Failed to fetch status for a long time. Fail the Run.
			rims := make(map[common.CLID]reviewInputMeta, len(rs.CLs))
			cg, err := prjcfg.GetConfigGroup(ctx, rs.ID.LUCIProject(), rs.ConfigGroupID)
			if err != nil {
				return nil, err
			}
			if rs.Mode != run.FullRun {
				panic(fmt.Errorf("impossible, %s runs cannot submit CLs", rs.Mode))
			}
			whoms := rs.Mode.GerritNotifyTargets()
			for _, id := range rs.CLs {
				rims[id] = reviewInputMeta{
					notify: whoms,
					// Add the same set of group/people to the attention set.
					addToAttention: whoms,
					reason:         treeStatusCheckFailedReason,
					message:        fmt.Sprintf(persistentTreeStatusAppFailureTemplate, cg.Content.GetVerifiers().GetTreeStatus().GetUrl()),
				}
			}
			scheduleTriggersCancellation(ctx, rs, rims, run.Status_FAILED)
			if err := releaseSubmitQueue(ctx, rs, impl.RM); err != nil {
				return nil, err
			}
			return &Result{
				State: rs,
			}, nil
		case treeErr != nil:
			logging.Warningf(ctx, "tree-status check failed with: %s, retrying in %s", treeErr, treeCheckInterval)
			fallthrough
		case !treeOpen:
			err := parallel.WorkPool(2, func(work chan<- func() error) {
				work <- func() error {
					// Tree is closed or status unknown, revisit after 1 minute.
					return impl.RM.PokeAfter(ctx, rs.ID, treeCheckInterval)
				}
				work <- func() error {
					// Give up the Submit Queue while waiting for tree to open.
					return releaseSubmitQueue(ctx, rs, impl.RM)
				}
			})
			if err != nil {
				return nil, common.MostSevereError(err)
			}
			return &Result{State: rs}, nil
		default:
			if err := markSubmitting(ctx, rs); err != nil {
				return nil, err
			}
			s := submit.NewSubmitter(ctx, rs.ID, rs.Submission, impl.RM, impl.GFactory)
			rs.SubmissionScheduled = true
			return &Result{
				State:         rs,
				PostProcessFn: s.Submit,
			}, nil
		}
	default:
		panic(fmt.Errorf("impossible status %s", status))
	}
}

// OnCLsSubmitted implements Handler interface.
func (*Impl) OnCLsSubmitted(ctx context.Context, rs *state.RunState, clids common.CLIDs) (*Result, error) {
	switch status := rs.Status; {
	case run.IsEnded(status):
		logging.Warningf(ctx, "received CLsSubmitted event when Run is %s", status)
		return &Result{State: rs}, nil
	case status != run.Status_SUBMITTING:
		return nil, errors.Reason("expected SUBMITTING status; got %s", status).Err()
	}
	rs = rs.ShallowCopy()
	rs.Submission = proto.Clone(rs.Submission).(*run.Submission)
	submitted := clids.Set()
	for _, cl := range rs.Submission.GetSubmittedCls() {
		submitted.AddI64(cl)
	}
	if rs.Submission.GetSubmittedCls() != nil {
		rs.Submission.SubmittedCls = rs.Submission.SubmittedCls[:0]
	}
	for _, cl := range rs.Submission.GetCls() {
		if submitted.HasI64(cl) {
			rs.Submission.SubmittedCls = append(rs.Submission.SubmittedCls, cl)
			submitted.DelI64(cl)
		}
	}
	rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
		Time: timestamppb.New(clock.Now(ctx)),
		Kind: &run.LogEntry_ClSubmitted{
			ClSubmitted: &run.LogEntry_CLSubmitted{
				NewlySubmittedCls: common.CLIDsAsInt64s(clids),
				TotalSubmitted:    int64(len(rs.Submission.SubmittedCls)),
			},
		},
	})
	if len(submitted) > 0 {
		unexpected := make(sort.IntSlice, 0, len(submitted))
		for clid := range submitted {
			unexpected = append(unexpected, int(clid))
		}
		unexpected.Sort()
		return nil, errors.Reason("received CLsSubmitted event for cls not belonging to this Run: %v", unexpected).Err()
	}
	return &Result{State: rs}, nil
}

// OnSubmissionCompleted implements Handler interface.
func (impl *Impl) OnSubmissionCompleted(ctx context.Context, rs *state.RunState, sc *eventpb.SubmissionCompleted) (*Result, error) {
	switch status := rs.Status; {
	case run.IsEnded(status):
		logging.Warningf(ctx, "received SubmissionCompleted event when Run is %s", status)
		rs = rs.ShallowCopy()
		if err := releaseSubmitQueueIfTaken(ctx, rs, impl.RM); err != nil {
			return nil, err
		}
		return &Result{State: rs}, nil
	case status != run.Status_SUBMITTING:
		return nil, errors.Reason("expected SUBMITTING status; got %s", status).Err()
	}

	rs = rs.ShallowCopy()
	if sc.GetQueueReleaseTimestamp() != nil {
		rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
			Time: sc.GetQueueReleaseTimestamp(),
			Kind: &run.LogEntry_ReleasedSubmitQueue_{
				ReleasedSubmitQueue: &run.LogEntry_ReleasedSubmitQueue{},
			},
		})
	}
	switch {
	case sc.GetResult() == eventpb.SubmissionResult_SUCCEEDED:
		se := impl.endRun(ctx, rs, run.Status_SUCCEEDED)
		return &Result{
			State:        rs,
			SideEffectFn: se,
		}, nil
	case sc.GetResult() == eventpb.SubmissionResult_FAILED_TRANSIENT:
		rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
			Time: timestamppb.New(clock.Now(ctx)),
			Kind: &run.LogEntry_SubmissionFailure_{
				SubmissionFailure: &run.LogEntry_SubmissionFailure{
					Event: sc,
				},
			},
		})
		return impl.tryResumeSubmission(ctx, rs, sc)
	case sc.GetResult() == eventpb.SubmissionResult_FAILED_PERMANENT:
		if clFailures := sc.GetClFailures(); clFailures != nil {
			failedCLs := make([]int64, len(clFailures.GetFailures()))
			for i, f := range clFailures.GetFailures() {
				failedCLs[i] = f.GetClid()
			}
			rs.Submission = proto.Clone(rs.Submission).(*run.Submission)
			rs.Submission.FailedCls = failedCLs
			rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
				Time: timestamppb.New(clock.Now(ctx)),
				Kind: &run.LogEntry_SubmissionFailure_{
					SubmissionFailure: &run.LogEntry_SubmissionFailure{
						Event: sc,
					},
				},
			})
		}
		cg, err := prjcfg.GetConfigGroup(ctx, rs.ID.LUCIProject(), rs.ConfigGroupID)
		if err != nil {
			return nil, err
		}
		if err := impl.cancelNotSubmittedCLTriggers(ctx, rs, sc, cg); err != nil {
			return nil, err
		}
		se := impl.endRun(ctx, rs, run.Status_FAILED)
		return &Result{
			State:        rs,
			SideEffectFn: se,
		}, nil
	default:
		panic(fmt.Errorf("impossible submission result %s", sc.GetResult()))
	}
}

// TryResumeSubmission implements Handler interface.
func (impl *Impl) TryResumeSubmission(ctx context.Context, rs *state.RunState) (*Result, error) {
	return impl.tryResumeSubmission(ctx, rs, nil)
}

func (impl *Impl) tryResumeSubmission(ctx context.Context, rs *state.RunState, sc *eventpb.SubmissionCompleted) (*Result, error) {
	switch {
	case rs.Status != run.Status_SUBMITTING || rs.SubmissionScheduled:
		return &Result{State: rs}, nil
	case sc != nil && sc.Result != eventpb.SubmissionResult_FAILED_TRANSIENT:
		panic(fmt.Errorf("submission can only be resumed on nil submission completed event or event reporting transient failure; got %s", sc))
	}

	deadline := rs.Submission.GetDeadline()
	taskID := rs.Submission.GetTaskId()
	switch {
	case deadline == nil:
		panic(fmt.Errorf("impossible: run %q is submitting but Run.Submission.Deadline is not set", rs.ID))
	case taskID == "":
		panic(fmt.Errorf("impossible: run %q is submitting but Run.Submission.TaskId is not set", rs.ID))
	}

	switch expired := clock.Now(ctx).After(deadline.AsTime()); {
	case expired:
		rs = rs.ShallowCopy()
		var status run.Status
		switch submittedCnt := len(rs.Submission.GetSubmittedCls()); {
		case submittedCnt > 0 && submittedCnt == len(rs.Submission.GetCls()):
			// Fully submitted
			status = run.Status_SUCCEEDED
		default: // None submitted or partially submitted
			status = run.Status_FAILED
			// Make a submission completed event with permanent failure.
			if clFailures := sc.GetClFailures(); clFailures != nil {
				rs.Submission = proto.Clone(rs.Submission).(*run.Submission)
				rs.Submission.FailedCls = make([]int64, len(clFailures.GetFailures()))
				sc = &eventpb.SubmissionCompleted{
					Result: eventpb.SubmissionResult_FAILED_PERMANENT,
					FailureReason: &eventpb.SubmissionCompleted_ClFailures{
						ClFailures: &eventpb.SubmissionCompleted_CLSubmissionFailures{
							Failures: make([]*eventpb.SubmissionCompleted_CLSubmissionFailure, len(clFailures.GetFailures())),
						},
					},
				}
				for i, f := range clFailures.GetFailures() {
					rs.Submission.FailedCls[i] = f.GetClid()
					sc.GetClFailures().Failures[i] = &eventpb.SubmissionCompleted_CLSubmissionFailure{
						Clid:    f.GetClid(),
						Message: fmt.Sprintf("CL failed to submit because of transient failure: %s. However, submission is running out of time to retry.", f.GetMessage()),
					}
				}
			} else {
				sc = &eventpb.SubmissionCompleted{
					Result: eventpb.SubmissionResult_FAILED_PERMANENT,
					FailureReason: &eventpb.SubmissionCompleted_Timeout{
						Timeout: true,
					},
				}
			}
			cg, err := prjcfg.GetConfigGroup(ctx, rs.ID.LUCIProject(), rs.ConfigGroupID)
			if err != nil {
				return nil, err
			}

			if err := impl.cancelNotSubmittedCLTriggers(ctx, rs, sc, cg); err != nil {
				return nil, err
			}
		}
		if err := releaseSubmitQueueIfTaken(ctx, rs, impl.RM); err != nil {
			return nil, err
		}
		se := impl.endRun(ctx, rs, status)
		return &Result{
			State:        rs,
			SideEffectFn: se,
		}, nil
	case taskID == mustTaskIDFromContext(ctx):
		// Matching taskID indicates current task is the retry of a previous
		// submitting task that has failed transiently. Continue the submission.
		rs = rs.ShallowCopy()
		s := submit.NewSubmitter(ctx, rs.ID, rs.Submission, impl.RM, impl.GFactory)
		rs.SubmissionScheduled = true
		return &Result{
			State:         rs,
			PostProcessFn: s.Submit,
		}, nil
	default:
		// Presumably another task is working on the submission at this time.
		// So, wake up RM as soon as the submission expires. Meanwhile, don't
		// consume the event as the retries of submission task will process
		// this event. It's probably a race condition that this task sees this
		// event first.
		if err := impl.RM.Invoke(ctx, rs.ID, deadline.AsTime()); err != nil {
			return nil, err
		}
		return &Result{
			State:          rs,
			PreserveEvents: true,
		}, nil
	}
}

func acquireSubmitQueue(ctx context.Context, rs *state.RunState, rm RM) (waitlisted bool, err error) {
	cg, err := prjcfg.GetConfigGroup(ctx, rs.ID.LUCIProject(), rs.ConfigGroupID)
	if err != nil {
		return false, err
	}
	now := clock.Now(ctx).UTC()
	var innerErr error
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		waitlisted, innerErr = submit.TryAcquire(ctx, rm.NotifyReadyForSubmission, rs.ID, cg.SubmitOptions)
		switch {
		case innerErr != nil:
			return innerErr
		case !waitlisted:
			// It is possible that RM fails before successfully completing the state
			// transition. In that case, this Run will block Submit Queue infinitely.
			// Sending a ReadyForSubmission event after 10s as a fail-safe to ensure
			// Run keeps making progress.
			return rm.NotifyReadyForSubmission(ctx, rs.ID, now.Add(10*time.Second))
		default:
			return nil
		}
	}, nil)
	switch {
	case innerErr != nil:
		return false, innerErr
	case err != nil:
		return false, errors.Annotate(err, "failed to run the transaction to acquire submit queue").Tag(transient.Tag).Err()
	case waitlisted:
		rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
			Time: timestamppb.New(clock.Now(ctx)),
			Kind: &run.LogEntry_Waitlisted_{
				Waitlisted: &run.LogEntry_Waitlisted{},
			},
		})
		logging.Debugf(ctx, "Waitlisted in Submit Queue")
		return true, nil
	default:
		rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
			Time: timestamppb.New(clock.Now(ctx)),
			Kind: &run.LogEntry_AcquiredSubmitQueue_{
				AcquiredSubmitQueue: &run.LogEntry_AcquiredSubmitQueue{},
			},
		})
		logging.Debugf(ctx, "Acquired Submit Queue")
		return false, nil
	}
}

// releaseSubmitQueueIfTaken checks if submit queue is occupied by the given
// Run before trying to release.
func releaseSubmitQueueIfTaken(ctx context.Context, rs *state.RunState, rm RM) error {
	switch current, waitlist, err := submit.LoadCurrentAndWaitlist(ctx, rs.ID); {
	case err != nil:
		return err
	case current == rs.ID:
		return releaseSubmitQueue(ctx, rs, rm)
	default:
		for _, w := range waitlist {
			if w == rs.ID {
				return releaseSubmitQueue(ctx, rs, rm)
			}
		}
	}
	return nil
}

func releaseSubmitQueue(ctx context.Context, rs *state.RunState, rm RM) error {
	var innerErr error
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		innerErr = submit.Release(ctx, rm.NotifyReadyForSubmission, rs.ID)
		return innerErr
	}, nil)
	switch {
	case innerErr != nil:
		return innerErr
	case err != nil:
		return errors.Annotate(err, "failed to release submit queue").Tag(transient.Tag).Err()
	}
	rs.LogEntries = append(rs.LogEntries, &run.LogEntry{
		Time: timestamppb.New(clock.Now(ctx)),
		Kind: &run.LogEntry_ReleasedSubmitQueue_{
			ReleasedSubmitQueue: &run.LogEntry_ReleasedSubmitQueue{},
		},
	})
	logging.Debugf(ctx, "Released Submit Queue")
	return nil
}

const submissionDuration = 20 * time.Minute

func markSubmitting(ctx context.Context, rs *state.RunState) error {
	rs.Status = run.Status_SUBMITTING
	var err error
	if rs.Submission.Cls, err = orderCLIDsInSubmissionOrder(ctx, rs.CLs, rs.ID, rs.Submission); err != nil {
		return err
	}
	rs.Submission.Deadline = timestamppb.New(clock.Now(ctx).UTC().Add(submissionDuration))
	rs.Submission.TaskId = mustTaskIDFromContext(ctx)
	return nil
}

func (impl *Impl) cancelNotSubmittedCLTriggers(ctx context.Context, rs *state.RunState, sc *eventpb.SubmissionCompleted, cg *prjcfg.ConfigGroup) error {
	allCLIDs := common.MakeCLIDs(rs.Submission.GetCls()...)
	allRunCLs, err := run.LoadRunCLs(ctx, rs.ID, allCLIDs)
	if err != nil {
		return err
	}
	whoms := rs.Mode.GerritNotifyTargets()
	meta := reviewInputMeta{
		notify: whoms,
		// Add the same set of group/people to the attention set.
		addToAttention: whoms,
		reason:         submissionFailureAttentionReason,
	}
	runCLExternalIDs := make([]changelist.ExternalID, len(allRunCLs))
	for i, runCL := range allRunCLs {
		runCLExternalIDs[i] = runCL.ExternalID
	}

	// Single-CL Run
	if len(allRunCLs) == 1 {
		switch {
		case sc.GetClFailures() != nil:
			failures := sc.GetClFailures().GetFailures()
			if len(failures) != 1 {
				panic(fmt.Errorf("expected exactly 1 failed CL, got %v", failures))
			}
			meta.message = failures[0].GetMessage()
		case sc.GetTimeout():
			meta.message = timeoutMsg
		default:
			meta.message = defaultMsg
		}
		return impl.cancelCLTriggers(ctx, rs.ID, allRunCLs, runCLExternalIDs, cg, meta)
	}

	// Multi-CL Run
	submitted, failed, pending := splitRunCLs(allRunCLs, rs.Submission, sc)
	msgSuffix := makeSubmissionMsgSuffix(submitted, failed, pending)
	switch {
	case sc.GetClFailures() != nil:
		var wg sync.WaitGroup
		errs := make(errors.MultiError, len(allRunCLs))
		// cancel triggers of CLs that fail to submit.
		messages := make(map[common.CLID]string, len(sc.GetClFailures().GetFailures()))
		for _, f := range sc.GetClFailures().GetFailures() {
			messages[common.CLID(f.GetClid())] = f.GetMessage()
		}
		for i, failedCL := range failed {
			i, failedCL := i, failedCL
			meta := meta
			wg.Add(1)
			go func() {
				defer wg.Done()
				meta.message = fmt.Sprintf("%s\n\n%s", messages[failedCL.ID], msgSuffix)
				errs[i] = impl.cancelCLTriggers(ctx, rs.ID, []*run.RunCL{failedCL}, runCLExternalIDs, cg, meta)
			}()
		}
		// Cancel triggers of CLs that CV won't try to submit.
		var sb strings.Builder
		// TODO(yiwzhang): Once CV learns how to submit multiple CLs in parallel,
		// this should be optimized to print out failed CLs that each pending CL
		// depends on instead of printing out all failed CLs.
		// Example: considering a CL group where CL B and CL C are submitted in
		// parallel and neither of them succeeds:
		//   A (submitted)
		//   |
		//   |--> B (failed) --> D (pending)
		//   |
		//   |--> C (failed) --> E (pending)
		// the message CV posts on CL D should only include the fact that CV fails
		// to submit CL B.
		for _, f := range failed {
			fmt.Fprintf(&sb, "\n* %s", f.ExternalID.MustURL())
		}
		fmt.Fprint(&sb, "\n\n")
		fmt.Fprint(&sb, msgSuffix)
		meta.message = fmt.Sprintf("%s%s", partiallySubmittedMsgForPendingCLs, sb.String())
		for i, pendingCL := range pending {
			i, pendingCL := i, pendingCL
			wg.Add(1)
			go func() {
				defer wg.Done()
				errs[len(failed)+i] = impl.cancelCLTriggers(ctx, rs.ID, []*run.RunCL{pendingCL}, runCLExternalIDs, cg, meta)
			}()
		}

		msg := fmt.Sprintf("%s%s", partiallySubmittedMsgForSubmittedCLs, sb.String())
		for i, rcl := range submitted {
			i, rcl := i, rcl
			wg.Add(1)
			go func() {
				defer wg.Done()
				errs[len(failed)+len(pending)+i] = postMsgForDependentFailures(ctx, impl.GFactory, rcl, msg)
			}()
		}

		wg.Wait()
		return common.MostSevereError(errs)
	case sc.GetTimeout():
		meta.message = fmt.Sprintf("%s\n\n%s", timeoutMsg, msgSuffix)
		return impl.cancelCLTriggers(ctx, rs.ID, pending, runCLExternalIDs, cg, meta)
	default:
		meta.message = fmt.Sprintf("%s\n\n%s", defaultMsg, msgSuffix)
		return impl.cancelCLTriggers(ctx, rs.ID, pending, runCLExternalIDs, cg, meta)
	}
}

func makeSubmissionMsgSuffix(submitted, failed, pending []*run.RunCL) string {
	submittedURLs := make([]string, len(submitted))
	for i, cl := range submitted {
		submittedURLs[i] = cl.ExternalID.MustURL()
	}
	notSubmittedURLs := make([]string, len(failed)+len(pending))
	for i, cl := range failed {
		notSubmittedURLs[i] = cl.ExternalID.MustURL()
	}
	for i, cl := range pending {
		notSubmittedURLs[len(failed)+i] = cl.ExternalID.MustURL()
	}
	if len(submittedURLs) > 0 { // partial submission
		return fmt.Sprintf(partiallySubmittedMsgSuffixFmt,
			strings.Join(notSubmittedURLs, "\n* "),
			strings.Join(submittedURLs, "\n* "),
		)
	}
	return fmt.Sprintf(noneSubmittedMsgSuffixFmt, strings.Join(notSubmittedURLs, "\n* "))
}

////////////////////////////////////////////////////////////////////////////////
// Helper methods

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

func splitRunCLs(cls []*run.RunCL, submission *run.Submission, sc *eventpb.SubmissionCompleted) (submitted, failed, pending []*run.RunCL) {
	submittedSet := common.MakeCLIDsSet(submission.GetSubmittedCls()...)
	failedSet := make(common.CLIDsSet, len(sc.GetClFailures().GetFailures()))
	for _, f := range sc.GetClFailures().GetFailures() {
		if submittedSet.HasI64(f.GetClid()) {
			panic(fmt.Errorf("impossible; cl %d is marked both submitted and failed", f.GetClid()))
		}
		failedSet.AddI64(f.GetClid())
	}

	submitted = make([]*run.RunCL, 0, len(submittedSet))
	failed = make([]*run.RunCL, 0, len(failedSet))
	pending = make([]*run.RunCL, 0, len(cls)-len(submittedSet)-len(failedSet))
	for _, cl := range cls {
		switch {
		case submittedSet.Has(cl.ID):
			submitted = append(submitted, cl)
		case failedSet.Has(cl.ID):
			failed = append(failed, cl)
		default:
			pending = append(pending, cl)
		}
	}
	return submitted, failed, pending
}

// TODO(crbug/1302119): Replace terms like "Project admin" with dedicated
// contact sourced from Project Config.
const (
	cvBugLink  = "https://bugs.chromium.org/p/chromium/issues/entry?components=Infra%3ELUCI%3EBuildService%3EPreSubmit%3ECV"
	defaultMsg = "Submission of this CL failed due to unexpected internal " +
		"error. Please contact LUCI team.\n\n" + cvBugLink
	noneSubmittedMsgSuffixFmt = "None of the CLs in the Run has been " +
		"submitted. CLs:\n* %s"
	partiallySubmittedMsgForPendingCLs = "This CL is not submitted because " +
		"submission has failed for the following CL(s) which this CL depends on."
	partiallySubmittedMsgForSubmittedCLs = "This CL is submitted. However, " +
		"submission has failed for the following CL(s) which depend on this CL."
	partiallySubmittedMsgSuffixFmt = "CLs in the Run have been submitted " +
		"partially.\nNot submitted:\n* %s\nSubmitted:\n* %s\n" +
		"Please, use your judgement to determine if already submitted CLs have " +
		"to be reverted, or if the remaining CLs could be manually submitted. " +
		"If you think the partially submitted CLs may have broken the " +
		"tip-of-tree of your project, consider notifying your infrastructure " +
		"team/gardeners/sheriffs."
	timeoutMsg = "Ran out of time to submit this CL. " +
		// TODO(yiwzhang): Generally, time out means CV is doing something
		// wrong and looping over internally, However, timeout could also
		// happen when submitting large CL stack and Gerrit is slow. In that
		// case, CV can't do anything about it. After launching m1, gather data
		// to see under what circumstance it may happen and revise this message
		// so that CV doesn't get blamed for timeout it isn't responsible for.
		"Please contact LUCI team.\n\n" + cvBugLink
	persistentTreeStatusAppFailureTemplate = "Could not submit this CL " +
		"because the tree status app at %s repeatedly returned failures. "
	treeStatusCheckFailedReason      = "Tree status check failed."
	submissionFailureAttentionReason = "Submission failed."
)

// postMsgForDependentFailures posts a review message to
// a given CL to notify submission failures of the dependent CLs.
func postMsgForDependentFailures(ctx context.Context, gf gerrit.Factory, rcl *run.RunCL, msg string) error {
	queryOpts := []gerritpb.QueryOption{gerritpb.QueryOption_MESSAGES}
	posted, err := util.IsActionTakenOnGerritCL(ctx, gf, rcl, queryOpts, func(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
		// In practice, Gerrit generally orders the messages from earliest to
		// latest. So iterating in the reverse order because it's more likely the
		// desired message is posted recently. Also don't visit any messages before
		// the run trigger as those messages should belong to previous Runs.
		clTriggeredAt := rcl.Trigger.Time.AsTime()
		for i := len(ci.GetMessages()) - 1; i >= 0; i-- {
			m := ci.GetMessages()[i]
			switch t := m.GetDate().AsTime(); {
			case t.Before(clTriggeredAt):
				// i-th message is too old, no need to check even older ones.
				return time.Time{}
			case strings.Contains(m.GetMessage(), msg):
				return t
			}
		}
		return time.Time{}
	})

	switch {
	case err != nil:
		return err
	case !posted.IsZero():
		return nil
	}

	ci := rcl.Detail.GetGerrit().GetInfo()
	ownerAndVotersAccounts := gerrit.Whoms{gerrit.Owner, gerrit.CQVoters}.ToAccountIDsSorted(ci)
	req := &gerritpb.SetReviewRequest{
		Number:     ci.GetNumber(),
		Project:    ci.GetProject(),
		RevisionId: ci.GetCurrentRevision(),
		Message:    msg,
		Tag:        run.FullRun.GerritMessageTag(),
		Notify:     gerritpb.Notify_NOTIFY_NONE,
		NotifyDetails: &gerritpb.NotifyDetails{
			Recipients: []*gerritpb.NotifyDetails_Recipient{
				{
					RecipientType: gerritpb.NotifyDetails_RECIPIENT_TYPE_TO,
					Info: &gerritpb.NotifyDetails_Info{
						Accounts: ownerAndVotersAccounts,
					},
				},
			},
		},
		AddToAttentionSet: make([]*gerritpb.AttentionSetInput, len(ownerAndVotersAccounts)),
	}
	reason := fmt.Sprintf("ps#%d: failed to submit dependent CLs",
		ci.GetRevisions()[ci.GetCurrentRevision()].GetNumber())
	for i, acct := range ownerAndVotersAccounts {
		req.AddToAttentionSet[i] = &gerritpb.AttentionSetInput{
			User:   strconv.Itoa(int(acct)),
			Reason: reason,
		}
	}
	return util.MutateGerritCL(ctx, gf, rcl, req, 1*time.Minute, "post-msg-for-dependent-failure")
}
