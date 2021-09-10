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

package e2e

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"

	cvbqpb "go.chromium.org/luci/cv/api/bigquery/v1"
	migrationpb "go.chromium.org/luci/cv/api/migration"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO(tandrii): this is a slow test (~0.6s on my laptop),
// but it will become faster once LoadGerritRuns is optimized.
func TestConcurentRunsSingular(t *testing.T) {
	t.Parallel()

	Convey("CV juggles a bunch of concurrent Runs", t, func() {
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChangeFirst = 1001
		const N = 50

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)

		// Prepare a bunch of actions to play out over time.
		actions := make([]struct {
			gChange     int
			user        string
			mode        run.Mode
			triggerTime time.Time // ever-increasing
			finishTime  time.Time // pseudo-random
			finalStatus run.Status
		}, N)
		var expectSubmitted, expectFinished, expectFailed []int
		for i := range actions {
			a := &actions[i]
			a.gChange = gChangeFirst + i
			a.user = fmt.Sprintf("user-%d", (i%10)+1)
			ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
				a.gChange, gf.Project(gRepo), gf.Ref(gRef), gf.PS(1), gf.Owner(a.user),
				gf.Updated(ct.Clock.Now()),
			)))
			a.mode = run.DryRun
			if i%3 == 0 {
				a.mode = run.FullRun
			}
			a.triggerTime = ct.Clock.Now().Add(time.Duration(i*3) * time.Second)
			a.finishTime = a.triggerTime.Add(time.Duration(i*13%5) * time.Minute)
			if i%2 == 0 {
				a.finalStatus = run.Status_SUCCEEDED
				if a.mode == run.FullRun {
					expectSubmitted = append(expectSubmitted, a.gChange)
				} else {
					expectFinished = append(expectFinished, a.gChange)
				}
			} else {
				a.finalStatus = run.Status_FAILED
				expectFailed = append(expectFailed, a.gChange)
			}
		}
		indexesByFinishTime := make([]int, len(actions))
		for i := range actions {
			indexesByFinishTime[i] = i
		}
		sort.Slice(indexesByFinishTime, func(i, j int) bool {
			idxI, idxJ := indexesByFinishTime[i], indexesByFinishTime[j]
			return actions[idxI].finishTime.Before(actions[idxJ].finishTime)
		})

		// Start CQDaemon and make it obey finishAt and finalStatus.
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				gChange := r.GetAttempt().GetGerritChanges()[0].GetChange()
				a := actions[gChange-gChangeFirst]
				if ct.Clock.Now().Before(a.finishTime) {
					return r
				}
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				if a.finalStatus == run.Status_SUCCEEDED {
					r.Attempt.Status = cvbqpb.AttemptStatus_SUCCESS
				} else {
					r.Attempt.Status = cvbqpb.AttemptStatus_FAILURE
				}
				return r
			},
		)

		ct.LogPhase(ctx, fmt.Sprintf("Triggering CQ on %d CLs", len(actions)))
		for i := range actions {
			a := &actions[i]
			if !ct.Clock.Now().After(a.triggerTime) {
				ct.RunUntil(ctx, func() bool { return ct.Clock.Now().After(a.triggerTime) })
			}
			ct.GFake.MutateChange(gHost, a.gChange, func(c *gf.Change) {
				val := 1
				if a.mode == run.FullRun {
					val = 2
				}
				gf.CQ(val, a.triggerTime, gf.U(a.user))(c.Info)
				gf.Updated(a.triggerTime)(c.Info)
			})
		}

		// Now iterate in increasing finishAt, checking state of Gerrit CL.
		var actualSubmitted, actualFinished, actualFailed, actualWeird []int
		for _, i := range indexesByFinishTime {
			a := actions[i]
			ct.LogPhase(ctx, fmt.Sprintf("Checking state of #%d %s expected state %s", a.gChange, a.mode, a.finalStatus))
			var runs []*run.Run
			ct.RunUntil(ctx, func() bool {
				if !ct.Clock.Now().After(a.finishTime) {
					return false
				}
				runs = ct.LoadGerritRuns(ctx, gHost, int64(a.gChange), lProject)
				return len(runs) > 0
			})
			So(runs, ShouldHaveLength, 1)
			r := runs[0]
			ct.RunUntil(ctx, func() bool {
				r = ct.LoadRun(ctx, runs[0].ID)
				return run.IsEnded(r.Status)
			})
			switch {
			case r.Status == run.Status_FAILED:
				actualFailed = append(actualFailed, a.gChange)
				So(ct.MaxCQVote(ctx, gHost, int64(a.gChange)), ShouldEqual, 0)
			case r.Status == run.Status_SUCCEEDED && a.mode == run.DryRun:
				actualFinished = append(actualFinished, a.gChange)
				So(ct.MaxCQVote(ctx, gHost, int64(a.gChange)), ShouldEqual, 0)
			case r.Status == run.Status_SUCCEEDED && a.mode == run.FullRun:
				actualSubmitted = append(actualSubmitted, a.gChange)
				So(ct.GFake.GetChange(gHost, a.gChange).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			default:
				actualWeird = append(actualWeird, a.gChange)
			}
			So(r.CreateTime, ShouldEqual, datastore.RoundTime(a.triggerTime.UTC()))
			So(r.EndTime, ShouldHappenAfter, a.finishTime)
		}

		sort.Sort(sort.IntSlice(actualSubmitted))
		sort.Sort(sort.IntSlice(actualFailed))
		sort.Sort(sort.IntSlice(actualFinished))
		So(actualSubmitted, ShouldResemble, expectSubmitted)
		So(actualFailed, ShouldResemble, expectFailed)
		So(actualFinished, ShouldResemble, expectFinished)
		So(actualWeird, ShouldBeEmpty)

		So(ct.LoadRunsOf(ctx, lProject), ShouldHaveLength, len(actions))

		ct.LogPhase(ctx, "Wait for all BQ exports to complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == len(actions) })

		ct.LogPhase(ctx, "Wait for all PubSub messages are sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == len(actions) })
	})
}
