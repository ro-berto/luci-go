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

package impl

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/tq/tqtesting"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/common/eventbox"
	"go.chromium.org/luci/cv/internal/cvtesting"
	"go.chromium.org/luci/cv/internal/gerrit/updater"
	"go.chromium.org/luci/cv/internal/prjmanager"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	submitpb "go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/handler"
	"go.chromium.org/luci/cv/internal/run/impl/state"
	"go.chromium.org/luci/cv/internal/run/runtest"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRunManager(t *testing.T) {
	t.Parallel()

	Convey("RunManager", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		const runID = "chromium/222-1-deadbeef"
		const initialEVersion = 10
		So(datastore.Put(ctx, &run.Run{
			ID:       runID,
			Status:   run.Status_RUNNING,
			EVersion: initialEVersion,
		}), ShouldBeNil)

		currentRun := func(ctx context.Context) *run.Run {
			ret := &run.Run{ID: runID}
			So(datastore.Get(ctx, ret), ShouldBeNil)
			return ret
		}

		notifier := run.NewNotifier(ct.TQDispatcher)
		pm := prjmanager.NewNotifier(ct.TQDispatcher)
		clMutator := changelist.NewMutator(ct.TQDispatcher, pm, notifier)
		u := updater.New(ct.TQDispatcher, ct.GFactory(), clMutator)
		_ = New(notifier, pm, clMutator, u, ct.GFactory(), ct.TreeFake.Client(), ct.BQFake)

		// sorted by the order of execution.
		eventTestcases := []struct {
			event                *eventpb.Event
			sendFn               func(context.Context) error
			invokedHandlerMethod string
		}{
			{
				&eventpb.Event{
					Event: &eventpb.Event_Cancel{
						Cancel: &eventpb.Cancel{},
					},
				},
				func(ctx context.Context) error {
					return notifier.Cancel(ctx, runID)
				},
				"Cancel",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_Start{
						Start: &eventpb.Start{},
					},
				},
				func(ctx context.Context) error {
					return notifier.Start(ctx, runID)
				},
				"Start",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_NewConfig{
						NewConfig: &eventpb.NewConfig{
							Hash:     "deadbeef",
							Eversion: 2,
						},
					},
				},
				func(ctx context.Context) error {
					return notifier.UpdateConfig(ctx, runID, "deadbeef", 2)
				},
				"UpdateConfig",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_ClsUpdated{
						ClsUpdated: &changelist.CLUpdatedEvents{
							Events: []*changelist.CLUpdatedEvent{
								{
									Clid:     int64(1),
									Eversion: int64(2),
								},
							},
						},
					},
				},
				func(ctx context.Context) error {
					return notifier.NotifyCLsUpdated(ctx, runID, &changelist.CLUpdatedEvents{
						Events: []*changelist.CLUpdatedEvent{
							{
								Clid:     int64(1),
								Eversion: int64(2),
							},
						},
					})
				},
				"OnCLsUpdated",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_CqdTryjobsUpdated{
						CqdTryjobsUpdated: &eventpb.CQDTryjobsUpdated{},
					},
				},
				func(ctx context.Context) error {
					return notifier.NotifyCQDTryjobsUpdated(ctx, runID)
				},
				"OnCQDTryjobsUpdated",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_CqdVerificationCompleted{
						CqdVerificationCompleted: &eventpb.CQDVerificationCompleted{},
					},
				},
				func(ctx context.Context) error {
					return notifier.NotifyCQDVerificationCompleted(ctx, runID)
				},
				"OnCQDVerificationCompleted",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_ClSubmitted{
						ClSubmitted: &eventpb.CLSubmitted{
							Clid: 1,
						},
					},
				},
				func(ctx context.Context) error {
					return notifier.SendNow(ctx, runID, &eventpb.Event{
						Event: &eventpb.Event_ClSubmitted{
							ClSubmitted: &eventpb.CLSubmitted{
								Clid: 1,
							},
						},
					})
				},
				"OnCLSubmitted",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_SubmissionCompleted{
						SubmissionCompleted: &submitpb.SubmissionCompleted{
							Result: submitpb.SubmissionResult_SUCCEEDED,
						},
					},
				},
				func(ctx context.Context) error {
					return notifier.SendNow(ctx, runID, &eventpb.Event{
						Event: &eventpb.Event_SubmissionCompleted{
							SubmissionCompleted: &submitpb.SubmissionCompleted{
								Result: submitpb.SubmissionResult_SUCCEEDED,
							},
						},
					})
				},
				"OnSubmissionCompleted",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_ReadyForSubmission{
						ReadyForSubmission: &eventpb.ReadyForSubmission{},
					},
				},
				func(ctx context.Context) error {
					return notifier.SendNow(ctx, runID, &eventpb.Event{
						Event: &eventpb.Event_ReadyForSubmission{
							ReadyForSubmission: &eventpb.ReadyForSubmission{},
						},
					})
				},
				"OnReadyForSubmission",
			},
			{
				&eventpb.Event{
					Event: &eventpb.Event_Poke{
						Poke: &eventpb.Poke{},
					},
				},
				func(ctx context.Context) error {
					return notifier.PokeNow(ctx, runID)
				},
				"Poke",
			},
		}
		for _, et := range eventTestcases {
			Convey(fmt.Sprintf("Can process Event %T", et.event.GetEvent()), func() {
				fh := &fakeHandler{}
				ctx = context.WithValue(ctx, &fakeHandlerKey, fh)
				So(et.sendFn(ctx), ShouldBeNil)
				runtest.AssertInEventbox(ctx, runID, et.event)
				So(runtest.Runs(ct.TQ.Tasks()), ShouldResemble, common.RunIDs{runID})
				ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
				So(fh.invocations[0], ShouldEqual, et.invokedHandlerMethod)
				So(currentRun(ctx).EVersion, ShouldEqual, initialEVersion+1)
				runtest.AssertNotInEventbox(ctx, runID, et.event) // consumed
			})
		}

		Convey("Process Events in order", func() {
			fh := &fakeHandler{}
			ctx = context.WithValue(ctx, &fakeHandlerKey, fh)

			var expectInvokedMethods []string
			for _, et := range eventTestcases {
				// skipping Cancel because when Start and Cancel are both present.
				// only Cancel will execute. See next test
				if et.event.GetCancel() == nil {
					expectInvokedMethods = append(expectInvokedMethods, et.invokedHandlerMethod)
				}
			}
			rand.Shuffle(len(eventTestcases), func(i, j int) {
				eventTestcases[i], eventTestcases[j] = eventTestcases[j], eventTestcases[i]
			})
			for _, etc := range eventTestcases {
				if etc.event.GetCancel() == nil {
					So(etc.sendFn(ctx), ShouldBeNil)
				}
			}
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			expectInvokedMethods = append(expectInvokedMethods, "TryResumeSubmission") // always invoked
			So(fh.invocations, ShouldResemble, expectInvokedMethods)
			So(currentRun(ctx).EVersion, ShouldEqual, initialEVersion+1)
		})

		Convey("Don't Start if received both Cancel and Start Event", func() {
			fh := &fakeHandler{}
			ctx = context.WithValue(ctx, &fakeHandlerKey, fh)
			notifier.Start(ctx, runID)
			notifier.Cancel(ctx, runID)
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			So(fh.invocations[0], ShouldEqual, "Cancel")
			for _, inv := range fh.invocations[1:] {
				So(inv, ShouldNotEqual, "Start")
			}
			So(currentRun(ctx).EVersion, ShouldEqual, initialEVersion+1)
			runtest.AssertNotInEventbox(ctx, runID, &eventpb.Event{
				Event: &eventpb.Event_Cancel{
					Cancel: &eventpb.Cancel{},
				},
			},
				&eventpb.Event{
					Event: &eventpb.Event_Start{
						Start: &eventpb.Start{},
					},
				},
			)
		})

		Convey("Can Preserve events", func() {
			fh := &fakeHandler{preserveEvents: true}
			ctx = context.WithValue(ctx, &fakeHandlerKey, fh)
			So(notifier.Start(ctx, runID), ShouldBeNil)
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			So(currentRun(ctx).EVersion, ShouldEqual, initialEVersion+1)
			runtest.AssertInEventbox(ctx, runID,
				&eventpb.Event{
					Event: &eventpb.Event_Start{
						Start: &eventpb.Start{},
					},
				},
			)
		})

		Convey("Can save RunLog", func() {
			fh := &fakeHandler{startAddsLogEntries: []*run.LogEntry{
				{
					Time: timestamppb.New(clock.Now(ctx)),
					Kind: &run.LogEntry_Created_{Created: &run.LogEntry_Created{
						ConfigGroupId: "deadbeef/main",
					}},
				},
			}}
			ctx = context.WithValue(ctx, &fakeHandlerKey, fh)
			So(notifier.Start(ctx, runID), ShouldBeNil)
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			So(currentRun(ctx).EVersion, ShouldEqual, initialEVersion+1)
			entries, err := run.LoadRunLogEntries(ctx, runID)
			So(err, ShouldBeNil)
			So(entries, ShouldResembleProto, fh.startAddsLogEntries)
		})

		Convey("Can run PostProcessFn", func() {
			var postProcessFnExecuted bool
			fh := &fakeHandler{
				postProcessFn: func(c context.Context) error {
					postProcessFnExecuted = true
					return nil
				},
			}
			ctx = context.WithValue(ctx, &fakeHandlerKey, fh)
			So(notifier.Start(ctx, runID), ShouldBeNil)
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			So(postProcessFnExecuted, ShouldBeTrue)
		})
	})

	Convey("Poke", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		const runID = "chromium/222-1-deadbeef"
		tCreate := ct.Clock.Now().UTC().Add(-2 * time.Minute)
		So(datastore.Put(ctx, &run.Run{
			ID:         runID,
			Status:     run.Status_RUNNING,
			CreateTime: tCreate,
			StartTime:  tCreate.Add(1 * time.Minute),
			EVersion:   10,
		}), ShouldBeNil)

		notifier := run.NewNotifier(ct.TQDispatcher)
		pm := prjmanager.NewNotifier(ct.TQDispatcher)
		clMutator := changelist.NewMutator(ct.TQDispatcher, pm, notifier)
		u := updater.New(ct.TQDispatcher, ct.GFactory(), clMutator)
		_ = New(notifier, pm, clMutator, u, ct.GFactory(), ct.TreeFake.Client(), ct.BQFake)

		Convey("Recursive", func() {
			So(notifier.PokeNow(ctx, runID), ShouldBeNil)
			So(runtest.Runs(ct.TQ.Tasks()), ShouldResemble, common.RunIDs{runID})
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			for i := 0; i < 10; i++ {
				now := clock.Now(ctx)
				runtest.AssertInEventbox(ctx, runID, &eventpb.Event{
					Event: &eventpb.Event_Poke{
						Poke: &eventpb.Poke{},
					},
					ProcessAfter: timestamppb.New(now.Add(pokeInterval)),
				})
				ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
			}

			Convey("Stops after Run is finalized", func() {
				So(datastore.Put(ctx, &run.Run{
					ID:         runID,
					Status:     run.Status_CANCELLED,
					CreateTime: tCreate,
					StartTime:  tCreate.Add(1 * time.Minute),
					EndTime:    ct.Clock.Now().UTC(),
					EVersion:   11,
				}), ShouldBeNil)
				ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))
				runtest.AssertEventboxEmpty(ctx, runID)
			})
		})

		Convey("Existing event due during the interval", func() {
			So(notifier.PokeNow(ctx, runID), ShouldBeNil)
			So(notifier.PokeAfter(ctx, runID, 30*time.Second), ShouldBeNil)
			ct.TQ.Run(ctx, tqtesting.StopAfterTask(eventpb.ManageRunTaskClass))

			runtest.AssertNotInEventbox(ctx, runID, &eventpb.Event{
				Event: &eventpb.Event_Poke{
					Poke: &eventpb.Poke{},
				},
				ProcessAfter: timestamppb.New(clock.Now(ctx).Add(pokeInterval)),
			})
			So(runtest.Tasks(ct.TQ.Tasks()), ShouldHaveLength, 1)
			task := runtest.Tasks(ct.TQ.Tasks())[0]
			So(task.ETA, ShouldResemble, clock.Now(ctx).UTC().Add(30*time.Second))
			So(task.Payload, ShouldResembleProto, &eventpb.ManageRunTask{RunId: string(runID)})
		})
	})
}

