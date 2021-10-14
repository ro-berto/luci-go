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
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"

	cvbqpb "go.chromium.org/luci/cv/api/bigquery/v1"
	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	migrationpb "go.chromium.org/luci/cv/api/migration"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	"go.chromium.org/luci/cv/internal/configs/validation"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/migration"
	"go.chromium.org/luci/cv/internal/migration/cqdfake"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/pubsub"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestCreatesSingularRun(t *testing.T) {
	t.Parallel()

	Convey("CV creates 1 CL Run, which gets canceled by the user", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 33

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		cfg.CqStatusHost = validation.CQStatusHostPublic
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart), gf.CQ(+2, tStart, gf.U("user-2")),
		)))

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV notices CL and starts the Run")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(ct.LoadGerritCL(ctx, gHost, gChange).IncompleteRuns.ContainsSorted(r.ID), ShouldBeTrue)

		// Under normal conditions, this would be done by the same TQ handler that
		// creates the Run, but with flaky Datastore this may come late.
		ct.LogPhase(ctx, "CV Project Manager incorporates Run into component")
		ct.RunUntil(ctx, func() bool {
			cs := ct.LoadProject(ctx, lProject).State.GetComponents()
			return len(cs) == 1 && len(cs[0].GetPruns()) == 1
		})
		So(ct.LoadProject(ctx, lProject).State.GetComponents()[0].GetPruns()[0].GetId(), ShouldResemble, string(r.ID))

		ct.LogPhase(ctx, "CV posts starting message")
		ct.RunUntil(ctx, func() bool {
			return ct.LastMessage(gHost, gChange) != nil
		})
		So(ct.LastMessage(gHost, gChange).GetMessage(), ShouldContainSubstring,
			fmt.Sprintf("CQ is trying the patch.\n\nFollow status at: https://luci-change-verifier.appspot.com/ui/run/%s\n\nBot data: {", r.ID))

		ct.LogPhase(ctx, "User cancels the Run")
		ct.GFake.MutateChange(gHost, gChange, func(c *gf.Change) {
			gf.CQ(0, ct.Clock.Now(), "user-2")(c.Info)
			gf.Updated(ct.Clock.Now())(c.Info)
		})

		ct.LogPhase(ctx, "CV cancels the Run")
		ct.RunUntil(ctx, func() bool {
			r2 := ct.LoadRun(ctx, r.ID)
			p := ct.LoadProject(ctx, lProject)
			return r2.Status == run.Status_CANCELLED && len(p.State.GetComponents()) == 0
		})

		/////////////////////////    Verify    ////////////////////////////////
		So(ct.LoadGerritCL(ctx, gHost, gChange).IncompleteRuns.ContainsSorted(r.ID), ShouldBeFalse)
		So(ct.LoadProject(ctx, lProject).State.GetPcls(), ShouldBeEmpty)
		So(ct.LoadProject(ctx, lProject).State.GetComponents(), ShouldBeEmpty)
	})
}

