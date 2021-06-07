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

	"go.chromium.org/luci/cv/internal/bq"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/eventbox"
	"go.chromium.org/luci/cv/internal/gerrit/updater"
	"go.chromium.org/luci/cv/internal/prjmanager"
	"go.chromium.org/luci/cv/internal/run"
	runbq "go.chromium.org/luci/cv/internal/run/bq"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/impl/handler"
	"go.chromium.org/luci/cv/internal/run/impl/state"
	"go.chromium.org/luci/cv/internal/tree"
)

// RunManager manages Runs.
type RunManager struct {
	runNotifier *run.Notifier
	pmNotifier  *prjmanager.Notifier
	clUpdater   *updater.Updater
	handler     handler.Handler
}

func New(n *run.Notifier, pm *prjmanager.Notifier, u *updater.Updater, tc tree.Client, bqc bq.Client) *RunManager {
	rm := &RunManager{n, pm, u, &handler.Impl{
		PM:         pm,
		RM:         n,
		TreeClient: tc,
		BQExporter: runbq.NewExporter(n.TaskRefs.Tqd, bqc),
		CLUpdater:  u,
	}}
	n.TaskRefs.ManageRun.AttachHandler(
		func(ctx context.Context, payload proto.Message) error {
			task := payload.(*eventpb.ManageRunTask)
			err := rm.manageRun(ctx, common.RunID(task.GetRunId()))
			// TODO(tandrii/yiwzhang): avoid retries iff we know a new task was
			// already scheduled for the next second.
			return common.TQIfy{
				KnownRetry: []error{
					handler.ErrTransientSubmissionFailure,
					eventbox.ErrConcurretMutation,
				},
			}.Error(ctx, err)
		},
	)
	return rm
}

var pokeInterval = 5 * time.Minute

var fakeHandlerKey = "Fake Run Events Handler"

func (rm *RunManager) manageRun(ctx context.Context, runID common.RunID) error {
	ctx = logging.SetField(ctx, "run", runID)
	recipient := datastore.MakeKey(ctx, run.RunKind, string(runID))
	proc := &runProcessor{
		runID:       runID,
		runNotifier: rm.runNotifier,
		pmNotifier:  rm.pmNotifier,
		clUpdater:   rm.clUpdater,
		handler:     rm.handler,
	}
	if h, ok := ctx.Value(&fakeHandlerKey).(handler.Handler); ok {
		proc.handler = h
	}
	postProcessFns, err := eventbox.ProcessBatch(ctx, recipient, proc)
	if err != nil {
		return errors.Annotate(err, "run: %q", runID).Err()
	}
	for _, postProcessFn := range postProcessFns {
		if err := postProcessFn(ctx); err != nil {
			return errors.Annotate(err, "run: %q", runID).Err()
		}
	}
	return nil
}

// runProcessor implements eventbox.Processor.
type runProcessor struct {
	runID common.RunID

	runNotifier *run.Notifier
	pmNotifier  *prjmanager.Notifier
	clUpdater   *updater.Updater

	handler handler.Handler
}

var _ eventbox.Processor = (*runProcessor)(nil)

// LoadState is called to load the state before a transaction.
func (rp *runProcessor) LoadState(ctx context.Context) (eventbox.State, eventbox.EVersion, error) {
	r := run.Run{ID: rp.runID}
	switch err := datastore.Get(ctx, &r); {
	case err == datastore.ErrNoSuchEntity:
		err = errors.Reason("CRITICAL: requested run entity %q is missing in datastore.", rp.runID).Err()
		common.LogError(ctx, err)
		panic(err)
	case err != nil:
		return nil, 0, errors.Annotate(err, "failed to get Run %q", rp.runID).Tag(transient.Tag).Err()
	}
	rs := &state.RunState{Run: r}
	return rs, eventbox.EVersion(r.EVersion), nil
}

