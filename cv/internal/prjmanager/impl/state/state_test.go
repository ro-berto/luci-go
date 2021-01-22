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

package state

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/tq/tqtesting"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/config"
	"go.chromium.org/luci/cv/internal/cvtesting"
	"go.chromium.org/luci/cv/internal/gerrit/cfgmatcher"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/gobmap"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/gerrit/updater"
	"go.chromium.org/luci/cv/internal/prjmanager"
	"go.chromium.org/luci/cv/internal/prjmanager/internal"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

type ctest struct {
	cvtesting.Test

	lProject string
	gHost    string
}

func (ct ctest) runCLUpdater(ctx context.Context, change int64) *changelist.CL {
	return ct.runCLUpdaterAs(ctx, change, ct.lProject)
}

func (ct ctest) runCLUpdaterAs(ctx context.Context, change int64, lProject string) *changelist.CL {
	So(updater.Schedule(ctx, &updater.RefreshGerritCL{
		LuciProject: lProject,
		Host:        ct.gHost,
		Change:      change,
	}), ShouldBeNil)
	ct.TQ.Run(ctx, tqtesting.StopAfterTask(updater.TaskClassID))
	eid, err := changelist.GobID(ct.gHost, change)
	So(err, ShouldBeNil)
	cl, err := eid.Get(ctx)
	So(err, ShouldBeNil)
	So(cl, ShouldNotBeNil)
	return cl
}

const cfgText1 = `
  config_groups {
    name: "g0"
    gerrit {
      url: "https://c-review.example.com"  # Must match gHost.
      projects {
        name: "repo/a"
        ref_regexp: "refs/heads/main"
      }
    }
  }
  config_groups {
    name: "g1"
    fallback: YES
    gerrit {
      url: "https://c-review.example.com"  # Must match gHost.
      projects {
        name: "repo/a"
        ref_regexp: "refs/heads/.+"
      }
    }
  }
`

func updateConfigToNoFallabck(ctx context.Context, ct *ctest) config.Meta {
	cfgText2 := strings.ReplaceAll(cfgText1, "fallback: YES", "fallback: NO")
	cfg2 := &cfgpb.Config{}
	So(prototext.Unmarshal([]byte(cfgText2), cfg2), ShouldBeNil)
	ct.Cfg.Update(ctx, ct.lProject, cfg2)
	gobmap.Update(ctx, ct.lProject)
	return ct.Cfg.MustExist(ctx, ct.lProject)
}

