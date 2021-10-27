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
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/gae/service/datastore"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/prjmanager"
	"go.chromium.org/luci/cv/internal/prjmanager/itriager"
	"go.chromium.org/luci/cv/internal/prjmanager/pmtest"
	"go.chromium.org/luci/cv/internal/prjmanager/prjpb"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/runcreator"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestEarliestDecisionTime(t *testing.T) {
	t.Parallel()

	Convey("earliestDecisionTime works", t, func() {
		now := testclock.TestRecentTimeUTC
		t0 := now.Add(time.Hour)

		earliest := func(cs []*prjpb.Component) time.Time {
			t, tPB, asap := earliestDecisionTime(cs)
			if asap {
				return now
			}
			if t.IsZero() {
				So(tPB, ShouldBeNil)
			} else {
				So(tPB.AsTime(), ShouldResemble, t)
			}
			return t
		}

		cs := []*prjpb.Component{
			{DecisionTime: nil},
		}
		So(earliest(cs), ShouldResemble, time.Time{})

		cs = append(cs, &prjpb.Component{DecisionTime: timestamppb.New(t0.Add(time.Second))})
		So(earliest(cs), ShouldResemble, t0.Add(time.Second))

		cs = append(cs, &prjpb.Component{})
		So(earliest(cs), ShouldResemble, t0.Add(time.Second))

		cs = append(cs, &prjpb.Component{DecisionTime: timestamppb.New(t0.Add(time.Hour))})
		So(earliest(cs), ShouldResemble, t0.Add(time.Second))

		cs = append(cs, &prjpb.Component{DecisionTime: timestamppb.New(t0)})
		So(earliest(cs), ShouldResemble, t0)

		cs = append(cs, &prjpb.Component{
			TriageRequired: true,
			// DecisionTime in this case doesn't matter.
			DecisionTime: timestamppb.New(t0.Add(10 * time.Hour)),
		})
		So(earliest(cs), ShouldResemble, now)
	})
}