// Mutate is called before a transaction to compute transitions based on a
// batch of events.
//
// All actions that must be done atomically with updating state must be
// encapsulated inside Transition.SideEffectFn callback.
func (rp *runProcessor) Mutate(ctx context.Context, events eventbox.Events, s eventbox.State) ([]eventbox.Transition, eventbox.Events, error) {
	tr := &triageResult{}
	for _, e := range events {
		tr.triage(ctx, e)
	}
	ts, err := rp.processTriageResults(ctx, tr, s.(*state.RunState))
	return ts, nil, err
}

// FetchEVersion is called at the beginning of a transaction.
//
// The returned EVersion is compared against the one associated with a state
// loaded via GetState. If different, the transaction is aborted and new state
// isn't saved.
func (rp *runProcessor) FetchEVersion(ctx context.Context) (eventbox.EVersion, error) {
	r := &run.Run{ID: rp.runID}
	if err := datastore.Get(ctx, r); err != nil {
		return 0, errors.Annotate(err, "failed to get %q", rp.runID).Tag(transient.Tag).Err()
	}
	return eventbox.EVersion(r.EVersion), nil
}

// SaveState is called in a transaction to save the state if it has changed.
//
// The passed eversion is incremented value of eversion of what GetState
// returned before.
func (rp *runProcessor) SaveState(ctx context.Context, st eventbox.State, ev eventbox.EVersion) error {
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
	newConfigEvents struct {
		events   eventbox.Events
		hash     string
		eversion int64
	}
	clUpdatedEvents struct {
		events eventbox.Events
		cls    common.CLIDs
	}
	readyForSubmissionEvents eventbox.Events
	clSubmittedEvents        struct {
		events eventbox.Events
		cls    common.CLIDs
	}
	submissionCompletedEvent struct {
		event eventbox.Event
		sc    *eventpb.SubmissionCompleted
	}
	cqdVerificationCompletedEvents eventbox.Events
	cqdFinished                    eventbox.Events
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
		// Record all events but only the latest config hash.
		tr.newConfigEvents.events = append(tr.newConfigEvents.events, item)
		if ev := e.GetNewConfig().GetEversion(); ev > tr.newConfigEvents.eversion {
			tr.newConfigEvents.eversion = ev
			tr.newConfigEvents.hash = e.GetNewConfig().GetHash()
		}
	case *eventpb.Event_ClUpdated:
		tr.clUpdatedEvents.events = append(tr.clUpdatedEvents.events, item)
		tr.clUpdatedEvents.cls = append(tr.clUpdatedEvents.cls, common.CLID(e.GetClUpdated().GetClid()))
	case *eventpb.Event_ReadyForSubmission:
		tr.readyForSubmissionEvents = append(tr.readyForSubmissionEvents, item)
	case *eventpb.Event_ClSubmitted:
		tr.clSubmittedEvents.events = append(tr.clSubmittedEvents.events, item)
		tr.clSubmittedEvents.cls = append(tr.clSubmittedEvents.cls, common.CLID(e.GetClSubmitted().GetClid()))
	case *eventpb.Event_SubmissionCompleted:
		if tr.submissionCompletedEvent.sc != nil {
			panic("received more than 1 SubmissionCompleted result")
		}
		tr.submissionCompletedEvent.event = item
		tr.submissionCompletedEvent.sc = e.GetSubmissionCompleted()
	case *eventpb.Event_CqdVerificationCompleted:
		tr.cqdVerificationCompletedEvents = append(tr.cqdVerificationCompletedEvents, item)
	case *eventpb.Event_CqdFinished:
		tr.cqdFinished = append(tr.cqdFinished, item)
	default:
		panic(fmt.Errorf("unknown event: %T [id=%q]", e.GetEvent(), item.ID))
	}
}

