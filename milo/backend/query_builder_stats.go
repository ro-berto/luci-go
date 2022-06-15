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

package backend

import (
	"context"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/buildbucket/bbperms"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"
	milopb "go.chromium.org/luci/milo/api/service/v1"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/milo/common/model/milostatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/realms"
	"google.golang.org/grpc/codes"
)

// QueryBuilderStats implements milopb.MiloInternal service
func (s *MiloInternalService) QueryBuilderStats(ctx context.Context, req *milopb.QueryBuilderStatsRequest) (_ *milopb.BuilderStats, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()

	// Validate request.
	err = validatesQueryBuilderStatsRequest(req)
	if err != nil {
		return nil, appstatus.BadRequest(err)
	}

	// Perform ACL check.
	realm := realms.Join(req.Builder.Project, req.Builder.Bucket)
	allowed, err := auth.HasPermission(ctx, bbperms.BuildsList, realm, nil)
	if err != nil {
		return nil, err
	}
	if !allowed {
		if auth.CurrentIdentity(ctx) == identity.AnonymousIdentity {
			return nil, appstatus.Error(codes.Unauthenticated, "not logged in")
		}
		return nil, appstatus.Error(codes.PermissionDenied, "no access to the bucket")
	}

	legacyBuilderID := common.LegacyBuilderIDString(req.Builder)
	stats := &milopb.BuilderStats{}

	err = parallel.FanOutIn(func(fetch chan<- func() error) {

		// Pending builds
		fetch <- func() error {
			q := datastore.NewQuery("BuildSummary").
				Eq("BuilderID", legacyBuilderID).
				Eq("Summary.Status", milostatus.NotRun)
			pending, err := datastore.Count(ctx, q)
			stats.PendingBuildsCount = int32(pending)
			return err
		}

		// Running builds
		fetch <- func() error {
			q := datastore.NewQuery("BuildSummary").
				Eq("BuilderID", legacyBuilderID).
				Eq("Summary.Status", milostatus.Running)
			running, err := datastore.Count(ctx, q)
			stats.RunningBuildsCount = int32(running)
			return err
		}
	})

	if err != nil {
		return nil, err
	}

	return stats, nil
}

func validatesQueryBuilderStatsRequest(req *milopb.QueryBuilderStatsRequest) error {
	if req.Builder == nil || req.Builder.Project == "" || req.Builder.Bucket == "" || req.Builder.Builder == "" {

		return errors.Reason("builder_id is required").Err()
	}
	return nil
}