func TestUpdateConfig(t *testing.T) {
	t.Parallel()

	Convey("updateConfig works", t, func() {
		ct := ctest{
			lProject: "test",
			gHost:    "c-review.example.com",
		}
		ctx, cancel := ct.SetUp()
		defer cancel()

		cfg1 := &cfgpb.Config{}
		So(prototext.Unmarshal([]byte(cfgText1), cfg1), ShouldBeNil)

		ct.Cfg.Create(ctx, ct.lProject, cfg1)
		meta := ct.Cfg.MustExist(ctx, ct.lProject)
		So(gobmap.Update(ctx, ct.lProject), ShouldBeNil)

		Convey("initializes newly started project", func() {
			// Newly started project doesn't have any CLs, yet, regardless of what CL
			// snapshots are stored in Datastore.
			s0 := NewInitial(ct.lProject)
			pb0 := backupPB(s0)
			s1, sideEffect, err := s0.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s0.PB, ShouldResembleProto, pb0) // s0 must not change.
			So(sideEffect, ShouldResemble, &UpdateIncompleteRunsConfig{
				Hash:     meta.Hash(),
				EVersion: meta.EVersion,
				RunIDs:   nil,
			})
			So(s1.Status, ShouldEqual, prjmanager.Status_STARTED)
			So(s1.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Components:       nil,
				Pcls:             nil,
				DirtyComponents:  true,
			})
		})

		// Add 3 CLs: 101 standalone and 202<-203 as a stack.
		ci101 := gf.CI(
			101, gf.PS(1), gf.Ref("refs/heads/main"), gf.Project("repo/a"),
			gf.CQ(+2, ct.Clock.Now(), gf.U("user-1")), gf.Updated(ct.Clock.Now()),
		)
		ci202 := gf.CI(
			202, gf.PS(3), gf.Ref("refs/heads/other"), gf.Project("repo/a"), gf.AllRevs(),
			gf.CQ(+1, ct.Clock.Now(), gf.U("user-2")), gf.Updated(ct.Clock.Now()),
		)
		ci203 := gf.CI(
			203, gf.PS(3), gf.Ref("refs/heads/other"), gf.Project("repo/a"), gf.AllRevs(),
			gf.CQ(+1, ct.Clock.Now(), gf.U("user-2")), gf.Updated(ct.Clock.Now()),
		)
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci101})
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci202})
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci203})
		ct.GFake.SetDependsOn(ct.gHost, "203_3" /* child */, "202_2" /*parent*/)
		cl101 := ct.runCLUpdater(ctx, 101)
		cl202 := ct.runCLUpdater(ctx, 202)
		cl203 := ct.runCLUpdater(ctx, 203)

		s1 := NewExisting(prjmanager.Status_STARTED, &internal.PState{
			LuciProject:      ct.lProject,
			ConfigHash:       meta.Hash(),
			ConfigGroupNames: []string{"g0", "g1"},
			Pcls: []*internal.PCL{
				{
					Clid:               int64(cl101.ID),
					Eversion:           1,
					ConfigGroupIndexes: []int32{0}, // g0
					Status:             internal.PCL_OK,
					Trigger:            trigger.Find(ci101),
				},
				{
					Clid:               int64(cl202.ID),
					Eversion:           1,
					ConfigGroupIndexes: []int32{1}, // g1
					Status:             internal.PCL_OK,
					Trigger:            trigger.Find(ci202),
				},
				{
					Clid:               int64(cl203.ID),
					Eversion:           1,
					ConfigGroupIndexes: []int32{1}, // g1
					Status:             internal.PCL_OK,
					Trigger:            trigger.Find(ci203),
					Deps:               []*changelist.Dep{{Clid: int64(cl202.ID), Kind: changelist.DepKind_HARD}},
				},
			},
			Components: []*internal.Component{
				{
					Clids: []int64{int64(cl101.ID)},
					Pruns: []*internal.PRun{
						{
							Id:    ct.lProject + "/" + "1111-v1-beef",
							Clids: []int64{int64(cl101.ID)},
						},
					},
				},
				{
					Clids: []int64{404},
				},
			},
		})
		pb1 := backupPB(s1)

		Convey("noop update is quick", func() {
			s2, sideEffect, err := s1.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s2, ShouldEqual, s1) // pointer comparison only.
			So(sideEffect, ShouldBeNil)
		})

		Convey("existing projects is updated without touching components", func() {
			meta2 := updateConfigToNoFallabck(ctx, &ct)
			s2, sideEffect, err := s1.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s1.PB, ShouldResembleProto, pb1) // s1 must not change.
			So(sideEffect, ShouldResemble, &UpdateIncompleteRunsConfig{
				Hash:     meta2.Hash(),
				EVersion: meta2.EVersion,
				RunIDs:   common.MakeRunIDs(ct.lProject + "/" + "1111-v1-beef"),
			})
			So(s2.Status, ShouldEqual, prjmanager.Status_STARTED)
			So(s2.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta2.Hash(), // changed
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					{
						Clid:               int64(cl101.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{0, 1}, // +g1, because g1 is no longer "fallback: YES"
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci101),
					},
					pb1.Pcls[1], // #202 didn't change.
					pb1.Pcls[2], // #203 didn't change.
				},
				Components:      pb1.Components, // no changes here.
				DirtyComponents: true,           // set to re-eval components
			})
		})

		Convey("disabled project updated with long ago deleted CL", func() {
			s1.Status = prjmanager.Status_STOPPED
			for _, c := range s1.PB.GetComponents() {
				c.Pruns = nil // disabled projects don't have incomplete runs.
			}
			pb1 = backupPB(s1)
			changelist.Delete(ctx, cl101.ID)

			meta2 := updateConfigToNoFallabck(ctx, &ct)
			s2, sideEffect, err := s1.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s1.PB, ShouldResembleProto, pb1) // s1 must not change.
			So(sideEffect, ShouldResemble, &UpdateIncompleteRunsConfig{
				Hash:     meta2.Hash(),
				EVersion: meta2.EVersion,
				// No runs to notify.
			})
			So(s2.Status, ShouldEqual, prjmanager.Status_STARTED)
			So(s2.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta2.Hash(), // changed
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					{
						Clid:     int64(cl101.ID),
						Eversion: 1,
						Status:   internal.PCL_DELETED,
					},
					pb1.Pcls[1], // #202 didn't change.
					pb1.Pcls[2], // #203 didn't change.
				},
				Components:      pb1.Components, // no changes here.
				DirtyComponents: true,           // set to re-eval components
			})
		})

		Convey("disabled project waits for incomplete Runs", func() {
			ct.Cfg.Disable(ctx, ct.lProject)
			s2, sideEffect, err := s1.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s2.Status, ShouldEqual, prjmanager.Status_STOPPING)
			So(s2.PB, ShouldResembleProto, s1.PB)
			So(sideEffect, ShouldResemble, &CancelIncompleteRuns{
				RunIDs: common.MakeRunIDs(ct.lProject + "/" + "1111-v1-beef"),
			})

		})

		Convey("disabled project stops iff there are no incomplete Runs", func() {
			for _, c := range s1.PB.GetComponents() {
				c.Pruns = nil
			}
			ct.Cfg.Disable(ctx, ct.lProject)
			s2, sideEffect, err := s1.UpdateConfig(ctx)
			So(err, ShouldBeNil)
			So(s2.Status, ShouldEqual, prjmanager.Status_STOPPED)
			So(s2.PB, ShouldResembleProto, s1.PB)
			So(sideEffect, ShouldBeNil)
		})

		// The rest of the test coverage of UpdateConfig is achieved by testing code
		// of makePCL.

		Convey("makePCL with full snapshot works", func() {
			var err error
			s1.cfgMatcher, err = cfgmatcher.LoadMatcherFrom(ctx, meta)
			So(err, ShouldBeNil)

			Convey("Status == OK", func() {
				expected := &internal.PCL{
					Clid:               int64(cl101.ID),
					Eversion:           int64(cl101.EVersion),
					ConfigGroupIndexes: []int32{0}, // g0
					Trigger: &run.Trigger{
						Email:           "user-1@example.com",
						GerritAccountId: 1,
						Mode:            string(run.FullRun),
						Time:            timestamppb.New(ct.Clock.Now()),
					},
				}
				Convey("CL snapshotted with current config", func() {
					So(s1.makePCL(ctx, cl101), ShouldResembleProto, expected)
				})
				Convey("CL snapshotted with an older config", func() {
					cl101.ApplicableConfig.GetProjects()[0].ConfigGroupIds = []string{"oldhash/g0"}
					So(s1.makePCL(ctx, cl101), ShouldResembleProto, expected)
				})
				Convey("not triggered CL", func() {
					delete(cl101.Snapshot.GetGerrit().GetInfo().GetLabels(), trigger.CQLabelName)
					expected.Trigger = nil
					So(s1.makePCL(ctx, cl101), ShouldResembleProto, expected)
				})
			})

			Convey("snapshot from diff project requires waiting", func() {
				cl101.Snapshot.LuciProject = "another"
				So(s1.makePCL(ctx, cl101), ShouldResembleProto, &internal.PCL{
					Clid:     int64(cl101.ID),
					Eversion: int64(cl101.EVersion),
					Status:   internal.PCL_UNKNOWN,
				})
			})

			Convey("CL from diff project is unwatched", func() {
				s1.PB.LuciProject = "another"
				So(s1.makePCL(ctx, cl101), ShouldResembleProto, &internal.PCL{
					Clid:     int64(cl101.ID),
					Eversion: int64(cl101.EVersion),
					Status:   internal.PCL_UNWATCHED,
				})
			})

			Convey("CL watched by several projects is unwatched", func() {
				cl101.ApplicableConfig.Projects = append(
					cl101.ApplicableConfig.GetProjects(),
					&changelist.ApplicableConfig_Project{
						ConfigGroupIds: []string{"g"},
						Name:           "another",
					})
				So(s1.makePCL(ctx, cl101), ShouldResembleProto, &internal.PCL{
					Clid:     int64(cl101.ID),
					Eversion: int64(cl101.EVersion),
					Status:   internal.PCL_UNWATCHED,
				})
			})
		})
	})
}

