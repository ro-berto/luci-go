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

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/errors"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/common"
)

// PanicIfNotValid checks that Snapshot stored has required fields set.
func (s *Snapshot) PanicIfNotValid() {
	switch {
	case s == nil:
	case s.GetExternalUpdateTime() == nil:
		panic("missing ExternalUpdateTime")
	case s.GetLuciProject() == "":
		panic("missing LuciProject")
	case s.GetMinEquivalentPatchset() == 0:
		panic("missing MinEquivalentPatchset")
	case s.GetPatchset() == 0:
		panic("missing Patchset")

	case s.GetGerrit() == nil:
		panic("Gerrit is required, until CV supports more code reviews")
	case s.GetGerrit().GetInfo() == nil:
		panic("Gerrit.Info is required, until CV supports more code reviews")
	}
}

// LoadCLsMap loads `CL` entities which are values in the provided map.
//
// Updates `CL` entities *in place*, but also returns them as a slice.
func LoadCLsMap(ctx context.Context, m map[common.CLID]*CL) ([]*CL, error) {
	cls := make([]*CL, 0, len(m))
	for _, cl := range m {
		cls = append(cls, cl)
	}
	return loadCLs(ctx, cls)
}

// LoadCLsByIDs loads `CL` entities of the provided list of clids.
func LoadCLsByIDs(ctx context.Context, clids common.CLIDs) ([]*CL, error) {
	cls := make([]*CL, len(clids))
	for i, clid := range clids {
		cls[i] = &CL{ID: clid}
	}
	return loadCLs(ctx, cls)
}

// LoadCLs loads given `CL` entities.
func LoadCLs(ctx context.Context, cls []*CL) error {
	_, err := loadCLs(ctx, cls)
	return err
}

func loadCLs(ctx context.Context, cls []*CL) ([]*CL, error) {
	err := datastore.Get(ctx, cls)
	switch merr, ok := err.(errors.MultiError); {
	case err == nil:
		return cls, nil
	case ok:
		for i, err := range merr {
			if err == datastore.ErrNoSuchEntity {
				return nil, errors.Reason("CL %d not found in Datastore", cls[i].ID).Err()
			}
		}
		count, err := merr.Summary()
		return nil, errors.Annotate(err, "failed to load %d out of %d CLs", count, len(cls)).Tag(transient.Tag).Err()
	default:
		return nil, errors.Annotate(err, "failed to load %d CLs", len(cls)).Tag(transient.Tag).Err()
	}
}

// RemoveUnusedGerritInfo mutates given ChangeInfo to remove what CV definitely
// doesn't need to reduce bytes shuffled to/from Datastore.
//
// Doesn't complain if anything is missing.
//
// NOTE: keep this function actions in sync with storage.proto doc for
// Gerrit.info field.
func RemoveUnusedGerritInfo(ci *gerritpb.ChangeInfo) {
	const keepEmail = true
	const removeEmail = false
	cleanUser := func(u *gerritpb.AccountInfo, keepEmail bool) {
		if u == nil {
			return
		}
		u.SecondaryEmails = nil
		u.Name = ""
		u.Username = ""
		if !keepEmail {
			u.Email = ""
		}
	}

	cleanRevision := func(r *gerritpb.RevisionInfo) {
		if r == nil {
			return
		}
		cleanUser(r.GetUploader(), keepEmail)
		r.Description = ""
		// TODO(crbug/1260615): erase commit message after CQDaemon is gone.
		r.Files = nil
	}

	cleanMessage := func(m *gerritpb.ChangeMessageInfo) {
		if m == nil {
			return
		}
		cleanUser(m.GetAuthor(), removeEmail)
		cleanUser(m.GetRealAuthor(), removeEmail)
	}

	cleanLabel := func(l *gerritpb.LabelInfo) {
		if l == nil {
			return
		}
		all := l.GetAll()[:0]
		for _, a := range l.GetAll() {
			if a.GetValue() == 0 {
				continue
			}
			cleanUser(a.GetUser(), keepEmail)
			all = append(all, a)
		}
		l.All = all
	}

	for _, r := range ci.GetRevisions() {
		cleanRevision(r)
	}
	for _, m := range ci.GetMessages() {
		cleanMessage(m)
	}
	for _, l := range ci.GetLabels() {
		cleanLabel(l)
	}
	cleanUser(ci.GetOwner(), keepEmail)
}

// OwnerIdentity is the identity of a user owning this CL.
//
// Snapshot must not be nil.
func (s *Snapshot) OwnerIdentity() (identity.Identity, error) {
	if s == nil {
		panic("Snapshot is nil")
	}

	g := s.GetGerrit()
	if g == nil {
		return "", errors.New("non-Gerrit CLs not supported")
	}
	owner := g.GetInfo().GetOwner()
	if owner == nil {
		panic("Snapshot Gerrit has no owner. Bug in gerrit/updater")
	}
	email := owner.GetEmail()
	if email == "" {
		return "", errors.Reason(
			"CL %s/%d owner email of account %d is unknown",
			g.GetHost(), g.GetInfo().GetNumber(),
			owner.GetAccountId(),
		).Err()
	}
	return identity.MakeIdentity("user:" + email)
}

// IsSubmittable returns whether the change has been approved
// by the project submit rules.
func (s *Snapshot) IsSubmittable() (bool, error) {
	if s == nil {
		panic("Snapshot is nil")
	}

	g := s.GetGerrit()
	if g == nil {
		return false, errors.New("non-Gerrit CLs not supported")
	}
	return g.GetInfo().GetSubmittable(), nil
}

// IsSubmitted returns whether the change has been submitted.
func (s *Snapshot) IsSubmitted() (bool, error) {
	if s == nil {
		panic("Snapshot is nil")
	}

	g := s.GetGerrit()
	if g == nil {
		return false, errors.New("non-Gerrit CLs not supported")
	}
	return g.GetInfo().GetStatus() == gerritpb.ChangeStatus_MERGED, nil
}

func (t *UpdateCLTask) getUpdateTimeHint() *timestamppb.Timestamp {
	// TODO(crbug.com/1358208: remove this function
	if t == nil {
		return nil
	}

	if t := t.GetHint().GetExternalUpdateTime(); t != nil {
		return t
	}
	return t.GetUpdatedHint()
}
