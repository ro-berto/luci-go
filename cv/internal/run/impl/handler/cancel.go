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

package handler

import (
	"context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/impl/state"
)

// Cancel implements Handler interface.
func (impl *Impl) Cancel(ctx context.Context, rs *state.RunState) (*Result, error) {
	switch status := rs.Run.Status; {
	case status == run.Status_STATUS_UNSPECIFIED:
		err := errors.Reason("CRITICAL: can't cancel a Run with unspecified status").Err()
		common.LogError(ctx, err)
		panic(err)
	case status == run.Status_SUBMITTING:
		logging.Debugf(ctx, "Run cancellation can't be fulfilled at this time as Run is currently submitting.")
		// Don't consume the events so that the RM executing the submission will
		// be able to read the Cancel events and attempt to cancel if the Run
		// failed to submit.
		return &Result{State: rs, PreserveEvents: true}, nil
	case run.IsEnded(status):
		logging.Debugf(ctx, "skip cancellation because Run has already ended.")
		return &Result{State: rs}, nil
	}

	rs = rs.ShallowCopy()
	se := endRun(ctx, rs, run.Status_CANCELLED, impl.PM)
	if rs.Run.StartTime.IsZero() {
		// This run has never started but already gets a cancelled event.
		rs.Run.StartTime = rs.Run.EndTime
	}
	return &Result{
		State:        rs,
		SideEffectFn: se,
	}, nil
}