func TestOnCLsUpdated(t *testing.T) {
	t.Parallel()

	Convey("OnCLsUpdated works", t, func() {
		ct := ctest{
			lProject: "test",
			gHost:    "c-review.example.com",
		}
		ctx, cancel := ct.SetUp()
		defer cancel()

		cfg1 := &cfgpb.Config{}
		So(prototext.Unmarshal([]byte(cfgText1), cfg1), ShouldBeNil)

		ct.Cfg.Create(ctx, ct.lProject, cfg1)
		meta := ct.Cfg.MustExist(ctx, ct.lProject)
		So(gobmap.Update(ctx, ct.lProject), ShouldBeNil)

		// Add 3 CLs: 101 standalone and 202<-203 as a stack.
		ci101 := gf.CI(
			101, gf.PS(1), gf.Ref("refs/heads/main"), gf.Project("repo/a"),
			gf.CQ(+2, ct.Clock.Now(), gf.U("user-1")), gf.Updated(ct.Clock.Now()),
		)
		ci202 := gf.CI(
			202, gf.PS(3), gf.Ref("refs/heads/other"), gf.Project("repo/a"), gf.AllRevs(),
			gf.CQ(+1, ct.Clock.Now(), gf.U("user-2")), gf.Updated(ct.Clock.Now()),
		)
		ci203 := gf.CI(
			203, gf.PS(3), gf.Ref("refs/heads/other"), gf.Project("repo/a"), gf.AllRevs(),
			gf.CQ(+1, ct.Clock.Now(), gf.U("user-2")), gf.Updated(ct.Clock.Now()),
		)
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci101})
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci202})
		ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: ci203})
		ct.GFake.SetDependsOn(ct.gHost, "203_3" /* child */, "202_2" /*parent*/)
		cl101 := ct.runCLUpdater(ctx, 101)
		cl202 := ct.runCLUpdater(ctx, 202)
		cl203 := ct.runCLUpdater(ctx, 203)

		s0 := NewExisting(prjmanager.Status_STARTED, &internal.PState{
			LuciProject:      ct.lProject,
			ConfigHash:       meta.Hash(),
			ConfigGroupNames: []string{"g0", "g1"},
		})
		pb0 := backupPB(s0)

		// NOTE: conversion of individual CL to PCL is in TestUpdateConfig.

		Convey("One simple CL", func() {
			s1, sideEffect, err := s0.OnCLsUpdated(ctx, []*internal.CLUpdated{
				{Clid: int64(cl101.ID), Eversion: int64(cl101.EVersion)},
			})
			So(err, ShouldBeNil)
			So(s0.PB, ShouldResembleProto, pb0)
			So(sideEffect, ShouldBeNil)
			So(s1.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					{
						Clid:               int64(cl101.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{0}, // g0
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci101),
					},
				},
				DirtyComponents: true,
			})
			Convey("Noop based on EVersion", func() {
				s2, sideEffect, err := s1.OnCLsUpdated(ctx, []*internal.CLUpdated{
					{Clid: int64(cl101.ID), Eversion: 1}, // already known
				})
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				So(s1, ShouldEqual, s2) // pointer comparison only.
			})

			Convey("Removes duplicates", func() {
				pb := backupPB(s1)
				bumpEVersion(ctx, cl101, 10)
				s2, sideEffect, err := s1.OnCLsUpdated(ctx, []*internal.CLUpdated{
					{Clid: int64(cl101.ID), Eversion: 5},
					{Clid: int64(cl101.ID), Eversion: 7},
				})
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				So(s1.PB, ShouldResembleProto, pb)
				pb.GetPcls()[0].Eversion = int64(cl101.EVersion)
				So(s2.PB, ShouldResembleProto, pb)
			})
		})

		Convey("One CL with a yet unknown dep", func() {
			s1, sideEffect, err := s0.OnCLsUpdated(ctx, []*internal.CLUpdated{
				{Clid: int64(cl203.ID), Eversion: 1},
			})
			So(err, ShouldBeNil)
			So(s0.PB, ShouldResembleProto, pb0)
			So(sideEffect, ShouldBeNil)
			So(s1.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					{
						Clid:               int64(cl203.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{1}, // g1
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci203),
						Deps:               []*changelist.Dep{{Clid: int64(cl202.ID), Kind: changelist.DepKind_HARD}},
					},
				},
				DirtyComponents: true,
			})
		})

		Convey("PCLs must remain sorted", func() {
			s1 := NewExisting(prjmanager.Status_STARTED, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					{
						Clid:               int64(cl101.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{0}, // g0
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci101),
					},
					{
						Clid:               int64(cl203.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{1}, // g1
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci203),
						Deps:               []*changelist.Dep{{Clid: int64(cl202.ID), Kind: changelist.DepKind_HARD}},
					},
				},
			})
			pb1 := backupPB(s1)
			bumpEVersion(ctx, cl203, 3)
			s2, sideEffect, err := s1.OnCLsUpdated(ctx, []*internal.CLUpdated{
				{Clid: 404, Eversion: 404},                               // doesn't even exist
				{Clid: int64(cl202.ID), Eversion: int64(cl202.EVersion)}, // new
				{Clid: int64(cl101.ID), Eversion: int64(cl101.EVersion)}, // unchanged
				{Clid: int64(cl203.ID), Eversion: 3},                     // updated
			})
			So(err, ShouldBeNil)
			So(s1.PB, ShouldResembleProto, pb1)
			So(sideEffect, ShouldBeNil)
			So(s2.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Pcls: []*internal.PCL{
					s1.PB.GetPcls()[0], // 101 is unchanged
					{ // new & inserted at the right spot
						Clid:               int64(cl202.ID),
						Eversion:           1,
						ConfigGroupIndexes: []int32{1}, // g1
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci202),
					},
					{ // updated
						Clid:               int64(cl203.ID),
						Eversion:           3,
						ConfigGroupIndexes: []int32{1}, // g1
						Status:             internal.PCL_OK,
						Trigger:            trigger.Find(ci203),
						Deps:               []*changelist.Dep{{Clid: int64(cl202.ID), Kind: changelist.DepKind_HARD}},
					},
					{
						Clid:     404,
						Eversion: 0,
						Status:   internal.PCL_DELETED,
					},
				},
				DirtyComponents: true,
			})
		})

		Convey("non-STARTED project ignores all CL events", func() {
			s0.Status = prjmanager.Status_STOPPING
			s1, sideEffect, err := s0.OnCLsUpdated(ctx, []*internal.CLUpdated{
				{Clid: int64(cl101.ID), Eversion: int64(cl101.EVersion)},
			})
			So(err, ShouldBeNil)
			So(sideEffect, ShouldBeNil)
			So(s0, ShouldEqual, s1) // pointer comparison only.
		})
	})
}