func TestCreatesSingularQuickDryRunSuccess(t *testing.T) {
	t.Parallel()

	Convey("CV creates 1 CL Quick Dry Run, which succeeds", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 33
		const quickLabel = "Quick-Label"

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		cfg.GetConfigGroups()[0].AdditionalModes = []*cfgpb.Mode{{
			Name:            string(run.QuickDryRun),
			CqLabelValue:    1,
			TriggeringValue: 1,
			TriggeringLabel: quickLabel,
		}}
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+1, tStart, gf.U("user-2")),
			gf.Vote(quickLabel, +1, tStart, gf.U("user-2")),
			// Spurious vote from user-3.
			gf.Vote(quickLabel, +5, tStart.Add(-10*time.Second), gf.U("user-3")),
		)))

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV notices CL and starts the Run")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(r.Mode, ShouldEqual, run.QuickDryRun)

		ct.LogPhase(ctx, "CQDaemon posts starting message to the Gerrit CL")
		ct.RunUntil(ctx, func() bool {
			m := ct.LastMessage(gHost, gChange).GetMessage()
			return strings.Contains(m, cqdfake.StartingMessage) && strings.Contains(m, string(r.ID))
		})

		ct.LogPhase(ctx, "CQDaemon decides that QuickDryRun has passed and notifies CV")
		ct.Clock.Add(time.Minute)
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				r.Attempt.Status = cvbqpb.AttemptStatus_SUCCESS
				r.Attempt.Substatus = cvbqpb.AttemptSubstatus_NO_SUBSTATUS
				return r
			},
		)
		ct.RunUntil(ctx, func() bool {
			return nil == datastore.Get(ctx, &migration.VerifiedCQDRun{ID: r.ID})
		})

		ct.LogPhase(ctx, "CV finalizes the run and sends BQ and pubsub events")
		var finalRun *run.Run
		ct.RunUntil(ctx, func() bool {
			finalRun = ct.LoadRun(ctx, r.ID)
			proj := ct.LoadProject(ctx, lProject)
			cl := ct.LoadGerritCL(ctx, gHost, gChange)
			return run.IsEnded(finalRun.Status) &&
				len(proj.State.GetComponents()) == 0 &&
				!cl.IncompleteRuns.ContainsSorted(r.ID)
		})

		// Fail quickly iff PM created a new Run on stale CL data.
		So(ct.LoadRunsOf(ctx, lProject), ShouldHaveLength, 1)
		So(finalRun.Status, ShouldEqual, run.Status_SUCCEEDED)
		So(ct.MaxCQVote(ctx, gHost, gChange), ShouldEqual, 0)
		So(ct.MaxVote(ctx, gHost, gChange, quickLabel), ShouldEqual, 0)
		So(ct.LastMessage(gHost, gChange).GetMessage(), ShouldContainSubstring, "Quick dry run: This CL passed")

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 1 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 1 })
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    r.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_SUCCEEDED,
			Eversion:    int64(finalRun.EVersion),
		})
	})
}

func TestCreatesSingularQuickDryRunThenUpgradeToFullRunFailed(t *testing.T) {
	t.Parallel()

	Convey("CV creates 1 CL Quick Dry Run first and then upgrades to Full Run", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 33
		const quickLabel = "Quick-Label"

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		cfg.GetConfigGroups()[0].AdditionalModes = []*cfgpb.Mode{{
			Name:            string(run.QuickDryRun),
			CqLabelValue:    1,
			TriggeringValue: 1,
			TriggeringLabel: quickLabel,
		}}
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+1, tStart, gf.U("user-2")),
			gf.Vote(quickLabel, +1, tStart, gf.U("user-2")),
		)))

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV notices CL and starts the Run")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var qdr *run.Run
		ct.RunUntil(ctx, func() bool {
			qdr = ct.EarliestCreatedRunOf(ctx, lProject)
			return qdr != nil && qdr.Status == run.Status_RUNNING
		})
		So(qdr.Mode, ShouldEqual, run.QuickDryRun)

		ct.LogPhase(ctx, "User upgrades to Full Run but doesn't unvote Quick-Label")
		tStart2 := ct.Clock.Now()
		ct.GFake.MutateChange(gHost, gChange, func(c *gf.Change) {
			gf.CQ(+2, tStart2, gf.U("user-2"))(c.Info)
			gf.Updated(tStart2)(c.Info)
		})
		So(ct.MaxVote(ctx, gHost, gChange, quickLabel), ShouldEqual, 1) // vote stays
		var fr *run.Run
		ct.RunUntil(ctx, func() bool {
			runs := ct.LoadRunsOf(ctx, lProject)
			if len(runs) > 2 {
				panic("impossible; more than 2 runs created.")
			}
			for _, r := range runs {
				if r.ID == qdr.ID {
					qdr = r
				} else {
					fr = r
				}
			}
			return qdr.Status == run.Status_CANCELLED && (fr != nil && fr.Status == run.Status_RUNNING)
		})
		So(fr.Mode, ShouldEqual, run.FullRun)

		ct.LogPhase(ctx, "CQDaemon decides that Full Run has failed and notifies CV")
		ct.Clock.Add(time.Minute)
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				r.Attempt.Status = cvbqpb.AttemptStatus_FAILURE
				r.Attempt.Substatus = cvbqpb.AttemptSubstatus_NO_SUBSTATUS
				return r
			},
		)
		ct.RunUntil(ctx, func() bool {
			return nil == datastore.Get(ctx, &migration.VerifiedCQDRun{ID: fr.ID})
		})

		ct.LogPhase(ctx, "CV finalizes the run")
		var finalRun *run.Run
		ct.RunUntil(ctx, func() bool {
			finalRun = ct.LoadRun(ctx, fr.ID)
			return finalRun != nil && run.IsEnded(finalRun.Status)
		})

		So(ct.LoadRunsOf(ctx, lProject), ShouldHaveLength, 2)
		So(finalRun.Status, ShouldEqual, run.Status_FAILED)
		So(ct.MaxCQVote(ctx, gHost, gChange), ShouldEqual, 0)
		// Removed the stale Quick-Run vote.
		So(ct.MaxVote(ctx, gHost, gChange, quickLabel), ShouldEqual, 0)

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 2 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 2 })
		// 1st message is for the dry-run cancelled.
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    qdr.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_CANCELLED,
			Eversion:    int64(ct.LoadRun(ctx, qdr.ID).EVersion),
		})
		// 2nd message is for the the full-run failed.
		So(ct.RunEndedPubSubTasks()[1].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    fr.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_FAILED,
			Eversion:    int64(ct.LoadRun(ctx, fr.ID).EVersion),
		})
	})
}

