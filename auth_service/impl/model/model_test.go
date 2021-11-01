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
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"
)

func TestGetAuthGroup(t *testing.T) {
	t.Parallel()
	Convey("Testing GetAuthGroup", t, func() {
		ctx := memory.Use(context.Background())

		authVersionedEntityMixin := &AuthVersionedEntityMixin{
			ModifiedTS:    time.Date(2021, time.August, 16, 12, 20, 0, 0, time.UTC),
			ModifiedBy:    "test-account",
			AuthDBRev:     1337,
			AuthDBPrevRev: 0,
		}

		authGroup := &AuthGroup{
			ID:                       "test-auth-group-1",
			Parent:                   RootKey(ctx),
			AuthVersionedEntityMixin: *authVersionedEntityMixin,
			Members: []string{
				"member1",
				"member2",
			},
			Globs: []string{
				"member1@example.com",
				"member2@example.com",
			},
			Nested: []string{
				"test/group-2",
			},
			Description: "This is a test auth group.",
			Owners:      "test-auth-group",
			CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
			CreatedBy:   "test-user",
		}

		_, err := GetAuthGroup(ctx, "test-auth-group-1")
		So(err, ShouldEqual, datastore.ErrNoSuchEntity)

		err = datastore.Put(ctx, authGroup)
		So(err, ShouldBeNil)

		actual, err := GetAuthGroup(ctx, "test-auth-group-1")
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, authGroup)
	})
}

func TestGetAllAuthGroups(t *testing.T) {
	t.Parallel()
	Convey("Testing GetAllAuthGroups", t, func() {
		ctx := memory.Use(context.Background())

		authVersionedEntityMixin := &AuthVersionedEntityMixin{
			ModifiedTS:    time.Date(2021, time.August, 16, 12, 20, 0, 0, time.UTC),
			ModifiedBy:    "test-account",
			AuthDBRev:     1337,
			AuthDBPrevRev: 0,
		}

		expectedAuthGroups := []*AuthGroup{
			{
				Kind:                     "AuthGroup",
				ID:                       "test-auth-group-1",
				Parent:                   RootKey(ctx),
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Members: []string{
					"member1",
					"member2",
				},
				Globs: []string{
					"member1@example.com",
					"member2@example.com",
				},
				Nested: []string{
					"test/group-2",
				},
				Description: "This is a test auth group.",
				Owners:      "test-auth-group",
				CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
				CreatedBy:   "test-user",
			}, {
				Kind:                     "AuthGroup",
				ID:                       "test-auth-group-2",
				Parent:                   RootKey(ctx),
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Members: []string{
					"member3",
				},
				Globs: []string{
					"member3@example.com",
				},
				Nested: []string{
					"test/group-2",
				},
				Description: "This is another test auth group.",
				Owners:      "test-auth-group-1",
				CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
				CreatedBy:   "test-user",
			}, {
				Kind:                     "AuthGroup",
				ID:                       "test-auth-group-3",
				Parent:                   RootKey(ctx),
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Members: []string{
					"member4",
					"member5",
					"member6",
				},
				Globs: []string{
					"member@example.com",
				},
				Nested: []string{
					"auth/group-3",
				},
				Description: "Yet another test auth group.",
				Owners:      "test-auth-group-1",
				CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
				CreatedBy:   "test-user",
			}, {
				Kind:                     "AuthGroup",
				ID:                       "test-auth-group-4",
				Parent:                   RootKey(ctx),
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Members: []string{
					"member4",
					"member5",
					"member6",
				},
				Globs: []string{
					"member@example.com",
					"member5@example.com",
				},
				Nested: []string{
					"test-auth-group-3",
				},
				Description: "Group to test AuthGroup.",
				Owners:      "test-auth-group-2",
				CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
				CreatedBy:   "test-user",
			},
		}

		err := datastore.Put(ctx, expectedAuthGroups)
		So(err, ShouldBeNil)

		actualAuthGroups, err := GetAllAuthGroups(ctx)
		So(err, ShouldBeNil)
		So(actualAuthGroups, ShouldResemble, expectedAuthGroups)
	})
}

func TestGetAuthIPAllowlist(t *testing.T) {
	t.Parallel()
	Convey("Testing GetAuthIPAllowlist", t, func() {
		ctx := memory.Use(context.Background())

		authVersionedEntityMixin := &AuthVersionedEntityMixin{
			ModifiedTS:    time.Date(2021, time.August, 16, 12, 20, 0, 0, time.UTC),
			ModifiedBy:    "test-account",
			AuthDBRev:     1337,
			AuthDBPrevRev: 0,
		}

		authIPAllowlist := &AuthIPAllowlist{
			ID:                       "test-auth-ip-allowlist-1",
			Parent:                   RootKey(ctx),
			AuthVersionedEntityMixin: *authVersionedEntityMixin,
			Subnets: []string{
				"123.456.789.101/24",
				"123.456.789.112/24",
			},
			Description: "This is a test AuthIPAllowlist!",
			CreatedTS:   time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC),
			CreatedBy:   "test-user",
		}

		_, err := GetAuthIPAllowlist(ctx, "test-auth-ip-allowlist-1")
		So(err, ShouldEqual, datastore.ErrNoSuchEntity)

		err = datastore.Put(ctx, authIPAllowlist)
		So(err, ShouldBeNil)

		actual, err := GetAuthIPAllowlist(ctx, "test-auth-ip-allowlist-1")
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, authIPAllowlist)
	})
}

