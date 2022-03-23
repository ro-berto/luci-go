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
	"testing"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGerritCLDeleted(t *testing.T) {
	t.Parallel()

	Convey("CV cancels a Run with some grace period after Gerrit CL is deleted", t, func() {
		ct := Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "infra"
		const gHost = "g-review.example.com"
		const gRepo = "re/po"
		const gRef = "refs/heads/main"
		const gChange = 404

		cfg := MakeCfgSingular("cg0", gHost, gRepo, gRef)
		prjcfgtest.Create(ctx, lProject, cfg)
		So(ct.PMNotifier.UpdateConfig(ctx, lProject), ShouldBeNil)

		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), gf.CI(
			gChange, gf.Project(gRepo), gf.Ref(gRef),
			gf.Owner("user-1"),
			gf.CQ(+1, ct.Clock.Now(), gf.U("user-2")),
			gf.Updated(ct.Clock.Now()),
		)))
		// Only a committer can trigger a DryRun for someone else' CL.
		ct.AddCommitter("user-2")
		ct.LogPhase(ctx, "CV starts a Run")
		var r *run.Run
		ct.RunUntil(ctx, func() bool {
			r = ct.EarliestCreatedRunOf(ctx, lProject)
			return r != nil && r.Status == run.Status_RUNNING
		})

		ct.LogPhase(ctx, "CL disappears")
		backup := ct.GFake.GetChange(gHost, gChange)
		ct.GFake.DeleteChange(gHost, gChange)
		// This will have to wait for the next full poll.
		ct.RunUntil(ctx, func() bool {
			return ct.LoadCL(ctx, r.CLs[0]).AccessKind(ctx, lProject) == changelist.AccessDeniedProbably
		})

		ct.LogPhase(ctx, "CL re-appears")
		ct.GFake.CreateChange(backup)
		// To avoid races in this test, "touch" the CL s.t. CV re-discovers it
		// immediately even in incremental (not full) poll.
		ct.GFake.MutateChange(gHost, gChange, func(c *gf.Change) { gf.Updated(ct.Clock.Now())(c.Info) })
		ct.RunUntil(ctx, func() bool {
			return ct.LoadCL(ctx, r.CLs[0]).AccessKind(ctx, lProject) == changelist.AccessGranted
		})

		r = ct.LoadRun(ctx, r.ID)
		So(r.Status, ShouldEqual, run.Status_RUNNING)
	})
}