func TestCreatesSingularFullRunSuccess(t *testing.T) {
	t.Parallel()

	Convey("CV creates 1 CL Full Run, which succeeds", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 33
		const gPatchSet = 6

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-2")),
		)))

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV notices CL and starts the Run")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(r.Mode, ShouldEqual, run.FullRun)

		ct.LogPhase(ctx, "CQDaemon posts starting message to the Gerrit CL")
		ct.RunUntil(ctx, func() bool {
			m := ct.LastMessage(gHost, gChange).GetMessage()
			return strings.Contains(m, cqdfake.StartingMessage) && strings.Contains(m, string(r.ID))
		})

		ct.LogPhase(ctx, "CQDaemon decides that FullRun has passed and notifies CV to submit")
		ct.Clock.Add(time.Minute)
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				r.Attempt.Status = cvbqpb.AttemptStatus_SUCCESS
				r.Attempt.Substatus = cvbqpb.AttemptSubstatus_NO_SUBSTATUS
				return r
			},
		)
		ct.RunUntil(ctx, func() bool {
			res, err := datastore.Exists(ctx, &migration.VerifiedCQDRun{ID: r.ID})
			return err == nil && res.All()
		})

		ct.LogPhase(ctx, "CV submits the run and sends BQ and Pubsub events")
		var finalRun *run.Run
		ct.RunUntil(ctx, func() bool {
			finalRun = ct.LoadRun(ctx, r.ID)
			proj := ct.LoadProject(ctx, lProject)
			cl := ct.LoadGerritCL(ctx, gHost, gChange)
			return run.IsEnded(finalRun.Status) &&
				len(proj.State.GetComponents()) == 0 &&
				!cl.IncompleteRuns.ContainsSorted(r.ID)
		})

		So(finalRun.Status, ShouldEqual, run.Status_SUCCEEDED)
		ci := ct.GFake.GetChange(gHost, gChange).Info
		So(ci.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
		So(ci.GetRevisions()[ci.GetCurrentRevision()].GetNumber(), ShouldEqual, int32(gPatchSet+1))

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 1 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 1 })
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    r.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_SUCCEEDED,
			Eversion:    int64(ct.LoadRun(ctx, r.ID).EVersion),
		})
	})
}

