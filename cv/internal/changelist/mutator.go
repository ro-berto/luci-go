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

package changelist

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/common"
)

// Mutator modifies CLs and guarantees at least once notification of relevant CV
// components.
//
// All CL entities in production code must be modified via the Mutator.
//
// Mutator notifies 2 CV components: Run and Project managers.
// In the future, it'll also notify Tryjob Manager.
//
// Run Manager is notified for each IncompleteRuns in the **new** CL version.
//
// Project manager is notified in following cases:
//  1. On the project in the context of which the CL is being modified.
//  2. On the project which owns the Snapshot of the *prior* CL version (if it
//     had any Snapshot).
//
// When the number of notifications is large, Mutator may chose to
// transactionally enqueue a TQ task, which will send notifications in turn.
type Mutator struct {
	pm pmNotifier
	rm rmNotifier
}

func NewMutator(pm pmNotifier, rm rmNotifier) *Mutator {
	return &Mutator{pm, rm}
}

// pmNotifier encapsulates interaction with Project Manager.
//
// In production, implemented by prjmanager.Notifier.
type pmNotifier interface {
	NotifyCLUpdated(ctx context.Context, project string, cl common.CLID, eversion int) error
}

// rmNotifier encapsulates interaction with Run Manager.
//
// In production, implemented by run.Notifier.
type rmNotifier interface {
	NotifyCLUpdated(ctx context.Context, rid common.RunID, cl common.CLID, eversion int) error
}

// ErrStopMutation is a special error used by MutateCallback to signal that no
// mutation is necessary.
//
// This is very useful because the datastore.RunInTransaction(ctx, f, ...)
// does retries by default which combined with submarine writes (transaction
// actually succeeded, but the client didn't get to know, e.g. due to network
// flake) means an idempotent MutateCallback can avoid noop updates yet still
// keep the code clean and readable. For example,
//
// ```
// cl, err := mu.Update(ctx, project, clid, func (cl *changelist.CL) error {
//   if cl.Snapshot == nil {
//     return ErrStopMutation // noop
//   }
//   cl.Snapshot = nil
//   return nil
// })
// if err != nil {
//   return errors.Annotate(err, "failed to reset Snapshot").Err()
// }
// doSomething(ctx, cl)
// ```
var ErrStopMutation = errors.New("stop CL mutation")

// MutateCallback is called by Mutator to mutate the CL inside a transaction.
//
// The function should be idempotent.
//
// If no error is returned, Mutator proceeds saving the CL.
//
// If special ErrStopMutation is returned, Mutator aborts the tranasction and
// returns existing CL read from Datastore and no error. In the special case of
// Upsert(), the returned CL may actually be nil if CL didn't exist.
//
// If any error is returned other than ErrStopMutation, Mutator aborts the
// transaction and returns nil CL and the exact same error.
type MutateCallback func(cl *CL) error

// Upsert creates new or updates existing CL via a dedicated transaction in the
// context of the given LUCI project.
//
// Prefer to use Update if CL ID is known.
//
// If CL didn't exist before, the callback is provided a CL with temporarily
// reserved ID. Until Upsert returns with success, this ID is not final,
// but it's fine to use it in other entities saved within the same transaction.
//
// If CL didn't exist before and the callback returns ErrStopMutation, then
// Upsert returns (nil, nil).
func (m *Mutator) Upsert(ctx context.Context, project string, eid ExternalID, clbk MutateCallback) (*CL, error) {
	// Quick path in case CL already exists, which is a common case,
	// and can usually be satisfied by dscache lookup.
	mapEntity := clMap{ExternalID: eid}
	switch err := datastore.Get(ctx, &mapEntity); {
	case err == datastore.ErrNoSuchEntity:
		// OK, proceed to slow path below.
	case err != nil:
		return nil, errors.Annotate(err, "failed to get clMap entity %q", eid).Tag(transient.Tag).Err()
	default:
		return m.Update(ctx, project, mapEntity.InternalID, clbk)
	}

	var result *CL
	var innerErr error
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) (err error) {
		defer func() { innerErr = err }()
		// Check if CL exists and prepare appropriate clMutation.
		var clMutation *CLMutation
		mapEntity := clMap{ExternalID: eid}
		switch err := datastore.Get(ctx, &mapEntity); {
		case err == datastore.ErrNoSuchEntity:
			clMutation, err = m.beginInsert(ctx, project, eid)
			if err != nil {
				return err
			}
		case err != nil:
			return errors.Annotate(err, "failed to get clMap entity %q", eid).Tag(transient.Tag).Err()
		default:
			clMutation, err = m.Begin(ctx, project, mapEntity.InternalID)
			if err != nil {
				return err
			}
			result = clMutation.CL
		}
		if err := clbk(clMutation.CL); err != nil {
			return err
		}
		result, err = clMutation.Finalize(ctx)
		return err
	}, nil)
	switch {
	case innerErr == ErrStopMutation:
		return result, nil
	case innerErr != nil:
		return nil, innerErr
	case err != nil:
		return nil, errors.Annotate(err, "failed to commit Upsert of CL %q", eid).Tag(transient.Tag).Err()
	default:
		return result, nil
	}
}

