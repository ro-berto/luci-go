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

package perm

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/auth/realms"
	"go.chromium.org/luci/server/caching"

	// TODO(crbug/1242998): Remove once safe get becomes datastore default.
	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"

	"go.chromium.org/luci/buildbucket/appengine/model"
	"go.chromium.org/luci/buildbucket/bbperms"
	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
)

func testingContext() context.Context {
	ctx := context.Background()
	ctx = memory.Use(ctx)
	ctx = caching.WithEmptyProcessCache(ctx)
	return ctx
}

func TestHasInBucketLegacy(t *testing.T) {
	t.Parallel()

	Convey("With mocked auth DB", t, func() {
		const (
			anon           = identity.AnonymousIdentity
			admin          = identity.Identity("user:admin@example.com")
			reader         = identity.Identity("user:reader@example.com")
			writer         = identity.Identity("user:writer@example.com")
			sameProject    = identity.Identity("project:some-project")
			anotherProject = identity.Identity("project:another-project")
		)

		s := &authtest.FakeState{
			FakeDB: authtest.NewFakeDB(
				authtest.MockMembership(admin, "administrators"),
				authtest.MockMembership(reader, "readers"),
				authtest.MockMembership(writer, "writers"),
			),
		}
		ctx := auth.WithState(testingContext(), s)

		makeBucket := func(acls []*pb.Acl) {
			So(datastore.Put(ctx, &model.Bucket{
				ID:     "some-bucket",
				Parent: model.ProjectKey(ctx, "some-project"),
				Proto:  &pb.Bucket{Acls: acls},
			}), ShouldBeNil)
		}

		check := func(perm realms.Permission, caller identity.Identity) codes.Code {
			s.Identity = caller
			err := HasInBucket(ctx, perm, "some-project", "some-bucket")
			if err == nil {
				return codes.OK
			}
			status, ok := appstatus.Get(err)
			if !ok {
				return codes.Internal
			}
			return status.Code()
		}

		Convey("Missing bucket", func() {
			So(check(bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
			So(check(bbperms.BuildsGet, admin), ShouldEqual, codes.NotFound)
			So(check(bbperms.BuildsGet, sameProject), ShouldEqual, codes.NotFound)
			So(check(bbperms.BuildsGet, reader), ShouldEqual, codes.NotFound)
		})

		Convey("Existing bucket, no ACLs in it", func() {
			makeBucket(nil)

			So(check(bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
			So(check(bbperms.BuildsGet, admin), ShouldEqual, codes.OK)
			So(check(bbperms.BuildsGet, sameProject), ShouldEqual, codes.OK)
			So(check(bbperms.BuildsGet, anotherProject), ShouldEqual, codes.NotFound)
			So(check(bbperms.BuildsGet, reader), ShouldEqual, codes.NotFound)
		})

		Convey("Existing bucket, with ACLs", func() {
			makeBucket([]*pb.Acl{
				{Role: pb.Acl_READER, Group: "readers"},
				{Role: pb.Acl_WRITER, Group: "writers"},
			})

			Convey("Read perm", func() {
				So(check(bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
				So(check(bbperms.BuildsGet, admin), ShouldEqual, codes.OK)
				So(check(bbperms.BuildsGet, sameProject), ShouldEqual, codes.OK)
				So(check(bbperms.BuildsGet, anotherProject), ShouldEqual, codes.NotFound)
				So(check(bbperms.BuildsGet, reader), ShouldEqual, codes.OK)
				So(check(bbperms.BuildsGet, writer), ShouldEqual, codes.OK)
			})

			Convey("Write perm", func() {
				So(check(bbperms.BuildsCancel, anon), ShouldEqual, codes.NotFound)
				So(check(bbperms.BuildsCancel, admin), ShouldEqual, codes.OK)
				So(check(bbperms.BuildsCancel, sameProject), ShouldEqual, codes.OK)
				So(check(bbperms.BuildsCancel, anotherProject), ShouldEqual, codes.NotFound)
				So(check(bbperms.BuildsCancel, reader), ShouldEqual, codes.PermissionDenied)
				So(check(bbperms.BuildsCancel, writer), ShouldEqual, codes.OK)
			})
		})
	})
}

func TestHasInBucketRealms(t *testing.T) {
	t.Parallel()

	Convey("With mocked auth DB", t, func() {
		const (
			anon         = identity.AnonymousIdentity
			admin        = identity.Identity("user:admin@example.com")
			reader       = identity.Identity("user:reader@example.com")
			writer       = identity.Identity("user:writer@example.com")
			legacyReader = identity.Identity("user:legacy-reader@example.com")
			legacyWriter = identity.Identity("user:legacy-writer@example.com")

			appID            = "buildbucket-app-id"
			projectID        = "some-project"
			existingBucketID = "existing-bucket"
			missingBucketID  = "missing-bucket"

			existingRealmID = projectID + ":" + existingBucketID
			missingRealmID  = projectID + ":" + missingBucketID
		)

		ctx := testingContext()

		So(datastore.Put(ctx, &model.Bucket{
			ID:     existingBucketID,
			Parent: model.ProjectKey(ctx, projectID),
			Proto: &pb.Bucket{
				Acls: []*pb.Acl{
					{Role: pb.Acl_READER, Group: "legacy-readers"},
					{Role: pb.Acl_WRITER, Group: "legacy-writers"},
				},
			},
		}), ShouldBeNil)

		s := &authtest.FakeState{
			FakeDB: authtest.NewFakeDB(
				authtest.MockMembership(admin, Administrators),
				authtest.MockMembership(legacyReader, "legacy-readers"),
				authtest.MockMembership(legacyWriter, "legacy-writers"),
			),
		}
		ctx = auth.WithState(ctx, s)

		check := func(bucketID string, perm realms.Permission, caller identity.Identity) codes.Code {
			s.Identity = caller
			err := HasInBucket(ctx, perm, projectID, bucketID)
			if err == nil {
				return codes.OK
			}
			status, ok := appstatus.Get(err)
			if !ok {
				return codes.Internal
			}
			return status.Code()
		}

		Convey("No realm ACLs", func() {
			So(check(existingBucketID, bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
			So(check(existingBucketID, bbperms.BuildsGet, admin), ShouldEqual, codes.OK)
			So(check(existingBucketID, bbperms.BuildsGet, reader), ShouldEqual, codes.NotFound)
			So(check(existingBucketID, bbperms.BuildsGet, writer), ShouldEqual, codes.NotFound)
			So(check(existingBucketID, bbperms.BuildsGet, legacyReader), ShouldEqual, codes.OK)
			So(check(existingBucketID, bbperms.BuildsGet, legacyWriter), ShouldEqual, codes.OK)

			So(check(missingBucketID, bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
			So(check(missingBucketID, bbperms.BuildsGet, admin), ShouldEqual, codes.NotFound)
			So(check(missingBucketID, bbperms.BuildsGet, reader), ShouldEqual, codes.NotFound)
			So(check(missingBucketID, bbperms.BuildsGet, writer), ShouldEqual, codes.NotFound)
			So(check(missingBucketID, bbperms.BuildsGet, legacyReader), ShouldEqual, codes.NotFound)
			So(check(missingBucketID, bbperms.BuildsGet, legacyWriter), ShouldEqual, codes.NotFound)
		})

		Convey("With realm ACLs", func() {
			s.FakeDB.(*authtest.FakeDB).AddMocks(
				authtest.MockPermission(reader, existingRealmID, bbperms.BuildersGet),
				authtest.MockPermission(reader, existingRealmID, bbperms.BuildsGet),
				authtest.MockPermission(writer, existingRealmID, bbperms.BuildersGet),
				authtest.MockPermission(writer, existingRealmID, bbperms.BuildsGet),
				authtest.MockPermission(writer, existingRealmID, bbperms.BuildsCancel),

				authtest.MockPermission(reader, missingRealmID, bbperms.BuildersGet),
				authtest.MockPermission(reader, missingRealmID, bbperms.BuildsGet),
				authtest.MockPermission(writer, missingRealmID, bbperms.BuildersGet),
				authtest.MockPermission(writer, missingRealmID, bbperms.BuildsGet),
				authtest.MockPermission(writer, missingRealmID, bbperms.BuildsCancel),
			)

			Convey("Read perm", func() {
				So(check(existingBucketID, bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
				So(check(existingBucketID, bbperms.BuildsGet, admin), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsGet, reader), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsGet, writer), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsGet, legacyReader), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsGet, legacyWriter), ShouldEqual, codes.OK)
			})

			Convey("Write perm", func() {
				So(check(existingBucketID, bbperms.BuildsCancel, anon), ShouldEqual, codes.NotFound)
				So(check(existingBucketID, bbperms.BuildsCancel, admin), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsCancel, reader), ShouldEqual, codes.PermissionDenied)
				So(check(existingBucketID, bbperms.BuildsCancel, writer), ShouldEqual, codes.OK)
				So(check(existingBucketID, bbperms.BuildsCancel, legacyReader), ShouldEqual, codes.PermissionDenied)
				So(check(existingBucketID, bbperms.BuildsCancel, legacyWriter), ShouldEqual, codes.OK)
			})

			Convey("Missing bucket", func() {
				So(check(missingBucketID, bbperms.BuildsGet, anon), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsGet, admin), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsGet, reader), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsGet, writer), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsGet, legacyReader), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsGet, legacyWriter), ShouldEqual, codes.NotFound)

				So(check(missingBucketID, bbperms.BuildsCancel, anon), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsCancel, admin), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsCancel, reader), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsCancel, writer), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsCancel, legacyReader), ShouldEqual, codes.NotFound)
				So(check(missingBucketID, bbperms.BuildsCancel, legacyWriter), ShouldEqual, codes.NotFound)
			})
		})
	})
}

