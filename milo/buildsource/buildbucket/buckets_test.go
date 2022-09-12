// Copyright 2022 The LUCI Authors.
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

package buildbucket

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/buildbucket/bbperms"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
)

const testIdentity = identity.Identity("user:test@example.com")

func TestFilterVisibleBuilders(t *testing.T) {
	t.Parallel()

	Convey(`FilterVisibleBuilders`, t, func() {
		s := &authtest.FakeState{
			FakeDB: authtest.NewFakeDB(
				authtest.MockPermission(testIdentity, "proj1:bucket1", bbperms.BuildersList),
				authtest.MockPermission(testIdentity, "proj2:bucket1", bbperms.BuildersList),
			),
			Identity: testIdentity,
		}
		ctx := auth.WithState(context.Background(), s)

		builders := []*buildbucketpb.BuilderID{
			{
				Project: "proj1",
				Bucket:  "bucket1",
				Builder: "builder1",
			},
			{
				Project: "proj1",
				Bucket:  "bucket2",
				Builder: "builder1",
			},
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder1",
			},
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder2",
			},
			{
				Project: "proj2",
				Bucket:  "bucket2",
				Builder: "builder1",
			},
		}

		visibleBuilders, err := FilterVisibleBuilders(ctx, builders, "")
		So(err, ShouldBeNil)
		So(visibleBuilders, ShouldResemble, []*buildbucketpb.BuilderID{
			{
				Project: "proj1",
				Bucket:  "bucket1",
				Builder: "builder1",
			},
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder1",
			},
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder2",
			},
		})

		visibleBuildersInProj2, err := FilterVisibleBuilders(ctx, builders, "proj2")
		So(err, ShouldBeNil)
		So(visibleBuildersInProj2, ShouldResemble, []*buildbucketpb.BuilderID{
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder1",
			},
			{
				Project: "proj2",
				Bucket:  "bucket1",
				Builder: "builder2",
			},
		})
	})
}
