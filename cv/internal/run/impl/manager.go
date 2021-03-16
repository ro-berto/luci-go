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
	"time"

	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/eventbox"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/handler"
	"go.chromium.org/luci/cv/internal/run/impl/state"
)

func init() {
	eventpb.ManageRunTaskRef.AttachHandler(
		func(ctx context.Context, payload proto.Message) error {
			task := payload.(*eventpb.ManageRunTask)
			err := manageRun(ctx, common.RunID(task.GetRunId()))
			// TODO(tandrii/yiwzhang): avoid retries iff we know a new task was
			// already scheduled for the next second.
			return common.TQifyError(ctx, err)
		},
	)
}

var pokeInterval = 5 * time.Minute

var fakeHandlerKey = "Fake Run Events Handler"

func manageRun(ctx context.Context, runID common.RunID) error {
	ctx = logging.SetField(ctx, "run", runID)
	recipient := datastore.MakeKey(ctx, run.RunKind, string(runID))
	rm := &runManager{runID: runID}
	if h, ok := ctx.Value(&fakeHandlerKey).(handler.Handler); ok {
		rm.handler = h
	} else {
		rm.handler = &handler.Impl{}
	}
	return eventbox.ProcessBatch(ctx, recipient, rm)
}

// runManager implements eventbox.Processor.
type runManager struct {
	runID   common.RunID
	handler handler.Handler
}

var _ eventbox.Processor = (*runManager)(nil)

// LoadState is called to load the state before a transaction.
func (rm *runManager) LoadState(ctx context.Context) (eventbox.State, eventbox.EVersion, error) {
	r := run.Run{ID: rm.runID}
	switch err := datastore.Get(ctx, &r); {
	case err == datastore.ErrNoSuchEntity:
		err = errors.Reason("CRITICAL: requested run entity %q is missing in datastore.", rm.runID).Err()
		common.LogError(ctx, err)
		panic(err)
	case err != nil:
		return nil, 0, errors.Annotate(err, "failed to get Run %q", rm.runID).Tag(transient.Tag).Err()
	}
	rs := &state.RunState{
		Run: r,
	}
	return rs, eventbox.EVersion(r.EVersion), nil
}

// Mutate is called before a transaction to compute transitions based on a
// batch of events.
//
// All actions that must be done atomically with updating state must be
// encapsulated inside Transition.SideEffectFn callback.
func (rm *runManager) Mutate(ctx context.Context, events eventbox.Events, s eventbox.State) ([]eventbox.Transition, eventbox.Events, error) {
	tr := &triageResult{}
	for _, e := range events {
		tr.triage(ctx, e)
	}
	ts, err := rm.processTriageResults(ctx, tr, s.(*state.RunState))
	return ts, nil, err
}

// FetchEVersion is called at the beginning of a transaction.
//
// The returned EVersion is compared against the one associated with a state
// loaded via GetState. If different, the transaction is aborted and new state
// isn't saved.
func (rm *runManager) FetchEVersion(ctx context.Context) (eventbox.EVersion, error) {
	r := &run.Run{ID: rm.runID}
	if err := datastore.Get(ctx, r); err != nil {
		return 0, errors.Annotate(err, "failed to get %q", rm.runID).Tag(transient.Tag).Err()
	}
	return eventbox.EVersion(r.EVersion), nil

}

// SaveState is called in a transaction to save the state if it has changed.
//
// The passed eversion is incremented value of eversion of what GetState
// returned before.
func (rm *runManager) SaveState(ctx context.Context, st eventbox.State, ev eventbox.EVersion) error {
	rs := st.(*state.RunState)
	rs.Run.EVersion = int(ev)
	rs.Run.UpdateTime = clock.Now(ctx).UTC()
	if err := datastore.Put(ctx, &(rs.Run)); err != nil {
		return errors.Annotate(err, "failed to put Run %q", rs.Run.ID).Tag(transient.Tag).Err()
	}
	return nil
}

// triageResult is the result of the triage of the incoming events.
type triageResult struct {
	startEvents     eventbox.Events
	cancelEvents    eventbox.Events
	pokeEvents      eventbox.Events
	newConfigEvents eventbox.Events
	clUpdatedEvents struct {
		events eventbox.Events
		cls    common.CLIDs
	}
	readyForSubmissionEvents eventbox.Events
	clSubmittedEvents        struct {
		events eventbox.Events
		cls    common.CLIDs
	}
	submissionCompletedEvents struct {
		events  eventbox.Events
		result  eventpb.SubmissionResult
		attempt int32
	}
	cqdVerificationCompletedEvents eventbox.Events
	nextReadyEventTime             time.Time
}