func TestRunsCreatedAndFinished(t *testing.T) {
	t.Parallel()

	Convey("OnRunsCreated and OnRunsFinished works", t, func() {
		ct := ctest{
			lProject: "test",
			gHost:    "c-review.example.com",
		}
		ctx, cancel := ct.SetUp()
		defer cancel()

		cfg1 := &cfgpb.Config{}
		So(prototext.Unmarshal([]byte(cfgText1), cfg1), ShouldBeNil)
		ct.Cfg.Create(ctx, ct.lProject, cfg1)
		meta := ct.Cfg.MustExist(ctx, ct.lProject)

		run1 := &run.Run{ID: common.RunID(ct.lProject + "/101-aaa"), CLs: []common.CLID{101}}
		run789 := &run.Run{ID: common.RunID(ct.lProject + "/789-efg"), CLs: []common.CLID{709, 707, 708}}
		So(datastore.Put(ctx, run1, run789), ShouldBeNil)

		s1 := NewExisting(prjmanager.Status_STARTED, &internal.PState{
			LuciProject:      ct.lProject,
			ConfigHash:       meta.Hash(),
			ConfigGroupNames: []string{"g0", "g1"},
			// For OnRunsFinished / OnRunsCreated PCLs don't matter, so omit them from
			// the test for brevity, even though valid State must have PCLs covering
			// all components.
			Pcls: nil,
			Components: []*internal.Component{
				{
					Clids: []int64{101},
					Pruns: []*internal.PRun{{Id: ct.lProject + "/101-aaa", Clids: []int64{1}}},
				},
				{
					Clids: []int64{202, 203, 204},
				},
			},
			CreatedPruns: []*internal.PRun{
				{Id: ct.lProject + "/789-efg", Clids: []int64{707, 708, 709}},
			},
		})
		pb1 := backupPB(s1)

		Convey("Noops", func() {
			Convey("OnRunsFinished on not tracked Run", func() {
				s2, sideEffect, err := s1.OnRunsFinished(ctx, common.MakeRunIDs(ct.lProject+"/999-zzz"))
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				// although s2 is cloned, it must be exact same as s1.
				So(s2.Status, ShouldEqual, s1.Status)
				So(s2.PB, ShouldResembleProto, pb1)
			})
			Convey("OnRunsCreated on already tracked Run", func() {
				s2, sideEffect, err := s1.OnRunsCreated(ctx, common.MakeRunIDs(ct.lProject+"/101-aaa"))
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				So(s2, ShouldEqual, s1)
				So(pb1, ShouldResembleProto, s1.PB)
			})
			Convey("OnRunsCreated on somehow already deleted run", func() {
				s2, sideEffect, err := s1.OnRunsCreated(ctx, common.MakeRunIDs(ct.lProject+"/404-nnn"))
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				// although s2 is cloned, it must be exact same as s1.
				So(s2.Status, ShouldEqual, s1.Status)
				So(s2.PB, ShouldResembleProto, pb1)
			})
		})

		Convey("OnRunsCreated", func() {
			runX := &run.Run{ // Run involving all of CLs and more.
				ID: common.RunID(ct.lProject + "/000-xxx"),
				// The order doesn't have to and is intentionally not sorted here.
				CLs: []common.CLID{404, 101, 202, 204, 203},
			}
			run2 := &run.Run{ID: common.RunID(ct.lProject + "/202-bbb"), CLs: []common.CLID{202}}
			run3 := &run.Run{ID: common.RunID(ct.lProject + "/203-ccc"), CLs: []common.CLID{203}}
			run23 := &run.Run{ID: common.RunID(ct.lProject + "/232-bcb"), CLs: []common.CLID{203, 202}}
			run234 := &run.Run{ID: common.RunID(ct.lProject + "/234-bcd"), CLs: []common.CLID{203, 204, 202}}
			So(datastore.Put(ctx, run2, run3, run23, run234, runX), ShouldBeNil)

			s2, sideEffect, err := s1.OnRunsCreated(ctx, common.RunIDs{
				run2.ID, run3.ID, run23.ID, run234.ID, runX.ID,
				// non-existing Run shouldn't derail others.
				common.RunID(ct.lProject + "/404-nnn"),
			})
			So(err, ShouldBeNil)
			So(pb1, ShouldResembleProto, s1.PB)
			So(sideEffect, ShouldBeNil)
			So(s2.PB, ShouldResembleProto, &internal.PState{
				LuciProject:      ct.lProject,
				ConfigHash:       meta.Hash(),
				ConfigGroupNames: []string{"g0", "g1"},
				Components: []*internal.Component{
					s1.PB.GetComponents()[0], // 101 is unchanged
					{
						Clids: []int64{202, 203, 204},
						Pruns: []*internal.PRun{
							// Runs & CLs must be sorted by their respective IDs.
							{Id: string(run2.ID), Clids: []int64{202}},
							{Id: string(run3.ID), Clids: []int64{203}},
							{Id: string(run23.ID), Clids: []int64{202, 203}},
							{Id: string(run234.ID), Clids: []int64{202, 203, 204}},
						},
						Dirty: true,
					},
				},
				CreatedPruns: []*internal.PRun{
					{Id: string(runX.ID), Clids: []int64{101, 202, 203, 204, 404}},
					{Id: ct.lProject + "/789-efg", Clids: []int64{707, 708, 709}}, // unchanged
				},
			})
		})

		Convey("OnRunsFinished", func() {
			s1.Status = prjmanager.Status_STOPPING
			pb1 := backupPB(s1)

			Convey("deletes from Components", func() {
				pb1 := backupPB(s1)
				s2, sideEffect, err := s1.OnRunsFinished(ctx, common.MakeRunIDs(ct.lProject+"/101-aaa"))
				So(err, ShouldBeNil)
				So(pb1, ShouldResembleProto, s1.PB)
				So(sideEffect, ShouldBeNil)
				So(s2.Status, ShouldEqual, prjmanager.Status_STOPPING)
				So(s2.PB, ShouldResembleProto, &internal.PState{
					LuciProject:      ct.lProject,
					ConfigHash:       meta.Hash(),
					ConfigGroupNames: []string{"g0", "g1"},
					Components: []*internal.Component{
						{
							Clids: []int64{101},
							Pruns: nil,  // removed
							Dirty: true, // marked dirty
						},
						s1.PB.GetComponents()[1], // unchanged
					},
					CreatedPruns: s1.PB.GetCreatedPruns(), // unchanged
				})
			})

			Convey("deletes from CreatedPruns", func() {
				s2, sideEffect, err := s1.OnRunsFinished(ctx, common.MakeRunIDs(ct.lProject+"/789-efg"))
				So(err, ShouldBeNil)
				So(pb1, ShouldResembleProto, s1.PB)
				So(sideEffect, ShouldBeNil)
				So(s2.Status, ShouldEqual, prjmanager.Status_STOPPING)
				So(s2.PB, ShouldResembleProto, &internal.PState{
					LuciProject:      ct.lProject,
					ConfigHash:       meta.Hash(),
					ConfigGroupNames: []string{"g0", "g1"},
					Components:       s1.PB.Components, // unchanged
					CreatedPruns:     nil,              // removed
				})
			})

			Convey("stops PM iff all runs finished", func() {
				s2, sideEffect, err := s1.OnRunsFinished(ctx, common.MakeRunIDs(
					ct.lProject+"/101-aaa",
					ct.lProject+"/789-efg",
				))
				So(err, ShouldBeNil)
				So(pb1, ShouldResembleProto, s1.PB)
				So(sideEffect, ShouldBeNil)
				So(s2.Status, ShouldEqual, prjmanager.Status_STOPPED)
				So(s2.PB, ShouldResembleProto, &internal.PState{
					LuciProject:      ct.lProject,
					ConfigHash:       meta.Hash(),
					ConfigGroupNames: []string{"g0", "g1"},
					Pcls:             s1.PB.GetPcls(),
					Components: []*internal.Component{
						{Clids: []int64{101}, Dirty: true},
						s1.PB.GetComponents()[1], // unchanged.
					},
					CreatedPruns: nil, // removed
				})
			})
		})

	})
}