func TestGetRole(t *testing.T) {
	t.Parallel()

	Convey("With mocked auth DB", t, func() {
		s := &authtest.FakeState{
			FakeDB: authtest.NewFakeDB(
				authtest.MockMembership("user:reader@example.com", "readers"),
				authtest.MockMembership("user:writer@example.com", "writers"),
				authtest.MockMembership("user:writer@example.com", "readers"), // also a reader
			),
		}
		ctx := auth.WithState(testingContext(), s)

		role := func(id identity.Identity, acls []*pb.Acl) pb.Acl_Role {
			r, err := getRole(ctx, id, acls)
			So(err, ShouldBeNil)
			return r
		}

		Convey("Empty", func() {
			So(role("user:someone@example.com", nil), ShouldEqual, -1)
		})

		Convey("Email", func() {
			acls := []*pb.Acl{
				{
					Identity: "some-email@example.com",
					Role:     pb.Acl_READER,
				},
			}

			So(role("user:some-email@example.com", acls), ShouldEqual, pb.Acl_READER)
			So(role("user:another-email@example.com", acls), ShouldEqual, -1)
		})

		Convey("User", func() {
			acls := []*pb.Acl{
				{
					Identity: "user:some-email@example.com",
					Role:     pb.Acl_READER,
				},
			}

			So(role("user:some-email@example.com", acls), ShouldEqual, pb.Acl_READER)
			So(role("user:another-email@example.com", acls), ShouldEqual, -1)
		})

		Convey("Group", func() {
			acls := []*pb.Acl{
				{
					Group: "readers",
					Role:  pb.Acl_READER,
				},
				{
					Group: "empty",
					Role:  pb.Acl_READER,
				},
			}

			So(role("user:reader@example.com", acls), ShouldEqual, pb.Acl_READER)
			So(role("user:unknown@example.com", acls), ShouldEqual, -1)
		})

		Convey("Highest role wins", func() {
			acls := []*pb.Acl{
				{
					Group: "readers",
					Role:  pb.Acl_READER,
				},
				{
					Group: "writers",
					Role:  pb.Acl_WRITER,
				},
			}

			So(role("user:reader@example.com", acls), ShouldEqual, pb.Acl_READER)
			So(role("user:writer@example.com", acls), ShouldEqual, pb.Acl_WRITER)
			So(role("user:unknown@example.com", acls), ShouldEqual, -1)
		})
	})
}