func TestCreatesSingularDryRunAborted(t *testing.T) {
	t.Parallel()

	Convey("CV creates 1 CL Run, which gets canceled by the user", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 33

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+1, tStart, gf.U("user-2")),
		)))

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV notices CL and starts the Run")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(r.Mode, ShouldEqual, run.DryRun)

		ct.LogPhase(ctx, "CQDaemon posts starting message to the Gerrit CL")
		ct.RunUntil(ctx, func() bool {
			m := ct.LastMessage(gHost, gChange).GetMessage()
			return strings.Contains(m, cqdfake.StartingMessage) && strings.Contains(m, string(r.ID))
		})

		ct.LogPhase(ctx, "User aborts the run by removing a vote")
		ct.Clock.Add(time.Minute)
		ct.GFake.MutateChange(gHost, gChange, func(c *gf.Change) {
			gf.ResetVotes(c.Info, trigger.CQLabelName)
			gf.Updated(ct.Clock.Now())(c.Info)
		})

		ct.LogPhase(ctx, "CQDaemon stops working on the Run")
		ct.RunUntil(ctx, func() bool {
			return len(ct.MustCQD(ctx, lProject).ActiveAttemptKeys()) == 0
		})

		ct.LogPhase(ctx, "CV finalizes the Run and sends BQ event")
		ct.RunUntil(ctx, func() bool {
			r = ct.LoadRun(ctx, r.ID)
			return run.IsEnded(r.Status)
		})
		So(r.Status, ShouldEqual, run.Status_CANCELLED)

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 1 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 1 })
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    r.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_CANCELLED,
			Eversion:    int64(ct.LoadRun(ctx, r.ID).EVersion),
		})
	})
}

func TestCreatesSingularRunWithDeps(t *testing.T) {
	Convey("CV creates Run in singular config in presence of Git dependencies.", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		ct.LogPhase(ctx, "Set Git chain of 13 depends on 12, and vote CQ+1 on 13")
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			12, gf.Project(gRepo), gf.Ref(gRef), gf.PS(2),
		)))
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			13, gf.Project(gRepo), gf.Ref(gRef), gf.PS(3),
			gf.Updated(tStart),
			gf.CQ(+1, tStart, gf.U("user-1")),
		)))
		ct.GFake.SetDependsOn(gHost, "13_3", "12_2")

		ct.LogPhase(ctx, "CV starts dry Run on just the 13")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		var r13 *run.Run
		ct.RunUntil(ctx, func() bool {
			r13 = ct.EarliestCreatedRunOf(ctx, lProject)
			return r13 != nil && r13.Status == run.Status_RUNNING
		})
		So(r13.Mode, ShouldResemble, run.DryRun)
		So(r13.CLs, ShouldResemble, common.CLIDs{ct.LoadGerritCL(ctx, gHost, 13).ID})

		ct.LogPhase(ctx, "User votes CQ+2 on 12, CV starts Run on just 12 and doesn't touch 13")
		ct.Clock.Add(time.Minute)
		ct.GFake.MutateChange(gHost, 12, func(c *gf.Change) {
			gf.CQ(+2, ct.Clock.Now(), "user-1")(c.Info)
			gf.Updated(ct.Clock.Now())(c.Info)
		})
		ct.RunUntil(ctx, func() bool {
			return len(ct.LoadGerritCL(ctx, gHost, 12).IncompleteRuns) > 0
		})
		r12 := ct.LoadRun(ctx, ct.LoadGerritCL(ctx, gHost, 12).IncompleteRuns[0])
		So(r12.Mode, ShouldResemble, run.FullRun)
		So(r12.CLs, ShouldResemble, common.CLIDs{ct.LoadGerritCL(ctx, gHost, 12).ID})
		So(ct.LoadRun(ctx, r13.ID).Status, ShouldEqual, run.Status_RUNNING)
		So(r12.CreateTime, ShouldHappenAfter, r13.CreateTime)

		ct.LogPhase(ctx, "User upgrades to CQ+2 on 13, which terminates old Run.\n"+
			"Due to 13 depending on not yet submitted 12, CV purges 13 by removing CQ+2 vote")
		ct.Clock.Add(time.Minute)
		ct.GFake.MutateChange(gHost, 13, func(c *gf.Change) {
			gf.CQ(+2, ct.Clock.Now(), "user-1")(c.Info)
			gf.Updated(ct.Clock.Now())(c.Info)
		})
		ct.RunUntil(ctx, func() bool {
			return ct.MaxCQVote(ctx, gHost, 13) == 0
		})
	})
}

