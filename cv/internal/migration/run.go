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

package migration

import (
	"context"
	"time"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"

	cvbqpb "go.chromium.org/luci/cv/api/bigquery/v1"
	migrationpb "go.chromium.org/luci/cv/api/migration"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/run"
)

func fetchActiveRuns(ctx context.Context, project string) ([]*migrationpb.Run, error) {
	runs := []run.Run{}
	q := run.NewQueryWithLUCIProject(ctx, project).Eq("Status", run.Status_RUNNING)
	if err := datastore.GetAll(ctx, q, &runs); err != nil {
		return nil, errors.Annotate(err, "fetch Run entities").Tag(transient.Tag).Err()
	}
	numRuns := len(runs)
	if numRuns == 0 {
		return nil, nil
	}
	poolSize := numRuns
	if poolSize > 20 {
		poolSize = 20
	}
	ret := make([]*migrationpb.Run, numRuns)
	err := parallel.WorkPool(poolSize, func(workCh chan<- func() error) {
		for i, r := range runs {
			i, r := i, r
			workCh <- func() error {
				runKey := datastore.MakeKey(ctx, run.RunKind, string(r.ID))
				runCLs := make([]run.RunCL, len(r.CLs))
				for i, cl := range r.CLs {
					runCLs[i] = run.RunCL{
						ID:  cl,
						Run: runKey,
					}
				}
				if err := datastore.Get(ctx, runCLs); err != nil {
					return errors.Annotate(err, "fetch CLs for run %q", r.ID).Tag(transient.Tag).Err()
				}
				mcls := make([]*migrationpb.RunCL, len(runCLs))
				mode := r.Mode.BQAttemptMode()
				for i, cl := range runCLs {
					trigger := &migrationpb.RunCL_Trigger{
						Email:     cl.Trigger.GetEmail(),
						Time:      cl.Trigger.GetTime(),
						AccountId: cl.Trigger.GetGerritAccountId(),
					}
					mcl := &migrationpb.RunCL{
						Id: int64(cl.ID),
						Gc: &cvbqpb.GerritChange{
							Host:                       cl.Detail.GetGerrit().GetHost(),
							Project:                    cl.Detail.GetGerrit().GetInfo().GetProject(),
							Change:                     cl.Detail.GetGerrit().GetInfo().GetNumber(),
							Patchset:                   int64(cl.Detail.GetPatchset()),
							EarliestEquivalentPatchset: int64(cl.Detail.GetMinEquivalentPatchset()),
							Mode:                       mode,
						},
						Files:   cl.Detail.GetGerrit().GetFiles(),
						Info:    cl.Detail.GetGerrit().GetInfo(),
						Trigger: trigger,
						Deps:    make([]*migrationpb.RunCL_Dep, len(cl.Detail.GetDeps())),
					}
					for i, dep := range cl.Detail.GetDeps() {
						mcl.Deps[i] = &migrationpb.RunCL_Dep{
							Id: dep.GetClid(),
						}
						if dep.GetKind() == changelist.DepKind_HARD {
							mcl.Deps[i].Hard = true
						}
					}
					mcls[i] = mcl
				}
				ret[i] = &migrationpb.Run{
					Attempt: &cvbqpb.Attempt{
						LuciProject: project,
					},
					Id:  string(r.ID),
					Cls: mcls,
				}
				return nil
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// VerifiedCQDRun is the Run reported by CQDaemon after verification
// completes.
type VerifiedCQDRun struct {
	_kind string `gae:"$kind,migration.VerifiedCQDRun"`
	// ID is ID of this Run in CV.
	ID common.RunID `gae:"$id"`
	// Payload is what CQDaemon has reported.
	Payload *migrationpb.ReportVerifiedRunRequest
}

// FinishedCQDRun contains info about a finished Run reported by the CQDaemon.
//
// To be removed after the first milestone is reached.
type FinishedCQDRun struct {
	_kind string `gae:"$kind,migration.FinishedCQDRun"`
	// AttemptKey is the CQD ID of the Run.
	//
	// Once CV starts creating Runs, the CV's Run for the same Run will contain
	// the AttemptKey as a substring.
	AttemptKey string `gae:"$id"`
	// RunID may be set if CQD is aware of the RunID.
	//
	// For example, if milestone 1 migration is rolled back, some attempts may
	// have associated RunID.
	//
	// Although the CV RunID, if known, is also stored in the Payload,
	// a separate field is necessary for Datastore indexing.
	RunID common.RunID
	// RecordTime is when this entity was inserted.
	UpdateTime time.Time `gae:",noindex"`
	// Everything that CQD has sent.
	Payload *migrationpb.Run
}

func saveFinishedCQDRun(ctx context.Context, mr *migrationpb.Run) error {
	key := mr.GetAttempt().GetKey()
	try := 0
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		try++
		f := FinishedCQDRun{AttemptKey: key}
		switch err := datastore.Get(ctx, &f); {
		case err == datastore.ErrNoSuchEntity:
			// expected.
		case err != nil:
			return err
		default:
			logging.Warningf(ctx, "Overwriting FinishedCQDRun %q in %d-th try", key, try)
		}
		f = FinishedCQDRun{
			AttemptKey: key,
			UpdateTime: datastore.RoundTime(clock.Now(ctx).UTC()),
			Payload:    mr,
		}
		if id := f.Payload.GetId(); id != "" {
			f.RunID = common.RunID(id)
		}
		return datastore.Put(ctx, &f)
	}, nil)
	return errors.Annotate(err, "failed to record FinishedCQDRun %q after %d tries", key, try).Tag(transient.Tag).Err()
}