// Update mutates one CL via a dedicated transaction in the context of the given
// LUCI project.
//
// If the callback returns ErrStopMutation, then Update returns the read CL
// entity and nil error.
func (m *Mutator) Update(ctx context.Context, project string, id common.CLID, clbk MutateCallback) (*CL, error) {
	var result *CL
	var innerErr error
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) (err error) {
		defer func() { innerErr = err }()
		clMutation, err := m.Begin(ctx, project, id)
		if err != nil {
			return err
		}
		result = clMutation.CL
		if err := clbk(clMutation.CL); err != nil {
			return err
		}
		result, err = clMutation.Finalize(ctx)
		return err
	}, nil)
	switch {
	case innerErr == ErrStopMutation:
		return result, nil
	case innerErr != nil:
		return nil, innerErr
	case err != nil:
		return nil, errors.Annotate(err, "failed to commit update on CL %d", id).Tag(transient.Tag).Err()
	default:
		return result, nil
	}
}

// CLMutation encapsulates one CL mutation.
type CLMutation struct {
	// CL can be modified except the following fields:
	//  * ID
	//  * ExternalID
	//  * EVersion
	//  * UpdateTime
	CL *CL

	// m is a back reference to its parent -- Mutator.
	m *Mutator

	// trans is only to detect incorrect usage.
	trans datastore.Transaction
	// project in the context of which CL is modified.
	project string

	id         common.CLID
	externalID ExternalID

	priorEversion   int
	priorUpdateTime time.Time
	priorProject    string
}

func (m *Mutator) beginInsert(ctx context.Context, project string, eid ExternalID) (*CLMutation, error) {
	clMutation := &CLMutation{
		CL:      &CL{ExternalID: eid},
		m:       m,
		trans:   datastore.CurrentTransaction(ctx),
		project: project,
	}
	if err := datastore.AllocateIDs(ctx, clMutation.CL); err != nil {
		return nil, errors.Annotate(err, "failed to allocate new CL ID for %q", eid).Tag(transient.Tag).Err()
	}
	if err := datastore.Put(ctx, &clMap{ExternalID: eid, InternalID: clMutation.CL.ID}); err != nil {
		return nil, errors.Annotate(err, "failed to insert clMap entity for %q", eid).Tag(transient.Tag).Err()
	}
	clMutation.backup()
	return clMutation, nil
}

// Begin starts mutation of one CL inside an existing transaction in the context of
// the given LUCI project.
func (m *Mutator) Begin(ctx context.Context, project string, id common.CLID) (*CLMutation, error) {
	clMutation := &CLMutation{
		CL:      &CL{ID: id},
		m:       m,
		trans:   datastore.CurrentTransaction(ctx),
		project: project,
	}
	if clMutation.trans == nil {
		panic(fmt.Errorf("changelist.Mutator.Begin must be called inside an existing Datastore transaction"))
	}
	switch err := datastore.Get(ctx, clMutation.CL); {
	case err == datastore.ErrNoSuchEntity:
		return nil, errors.Annotate(err, "CL %d doesn't exist", id).Err()
	case err != nil:
		return nil, errors.Annotate(err, "failed to get CL %d", id).Tag(transient.Tag).Err()
	}
	clMutation.backup()
	return clMutation, nil
}