func TestCreatesMultiCLsFullRunSuccess(t *testing.T) {
	t.Parallel()

	Convey("CV creates 3 CLs Full Run, which succeeds", t, func() {
		/////////////////////////    Setup   ////////////////////////////////
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange1 = 11
		const gChange2 = 22
		const gChange3 = 33
		const gPatchSet = 6

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgCombinable("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		// Git Relationship: <base> -> ci2 -> ci3 -> ci1
		// and ci2 depends on ci1 using Cq-Depend git footer which forms a cycle.
		ci1 := gf.CI(
			gChange1, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-2")),
			gf.Desc(fmt.Sprintf("This is the first CL\nCq-Depend: %d", gChange3)),
		)
		ci2 := gf.CI(
			gChange2, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-2")),
		)
		ci3 := gf.CI(
			gChange3, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-2")),
		)
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), ci1, ci2, ci3))
		ct.GFake.SetDependsOn(gHost, ci3, ci2)
		ct.GFake.SetDependsOn(gHost, ci1, ci3)

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV discovers all CLs")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		clids := make(common.CLIDs, 3)
		ct.RunUntil(ctx, func() bool {
			for i, change := range []int64{gChange1, gChange2, gChange3} {
				cl := ct.LoadGerritCL(ctx, gHost, change)
				if cl == nil {
					return false
				}
				clids[i] = cl.ID
			}
			sort.Sort(clids)
			return true
		})

		ct.LogPhase(ctx, "CV starts a Run with all 3 CLs")
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(r.Mode, ShouldEqual, run.FullRun)
		runCLIDs := r.CLs
		sort.Sort(runCLIDs)
		So(r.CLs, ShouldResemble, clids)

		ct.LogPhase(ctx, "CQDaemon posts starting message to each Gerrit CLs")
		ct.RunUntil(ctx, func() bool {
			for _, change := range []int64{gChange1, gChange2, gChange3} {
				m := ct.LastMessage(gHost, change).GetMessage()
				if !strings.Contains(m, cqdfake.StartingMessage) || !strings.Contains(m, string(r.ID)) {
					return false
				}
			}
			return true
		})

		ct.LogPhase(ctx, "CQDaemon decides that FullRun has passed and notifies CV to submit")
		ct.Clock.Add(time.Minute)
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				r.Attempt.Status = cvbqpb.AttemptStatus_SUCCESS
				r.Attempt.Substatus = cvbqpb.AttemptSubstatus_NO_SUBSTATUS
				return r
			},
		)
		ct.RunUntil(ctx, func() bool {
			res, err := datastore.Exists(ctx, &migration.VerifiedCQDRun{ID: r.ID})
			return err == nil && res.All()
		})

		ct.LogPhase(ctx, "CV submits the run and sends BQ and pubsub events")
		var finalRun *run.Run
		ct.RunUntil(ctx, func() bool {
			finalRun = ct.LoadRun(ctx, r.ID)
			if !run.IsEnded(finalRun.Status) {
				return false
			}
			if proj := ct.LoadProject(ctx, lProject); len(proj.State.GetComponents()) > 0 {
				return false
			}
			for _, change := range []int64{gChange1, gChange2, gChange3} {
				cl := ct.LoadGerritCL(ctx, gHost, change)
				if cl.IncompleteRuns.ContainsSorted(r.ID) {
					return false
				}
			}
			return true
		})

		So(finalRun.Status, ShouldEqual, run.Status_SUCCEEDED)
		ci1 = ct.GFake.GetChange(gHost, gChange1).Info
		ci2 = ct.GFake.GetChange(gHost, gChange2).Info
		ci3 = ct.GFake.GetChange(gHost, gChange3).Info
		for _, ci := range []*gerritpb.ChangeInfo{ci1, ci2, ci3} {
			So(ci.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			So(ci.GetRevisions()[ci.GetCurrentRevision()].GetNumber(), ShouldEqual, int32(gPatchSet+1))
		}
		// verify submission order: [ci2, ci3, ci1]
		So(ci2.GetUpdated().AsTime(), ShouldHappenBefore, ci3.GetUpdated().AsTime())
		So(ci3.GetUpdated().AsTime(), ShouldHappenBefore, ci1.GetUpdated().AsTime())

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 1 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 1 })
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    r.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_SUCCEEDED,
			Eversion:    int64(ct.LoadRun(ctx, r.ID).EVersion),
		})
	})
}