func TestLoadActiveIntoPCLs(t *testing.T) {
	t.Parallel()

	Convey("loadActiveIntoPCLs works", t, func() {
		ct := ctest{
			lProject: "test",
			gHost:    "c-review.example.com",
		}
		ctx, cancel := ct.SetUp()
		defer cancel()

		cfg := &cfgpb.Config{}
		So(prototext.Unmarshal([]byte(cfgText1), cfg), ShouldBeNil)
		ct.Cfg.Create(ctx, ct.lProject, cfg)
		meta := ct.Cfg.MustExist(ctx, ct.lProject)
		gobmap.Update(ctx, ct.lProject)

		// Simulate existence of "test-b" project watching the same Gerrit host but
		// diff repo.
		const lProjectB = "test-b"
		cfgTextB := strings.ReplaceAll(cfgText1, "repo/a", "repo/b")
		cfgB := &cfgpb.Config{}
		So(prototext.Unmarshal([]byte(cfgTextB), cfgB), ShouldBeNil)
		ct.Cfg.Create(ctx, lProjectB, cfgB)
		gobmap.Update(ctx, lProjectB)

		cis := make(map[int]*gerritpb.ChangeInfo, 20)
		makeCI := func(i int, project string, cq int, extra ...gf.CIModifier) {
			mods := []gf.CIModifier{
				gf.Ref("refs/heads/main"),
				gf.Project(project),
				gf.Updated(ct.Clock.Now()),
			}
			if cq > 0 {
				mods = append(mods, gf.CQ(cq, ct.Clock.Now(), gf.U("user-1")))
			}
			mods = append(mods, extra...)
			cis[i] = gf.CI(i, mods...)
			ct.GFake.CreateChange(&gf.Change{Host: ct.gHost, ACLs: gf.ACLPublic(), Info: cis[i]})
		}
		makeStack := func(ids []int, project string, cq int) {
			for i, child := range ids {
				makeCI(child, project, cq)
				for _, parent := range ids[:i] {
					ct.GFake.SetDependsOn(ct.gHost, cis[child], cis[parent])
				}
			}
		}
		// Simulate the following CLs state in Gerrit:
		//   In this project:
		//     CQ+1
		//       1 <- 2       form a stack (2 depends on 1)
		//       3            depends on 2 via Cq-Depend.
		//     CQ+2
		//       4            standalone
		//       5 <- 6       form a stack (6 depends on 5)
		//       7 <- 8 <- 9  form a stack (9 depends on 7,8)
		//       13           CQ-Depend on 11 (diff project) and 12 (not existing).
		//   In another project:
		//     CQ+1
		//       10 <- 11     form a stack (11 depends on 10)
		makeStack([]int{1, 2}, "repo/a", +1)
		makeCI(3, "repo/a", +1, gf.Desc("T\n\nCq-Depend: 2"))
		makeStack([]int{7, 8, 9}, "repo/a", +2)
		makeStack([]int{5, 6}, "repo/a", +2)
		makeCI(4, "repo/a", +2)
		makeCI(13, "repo/a", +2, gf.Desc("T\n\nCq-Depend: 11,12"))
		makeStack([]int{10, 11}, "repo/b", +1)

		// Import into DS all CLs in their respective LUCI projects.
		// Do this in-order such that they match auto-assigned CLIDs by fake
		// Datastore as this helps test readability. Note that importing CL 13 would
		// create CL entity for dep #12 before creating CL 13th own entity.
		cls := make(map[int]*changelist.CL, 20)
		for i := 1; i < 14; i++ {
			if i == 12 {
				continue // skipped. will be done after 13
			}
			pr := ct.lProject
			if i == 10 || i == 11 {
				pr = lProjectB
			}
			cls[i] = ct.runCLUpdaterAs(ctx, int64(i), pr)
			// If Datastore fake changes, this check and comment above can be removed,
			// but the rest of the test will remain valid.
			So(cls[i].ID, ShouldEqual, i)
		}
		// This will get 404 from Gerrit.
		cls[12] = ct.runCLUpdater(ctx, 12)

		run4 := &run.Run{
			ID:  common.RunID(ct.lProject + "/1-a"),
			CLs: []common.CLID{cls[4].ID},
		}
		run56 := &run.Run{
			ID:  common.RunID(ct.lProject + "/56-bb"),
			CLs: []common.CLID{cls[5].ID, cls[6].ID},
		}
		run789 := &run.Run{
			ID:  common.RunID(ct.lProject + "/789-ccc"),
			CLs: []common.CLID{cls[9].ID, cls[7].ID, cls[8].ID},
		}
		So(datastore.Put(ctx, run4, run56, run789), ShouldBeNil)

		state := NewExisting(prjmanager.Status_STARTED, &internal.PState{
			LuciProject:      ct.lProject,
			ConfigHash:       meta.Hash(),
			ConfigGroupNames: []string{"g0", "g1"},
			DirtyComponents:  true,
		})

		Convey("just categorization", func() {
			state.PB.Pcls = sortPCLs([]*internal.PCL{
				defaultPCL(cls[5]),
				defaultPCL(cls[6]),
				defaultPCL(cls[7]),
				defaultPCL(cls[8]),
				defaultPCL(cls[9]),
				{Clid: 12, Eversion: 1, Status: internal.PCL_UNKNOWN},
			})
			state.PB.Components = []*internal.Component{
				{
					Clids: i64sorted(cls[5].ID, cls[6].ID),
					Pruns: []*internal.PRun{internal.MakePRun(run56)},
				},
				// Simulate 9 previously not depending on 7, 8.
				{Clids: i64sorted(cls[7].ID, cls[8].ID)},
				{Clids: i64s(cls[9].ID)},
			}
			// 789 doesn't match any 1 component, even though 7,8,9 CLs are in PCLs.
			state.PB.CreatedPruns = []*internal.PRun{internal.MakePRun(run789)}
			pbBefore := backupPB(state)

			cat := state.categorizeCLs()
			So(state.loadActiveIntoPCLs(ctx, cat), ShouldBeNil)
			So(cat, ShouldResemble, &categorizedCLs{
				active:   mkClidsSet(cls[5], cls[6], cls[7], cls[8], cls[9]),
				deps:     clidsSet{},
				unused:   mkClidsSet(cls[12]),
				unloaded: clidsSet{},
			})
			So(state.PB, ShouldResembleProto, pbBefore)
		})

		Convey("loads unloaded dependencies and active CLs without recursion", func() {
			state.PB.Pcls = []*internal.PCL{
				defaultPCL(cls[3]), // depends on 2, which in turns depends on 1.
			}
			state.PB.CreatedPruns = []*internal.PRun{internal.MakePRun(run56)}
			pb := backupPB(state)

			cat := state.categorizeCLs()
			So(cat, ShouldResemble, &categorizedCLs{
				active:   mkClidsSet(cls[3], cls[5], cls[6]),
				deps:     mkClidsSet(cls[2]),
				unused:   clidsSet{},
				unloaded: mkClidsSet(cls[2], cls[5], cls[6]),
			})
			So(state.loadActiveIntoPCLs(ctx, cat), ShouldBeNil)
			So(cat, ShouldResemble, &categorizedCLs{
				active:   mkClidsSet(cls[3], cls[2], cls[5], cls[6]),
				deps:     mkClidsSet(cls[1]),
				unused:   clidsSet{},
				unloaded: mkClidsSet(cls[1]),
			})
			pb.Pcls = sortPCLs([]*internal.PCL{
				defaultPCL(cls[2]),
				defaultPCL(cls[3]),
				defaultPCL(cls[5]),
				defaultPCL(cls[6]),
			})
			So(state.PB, ShouldResembleProto, pb)
		})

		Convey("loads incomplete Run with unloaded deps", func() {
			// This case shouldn't normally happen in practice. This case simulates a
			// runStale created a while ago of just (11, 13), presumably when current
			// project had CL #11 in scope.
			// Now, 11 and 13 depend on 10 and 12, respectively, and 10 and 11 are no
			// longer watched by current project.
			runStale := &run.Run{
				ID:  common.RunID(ct.lProject + "/111-s"),
				CLs: []common.CLID{cls[13].ID, cls[11].ID},
			}
			So(datastore.Put(ctx, runStale), ShouldBeNil)
			state.PB.CreatedPruns = []*internal.PRun{internal.MakePRun(runStale)}
			pb := backupPB(state)

			cat := state.categorizeCLs()
			So(cat, ShouldResemble, &categorizedCLs{
				active:   mkClidsSet(cls[11], cls[13]),
				deps:     clidsSet{},
				unused:   clidsSet{},
				unloaded: mkClidsSet(cls[11], cls[13]),
			})
			So(state.loadActiveIntoPCLs(ctx, cat), ShouldBeNil)
			So(cat, ShouldResemble, &categorizedCLs{
				active: mkClidsSet(cls[11], cls[13]),
				// 10 isn't in deps because this project has no visibility into CL 11.
				deps:     mkClidsSet(cls[12]),
				unused:   clidsSet{},
				unloaded: mkClidsSet(cls[12]),
			})
			pb.Pcls = sortPCLs([]*internal.PCL{
				defaultPCL(cls[13]),
				{
					Clid:     int64(cls[11].ID),
					Eversion: int64(cls[11].EVersion),
					Status:   internal.PCL_UNWATCHED,
					Deps:     nil, // not visible to this project
					Trigger:  nil, // not visible to this project
				},
			})
			So(state.PB, ShouldResembleProto, pb)
		})

		Convey("loads incomplete Run with non-existent CLs", func() {
			// This case shouldn't happen in practice, but it can't be ruled out.
			// In order to incorporate just added .CreatedRun into State,
			// Run's CLs must have PCL entries.
			runStale := &run.Run{
				ID:  common.RunID(ct.lProject + "/404-s"),
				CLs: []common.CLID{cls[4].ID, 404},
			}
			So(datastore.Put(ctx, runStale), ShouldBeNil)
			state.PB.CreatedPruns = []*internal.PRun{internal.MakePRun(runStale)}
			pb := backupPB(state)

			cat := state.categorizeCLs()
			So(cat, ShouldResemble, &categorizedCLs{
				active:   clidsSet{cls[4].ID: struct{}{}, 404: struct{}{}},
				deps:     clidsSet{},
				unused:   clidsSet{},
				unloaded: clidsSet{cls[4].ID: struct{}{}, 404: struct{}{}},
			})
			So(state.loadActiveIntoPCLs(ctx, cat), ShouldBeNil)
			So(cat, ShouldResemble, &categorizedCLs{
				active:   clidsSet{cls[4].ID: struct{}{}, 404: struct{}{}},
				deps:     clidsSet{},
				unused:   clidsSet{},
				unloaded: clidsSet{},
			})
			pb.Pcls = sortPCLs([]*internal.PCL{
				defaultPCL(cls[4]),
				{
					Clid:     404,
					Eversion: 0,
					Status:   internal.PCL_DELETED,
				},
			})
			So(state.PB, ShouldResembleProto, pb)
		})

		Convey("noop", func() {
			cat := state.categorizeCLs()
			So(state.loadActiveIntoPCLs(ctx, cat), ShouldBeNil)
			So(cat, ShouldResemble, &categorizedCLs{
				active:   clidsSet{},
				deps:     clidsSet{},
				unused:   clidsSet{},
				unloaded: clidsSet{},
			})
		})
	})
}

