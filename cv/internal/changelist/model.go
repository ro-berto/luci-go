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

package changelist

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/common"
)

// ExternalID is a unique CL ID deterministically constructed based on CL data.
//
// Currently, only Gerrit is supported.
type ExternalID string

// GobID makes an ExternalID for a Gerrit CL.
//
// Host is typically "something-review.googlesource.com".
// Change is a number, e.g. 2515619 for
// https://chromium-review.googlesource.com/c/infra/luci/luci-go/+/2515619
func GobID(host string, change int64) (ExternalID, error) {
	if strings.ContainsRune(host, '/') {
		return "", errors.Reason("invalid host %q: must not contain /", host).Err()
	}
	return ExternalID(fmt.Sprintf("gerrit/%s/%d", host, change)), nil
}

// ParseGobID returns Gerrit host and change if this is a GobID.
func (e ExternalID) ParseGobID() (host string, change int64, err error) {
	parts := strings.Split(string(e), "/")
	if len(parts) != 3 || parts[0] != "gerrit" {
		err = errors.Reason("%q is not a valid GobID", e).Err()
		return
	}
	host = parts[1]
	change, err = strconv.ParseInt(parts[2], 10, 63)
	if err != nil {
		err = errors.Annotate(err, "%q is not a valid GobID", e).Err()
	}
	return
}

