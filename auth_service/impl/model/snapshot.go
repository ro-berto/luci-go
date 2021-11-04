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

package model

import (
	"context"

	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/common/trace"
	"go.chromium.org/luci/gae/service/datastore"
)

// Snapshot contains transactionally captured AuthDB entities.
type Snapshot struct {
	ReplicationState *AuthReplicationState
	GlobalConfig     *AuthGlobalConfig
	Groups           []*AuthGroup
	IPAllowlists     []*AuthIPAllowlist

	// TODO:
	//   IPAllowlistAssignments
	//   RealmsGlobals
	//   ProjectRealms
}

// TakeSnapshot takes a consistent snapshot of the replicated subset of AuthDB
// entities.
//
// Runs a read-only transaction internally.
func TakeSnapshot(ctx context.Context) (snap *Snapshot, err error) {
	// This is a potentially slow operation. Capture it in the trace.
	ctx, span := trace.StartSpan(ctx, "go.chromium.org/luci/auth_service/impl/model/TakeSnapshot")
	defer func() { span.End(err) }()

	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		snap = &Snapshot{
			ReplicationState: &AuthReplicationState{
				Kind:   "AuthReplicationState",
				ID:     "self",
				Parent: RootKey(ctx),
			},
			GlobalConfig: &AuthGlobalConfig{
				Kind: "AuthGlobalConfig",
				ID:   "root",
			},
		}

		gr, ctx := errgroup.WithContext(ctx)
		gr.Go(func() error {
			return datastore.Get(ctx, snap.GlobalConfig, snap.ReplicationState)
		})
		gr.Go(func() (err error) {
			snap.Groups, err = GetAllAuthGroups(ctx)
			return
		})
		gr.Go(func() (err error) {
			snap.IPAllowlists, err = GetAllAuthIPAllowlists(ctx)
			return
		})

		// TODO:
		//  IPAllowlistAssignments
		//  RealmsGlobals
		//  ProjectRealms

		return gr.Wait()
	}, &datastore.TransactionOptions{ReadOnly: true})

	if err != nil {
		return nil, err
	}
	return snap, nil
}