func TestRepartition(t *testing.T) {
	t.Parallel()

	Convey("repartition works", t, func() {
		state := NewExisting(prjmanager.Status_STARTED, &internal.PState{
			DirtyComponents: true,
		})
		cat := &categorizedCLs{
			active:   clidsSet{},
			deps:     clidsSet{},
			unused:   clidsSet{},
			unloaded: clidsSet{},
		}

		defer func() {
			// Assert guarantees of repartition()
			So(state.PB.GetDirtyComponents(), ShouldBeFalse)
			So(state.PB.GetCreatedPruns(), ShouldBeNil)
			actual := state.pclIndex
			state.pclIndex = nil
			state.ensurePCLIndex()
			So(actual, ShouldResemble, state.pclIndex)
		}()

		Convey("nothing to do, except resetting DirtyComponents", func() {
			Convey("totally empty", func() {
				state.repartition(cat)
				So(state.PB, ShouldResembleProto, &internal.PState{})
			})
			Convey("1 active CL in 1 component", func() {
				cat.active.resetI64(1)
				state.PB.Components = []*internal.Component{{Clids: []int64{1}}}
				state.PB.Pcls = []*internal.PCL{{Clid: 1}}
				pb := backupPB(state)

				state.repartition(cat)
				pb.DirtyComponents = false
				So(state.PB, ShouldResembleProto, pb)
			})
			Convey("1 active CL in 1 dirty component with 1 Run", func() {
				cat.active.resetI64(1)
				state.PB.Components = []*internal.Component{{
					Clids: []int64{1},
					Pruns: []*internal.PRun{{Clids: []int64{1}, Id: "id"}},
					Dirty: true,
				}}
				state.PB.Pcls = []*internal.PCL{{Clid: 1}}
				pb := backupPB(state)

				state.repartition(cat)
				pb.DirtyComponents = false
				So(state.PB, ShouldResembleProto, pb)
			})
		})

		Convey("Compacts out unused PCLs", func() {
			cat.active.resetI64(1, 3)
			cat.unused.resetI64(2)
			state.PB.Pcls = []*internal.PCL{
				{Clid: 1},
				{Clid: 2},
				{Clid: 3, Deps: []*changelist.Dep{{Clid: 1}}},
			}

			state.repartition(cat)
			So(state.PB, ShouldResembleProto, &internal.PState{
				Pcls: []*internal.PCL{
					{Clid: 1},
					{Clid: 3, Deps: []*changelist.Dep{{Clid: 1}}},
				},
				Components: []*internal.Component{{
					Clids: []int64{1, 3},
					Dirty: true,
				}},
			})
		})

		Convey("Creates new components", func() {
			Convey("1 active CL converted into 1 new dirty component", func() {
				cat.active.resetI64(1)
				state.PB.Pcls = []*internal.PCL{{Clid: 1}}

				state.repartition(cat)
				So(state.PB, ShouldResembleProto, &internal.PState{
					Pcls: []*internal.PCL{{Clid: 1}},
					Components: []*internal.Component{{
						Clids: []int64{1},
						Dirty: true,
					}},
				})
			})
			Convey("Deps respected during conversion", func() {
				cat.active.resetI64(1, 2, 3)
				state.PB.Pcls = []*internal.PCL{
					{Clid: 1},
					{Clid: 2},
					{Clid: 3, Deps: []*changelist.Dep{{Clid: 1}}},
				}
				orig := backupPB(state)

				state.repartition(cat)
				sortByFirstCL(state.PB.Components)
				So(state.PB, ShouldResembleProto, &internal.PState{
					Pcls: orig.Pcls,
					Components: []*internal.Component{
						{
							Clids: []int64{1, 3},
							Dirty: true,
						},
						{
							Clids: []int64{2},
							Dirty: true,
						},
					},
				})
			})
		})

		Convey("Components splitting works", func() {
			Convey("Crossing-over 12, 34 => 13, 24", func() {
				cat.active.resetI64(1, 2, 3, 4)
				state.PB.Pcls = []*internal.PCL{
					{Clid: 1},
					{Clid: 2},
					{Clid: 3, Deps: []*changelist.Dep{{Clid: 1}}},
					{Clid: 4, Deps: []*changelist.Dep{{Clid: 2}}},
				}
				state.PB.Components = []*internal.Component{
					{Clids: []int64{1, 2}},
					{Clids: []int64{3, 4}},
				}
				orig := backupPB(state)

				state.repartition(cat)
				sortByFirstCL(state.PB.Components)
				So(state.PB, ShouldResembleProto, &internal.PState{
					Pcls: orig.Pcls,
					Components: []*internal.Component{
						{Clids: []int64{1, 3}, Dirty: true},
						{Clids: []int64{2, 4}, Dirty: true},
					},
				})
			})
			Convey("Loaded and unloaded deps can be shared by several components", func() {
				cat.active.resetI64(1, 2, 3)
				cat.deps.resetI64(4, 5)
				cat.unloaded.resetI64(5)
				state.PB.Pcls = []*internal.PCL{
					{Clid: 1, Deps: []*changelist.Dep{{Clid: 3}, {Clid: 4}, {Clid: 5}}},
					{Clid: 2, Deps: []*changelist.Dep{{Clid: 4}, {Clid: 5}}},
					{Clid: 3},
					{Clid: 4},
				}
				orig := backupPB(state)

				state.repartition(cat)
				sortByFirstCL(state.PB.Components)
				So(state.PB, ShouldResembleProto, &internal.PState{
					Pcls: orig.Pcls,
					Components: []*internal.Component{
						{Clids: []int64{1, 3}, Dirty: true},
						{Clids: []int64{2}, Dirty: true},
					},
				})
			})
		})

		Convey("CreatedRuns are moved into components", func() {
			Convey("Simple", func() {
				cat.active.resetI64(1, 2)
				state.PB.Pcls = []*internal.PCL{
					{Clid: 1},
					{Clid: 2, Deps: []*changelist.Dep{{Clid: 1}}},
				}
				state.PB.CreatedPruns = []*internal.PRun{{Clids: []int64{1, 2}, Id: "id"}}
				orig := backupPB(state)

				state.repartition(cat)
				So(state.PB, ShouldResembleProto, &internal.PState{
					CreatedPruns: nil,
					Pcls:         orig.Pcls,
					Components: []*internal.Component{
						{
							Clids: []int64{1, 2},
							Pruns: []*internal.PRun{{Clids: []int64{1, 2}, Id: "id"}},
							Dirty: true,
						},
					},
				})
			})
			Convey("Force-merge 2 existing components", func() {
				cat.active.resetI64(1, 2)
				state.PB.Pcls = []*internal.PCL{
					{Clid: 1},
					{Clid: 2},
				}
				state.PB.Components = []*internal.Component{
					{Clids: []int64{1}, Pruns: []*internal.PRun{{Clids: []int64{1}, Id: "1"}}},
					{Clids: []int64{2}, Pruns: []*internal.PRun{{Clids: []int64{2}, Id: "2"}}},
				}
				state.PB.CreatedPruns = []*internal.PRun{{Clids: []int64{1, 2}, Id: "12"}}
				orig := backupPB(state)

				state.repartition(cat)
				sortByFirstCL(state.PB.Components)
				So(state.PB, ShouldResembleProto, &internal.PState{
					CreatedPruns: nil,
					Pcls:         orig.Pcls,
					Components: []*internal.Component{
						{
							Clids: []int64{1, 2},
							Pruns: []*internal.PRun{ // must be sorted by ID
								{Clids: []int64{1}, Id: "1"},
								{Clids: []int64{1, 2}, Id: "12"},
								{Clids: []int64{2}, Id: "2"},
							},
							Dirty: true,
						},
					},
				})
			})
		})

		Convey("Does all at once", func() {
			// This test adds more test coverage for a busy project where components
			// are created, split, merged, and CreatedRuns are incorporated during
			// repartition(), especially likely after a config update.
			cat.active.resetI64(1, 2, 4, 5, 6)
			cat.deps.resetI64(7)
			cat.unused.resetI64(3)
			cat.unloaded.resetI64(7)
			state.PB.Pcls = []*internal.PCL{
				{Clid: 1},
				{Clid: 2, Deps: []*changelist.Dep{{Clid: 1}}},
				{Clid: 3, Deps: []*changelist.Dep{{Clid: 1}, {Clid: 2}}}, // but unused
				{Clid: 4},
				{Clid: 5, Deps: []*changelist.Dep{{Clid: 4}}},
				{Clid: 6, Deps: []*changelist.Dep{{Clid: 7}}},
			}
			state.PB.Components = []*internal.Component{
				{Clids: []int64{1, 2, 3}, Pruns: []*internal.PRun{{Clids: []int64{1}, Id: "1"}}},
				{Clids: []int64{4}, Pruns: []*internal.PRun{{Clids: []int64{4}, Id: "4"}}},
				{Clids: []int64{5}, Pruns: []*internal.PRun{{Clids: []int64{5}, Id: "5"}}},
			}
			state.PB.CreatedPruns = []*internal.PRun{
				{Clids: []int64{4, 5}, Id: "45"}, // so, merge component with {4}, {5}.
				{Clids: []int64{6}, Id: "6"},
			}

			state.repartition(cat)
			sortByFirstCL(state.PB.Components)
			So(state.PB, ShouldResembleProto, &internal.PState{
				Pcls: []*internal.PCL{
					{Clid: 1},
					{Clid: 2, Deps: []*changelist.Dep{{Clid: 1}}},
					// 3 was deleted
					{Clid: 4},
					{Clid: 5, Deps: []*changelist.Dep{{Clid: 4}}},
					{Clid: 6, Deps: []*changelist.Dep{{Clid: 7}}},
				},
				Components: []*internal.Component{
					{Clids: []int64{1, 2}, Dirty: true, Pruns: []*internal.PRun{{Clids: []int64{1}, Id: "1"}}},
					{Clids: []int64{4, 5}, Dirty: true, Pruns: []*internal.PRun{
						{Clids: []int64{4}, Id: "4"},
						{Clids: []int64{4, 5}, Id: "45"},
						{Clids: []int64{5}, Id: "5"},
					}},
					{Clids: []int64{6}, Dirty: true, Pruns: []*internal.PRun{{Clids: []int64{6}, Id: "6"}}},
				},
			})
		})
	})
}

