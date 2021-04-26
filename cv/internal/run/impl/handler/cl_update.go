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

	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/impl/state"
)

// OnCLUpdated implements Handler interface.
func (impl *Impl) OnCLUpdated(ctx context.Context, rs *state.RunState, clids common.CLIDs) (*Result, error) {
	switch status := rs.Run.Status; {
	case status == run.Status_STATUS_UNSPECIFIED:
		err := errors.Reason("CRITICAL: Received CLUpdated events but Run is in unspecified status").Err()
		common.LogError(ctx, err)
		panic(err)
	case status == run.Status_SUBMITTING:
		// Don't consume the events so that the RM executing the submission will
		// be able to read the CLUpdated events and take necessary actions after
		// submission completes. For example, a new PS is uploaded for one of
		// the unsubmitted CLs and cause the run submission to fail. RM should
		// cancel this Run instead of retrying.
		return &Result{State: rs, PreserveEvents: true}, nil
	case run.IsEnded(status):
		// Run is ended, update on CL shouldn't change the Run state.
		return &Result{State: rs}, nil
	}
	clids.Dedupe()

	cls := make([]*changelist.CL, len(clids))
	runCLs := make([]*run.RunCL, len(clids))
	runKey := datastore.MakeKey(ctx, run.RunKind, string(rs.Run.ID))
	for i, clid := range clids {
		cls[i] = &changelist.CL{ID: clid}
		runCLs[i] = &run.RunCL{ID: clid, Run: runKey}
	}

	if err := loadCLsAndRunCLs(ctx, cls, runCLs); err != nil {
		return nil, err
	}

	cg, err := rs.LoadConfigGroup(ctx)
	if err != nil {
		return nil, err
	}

	for i := range clids {
		switch cl, runCL := cls[i], runCLs[i]; {
		case cl.Snapshot.GetPatchset() > runCL.Detail.GetPatchset():
			// New PS discovered.
			return impl.Cancel(ctx, rs)
		case trigger.Find(cl.Snapshot.GetGerrit().GetInfo(), cg.Content) == nil:
			// Trigger has been removed.
			return impl.Cancel(ctx, rs)

		default:
			// TODO(crbug/1202270): handle some or all cases of changing trigger,
			// e.g. changing mode OR changing user.
		}
	}
	return &Result{State: rs}, nil
}

func loadCLsAndRunCLs(ctx context.Context, cls []*changelist.CL, runCLs []*run.RunCL) error {
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := datastore.Get(ectx, cls)
		switch merr, ok := err.(errors.MultiError); {
		case ok:
			for i, err := range merr {
				if err == datastore.ErrNoSuchEntity {
					return errors.Reason("CL(%d) doesn't exist", cls[i].ID).Err()
				}
			}
			n, first := merr.Summary()
			return errors.Annotate(first, "failed to load %d/%d CLs", n, len(cls)).Tag(transient.Tag).Err()
		case err != nil:
			return errors.Annotate(err, "failed to load CLs").Tag(transient.Tag).Err()
		}
		return nil
	})
	eg.Go(func() error {
		err := datastore.Get(ectx, runCLs)
		switch merr, ok := err.(errors.MultiError); {
		case ok:
			for i, err := range merr {
				if err == datastore.ErrNoSuchEntity {
					return errors.Reason("RunCL(%d) doesn't exist", runCLs[i].ID).Err()
				}
			}
			n, first := merr.Summary()
			return errors.Annotate(first, "failed to load %d/%d RunCLs", n, len(runCLs)).Tag(transient.Tag).Err()
		case err != nil:
			return errors.Annotate(err, "failed to load RunCLs").Tag(transient.Tag).Err()
		}
		return nil
	})
	return eg.Wait()
}
