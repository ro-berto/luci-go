// Copyright 2018 The LUCI Authors.
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

package model

import (
	"context"
	"sort"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/common"
)

// Instance represents a package instance as it is stored in the datastore.
//
// Contains some instance metadata, in particular a list of processors that
// scanned the instance (see below).
//
// The parent entity is the corresponding package entity. ID is derived from
// package instance file digest, see common.ObjectRefToInstanceID().
//
// Compatible with the python version of the backend.
type Instance struct {
	_kind  string                `gae:"$kind,PackageInstance"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	InstanceID string         `gae:"$id"`     // see common.ObjectRefToInstanceID()
	Package    *datastore.Key `gae:"$parent"` // see PackageKey()

	RegisteredBy string    `gae:"registered_by"` // who registered it
	RegisteredTs time.Time `gae:"registered_ts"` // when it was registered

	// Names of currently running or scheduled processors.
	ProcessorsPending []string `gae:"processors_pending"`
	// Names of processors that successfully finished the processing.
	ProcessorsSuccess []string `gae:"processors_success"`
	// Names of processors that returned fatal errors.
	ProcessorsFailure []string `gae:"processors_failure"`
}

// Proto returns cipd.Instance proto with information from this entity.
func (e *Instance) Proto() *api.Instance {
	var t *timestamppb.Timestamp
	if !e.RegisteredTs.IsZero() {
		t = timestamppb.New(e.RegisteredTs)
	}
	return &api.Instance{
		Package:      e.Package.StringID(),
		Instance:     common.InstanceIDToObjectRef(e.InstanceID),
		RegisteredBy: e.RegisteredBy,
		RegisteredTs: t,
	}
}

// FromProto fills in the entity based on the proto message.
//
// Returns the entity itself for easier chaining.
//
// Doesn't touch output-only fields at all.
func (e *Instance) FromProto(c context.Context, p *api.Instance) *Instance {
	e.InstanceID = common.ObjectRefToInstanceID(p.Instance)
	e.Package = PackageKey(c, p.Package)
	return e
}

// RegisterInstance transactionally registers an instance (and the corresponding
// package), if it isn't registered already.
//
// Calls the given callback (inside the transaction) if it is indeed registering
// a new instance. The callback may mutate the instance entity before it is
// stored. The callback may be called multiple times in case of retries (each
// time it will be given a fresh instance to be mutated).
//
// Returns (true, entity, nil) if the instance was registered just now or
// (false, entity, nil) if it existed before.
//
// In either case, it returns the entity that is stored now in the datastore.
// It is either the new instance, or something that existed there before.
func RegisterInstance(c context.Context, inst *Instance, cb func(context.Context, *Instance) error) (reg bool, out *Instance, err error) {
	err = Txn(c, "RegisterInstance", func(c context.Context) error {
		// Reset the state in case of a txn retry.
		reg = false
		out = nil
		events := Events{}

		// Such instance is already registered?
		existing := Instance{
			InstanceID: inst.InstanceID,
			Package:    inst.Package,
		}
		switch err := datastore.Get(c, &existing); {
		case err == nil:
			out = &existing
			return nil
		case err != datastore.ErrNoSuchEntity:
			return errors.Annotate(err, "failed to fetch the instance entity").Tag(transient.Tag).Err()
		}

		// Register the package entity too, if missing.
		switch res, err := datastore.Exists(c, inst.Package); {
		case err != nil:
			return errors.Annotate(err, "failed to fetch the package entity existence").Tag(transient.Tag).Err()
		case !res.Any():
			err := datastore.Put(c, &Package{
				Name:         inst.Package.StringID(),
				RegisteredBy: inst.RegisteredBy,
				RegisteredTs: inst.RegisteredTs,
			})
			if err != nil {
				return errors.Annotate(err, "failed to create the package entity").Tag(transient.Tag).Err()
			}
			events.Emit(&api.Event{
				Kind:    api.EventKind_PACKAGE_CREATED,
				Package: inst.Package.StringID(),
			})
		}

		// Let the caller do more stuff inside this txn, e.g start TQ tasks.
		toPut := *inst
		if cb != nil {
			if err := cb(c, &toPut); err != nil {
				return errors.Annotate(err, "instance registration callback error").Err()
			}
		}

		// Finally register the package instance entity.
		if err := datastore.Put(c, &toPut); err != nil {
			return errors.Annotate(err, "failed to create the package instance entity").Tag(transient.Tag).Err()
		}
		events.Emit(&api.Event{
			Kind:     api.EventKind_INSTANCE_CREATED,
			Package:  inst.Package.StringID(),
			Instance: inst.InstanceID,
		})

		reg = true
		out = &toPut
		return events.Flush(c)
	})
	return
}

// ListInstances lists instances of a package, more recent first.
//
// Only does a query over Instances entities. Doesn't check whether the Package
// entity exists. Returns up to pageSize entities, plus non-nil cursor (if
// there are more results). pageSize must be positive.
func ListInstances(c context.Context, pkg string, pageSize int32, cursor datastore.Cursor) (out []*Instance, nextCur datastore.Cursor, err error) {
	if pageSize <= 0 {
		panic("pageSize must be positive")
	}

	q := datastore.NewQuery("PackageInstance").
		Ancestor(PackageKey(c, pkg)).
		Order("-registered_ts").
		Limit(pageSize)
	if cursor != nil {
		q = q.Start(cursor)
	}

	err = datastore.Run(c, q, func(ent *Instance, cb datastore.CursorCB) error {
		out = append(out, ent)
		if len(out) >= int(pageSize) {
			if nextCur, err = cb(); err != nil {
				return err
			}
			return datastore.Stop
		}
		return nil
	})
	if err != nil {
		return nil, nil, transient.Tag.Apply(err)
	}
	return
}

// CheckInstanceExists fetches the instance and verifies it exists.
//
// Can be called as part of a transaction. Updates 'inst' in place.
//
// Returns gRPC-tagged NotFound error if there's no such instance or package.
func CheckInstanceExists(c context.Context, inst *Instance) error {
	switch err := datastore.Get(c, inst); {
	case err == datastore.ErrNoSuchEntity:
		// Maybe the package is missing completely?
		if err := CheckPackageExists(c, inst.Package.StringID()); err != nil {
			return err
		}
		return errors.Reason("no such instance").Tag(grpcutil.NotFoundTag).Err()
	case err != nil:
		return errors.Annotate(err, "failed to check the instance presence").Tag(transient.Tag).Err()
	default:
		return nil
	}
}

// CheckInstanceReady fetches the instance and verifies it exists and has zero
// pending or failed processors.
//
// Can be called as part of a transaction. Updates 'inst' in place.
//
// Returns gRPC-tagged errors:
//    NotFound if there's no such instance or package.
//    FailedPrecondition if some processors are still running.
//    Aborted if some processors have failed.
func CheckInstanceReady(c context.Context, inst *Instance) error {
	switch err := CheckInstanceExists(c, inst); {
	case err != nil:
		return err
	case len(inst.ProcessorsFailure) != 0:
		return errors.Reason("some processors failed to process this instance: %s",
			strings.Join(inst.ProcessorsFailure, ", ")).Tag(grpcutil.AbortedTag).Err()
	case len(inst.ProcessorsPending) != 0:
		return errors.Reason("the instance is not ready yet, pending processors: %s",
			strings.Join(inst.ProcessorsPending, ", ")).Tag(grpcutil.FailedPreconditionTag).Err()
	default:
		return nil
	}
}

// FetchProcessors fetches results of all processors assigned to the instance
// and returns them as cipd.Processor proto messages (sorted by processor ID).
func FetchProcessors(c context.Context, inst *Instance) ([]*api.Processor, error) {
	count := len(inst.ProcessorsPending) +
		len(inst.ProcessorsSuccess) +
		len(inst.ProcessorsFailure)
	if count == 0 {
		return nil, nil
	}

	key := datastore.KeyForObj(c, inst)

	// Fetch results of all finished processors. All entities should exist, since
	// ProcessorsSuccess/ProcessorsFailure is updated transactionally when
	// entities are created.
	finished := make([]*ProcessingResult, 0, len(inst.ProcessorsSuccess)+len(inst.ProcessorsFailure))
	for _, p := range inst.ProcessorsSuccess {
		finished = append(finished, &ProcessingResult{ProcID: p, Instance: key})
	}
	for _, p := range inst.ProcessorsFailure {
		finished = append(finished, &ProcessingResult{ProcID: p, Instance: key})
	}
	if err := datastore.Get(c, finished); err != nil {
		return nil, errors.Annotate(err, "failed to fetch finished processors").Tag(transient.Tag).Err()
	}

	// Convert all them to proto.
	out := make([]*api.Processor, 0, count)
	for _, p := range finished {
		proc, err := p.Proto()
		if err != nil {
			return nil, errors.Annotate(err, "failed to read results of processor %q", p.ProcID).Err()
		}
		out = append(out, proc)
	}

	// Add pending processors to the output too. They don't have entities in the
	// datastore, so we haven't fetched anything for them.
	for _, p := range inst.ProcessorsPending {
		proc, err := (&ProcessingResult{ProcID: p, Instance: key}).Proto()
		if err != nil {
			panic(err) // impossible for empty ProcessingResult{}
		}
		out = append(out, proc)
	}

	// Sort by processor ID, as promised by the API.
	sort.Slice(out, func(i, j int) bool { return out[i].Id < out[j].Id })
	return out, nil
}
