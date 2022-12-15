// Copyright 2022 The LUCI Authors.
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

package aggrmetrics

import (
	"context"

	bbutil "go.chromium.org/luci/buildbucket/protoutil"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/common/tsmon/types"
	"go.chromium.org/luci/hardcoded/chromeinfra"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/metrics"
	"go.chromium.org/luci/cv/internal/tryjob"
)

type builderPresenceAggregator struct {
	env *common.Env
}

// metrics implements aggregator interface.
func (t *builderPresenceAggregator) metrics() []types.Metric {
	return []types.Metric{metrics.Public.TryjobBuilderPresence}
}

// report implements aggregator interface.
func (t *builderPresenceAggregator) report(ctx context.Context, projects []string) error {
	err := parallel.WorkPool(min(8, len(projects)), func(work chan<- func() error) {
		for _, project := range projects {
			project := project
			work <- func() error {
				meta, err := prjcfg.GetLatestMeta(ctx, project)
				switch {
				case err != nil:
					return err
				case meta.Status != prjcfg.StatusEnabled:
					// race condition: project gets disabled right before loading the
					// config
					return nil
				}
				cgs, err := meta.GetConfigGroups(ctx)
				if err != nil {
					return err
				}
				for _, cg := range cgs {
					if err := reportBuilders(ctx, t.env, project, cg); err != nil {
						return err
					}
				}
				return nil
			}
		}
	})
	return err
}

func reportBuilders(ctx context.Context, env *common.Env, project string, cg *prjcfg.ConfigGroup) error {
	cgName := cg.Content.GetName()
	for _, b := range cg.Content.GetVerifiers().GetTryjob().GetBuilders() {
		builderID, err := bbutil.ParseBuilderID(b.GetName())
		if err != nil {
			return err
		}
		// Synthesize definition.
		def := &tryjob.Definition{
			Backend: &tryjob.Definition_Buildbucket_{
				Buildbucket: &tryjob.Definition_Buildbucket{
					Host:    chromeinfra.BuildbucketHost,
					Builder: builderID,
				},
			},
		}
		tryjob.RunWithBuilderMetricsTarget(ctx, env, def, func(ctx context.Context) {
			metrics.Public.TryjobBuilderPresence.Set(ctx, true,
				project,
				cgName,
				b.GetIncludableOnly(),
				len(b.GetLocationFilters()) > 0,
				b.GetExperimentPercentage() > 0,
			)
		})
	}
	return nil
}