func TestGetAllAuthIPAllowlists(t *testing.T) {
	t.Parallel()
	Convey("Testing GetAllAuthIPAllowlists", t, func() {
		ctx := memory.Use(context.Background())
		createdTime := time.Date(2021, time.August, 16, 15, 20, 0, 0, time.UTC)

		authVersionedEntityMixin := &AuthVersionedEntityMixin{
			ModifiedTS:    time.Date(2021, time.August, 16, 12, 20, 0, 0, time.UTC),
			ModifiedBy:    "test-account",
			AuthDBRev:     1337,
			AuthDBPrevRev: 0,
		}

		// Out of order alphabetically by ID.
		unorderedAllowlists := []*AuthIPAllowlist{
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "test-allowlist-1",
				Parent:                   RootKey(ctx),
				Subnets:                  []string(nil),
				Description:              "This is a test allowlist.",
				CreatedTS:                createdTime,
				CreatedBy:                "test-user",
			},
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "abc-allowlist",
				Parent:                   RootKey(ctx),
				Subnets: []string{
					"0.0.0.0/24",
					"0.0.0.128/24",
				},
				Description: "This is a test allowlist.",
				CreatedTS:   createdTime,
				CreatedBy:   "test-user-1",
			},
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "bcd-allowlist",
				Parent:                   RootKey(ctx),
				Subnets: []string{
					"127.0.0.1/24",
				},
				Description: "This is a test allowlist.",
				CreatedTS:   createdTime,
				CreatedBy:   "test-user-2",
			},
		}

		// Expected response is ordered alphabetically by ID.
		expectedAllowlists := []*AuthIPAllowlist{
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "abc-allowlist",
				Parent:                   RootKey(ctx),
				Subnets: []string{
					"0.0.0.0/24",
					"0.0.0.128/24",
				},
				Description: "This is a test allowlist.",
				CreatedTS:   createdTime,
				CreatedBy:   "test-user-1",
			},
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "bcd-allowlist",
				Parent:                   RootKey(ctx),
				Subnets: []string{
					"127.0.0.1/24",
				},
				Description: "This is a test allowlist.",
				CreatedTS:   createdTime,
				CreatedBy:   "test-user-2",
			},
			{
				AuthVersionedEntityMixin: *authVersionedEntityMixin,
				Kind:                     "AuthIPWhitelist",
				ID:                       "test-allowlist-1",
				Parent:                   RootKey(ctx),
				Subnets:                  []string(nil),
				Description:              "This is a test allowlist.",
				CreatedTS:                createdTime,
				CreatedBy:                "test-user",
			},
		}

		err := datastore.Put(ctx, unorderedAllowlists)
		So(err, ShouldBeNil)

		actualAuthIPAllowlists, err := GetAllAuthIPAllowlists(ctx)
		So(err, ShouldBeNil)
		So(actualAuthIPAllowlists, ShouldResemble, expectedAllowlists)
	})
}

func TestGetAuthGlobalConfig(t *testing.T) {
	t.Parallel()
	Convey("Testing GetAuthGlobalConfig", t, func() {
		ctx := memory.Use(context.Background())

		authVersionedEntityMixin := &AuthVersionedEntityMixin{
			ModifiedTS:    time.Date(2021, time.August, 16, 12, 20, 0, 0, time.UTC),
			ModifiedBy:    "test-account",
			AuthDBRev:     1337,
			AuthDBPrevRev: 0,
		}

		authGlobalConfig := &AuthGlobalConfig{
			AuthVersionedEntityMixin: *authVersionedEntityMixin,
			OAuthAdditionalClientIDs: []string{"add-client-id-1", "add-client-id-2"},
			OAuthClientID:            "test.client.id",
			OAuthClientSecret:        "test.client.secret",
			TokenServerURL:           "test.token.server.url",
			SecurityConfig:           []byte{1, 2, 3, 4},
		}

		_, err := GetAuthGlobalConfig(ctx)
		So(err, ShouldEqual, datastore.ErrNoSuchEntity)

		err = datastore.Put(ctx, authGlobalConfig)
		So(err, ShouldBeNil)

		actual, err := GetAuthGlobalConfig(ctx)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, authGlobalConfig)
	})
}