type fakeHandler struct {
	invocations         []string
	preserveEvents      bool
	postProcessFn       eventbox.PostProcessFn
	startAddsLogEntries []*run.LogEntry
}

var _ handler.Handler = &fakeHandler{}

func (fh *fakeHandler) Start(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("Start")
	rs = rs.ShallowCopy()
	if len(fh.startAddsLogEntries) > 0 {
		rs.LogEntries = append(rs.LogEntries, fh.startAddsLogEntries...)
	}
	return &handler.Result{
		State:          rs,
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) Cancel(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("Cancel")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) OnCLsUpdated(ctx context.Context, rs *state.RunState, _ common.CLIDs) (*handler.Result, error) {
	fh.addInvocation("OnCLsUpdated")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) OnReadyForSubmission(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("OnReadyForSubmission")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

// OnCLSubmitted records provided CLs have been submitted.
func (fh *fakeHandler) OnCLSubmitted(ctx context.Context, rs *state.RunState, clids common.CLIDs) (*handler.Result, error) {
	fh.addInvocation("OnCLSubmitted")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) OnSubmissionCompleted(ctx context.Context, rs *state.RunState, sc *submitpb.SubmissionCompleted) (*handler.Result, error) {
	fh.addInvocation("OnSubmissionCompleted")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) TryResumeSubmission(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("TryResumeSubmission")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) OnCQDTryjobsUpdated(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("OnCQDTryjobsUpdated")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) OnCQDVerificationCompleted(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("OnCQDVerificationCompleted")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) Poke(ctx context.Context, rs *state.RunState) (*handler.Result, error) {
	fh.addInvocation("Poke")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) UpdateConfig(ctx context.Context, rs *state.RunState, hash string) (*handler.Result, error) {
	fh.addInvocation("UpdateConfig")
	return &handler.Result{
		State:          rs.ShallowCopy(),
		PreserveEvents: fh.preserveEvents,
		PostProcessFn:  fh.postProcessFn,
	}, nil
}

func (fh *fakeHandler) addInvocation(method string) {
	fh.invocations = append(fh.invocations, method)
}