func TestBucketsByPerm(t *testing.T) {
	t.Parallel()

	Convey("GetAccessibleBuckets", t, func() {
		ctx := testingContext()
		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.Identity("user:user"),
		})
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)

		So(datastore.Put(ctx, &model.Bucket{
			ID:     "bucket1",
			Parent: model.ProjectKey(ctx, "project"),
			Proto: &pb.Bucket{
				Acls: []*pb.Acl{
					{
						Identity: "user:user",
						Role:     pb.Acl_READER,
					},
				},
			},
		}), ShouldBeNil)
		So(datastore.Put(ctx, &model.Bucket{
			ID:     "bucket1",
			Parent: model.ProjectKey(ctx, "project2"),
			Proto: &pb.Bucket{
				Acls: []*pb.Acl{
					{
						Identity: "user:user",
						Role:     pb.Acl_READER,
					},
				},
			},
		}), ShouldBeNil)
		So(datastore.Put(ctx, &model.Bucket{
			ID:     "bucket2",
			Parent: model.ProjectKey(ctx, "project"),
			Proto: &pb.Bucket{
				Acls: []*pb.Acl{
					{
						Identity: "user:user",
						Role:     pb.Acl_WRITER,
					},
				},
			},
		}), ShouldBeNil)

		buckets1, err := BucketsByPerm(ctx, bbperms.BuildersList, "")
		So(err, ShouldBeNil)
		So(buckets1, ShouldResemble, []string{"project/bucket1", "project/bucket2", "project2/bucket1"})

		buckets2, err := BucketsByPerm(ctx, bbperms.BuildsCancel, "")
		So(err, ShouldBeNil)
		So(buckets2, ShouldResemble, []string{"project/bucket2"})

		buckets3, err := BucketsByPerm(ctx, bbperms.BuildersList, "project2")
		So(err, ShouldBeNil)
		So(buckets3, ShouldResemble, []string{"project2/bucket1"})

		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.Identity("user:no_any_permission"),
		})
		buckets4, err := BucketsByPerm(ctx, bbperms.BuildersList, "")
		So(err, ShouldBeNil)
		So(buckets4, ShouldBeNil)
	})
}

func TestCanUpdateBuild(t *testing.T) {
	t.Parallel()

	Convey("With mocked auth DB", t, func() {
		member := identity.Identity("member@example.com")
		notMember := identity.Identity("not-member@example.com")
		s := &authtest.FakeState{
			FakeDB: authtest.NewFakeDB(
				authtest.MockMembership(member, UpdateBuildAllowedUsers),
			),
		}
		ctx := auth.WithState(testingContext(), s)

		Convey("With a member of the updater group", func() {
			s.Identity = member
			can, err := CanUpdateBuild(ctx)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)
		})

		Convey("With a non-member of the updater group", func() {
			s.Identity = notMember
			can, err := CanUpdateBuild(ctx)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})
	})
}