// CL is a CL entity in Datastore.
type CL struct {
	_kind  string                `gae:"$kind,CL"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	// ID is auto-generated by Datastore.
	ID common.CLID `gae:"$id"` // int64
	// ExternalID must not be modified once entity is created.
	ExternalID ExternalID `gae:",noindex"` // string. Indexed in CLMap entities.

	// EVersion is entity version. Every update should increment it by 1.
	// See Update() function.
	EVersion int `gae:",noindex"`

	// Snapshot is latest known state of a CL.
	// It may and often is behind the source of truth -- the code reveview site
	// (e.g. Gerrit).
	Snapshot *Snapshot

	// ApplicableConfig keeps track of configs applicable to the CL.
	ApplicableConfig *ApplicableConfig

	// DependentMeta stores metadata per LUCI project about this CL being a dependency
	// of another CL in the context of the specific LUCI project.
	//
	// See description in protobuf type with the same name.
	DependentMeta *DependentMeta

	// UpdateTime is exact time of when this entity was last updated.
	//
	// It's not indexed to avoid hot areas in the index.
	UpdateTime time.Time `gae:",noindex"`

	// TODO(tandrii): implement deletion of the oldest entities via additional
	// indexed field based on UpdateTime but with entropy in the lowest bits to
	// avoid hotspots.
}

// clMap is CLMap entity in Datastore which ensures strict 1:1 mapping
// between internal and external IDs.
type clMap struct {
	_kind string `gae:"$kind,CLMap"`

	// ExternalID as entity ID ensures uniqueness.
	ExternalID ExternalID `gae:"$id"` // string
	// InternalID is auto-generated by Datastore for CL entity.
	InternalID common.CLID `gae:",noindex"` // int64. Indexed in CL entities.
}

// Get reads a CL from datastore.
//
// Returns datastore.ErrNoSuchEntity if it doesn't exist.
func (eid ExternalID) Get(ctx context.Context) (*CL, error) {
	m := clMap{ExternalID: eid}
	switch err := datastore.Get(ctx, &m); {
	case err == datastore.ErrNoSuchEntity:
		return nil, err
	case err != nil:
		return nil, errors.Annotate(err, "failed to get CLMap").Tag(transient.Tag).Err()
	}
	return getExisting(ctx, m.InternalID, eid)
}

// GetOrInsert reads a CL from datastore, creating a new one if not exists yet.
//
// populate is called within a transaction to populate fields of a new entity.
// It should be a fast function.
//
// Warning:
//  * populate may be called several times since transaction can be retried.
//  * cl.ExternalID and cl.ID must not be changed by populate.
func (eid ExternalID) GetOrInsert(ctx context.Context, populate func(cl *CL)) (*CL, error) {
	// Fast path without transaction.
	if cl, err := eid.Get(ctx); err != datastore.ErrNoSuchEntity {
		return cl, err
	}
	var cl *CL
	m := clMap{ExternalID: eid}
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) (err error) {
		cl = nil
		switch err = datastore.Get(ctx, &m); {
		case err == nil:
			// Has just been created by someone else.
			return nil
		case err != datastore.ErrNoSuchEntity:
			return err
		}
		cl, err = insert(ctx, eid, populate)
		return
	}, nil)

	switch {
	case err != nil:
		return nil, errors.Annotate(err, "failed to getOrInsert a CL").Tag(transient.Tag).Err()
	case cl == nil:
		return getExisting(ctx, m.InternalID, eid)
	}
	return cl, nil
}

// Delete deletes CL and its CLMap entities trasactionally.
//
// Thus, Delete and insertion (part of ExternalID.getOrInsert)
// are atomic with respect to one another.
//
// However, ExternalID.get and fast path of ExternalID.getOrInsert if called
// concurrently with Delete may return temporary error, but on retry they would
// return ErrNoSuchEntity.
func Delete(ctx context.Context, id common.CLID) error {
	cl := CL{ID: id}
	switch err := datastore.Get(ctx, &cl); {
	case err == datastore.ErrNoSuchEntity:
		return nil // Nothing to do.
	case err != nil:
		return errors.Annotate(err, "failed to get a CL").Tag(transient.Tag).Err()
	}

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		m := clMap{ExternalID: cl.ExternalID}
		return datastore.Delete(ctx, &cl, &m)
	}, nil)
	if err != nil {
		return errors.Annotate(err, "failed to delete a CL").Tag(transient.Tag).Err()
	}
	return nil
}

// UpdateFields defines what parts of CL metadata to update.
//
// At least one field must be specified.
type UpdateFields struct {
	// Snapshot overwrites existing CL snapshot if newer according to its
	// .ExternalUpdateTime.
	Snapshot *Snapshot
	// ApplicableConfig overwrites existing CL ApplicableConfig if newer
	// accordingto to its .UpdateTime.
	ApplicableConfig *ApplicableConfig

	// AddDependentMeta adds or overwrites metadata per LUCI project in CL AsDepMeta.
	// Doesn't affect metadata stored for projects not referenced here.
	AddDependentMeta *DependentMeta
}

// IsEmpty returns true if no updates are necessary.
func (u UpdateFields) IsEmpty() bool {
	return (u.Snapshot == nil &&
		u.ApplicableConfig == nil &&
		len(u.AddDependentMeta.GetByProject()) == 0)
}

func (u UpdateFields) apply(cl *CL) (changed bool) {
	if u.ApplicableConfig != nil && !cl.ApplicableConfig.IsUpToDate(
		u.ApplicableConfig.GetUpdateTime().AsTime()) {
		cl.ApplicableConfig = u.ApplicableConfig
		changed = true
	}

	if s := u.Snapshot; s != nil && !cl.Snapshot.IsUpToDate(
		s.GetLuciProject(), s.GetExternalUpdateTime().AsTime()) {
		cl.Snapshot = s
		changed = true
		// Wipe out corresponding AsDepMeta entry if any.
		if m := cl.DependentMeta.GetByProject(); m != nil {
			if _, exists := m[s.GetLuciProject()]; exists {
				delete(m, s.GetLuciProject())
				if len(m) == 0 {
					cl.DependentMeta = nil
				}
			}
		}
	}

	switch {
	case u.AddDependentMeta == nil:
	case cl.DependentMeta == nil || cl.DependentMeta.GetByProject() == nil:
		cl.DependentMeta = u.AddDependentMeta
		changed = true
	default:
		e := cl.DependentMeta.GetByProject()
		for lProject, ameta := range u.AddDependentMeta.GetByProject() {
			emeta, exists := e[lProject]
			if !exists || emeta.GetUpdateTime().AsTime().Before(ameta.GetUpdateTime().AsTime()) {
				e[lProject] = ameta
				changed = true
			}
		}
	}
	return
}

// Update updates CL entity.
//
// Either ExternalID or a known common.CLID must be provided.
//
// If common.CLID is not known and CL for provided ExternalID doesn't exist,
// then a new CL is created with values from UpdateFields.
// Otherwise, an existing CL entity will be updated as documented in UpdateFields.
//
// TODO(tandrii): emit notification events.
func Update(ctx context.Context, eid ExternalID, knownCLID common.CLID, fields UpdateFields) error {
	if eid == "" && knownCLID == 0 {
		panic("either ExternalID or known common.CLID must be provided")
	}
	if fields.IsEmpty() {
		panic("must specify at least one UpdateFields field")
	}

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		if knownCLID == 0 {
			m := clMap{ExternalID: eid}
			switch err := datastore.Get(ctx, &m); {
			case err == datastore.ErrNoSuchEntity:
				// Insert new entity.
				_, err = insert(ctx, eid, func(cl *CL) {
					cl.Snapshot = fields.Snapshot
					cl.ApplicableConfig = fields.ApplicableConfig
					cl.DependentMeta = fields.AddDependentMeta
				})
				return err
			case err != nil:
				return errors.Annotate(err, "failed to get CLMap entity").Tag(transient.Tag).Err()
			}
			knownCLID = m.InternalID
		}
		cl, err := getExisting(ctx, knownCLID, eid)
		if err != nil {
			return err
		}
		// Update exsting entity.
		return update(ctx, cl, fields.apply)
	}, nil)
	return errors.Annotate(err, "failed to update CL").Tag(transient.Tag).Err()
}

func getExisting(ctx context.Context, clid common.CLID, eid ExternalID) (*CL, error) {
	cl := &CL{ID: clid}
	switch err := datastore.Get(ctx, cl); {
	case err == datastore.ErrNoSuchEntity:
		// This should not happen in practice except in the case of a very old CL
		// which is being deleted due to retention policy. Log error but return it
		// as transient as it's expected that CLMap entity would be removed soon,
		// and so a retry would be produce proper datastore.ErrNoSuchEntity error.
		msg := fmt.Sprintf("unexpectedly failed to get CL#%d given existing CLMap%q", clid, eid)
		logging.Errorf(ctx, msg)
		return nil, errors.Reason(msg).Tag(transient.Tag).Err()
	case err != nil:
		return nil, errors.Annotate(err, "failed to get CL").Tag(transient.Tag).Err()
	}
	return cl, nil
}

// insert creates new CL entity for given external ID.
//
// Must be called after verifying that such CLMap record doesn't exist.
func insert(ctx context.Context, eid ExternalID, populate func(*CL)) (*CL, error) {
	if datastore.CurrentTransaction(ctx) == nil {
		panic("must be called in transaction context")
	}
	// Create new CL and CLMap entry atomically.
	cl := &CL{
		ID:         0, // autogenerate by Datastore
		ExternalID: eid,
		EVersion:   1,
	}
	populate(cl)
	if cl.ID != 0 || cl.ExternalID != eid || cl.EVersion != 1 {
		panic(errors.New("populate changed ID or ExternalID or EVersion, but must not do this."))
	}
	cl.Snapshot.PanicIfNotValid()
	// datastore.Put will do RoundTime on its own, but without affecting our `cl`
	// object. Since `cl` object is passed outside, do rounding here s.t. it has
	// exact same data as would have been read from datastore right after the Put.
	cl.UpdateTime = datastore.RoundTime(clock.Now(ctx).UTC())

	if err := datastore.Put(ctx, cl); err != nil {
		return nil, errors.Annotate(err, "failed to save CL entity").Tag(transient.Tag).Err()
	}
	if err := datastore.Put(ctx, &clMap{ExternalID: eid, InternalID: cl.ID}); err != nil {
		return nil, errors.Annotate(err, "failed to save CLMap entity").Tag(transient.Tag).Err()
	}
	return cl, nil
}

func update(ctx context.Context, justRead *CL, mut func(*CL) (update bool)) error {
	if datastore.CurrentTransaction(ctx) == nil {
		panic("must be called in transaction context")
	}

	before := *justRead // shallow copy, avoiding cloning Snapshot.
	if !mut(justRead) {
		return nil
	}
	justRead.Snapshot.PanicIfNotValid()
	justRead.EVersion = before.EVersion + 1
	justRead.UpdateTime = clock.Now(ctx).UTC()
	if err := datastore.Put(ctx, justRead); err != nil {
		return errors.Annotate(err, "failed to put CL entity").Tag(transient.Tag).Err()
	}
	return nil
}