func TestCreatesSingularFullRunWithAllowOpenDeps(t *testing.T) {
	t.Parallel()

	Convey("CV submits stack of 3 CLs in singular mode with allow_submit_with_open_deps=true", t, func() {
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange1 = 11
		const gChange2 = 22
		const gChange3 = 33
		const gChange4 = 44
		const gPatchSet = 1

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		cfg.GetConfigGroups()[0].Verifiers = &cfgpb.Verifiers{
			GerritCqAbility: &cfgpb.Verifiers_GerritCQAbility{
				AllowSubmitWithOpenDeps: true,
			},
		}
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		// Git Relationship: <base> <- ci1 <- ci2 <- ci3 <- ci4
		ci1 := gf.CI(
			gChange1, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.Desc("This is the bottom of the stack"),
		)
		ci2 := gf.CI(
			gChange2, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.Desc("This is the 2nd CL of the stack"),
		)
		ci3 := gf.CI(
			gChange3, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.Desc("This is the 3nd CL of the stack, which was CQ-ed"),
			gf.CQ(+2, tStart, gf.U("user-2")),
		)
		ci4 := gf.CI(
			gChange4, gf.Project(gRepo), gf.Ref(gRef),
			gf.PS(gPatchSet),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.Desc("This is the top of the stack, not CQ-ed"),
		)
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), ci1, ci2, ci3, ci4))
		ct.GFake.SetDependsOn(gHost, ci2, ci1)
		ct.GFake.SetDependsOn(gHost, ci3, ci2)
		ct.GFake.SetDependsOn(gHost, ci4, ci3)

		/////////////////////////    Run CV   ////////////////////////////////
		ct.LogPhase(ctx, "CV discovers the CQ-ed CL")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		ct.RunUntil(ctx, func() bool { return ct.LoadGerritCL(ctx, gHost, gChange3) != nil })
		ct.LogPhase(ctx, "CV loads deps of the CQ-ed CL")
		ct.RunUntil(ctx, func() bool { return ct.LoadGerritCL(ctx, gHost, gChange2) != nil })
		ct.RunUntil(ctx, func() bool { return ct.LoadGerritCL(ctx, gHost, gChange1) != nil })

		ct.LogPhase(ctx, "CV starts a Run on the CQ-ed CL only")
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})
		So(r.Mode, ShouldEqual, run.FullRun)
		So(r.CLs, ShouldResemble, common.CLIDs{ct.LoadGerritCL(ctx, gHost, gChange3).ID})

		ct.LogPhase(ctx, "CQDaemon decides that FullRun has passed and notifies CV to submit")
		ct.Clock.Add(time.Minute)
		ct.MustCQD(ctx, lProject).SetVerifyClbk(
			func(r *migrationpb.ReportedRun) *migrationpb.ReportedRun {
				r = proto.Clone(r).(*migrationpb.ReportedRun)
				r.Attempt.Status = cvbqpb.AttemptStatus_SUCCESS
				r.Attempt.Substatus = cvbqpb.AttemptSubstatus_NO_SUBSTATUS
				return r
			},
		)
		ct.RunUntil(ctx, func() bool {
			res, err := datastore.Exists(ctx, &migration.VerifiedCQDRun{ID: r.ID})
			return err == nil && res.All()
		})

		ct.LogPhase(ctx, "CV submits the Run, which results in 3 CLs actually landing")
		ct.RunUntil(ctx, func() bool {
			r = ct.LoadRun(ctx, r.ID)
			return run.IsEnded(r.Status)
		})
		// CQ-ed was submitted directly.
		So(ct.GFake.GetChange(gHost, gChange3).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
		// Its deps were submitted implicitly.
		So(ct.GFake.GetChange(gHost, gChange1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
		So(ct.GFake.GetChange(gHost, gChange2).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)

		ct.LogPhase(ctx, "BQ export must complete")
		ct.RunUntil(ctx, func() bool { return ct.ExportedBQAttemptsCount() == 1 })

		ct.LogPhase(ctx, "RunEnded pubsub message must be sent")
		ct.RunUntil(ctx, func() bool { return len(ct.RunEndedPubSubTasks()) == 1 })
		So(ct.RunEndedPubSubTasks()[0].Payload, ShouldResembleProto, &pubsub.PublishRunEndedTask{
			PublicId:    r.ID.PublicID(),
			LuciProject: lProject,
			Status:      run.Status_SUCCEEDED,
			Eversion:    int64(ct.LoadRun(ctx, r.ID).EVersion),
		})
	})
}

