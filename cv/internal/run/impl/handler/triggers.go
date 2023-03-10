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
	"time"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/state"
)

const (
	// maxTriggersCancellationDuration is the maximum duration allowed for a Run
	// to cancel the triggers of all CLs.
	maxTriggersCancellationDuration = 5 * time.Minute

	logEntryLabelTriggerCancellation = "Trigger Cancellation"
)

func (impl *Impl) onCompletedCancelTriggers(ctx context.Context, rs *state.RunState, op *run.OngoingLongOps_Op, opCompleted *eventpb.LongOpCompleted) (*Result, error) {
	opID := opCompleted.GetOperationId()
	rs = rs.ShallowCopy()
	rs.RemoveCompletedLongOp(opID)
	if status := rs.Status; run.IsEnded(status) || status == run.Status_SUBMITTING {
		logging.Warningf(ctx, "long operation to cancel triggers has completed but Run is %s. Cancellation Result: %s", rs.Status, opCompleted)
		return &Result{State: rs}, nil
	}
	runStatus := op.GetCancelTriggers().GetRunStatusIfSucceeded()
	switch opCompleted.GetStatus() {
	case eventpb.LongOpCompleted_EXPIRED:
		runStatus = run.Status_FAILED
		rs.LogInfof(ctx, logEntryLabelTriggerCancellation, "failed to cancel the triggers of CLs within the %s deadline", maxTriggersCancellationDuration)
	case eventpb.LongOpCompleted_FAILED:
		runStatus = run.Status_FAILED
		fallthrough
	case eventpb.LongOpCompleted_SUCCEEDED:
		for _, result := range opCompleted.GetCancelTriggers().GetResults() {
			changeURL := changelist.ExternalID(result.GetExternalId()).MustURL()
			switch result.GetDetail().(type) {
			case *eventpb.LongOpCompleted_CancelTriggers_Result_SuccessInfo:
				rs.LogInfofAt(result.GetSuccessInfo().GetCancelledAt().AsTime(), logEntryLabelTriggerCancellation, "successfully cancelled the trigger of change %s", changeURL)
			case *eventpb.LongOpCompleted_CancelTriggers_Result_FailureInfo:
				rs.LogInfof(ctx, logEntryLabelTriggerCancellation, "failed to cancel the trigger of change %s. Reason: %s", changeURL, result.GetFailureInfo().GetFailureMessage())
			default:
				panic(fmt.Errorf("unexpected long op result status: %s", opCompleted.GetStatus()))
			}
		}
	default:
		panic(fmt.Errorf("unexpected LongOpCompleted status: %s", opCompleted.GetStatus()))
	}
	return &Result{
		State:        rs,
		SideEffectFn: impl.endRun(ctx, rs, runStatus),
	}, nil
}