func (rp *runProcessor) processTriageResults(ctx context.Context, tr *triageResult, rs *state.RunState) ([]eventbox.Transition, error) {
	statingState := rs
	var transitions []eventbox.Transition

	switch {
	case len(tr.cancelEvents) > 0:
		res, err := rp.handler.Cancel(ctx, rs)
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
		res, err := rp.handler.Start(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.startEvents, transitions)
	}

	if len(tr.newConfigEvents.events) > 0 {
		// TODO(tandrii,yiwzhang): update config.
		res, err := rp.handler.UpdateConfig(ctx, rs, tr.newConfigEvents.hash)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.newConfigEvents.events, transitions)
	}
	if len(tr.clUpdatedEvents.events) > 0 {
		res, err := rp.handler.OnCLUpdated(ctx, rs, tr.clUpdatedEvents.cls)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.clUpdatedEvents.events, transitions)
	}

	if len(tr.cqdVerificationCompletedEvents) > 0 {
		res, err := rp.handler.OnCQDVerificationCompleted(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.cqdVerificationCompletedEvents, transitions)
	}
	if len(tr.cqdFinished) > 0 {
		res, err := rp.handler.OnCQDFinished(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.cqdFinished, transitions)
	}

	if len(tr.clSubmittedEvents.events) > 0 {
		res, err := rp.handler.OnCLSubmitted(ctx, rs, tr.clSubmittedEvents.cls)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.clSubmittedEvents.events, transitions)
	}
	if sc := tr.submissionCompletedEvent.sc; sc != nil {
		res, err := rp.handler.OnSubmissionCompleted(ctx, rs, sc)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, eventbox.Events{tr.submissionCompletedEvent.event}, transitions)
	}
	if len(tr.readyForSubmissionEvents) > 0 {
		res, err := rp.handler.OnReadyForSubmission(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.readyForSubmissionEvents, transitions)
	}

	if len(tr.pokeEvents) > 0 {
		// TODO(crbug/1178658): trigger CL updater to refetch Run's CLs.
		res, err := rp.handler.Poke(ctx, rs)
		if err != nil {
			return nil, err
		}
		rs, transitions = applyResult(res, tr.pokeEvents, transitions)
	}
	// Sumbission runs as PostProcessFn after event handling/state transition
	// is done. It is possible that submission never reports the result back
	// to RM (e.g. app crash in the middle, task timeout and etc.). In that
	// case, when task retries, CV should try to resume the submission or fail
	// the submission if deadline is exceeded even though no event was received.
	// Therefore, always run TryResumeSubmission at the end regardless.
	res, err := rp.handler.TryResumeSubmission(ctx, rs)
	if err != nil {
		return nil, err
	}
	_, transitions = applyResult(res, nil, transitions)

	if err := rp.enqueueNextPoke(ctx, statingState.Run.Status, tr.nextReadyEventTime); err != nil {
		return nil, err
	}
	return transitions, nil
}

func applyResult(res *handler.Result, events eventbox.Events, transitions []eventbox.Transition) (*state.RunState, []eventbox.Transition) {
	t := eventbox.Transition{
		TransitionTo:  res.State,
		SideEffectFn:  res.SideEffectFn,
		PostProcessFn: res.PostProcessFn,
	}
	if !res.PreserveEvents {
		t.Events = events
	}
	return res.State, append(transitions, t)
}

func (rp *runProcessor) enqueueNextPoke(ctx context.Context, startingStatus run.Status, nextReadyEventTime time.Time) error {
	switch now := clock.Now(ctx); {
	case run.IsEnded(startingStatus):
		// Do not enqueue the next poke if run is ended at the beginning of the
		// state transition. Not using the end state after the state transition
		// here because CV may fail to save the state which may require the
		// recursive poke to unblock the Run.
		return nil
	case nextReadyEventTime.IsZero():
		return rp.runNotifier.PokeAfter(ctx, rp.runID, pokeInterval)
	case now.After(nextReadyEventTime):
		// It is possible that by this time, next ready event is already overdue.
		// Invoke Run Manager immediately.
		return rp.runNotifier.Invoke(ctx, rp.runID, time.Time{})
	case nextReadyEventTime.Before(now.Add(pokeInterval)):
		return rp.runNotifier.Invoke(ctx, rp.runID, nextReadyEventTime)
	default:
		return rp.runNotifier.PokeAfter(ctx, rp.runID, pokeInterval)
	}
}