func TestCreatesMultiCLsFailPostStartMessage(t *testing.T) {
	t.Parallel()

	Convey("CV creates 2 CLs Full Run, but fails to post start message on one of the CLs", t, func() {
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review.example.com"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange1 = 11
		const gChange2 = 22

		// Start CQDaemon.
		ct.MustCQD(ctx, lProject)

		cfg := MakeCfgCombinable("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)

		tStart := ct.Clock.Now()

		// Create mutually dependent CLs using Cq-Depend git footer.
		ci1 := gf.CI(
			gChange1, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-1")),
			gf.Desc(fmt.Sprintf("This is the first CL\n\nCq-Depend: %d", gChange2)),
		)
		ci2 := gf.CI(
			gChange2, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.Updated(tStart),
			gf.CQ(+2, tStart, gf.U("user-1")),
			gf.Desc(fmt.Sprintf("This is the second CL\n\nCq-Depend: %d", gChange1)),
		)
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), ci1, ci2))
		ct.LogPhase(ctx, "Set up gChange2 with no permission to post messages")
		ct.GFake.MutateChange(gHost, gChange2, func(c *gf.Change) {
			cnt := 0
			c.ACLs = func(op gf.Operation, luciProject string) *status.Status {
				switch op {
				case gf.OpRead, gf.OpAlterVotesOfOthers:
					return status.New(codes.OK, "")
				case gf.OpReview:
					cnt++
					if cnt == 1 {
						return status.New(codes.FailedPrecondition, "unlikely fluke rejecting the start message")
					}
					return status.New(codes.OK, "")
				default:
					panic(fmt.Errorf("Unknown op: %d", op))
				}
			}
		})

		ct.LogPhase(ctx, "CV discovers all CLs")
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)
		ct.RunUntil(ctx, func() bool {
			b1 := ct.LoadGerritCL(ctx, gHost, gChange1) != nil
			b2 := ct.LoadGerritCL(ctx, gHost, gChange2) != nil
			return b1 && b2
		})

		ct.LogPhase(ctx, "CV starts a new Run")
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil
		})
		So(r.Mode, ShouldEqual, run.FullRun)
		So(r.CLs, ShouldHaveLength, 2)

		ct.LogPhase(ctx, "CV posts a message on the first CL, but fails on another CL")
		ct.RunUntil(ctx, func() bool {
			return ct.LastMessage(gHost, gChange1) != nil
		})

		ct.LogPhase(ctx, "CV fails the Run")
		ct.RunUntil(ctx, func() bool {
			r = ct.LoadRun(ctx, r.ID)
			return run.IsEnded(r.Status)
		})
		So(r.Status, ShouldEqual, run.Status_FAILED)
		// Both CLs should have the failure message, but first CL should also have
		// the starting message.
		const expecteFailMsg = "Full run: This CL failed the CQ full run.\n\nFailed to post the starting message"
		So(ct.LastMessage(gHost, gChange2).GetMessage(), ShouldContainSubstring, expecteFailMsg)
		msgs1 := ct.GFake.GetChange(gHost, gChange1).Info.GetMessages()
		So(msgs1[len(msgs1)-1].GetMessage(), ShouldContainSubstring, expecteFailMsg)
		So(msgs1[len(msgs1)-2].GetMessage(), ShouldContainSubstring, "CQ is trying the patch.")
	})
}