// backupPB returns a deep copy of State.PB for future assertion that State
// wasn't modified.
func backupPB(s *State) *internal.PState {
	ret := &internal.PState{}
	proto.Merge(ret, s.PB)
	return ret
}

func bumpEVersion(ctx context.Context, cl *changelist.CL, desired int) {
	if cl.EVersion >= desired {
		panic(fmt.Errorf("can't go %d to %d", cl.EVersion, desired))
	}
	cl.EVersion = desired
	So(datastore.Put(ctx, cl), ShouldBeNil)
}

func defaultPCL(cl *changelist.CL) *internal.PCL {
	p := &internal.PCL{
		Clid:               int64(cl.ID),
		Eversion:           int64(cl.EVersion),
		ConfigGroupIndexes: []int32{0},
		Status:             internal.PCL_OK,
		Deps:               cl.Snapshot.GetDeps(),
	}
	ci := cl.Snapshot.GetGerrit().GetInfo()
	if ci != nil {
		p.Trigger = trigger.Find(ci)
	}
	return p
}

func customPCL(cl *changelist.CL, override *internal.PCL) *internal.PCL {
	p := defaultPCL(cl)
	proto.Merge(p, override)
	return p
}

func i64s(vs ...interface{}) []int64 {
	res := make([]int64, len(vs))
	for i, v := range vs {
		switch x := v.(type) {
		case int64:
			res[i] = x
		case common.CLID:
			res[i] = int64(x)
		case int:
			res[i] = int64(x)
		default:
			panic(fmt.Errorf("unknown type: %T %v", v, v))
		}
	}
	return res
}

func i64sorted(vs ...interface{}) []int64 {
	res := i64s(vs...)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func sortPCLs(vs []*internal.PCL) []*internal.PCL {
	sort.Slice(vs, func(i, j int) bool { return vs[i].GetClid() < vs[j].GetClid() })
	return vs
}

func mkClidsSet(cls ...*changelist.CL) clidsSet {
	res := make(clidsSet, len(cls))
	for _, cl := range cls {
		res[cl.ID] = struct{}{}
	}
	return res
}

func sortByFirstCL(cs []*internal.Component) []*internal.Component {
	sort.Slice(cs, func(i, j int) bool { return cs[i].GetClids()[0] < cs[j].GetClids()[0] })
	return cs
}