func TestComponentsActions(t *testing.T) {
	t.Parallel()

	Convey("Component actions logic work in the abstract", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		now := ct.Clock.Now()

		const lProject = "luci-project"

		prjcfgtest.Create(ctx, lProject, &cfgpb.Config{ConfigGroups: []*cfgpb.ConfigGroup{{Name: "main"}}})
		meta := prjcfgtest.MustExist(ctx, lProject)
		pmNotifier := prjmanager.NewNotifier(ct.TQDispatcher)
		runNotifier := run.NewNotifier(ct.TQDispatcher)
		h := Handler{
			PMNotifier:  pmNotifier,
			RunNotifier: runNotifier,
			CLMutator:   changelist.NewMutator(ct.TQDispatcher, pmNotifier, runNotifier),
		}
		state := &State{
			PB: &prjpb.PState{
				LuciProject: lProject,
				Status:      prjpb.Status_STARTED,
				ConfigHash:  meta.Hash(),
				Pcls: []*prjpb.PCL{
					{Clid: 1},
					{Clid: 2},
					{Clid: 3},
					{Clid: 999},
				},
				Components: []*prjpb.Component{
					{Clids: []int64{999}}, // never sees any action.
					{Clids: []int64{1}, DecisionTime: timestamppb.New(now.Add(1 * time.Minute))},
					{Clids: []int64{2}, DecisionTime: timestamppb.New(now.Add(2 * time.Minute))},
					{Clids: []int64{3}, DecisionTime: timestamppb.New(now.Add(3 * time.Minute))},
				},
				NextEvalTime: timestamppb.New(now.Add(1 * time.Minute)),
			},
		}

		pb := backupPB(state)

		markComponentsForTriage := func(indexes ...int) {
			for _, i := range indexes {
				state.PB.GetComponents()[i].TriageRequired = true
			}
			pb = backupPB(state)
		}

		markTriaged := func(c *prjpb.Component) *prjpb.Component {
			if !c.GetTriageRequired() {
				panic(fmt.Errorf("must required triage"))
			}
			o := c.CloneShallow()
			o.TriageRequired = false
			return o
		}

		calledOn := make(chan *prjpb.Component, len(state.PB.Components))
		collectCalledOn := func() []int {
			var out []int
		loop:
			for {
				select {
				case c := <-calledOn:
					out = append(out, int(c.GetClids()[0]))
				default:
					break loop
				}
			}
			sort.Ints(out)
			return out
		}

		Convey("noop at triage", func() {
			h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				calledOn <- c
				return itriager.Result{}, nil
			}
			actions, saveForDebug, err := h.triageComponents(ctx, state)
			So(err, ShouldBeNil)
			So(saveForDebug, ShouldBeFalse)
			So(actions, ShouldBeNil)
			So(state.PB, ShouldResembleProto, pb)
			So(collectCalledOn(), ShouldBeEmpty)

			Convey("ExecDeferred", func() {
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				So(state.PB, ShouldResembleProto, pb)
				So(state2, ShouldEqual, state) // pointer comparison
				So(sideEffect, ShouldBeNil)
				// Always creates new task iff there is NextEvalTime.
				So(pmtest.ETAsOF(ct.TQ.Tasks(), lProject), ShouldNotBeEmpty)
			})
		})

		Convey("triage called on TriageRequired components or when decision time is <= now", func() {
			ct.Clock.Set(state.PB.Components[1].DecisionTime.AsTime())
			c1next := state.PB.Components[1].DecisionTime.AsTime().Add(time.Hour)
			markComponentsForTriage(3)
			h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				calledOn <- c
				switch c.GetClids()[0] {
				case 1:
					c = c.CloneShallow()
					c.DecisionTime = timestamppb.New(c1next)
					return itriager.Result{NewValue: c}, nil
				case 3:
					return itriager.Result{NewValue: markTriaged(c)}, nil
				}
				panic("unreachable")
			}
			actions, saveForDebug, err := h.triageComponents(ctx, state)
			So(err, ShouldBeNil)
			So(saveForDebug, ShouldBeFalse)
			So(actions, ShouldHaveLength, 2)
			So(collectCalledOn(), ShouldResemble, []int{1, 3})

			Convey("ExecDeferred", func() {
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				pb.NextEvalTime = timestamppb.New(now.Add(2 * time.Minute))
				pb.Components[1].DecisionTime = timestamppb.New(c1next)
				pb.Components[3].TriageRequired = false
				So(state2.PB, ShouldResembleProto, pb)
				So(pmtest.ETAsWithin(ct.TQ.Tasks(), lProject, time.Second, now.Add(2*time.Minute)), ShouldNotBeEmpty)
			})
		})

		Convey("purges CLs", func() {
			markComponentsForTriage(1, 2, 3)
			h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				switch clid := c.GetClids()[0]; clid {
				case 1, 3:
					return itriager.Result{CLsToPurge: []*prjpb.PurgeCLTask{
						{
							PurgingCl: &prjpb.PurgingCL{Clid: clid},
							Reasons: []*changelist.CLError{
								{Kind: &changelist.CLError_OwnerLacksEmail{OwnerLacksEmail: true}},
							},
						},
					}}, nil
				case 2:
					return itriager.Result{}, nil
				}
				panic("unreachable")
			}
			actions, saveForDebug, err := h.triageComponents(ctx, state)
			So(err, ShouldBeNil)
			So(saveForDebug, ShouldBeFalse)
			So(actions, ShouldHaveLength, 3)
			So(state.PB, ShouldResembleProto, pb)

			Convey("ExecDeferred", func() {
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				expectedDeadline := timestamppb.New(now.Add(maxPurgingCLDuration))
				So(state2.PB.GetPurgingCls(), ShouldResembleProto, []*prjpb.PurgingCL{
					{Clid: 1, OperationId: "1580640000-1", Deadline: expectedDeadline},
					{Clid: 3, OperationId: "1580640000-3", Deadline: expectedDeadline},
				})

				So(sideEffect, ShouldHaveSameTypeAs, &TriggerPurgeCLTasks{})
				ps := sideEffect.(*TriggerPurgeCLTasks).payloads
				So(ps, ShouldHaveLength, 2)
				// Unlike PB.PurgingCls, the tasks aren't necessarily sorted.
				sort.Slice(ps, func(i, j int) bool { return ps[i].GetPurgingCl().GetClid() < ps[j].GetPurgingCl().GetClid() })
				So(ps[0].GetPurgingCl(), ShouldResembleProto, state2.PB.GetPurgingCls()[0]) // CL#1
				So(ps[0].GetTrigger(), ShouldResembleProto, state2.PB.GetPcls()[1 /*CL#1*/].GetTrigger())
				So(ps[0].GetLuciProject(), ShouldEqual, lProject)
				So(ps[1].GetPurgingCl(), ShouldResembleProto, state2.PB.GetPurgingCls()[1]) // CL#3
			})
		})

		Convey("partial failure in triage", func() {
			markComponentsForTriage(1, 2, 3)
			h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				switch c.GetClids()[0] {
				case 1:
					return itriager.Result{}, errors.New("oops1")
				case 2, 3:
					return itriager.Result{NewValue: markTriaged(c)}, nil
				}
				panic("unreachable")
			}
			actions, saveForDebug, err := h.triageComponents(ctx, state)
			So(err, ShouldBeNil)
			So(saveForDebug, ShouldBeFalse)
			So(actions, ShouldHaveLength, 2)
			So(state.PB, ShouldResembleProto, pb)

			Convey("ExecDeferred", func() {
				// Execute slightly after #1 component decision time.
				ct.Clock.Set(pb.Components[1].DecisionTime.AsTime().Add(time.Microsecond))
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				pb.Components[2].TriageRequired = false
				pb.Components[3].TriageRequired = false
				pb.NextEvalTime = timestamppb.New(ct.Clock.Now()) // re-triage ASAP.
				So(state2.PB, ShouldResembleProto, pb)
				// Self-poke task must be scheduled for earliest possible from now.
				So(pmtest.ETAsWithin(ct.TQ.Tasks(), lProject, time.Second, ct.Clock.Now().Add(prjpb.PMTaskInterval)), ShouldNotBeEmpty)
			})
		})

		Convey("outdated PMState detected during triage", func() {
			markComponentsForTriage(1, 2, 3)
			h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				switch c.GetClids()[0] {
				case 1:
					return itriager.Result{}, errors.Annotate(itriager.ErrOutdatedPMState, "smth changed").Err()
				case 2, 3:
					return itriager.Result{NewValue: markTriaged(c)}, nil
				}
				panic("unreachable")
			}
			actions, saveForDebug, err := h.triageComponents(ctx, state)
			So(err, ShouldBeNil)
			So(saveForDebug, ShouldBeFalse)
			So(actions, ShouldHaveLength, 2)
			So(state.PB, ShouldResembleProto, pb)

			Convey("ExecDeferred", func() {
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				pb.Components[2].TriageRequired = false
				pb.Components[3].TriageRequired = false
				pb.NextEvalTime = timestamppb.New(ct.Clock.Now()) // re-triage ASAP.
				So(state2.PB, ShouldResembleProto, pb)
				// Self-poke task must be scheduled for earliest possible from now.
				So(pmtest.ETAsWithin(ct.TQ.Tasks(), lProject, time.Second, ct.Clock.Now().Add(prjpb.PMTaskInterval)), ShouldNotBeEmpty)
			})
		})

		Convey("100% failure in triage", func() {
			markComponentsForTriage(1, 2)
			h.ComponentTriage = func(_ context.Context, _ *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				return itriager.Result{}, errors.New("oops")
			}
			_, _, err := h.triageComponents(ctx, state)
			So(err, ShouldErrLike, "failed to triage 2 components")
			So(state.PB, ShouldResembleProto, pb)

			Convey("ExecDeferred", func() {
				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldNotBeNil)
				So(sideEffect, ShouldBeNil)
				So(state2, ShouldBeNil)
			})
		})

		Convey("Catches panic in triage", func() {
			markComponentsForTriage(1)
			h.ComponentTriage = func(_ context.Context, _ *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
				panic(errors.New("oops"))
			}
			_, _, err := h.ExecDeferred(ctx, state)
			So(err, ShouldErrLike, errCaughtPanic)
			So(state.PB, ShouldResembleProto, pb)
		})

		Convey("With Run Creation", func() {
			// Run creation requires ProjectStateOffload entity to exist.
			So(datastore.Put(ctx, &prjmanager.ProjectStateOffload{
				ConfigHash: prjcfgtest.MustExist(ctx, lProject).ConfigGroupIDs[0].Hash(),
				Project:    datastore.MakeKey(ctx, prjmanager.ProjectKind, lProject),
				Status:     prjpb.Status_STARTED,
			}), ShouldBeNil)

			makeRunCreator := func(clid int64, fail bool) *runcreator.Creator {
				cfgGroups, err := prjcfgtest.MustExist(ctx, lProject).GetConfigGroups(ctx)
				if err != nil {
					panic(err)
				}
				ci := gf.CI(int(clid), gf.PS(1), gf.CQ(+1, ct.Clock.Now(), gf.U("user-1")))
				cl := &changelist.CL{
					ID:       common.CLID(clid),
					EVersion: 1,
					Snapshot: &changelist.Snapshot{Kind: &changelist.Snapshot_Gerrit{Gerrit: &changelist.Gerrit{
						Host: "gerrit-review.example.com",
						Info: ci,
					}}},
				}
				if fail {
					// Simulate EVersion mismatch to fail run creation.
					cl.EVersion = 2
				}
				err = datastore.Put(ctx, cl)
				if err != nil {
					panic(err)
				}
				cl.EVersion = 1

				return &runcreator.Creator{
					LUCIProject:   lProject,
					ConfigGroupID: cfgGroups[0].ID,
					Mode:          run.DryRun,
					OperationID:   fmt.Sprintf("op-%d-%t", clid, fail),
					Owner:         identity.Identity("user:user-1@example.com"),
					Options:       &run.Options{},
					InputCLs: []runcreator.CL{{
						ID:               common.CLID(clid),
						ExpectedEVersion: 1,
						Snapshot:         cl.Snapshot,
						TriggerInfo:      trigger.Find(ci, cfgGroups[0].Content),
					}},
				}
			}

			findRunOf := func(clid int) *run.Run {
				switch runs, _, err := (run.CLQueryBuilder{CLID: common.CLID(clid)}).LoadRuns(ctx); {
				case err != nil:
					panic(err)
				case len(runs) == 0:
					return nil
				case len(runs) > 1:
					panic(fmt.Errorf("%d Runs for given CL", len(runs)))
				default:
					return runs[0]
				}
			}

			Convey("100% success", func() {
				markComponentsForTriage(1)
				h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
					rc := makeRunCreator(1, false /* succeed */)
					return itriager.Result{NewValue: markTriaged(c), RunsToCreate: []*runcreator.Creator{rc}}, nil
				}

				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				So(sideEffect, ShouldBeNil)
				pb.Components[1].TriageRequired = false // must be saved, since Run Creation succeeded.
				So(state2.PB, ShouldResembleProto, pb)
				So(findRunOf(1), ShouldNotBeNil)
			})

			Convey("100% failure", func() {
				markComponentsForTriage(1)
				h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
					rc := makeRunCreator(1, true /* fail */)
					return itriager.Result{NewValue: markTriaged(c), RunsToCreate: []*runcreator.Creator{rc}}, nil
				}

				_, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldErrLike, "failed to actOnComponents")
				So(sideEffect, ShouldBeNil)
				So(findRunOf(1), ShouldBeNil)
			})

			Convey("Partial failure", func() {
				markComponentsForTriage(1, 2, 3)
				h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
					clid := c.GetClids()[0]
					// Set up each component trying to create a Run,
					// and #2 and #3 additionally purging a CL,
					// but #2 failing to create a Run.
					failIf := clid == 2
					rc := makeRunCreator(clid, failIf)
					res := itriager.Result{NewValue: markTriaged(c), RunsToCreate: []*runcreator.Creator{rc}}
					if clid != 1 {
						// Contrived example, since in practice purging a CL concurrently
						// with Run creation in the same component ought to happen only iff
						// there are several CLs and presumably on different CLs.
						res.CLsToPurge = []*prjpb.PurgeCLTask{
							{
								PurgingCl: &prjpb.PurgingCL{Clid: clid},
								Reasons: []*changelist.CLError{
									{Kind: &changelist.CLError_OwnerLacksEmail{OwnerLacksEmail: true}},
								},
							},
						}
					}
					return res, nil
				}

				state2, sideEffect, err := h.ExecDeferred(ctx, state)
				So(err, ShouldBeNil)
				// Only #3 component purge must be a SideEffect.
				So(sideEffect, ShouldHaveSameTypeAs, &TriggerPurgeCLTasks{})
				ps := sideEffect.(*TriggerPurgeCLTasks).payloads
				So(ps, ShouldHaveLength, 1)
				So(ps[0].GetPurgingCl().GetClid(), ShouldEqual, 3)

				So(findRunOf(1), ShouldNotBeNil)
				pb.Components[1].TriageRequired = false
				// Component #2 must remain unchanged.
				So(findRunOf(3), ShouldNotBeNil)
				pb.Components[3].TriageRequired = false
				pb.PurgingCls = []*prjpb.PurgingCL{
					{
						Clid: 3, OperationId: "1580640000-3",
						Deadline: timestamppb.New(ct.Clock.Now().Add(maxPurgingCLDuration)),
					},
				}
				pb.NextEvalTime = timestamppb.New(ct.Clock.Now()) // re-triage ASAP.
				So(state2.PB, ShouldResembleProto, pb)
			})

			Convey("Catches panic", func() {
				markComponentsForTriage(1)
				h.ComponentTriage = func(_ context.Context, c *prjpb.Component, _ itriager.PMState) (itriager.Result, error) {
					rc := makeRunCreator(1, false)
					rc.LUCIProject = "" // causes panic because of incorrect usage.
					return itriager.Result{NewValue: markTriaged(c), RunsToCreate: []*runcreator.Creator{rc}}, nil
				}

				_, _, err := h.ExecDeferred(ctx, state)
				So(err, ShouldErrLike, errCaughtPanic)
				So(state.PB, ShouldResembleProto, pb)
			})
		})
	})
}