func (tr *triageResult) triage(ctx context.Context, item eventbox.Event) {
	e := &eventpb.Event{}
	if err := proto.Unmarshal(item.Value, e); err != nil {
		// This is a bug in code or data corruption.
		// There is no way to recover on its own.
		logging.Errorf(ctx, "CRITICAL: failed to deserialize event %q: %s", item.ID, err)
		panic(err)
	}
	if pa := e.GetProcessAfter().AsTime(); pa.After(clock.Now(ctx)) {
		if tr.nextReadyEventTime.IsZero() || pa.Before(tr.nextReadyEventTime) {
			tr.nextReadyEventTime = pa
		}
		return
	}
	switch e.GetEvent().(type) {
	case *eventpb.Event_Start:
		tr.startEvents = append(tr.startEvents, item)
	case *eventpb.Event_Cancel:
		tr.cancelEvents = append(tr.cancelEvents, item)
	case *eventpb.Event_Poke:
		tr.pokeEvents = append(tr.pokeEvents, item)
	case *eventpb.Event_NewConfig:
		tr.newConfigEvents = append(tr.newConfigEvents, item)
	case *eventpb.Event_ClUpdated:
		tr.clUpdatedEvents.events = append(tr.clUpdatedEvents.events, item)
		tr.clUpdatedEvents.cls = append(tr.clUpdatedEvents.cls, common.CLID(e.GetClUpdated().GetClid()))
	case *eventpb.Event_ReadyForSubmission:
		tr.readyForSubmissionEvents = append(tr.readyForSubmissionEvents, item)
	case *eventpb.Event_ClSubmitted:
		tr.clSubmittedEvents.events = append(tr.clSubmittedEvents.events, item)
		tr.clSubmittedEvents.cls = append(tr.clSubmittedEvents.cls, common.CLID(e.GetClSubmitted().GetClid()))
	case *eventpb.Event_SubmissionCompleted:
		tr.submissionCompletedEvents.events = append(tr.submissionCompletedEvents.events, item)
		switch sc := e.GetSubmissionCompleted(); {
		case sc.GetAttempt() > tr.submissionCompletedEvents.attempt:
			tr.submissionCompletedEvents.attempt = sc.GetAttempt()
			tr.submissionCompletedEvents.result = sc.GetResult()
		case sc.GetAttempt() == tr.submissionCompletedEvents.attempt:
			panic(fmt.Errorf("received more than 1 SubmissionCompleted result for same submission attempt %d ", sc.GetAttempt()))
		}
	case *eventpb.Event_CqdVerificationCompleted:
		tr.cqdVerificationCompletedEvents = append(tr.cqdVerificationCompletedEvents, item)
	default:
		panic(fmt.Errorf("unknown event: %T [id=%q]", e.GetEvent(), item.ID))
	}
}

func (rm *runManager) processTriageResults(ctx context.Context, tr *triageResult, rs *state.RunState) ([]eventbox.Transition, error) {
	var transitions []eventbox.Transition
	if len(tr.clSubmittedEvents.events) > 0 {
		res, err := rm.handler.OnCLSubmitted(ctx, rs, tr.clSubmittedEvents.cls)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.clSubmittedEvents.events, transitions)
	}
	if sc := tr.submissionCompletedEvents; len(sc.events) > 0 {
		res, err := rm.handler.OnSubmissionCompleted(ctx, rs, sc.result, sc.attempt)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, sc.events, transitions)
	}
	if len(tr.readyForSubmissionEvents) > 0 {
		res, err := rm.handler.OnReadyForSubmission(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.readyForSubmissionEvents, transitions)
	}
	if len(tr.cqdVerificationCompletedEvents) > 0 {
		res, err := rm.handler.OnCQDVerificationCompleted(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.cqdVerificationCompletedEvents, transitions)
	}
	switch {
	case len(tr.cancelEvents) > 0:
		res, err := rm.handler.Cancel(ctx, rs)
		if err != nil {
			return nil, err
		}
		// Consume all the start events here as well because it is possible
		// that Run Manager receives start and cancel events at the same time.
		// For example, user requests to start a Run and immediately cancels
		// it. But the duration is long enough for Project Manager to create
		// this Run in CV. In that case, Run Manager should just move this Run
		// to cancelled state directly.
		events := append(tr.cancelEvents, tr.startEvents...)
		rs, transitions = applyResult(res, events, transitions)
	case len(tr.startEvents) > 0:
		res, err := rm.handler.Start(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.startEvents, transitions)
	}
	if len(tr.clUpdatedEvents.events) > 0 {
		res, err := rm.handler.OnCLUpdated(ctx, rs, tr.clUpdatedEvents.cls)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.clUpdatedEvents.events, transitions)
	}
	if len(tr.newConfigEvents) > 0 {
		// TODO(tandrii,yiwzhang): update config.
		transitions = append(transitions, eventbox.Transition{Events: tr.newConfigEvents, TransitionTo: rs})
	}
	if len(tr.pokeEvents) > 0 {
		// TODO(tandrii,yiwzhang): implement poke.
		// TODO(crbug/1178658): trigger CL updater to refetch Run's CLs.
		transitions = append(transitions, eventbox.Transition{Events: tr.pokeEvents, TransitionTo: rs})
	}

	if err := enqueueNextPoke(ctx, rm.runID, tr.nextReadyEventTime); err != nil {
		return nil, err
	}
	return transitions, nil
}

func applyResult(res *handler.Result, events eventbox.Events, transitions []eventbox.Transition) (*state.RunState, []eventbox.Transition) {
	t := eventbox.Transition{
		TransitionTo: res.State,
		SideEffectFn: res.SideEffectFn,
	}
	if !res.PreserveEvents {
		t.Events = events
	}
	return res.State, append(transitions, t)
}

func enqueueNextPoke(ctx context.Context, runID common.RunID, nextReadyEventTime time.Time) error {
	switch now := clock.Now(ctx); {
	case nextReadyEventTime.IsZero():
		return run.Poke(ctx, runID, pokeInterval)
	case now.After(nextReadyEventTime):
		// It is possible that by this time, next ready event is already overdue.
		// Invoke Run Manager immediately.
		return eventpb.Dispatch(ctx, string(runID), time.Time{})
	case nextReadyEventTime.Before(now.Add(pokeInterval)):
		return eventpb.Dispatch(ctx, string(runID), nextReadyEventTime)
	default:
		return run.Poke(ctx, runID, pokeInterval)
	}
}
