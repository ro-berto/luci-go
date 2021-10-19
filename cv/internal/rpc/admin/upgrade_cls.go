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

package admin

import (
	"context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/dsmapper"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/run"
)

var upgradeCLConfig = dsmapper.JobConfig{
	Mapper: "runcl-description",
	Query: dsmapper.Query{
		Kind: "RunCL",
	},
	PageSize:   32,
	ShardCount: 4,
}

var upgradeCLFactory = func(_ context.Context, j *dsmapper.Job, _ int) (dsmapper.Mapper, error) {
	tsJobName := string(j.Config.Mapper)
	tsJobID := int64(j.ID)

	upgradeCLs := func(ctx context.Context, keys []*datastore.Key) error {
		needUpgrade := func(cls []*run.RunCL) []*run.RunCL {
			toUpdate := cls[:0]
			for _, cl := range cls {
				ci := cl.Detail.GetGerrit().GetInfo()
				if ci == nil {
					continue
				}
				revInfo := ci.GetRevisions()[ci.GetCurrentRevision()]
				if revInfo == nil {
					continue
				}
				if revInfo.GetCommit().GetMessage() == "" {
					continue
				}
				toUpdate = append(toUpdate, cl)
			}
			return toUpdate
		}

		cls := make([]*run.RunCL, len(keys))
		for i, k := range keys {
			cls[i] = &run.RunCL{
				ID:  common.CLID(k.IntID()),
				Run: k.Parent(),
			}
		}

		// Check before a transaction if an update is even necessary.
		if err := datastore.Get(ctx, cls); err != nil {
			return errors.Annotate(err, "failed to fetch RunCLs").Tag(transient.Tag).Err()
		}
		cls = needUpgrade(cls)
		if len(cls) == 0 {
			return nil
		}

		err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
			// Reload inside transaction to avoid races with other CV parts.
			if err := datastore.Get(ctx, cls); err != nil {
				return errors.Annotate(err, "failed to fetch RunCLs").Tag(transient.Tag).Err()
			}
			cls = needUpgrade(cls)
			if len(cls) == 0 {
				return nil
			}
			for _, cl := range cls {
				changelist.RemoveUnusedGerritInfo(cl.Detail.GetGerrit().GetInfo())
			}
			return datastore.Put(ctx, cls)
		}, nil)
		if err != nil {
			return errors.Annotate(err, "failed to update RunCLs").Tag(transient.Tag).Err()
		}
		metricUpgraded.Add(ctx, int64(len(cls)), tsJobName, tsJobID, "RunCL")
		return nil
	}

	return upgradeCLs, nil
}
