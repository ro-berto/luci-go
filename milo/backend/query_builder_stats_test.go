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
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/buildbucket/access"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	milopb "go.chromium.org/luci/milo/api/service/v1"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/milo/common/model"
	"go.chromium.org/luci/milo/common/model/milostatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/caching/cachingtest"
)

func TestQueryBuilderStats(t *testing.T) {
	t.Parallel()
	Convey(`TestQueryBuilderStats`, t, func() {
		ctx := memory.Use(context.Background())

		caches := make(map[string]caching.BlobCache)
		ctx = caching.WithGlobalCache(ctx, func(namespace string) caching.BlobCache {
			cache, ok := caches[namespace]
			if !ok {
				cache = &cachingtest.BlobCache{LRU: lru.New(0)}
				caches[namespace] = cache
			}
			return cache
		})

		datastore.GetTestable(ctx).AddIndexes(&datastore.IndexDefinition{
			Kind: "BuildSummary",
			SortBy: []datastore.IndexColumn{
				{Property: "BuilderID"},
				{Property: "Created", Descending: true},
			},
		})
		datastore.GetTestable(ctx).Consistent(true)

		accessClient := access.TestClient{}
		srv := &MiloInternalService{
			GetCachedAccessClient: func(c context.Context) (*common.CachedAccessClient, error) {
				return common.NewTestCachedAccessClient(&accessClient, caching.GlobalCache(c, "TestQueryBuilderStats")), nil
			},
		}

		builder1 := &buildbucketpb.BuilderID{
			Project: "fake_project",
			Bucket:  "fake_bucket",
			Builder: "fake_builder1",
		}
		builder2 := &buildbucketpb.BuilderID{
			Project: "fake_project",
			Bucket:  "fake_bucket",
			Builder: "fake_builder2",
		}

		createFakeBuild := func(builder *buildbucketpb.BuilderID, buildNum int, createdAt time.Time, status milostatus.Status) *model.BuildSummary {
			builderID := common.LegacyBuilderIDString(builder)
			buildID := fmt.Sprintf("%s/%d", builderID, buildNum)
			return &model.BuildSummary{
				BuildKey:  datastore.MakeKey(ctx, "build", buildID),
				ProjectID: builder.Project,
				BuilderID: builderID,
				BuildID:   buildID,
				Summary: model.Summary{
					Status: status,
				},
				Created: createdAt,
			}
		}

		baseTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
		builds := []*model.BuildSummary{
			createFakeBuild(builder1, 1, baseTime.AddDate(0, 0, -6), milostatus.Running),
			createFakeBuild(builder1, 2, baseTime.AddDate(0, 0, -5), milostatus.Running),
			createFakeBuild(builder1, 3, baseTime.AddDate(0, 0, -4), milostatus.Success),
			createFakeBuild(builder2, 4, baseTime.AddDate(0, 0, -3), milostatus.Running),
			createFakeBuild(builder1, 5, baseTime.AddDate(0, 0, -2), milostatus.Failure),
			createFakeBuild(builder2, 6, baseTime.AddDate(0, 0, -3), milostatus.NotRun),
			createFakeBuild(builder1, 7, baseTime.AddDate(0, 0, -1), milostatus.NotRun),
		}

		err := datastore.Put(ctx, builds)
		So(err, ShouldBeNil)

		err = datastore.Put(ctx, &common.Project{
			ID:      "fake_project",
			ACL:     common.ACL{Identities: []identity.Identity{"user"}},
			LogoURL: "https://logo.com",
		})
		So(err, ShouldBeNil)

		Convey(`get build stats`, func() {
			c := auth.WithState(ctx, &authtest.FakeState{Identity: "user"})

			// Mock the access client response.
			accessClient.PermittedActionsResponse = access.Permissions{
				"luci.fake_project.fake_bucket": access.AccessBucket,
			}.ToProto(time.Hour)

			res, err := srv.QueryBuilderStats(c, &milopb.QueryBuilderStatsRequest{
				Builder: builder1,
			})
			So(err, ShouldBeNil)
			So(res.PendingBuildsCount, ShouldEqual, 1)
			So(res.RunningBuildsCount, ShouldEqual, 2)
		})

		Convey(`reject users with no access`, func() {
			accessClient.PermittedActionsResponse = access.Permissions{}.ToProto(time.Hour)

			_, err := srv.QueryBuilderStats(ctx, &milopb.QueryBuilderStatsRequest{
				Builder: builder1,
			})
			So(err, ShouldNotBeNil)
		})
	})
}
