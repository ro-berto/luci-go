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

package configcron

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
)

// PM encapsulates Project Manager notified by the ConfigRefresher.
//
// In production, this will be prjmanager.Notifier.
type PM interface {
	Poke(ctx context.Context, luciProject string) error
	UpdateConfig(ctx context.Context, luciProject string) error
}

type ProjectConfigRefresher struct {
	pm  PM
	tqd *tq.Dispatcher
}

// New creates new ProjectConfigRefresher and registers its TQ tasks.
func New(tqd *tq.Dispatcher, pm PM) *ProjectConfigRefresher {
	pcr := &ProjectConfigRefresher{pm, tqd}
	pcr.tqd.RegisterTaskClass(tq.TaskClass{
		ID:        "refresh-project-config",
		Prototype: &RefreshProjectConfigTask{},
		Queue:     "refresh-project-config",
		Kind:      tq.NonTransactional,
		Quiet:     true,
		Handler: func(ctx context.Context, payload proto.Message) error {
			task := payload.(*RefreshProjectConfigTask)
			if err := pcr.refreshProject(ctx, task.GetProject(), task.GetDisable()); err != nil {
				// Never retry tasks because the refresh task is submitted every minute
				// by the AppEngine Cron.
				err = tq.Fatal.Apply(err)
				return common.TQIfy{}.Error(ctx, err)
			}
			return nil
		},
	})
	return pcr
}

// SubmitRefreshTasks submits tasks that update config for LUCI projects
// or disable projects that do not have CV config in LUCI Config.
//
// It's expected to be called by a cron.
func (pcr *ProjectConfigRefresher) SubmitRefreshTasks(ctx context.Context) error {
	projects, err := prjcfg.ProjectsWithConfig(ctx)
	if err != nil {
		return err
	}

	// Consider only some projects, regardless of which projects are registered.
	// TODO(crbug/1158505): switch to -dev configs.
	if common.IsDev(ctx) {
		projects = []string{"infra", "cq-test"}
	} else {
		for i, p := range projects {
			if p == "cq-test" {
				projects = append(projects[:i], projects[i+1:]...)
				break
			}
		}
	}

	tasks := make([]*tq.Task, len(projects))
	for i, p := range projects {
		tasks[i] = &tq.Task{
			Title: "update/" + p,
			Payload: &RefreshProjectConfigTask{
				Project: p,
			},
		}
	}

	curEnabledProjects, err := prjcfg.GetAllProjectIDs(ctx, true)
	if err != nil {
		return err
	}
	projectsInLUCIConfig := stringset.NewFromSlice(projects...)
	for _, p := range curEnabledProjects {
		if !projectsInLUCIConfig.Has(p) {
			tasks = append(tasks, &tq.Task{
				Title: "disable/" + p,
				Payload: &RefreshProjectConfigTask{
					Project: p,
					Disable: true,
				},
			})
		}
	}

	err = parallel.WorkPool(32, func(workCh chan<- func() error) {
		for _, task := range tasks {
			task := task
			workCh <- func() (err error) {
				if err = pcr.tqd.AddTask(ctx, task); err != nil {
					logging.Errorf(ctx, "Failed to submit task for %q: %s", task.Title, err)
				}
				return
			}
		}
	})

	if err != nil {
		return err.(errors.MultiError).First()
	}
	return nil
}

func (pcr *ProjectConfigRefresher) refreshProject(ctx context.Context, project string, disable bool) error {
	action, actionFn := "update", prjcfg.UpdateProject
	if disable {
		action, actionFn = "disable", prjcfg.DisableProject
	}
	err := actionFn(ctx, project, func(ctx context.Context) error {
		return pcr.pm.UpdateConfig(ctx, project)
	})
	if err != nil {
		return errors.Annotate(err, "failed to %s project %q", action, project).Err()
	}
	if !disable {
		return pcr.maybePokePM(ctx, project)
	}
	return nil
}

const pokePMInterval = 10 * time.Minute

func (pcr *ProjectConfigRefresher) maybePokePM(ctx context.Context, project string) error {
	now := clock.Now(ctx).UTC()
	offset := common.ProjectOffset("cron-poke", pokePMInterval, project)
	nextPokeETA := now.Truncate(pokePMInterval).Add(offset)
	if nextPokeETA.Before(now) {
		nextPokeETA = nextPokeETA.Add(pokePMInterval)
	}

	// Cron runs every minute on average and triggers RefreshProjectConfigTask,
	// which may be delayed, so send iff it's less than 1.5 minutes before next
	// poke. This will sometimes result in 2 pokes sent instead of 1, but pokes
	// are less likely to not be sent at all.
	if nextPokeETA.Sub(now) < 90*time.Second {
		return pcr.pm.Poke(ctx, project)
	}
	return nil
}