func (clm *CLMutation) backup() {
	clm.id = clm.CL.ID
	clm.externalID = clm.CL.ExternalID
	clm.priorEversion = clm.CL.EVersion
	clm.priorUpdateTime = clm.CL.UpdateTime
	if p := clm.CL.Snapshot.GetLuciProject(); p != "" {
		clm.priorProject = p
	}
}

// Finalize finalizes CL mutation.
//
// Must be called at most once.
// Must be called in the same Datastore transaction as Begin() which began the
// CL mutation.
func (clm *CLMutation) Finalize(ctx context.Context) (*CL, error) {
	clm.finalize(ctx)
	if err := datastore.Put(ctx, clm.CL); err != nil {
		return nil, errors.Annotate(err, "failed to put CL %d", clm.id).Tag(transient.Tag).Err()
	}
	if err := clm.m.notifyOne(ctx, clm); err != nil {
		return nil, err
	}
	return clm.CL, nil
}

func (clm *CLMutation) finalize(ctx context.Context) {
	switch t := datastore.CurrentTransaction(ctx); {
	case clm.trans == nil:
		panic(fmt.Errorf("changelist.CLMutation.Finalize called the second time"))
	case t == nil:
		panic(fmt.Errorf("changelist.CLMutation.Finalize must be called inside an existing Datastore transaction"))
	case t != clm.trans:
		panic(fmt.Errorf("changelist.CLMutation.Finalize called inside a different Datastore transaction"))
	}
	clm.trans = nil

	switch {
	case clm.id != clm.CL.ID:
		panic(fmt.Errorf("CL.ID must be not be modified"))
	case clm.externalID != clm.CL.ExternalID:
		panic(fmt.Errorf("CL.ExternalID must be not be modified"))
	case clm.priorEversion != clm.CL.EVersion:
		panic(fmt.Errorf("CL.EVersion must be not be modified"))
	case !clm.priorUpdateTime.Equal(clm.CL.UpdateTime):
		panic(fmt.Errorf("CL.UpdateTime must be not be modified"))
	}
	clm.CL.EVersion++
	clm.CL.UpdateTime = datastore.RoundTime(clock.Now(ctx).UTC())
}

func (clm *CLMutation) BeginBatch(ctx context.Context, project string, ids common.CLIDs) (*CL, error) {
	panic("not implemented")
}

func (clm *CLMutation) FinalizeBatch(ctx context.Context) ([]*CL, error) {
	panic("not implemented")
}

type notification struct {
	id       common.CLID
	ev       int
	projects []string
	runs     common.RunIDs
}

func (clm *CLMutation) notification() *notification {
	n := &notification{
		id:       clm.id,
		ev:       clm.CL.EVersion,
		runs:     append(common.RunIDs(nil), clm.CL.IncompleteRuns...), // copy
		projects: make([]string, 1, 2),
	}
	n.projects[0] = clm.project
	if clm.priorProject != "" && clm.project != clm.priorProject {
		n.projects = append(n.projects, clm.priorProject)
	}
	return n
}

func (m *Mutator) notifyOne(ctx context.Context, clm *CLMutation) error {
	n := clm.notification()
	eg, ctx := errgroup.WithContext(ctx)
	for _, p := range n.projects {
		p := p
		eg.Go(func() error { return m.pm.NotifyCLUpdated(ctx, p, n.id, n.ev) })
	}
	// One CL should have very few Runs, so it's fine to process each within the
	// transaction in parallel.
	for _, r := range n.runs {
		r := r
		eg.Go(func() error { return m.rm.NotifyCLUpdated(ctx, r, n.id, n.ev) })
	}
	return eg.Wait()
}
