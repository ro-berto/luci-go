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

package rpc

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/secrets"
	"go.chromium.org/luci/server/secrets/testsecrets"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"

	"go.chromium.org/luci/analysis/internal/config"
	"go.chromium.org/luci/analysis/internal/perms"
	configpb "go.chromium.org/luci/analysis/proto/config"
	pb "go.chromium.org/luci/analysis/proto/v1"
)

func TestProjects(t *testing.T) {
	Convey("Given a projects server", t, func() {
		ctx := context.Background()

		// For user identification.
		ctx = authtest.MockAuthConfig(ctx)
		authState := &authtest.FakeState{
			Identity:       "user:someone@example.com",
			IdentityGroups: []string{"weetbix-access"},
		}
		ctx = auth.WithState(ctx, authState)
		ctx = secrets.Use(ctx, &testsecrets.Store{})

		// Provides datastore implementation needed for project config.
		ctx = memory.Use(ctx)
		server := NewProjectsServer()

		Convey("Unauthorised requests are rejected", func() {
			ctx = auth.WithState(ctx, &authtest.FakeState{
				Identity: "user:someone@example.com",
				// Not a member of weetbix-access.
				IdentityGroups: []string{"other-group"},
			})

			// Make some request (the request should not matter, as
			// a common decorator is used for all requests.)
			request := &pb.ListProjectsRequest{}

			rule, err := server.List(ctx, request)
			st, _ := grpcStatus.FromError(err)
			So(st.Code(), ShouldEqual, codes.PermissionDenied)
			So(st.Message(), ShouldEqual, "not a member of weetbix-access")
			So(rule, ShouldBeNil)
		})
		Convey("GetConfig", func() {
			authState.IdentityPermissions = []authtest.RealmPermission{
				{
					Realm:      "testproject:@root",
					Permission: perms.PermGetConfig,
				},
			}

			// Setup.
			configs := make(map[string]*configpb.ProjectConfig)
			projectTest := config.CreatePlaceholderProjectConfig()
			projectTest.Monorail.Project = "monorailproject"
			projectTest.Monorail.DisplayPrefix = "displayprefix.com"
			configs["testproject"] = projectTest
			config.SetTestProjectConfig(ctx, configs)

			request := &pb.GetProjectConfigRequest{
				Name: "projects/testproject/config",
			}

			Convey("No permission to get project config", func() {
				authState.IdentityPermissions = removePermission(authState.IdentityPermissions, perms.PermGetConfig)

				response, err := server.GetConfig(ctx, request)
				So(err, ShouldBeRPCPermissionDenied, "caller does not have permission analysis.config.get")
				So(response, ShouldBeNil)
			})
			Convey("Valid request", func() {
				response, err := server.GetConfig(ctx, request)
				So(err, ShouldBeNil)
				So(response, ShouldResembleProto, &pb.ProjectConfig{
					Name: "projects/testproject/config",
					Monorail: &pb.ProjectConfig_Monorail{
						Project:       "monorailproject",
						DisplayPrefix: "displayprefix.com",
					},
				})
			})
			Convey("Invalid request", func() {
				request.Name = "blah"

				// Run
				response, err := server.GetConfig(ctx, request)

				// Verify
				So(err, ShouldBeRPCInvalidArgument, "name: invalid project config name, expected format: projects/{project}/config")
				So(response, ShouldBeNil)
			})
			Convey("With project not configured", func() {
				err := config.SetTestProjectConfig(ctx, map[string]*configpb.ProjectConfig{})
				So(err, ShouldBeNil)

				// Run
				response, err := server.GetConfig(ctx, request)

				// Verify
				So(err, ShouldBeRPCFailedPrecondition, "project does not exist in LUCI Analysis")
				So(response, ShouldBeNil)
			})
		})
		Convey("List", func() {
			authState.IdentityPermissions = []authtest.RealmPermission{
				{
					Realm:      "chromium:@root",
					Permission: perms.PermGetConfig,
				},
				{
					Realm:      "chrome:@root",
					Permission: perms.PermGetConfig,
				},
			}

			// Setup
			projectChromium := config.CreatePlaceholderProjectConfig()
			projectChrome := config.CreatePlaceholderProjectConfig()
			projectSecret := config.CreatePlaceholderProjectConfig()

			configs := make(map[string]*configpb.ProjectConfig)
			configs["chromium"] = projectChromium
			configs["chrome"] = projectChrome
			configs["secret"] = projectSecret
			config.SetTestProjectConfig(ctx, configs)

			request := &pb.ListProjectsRequest{}

			Convey("No permission to view any project", func() {
				authState.IdentityPermissions = removePermission(authState.IdentityPermissions, perms.PermGetConfig)

				// Run
				projectsResponse, err := server.List(ctx, request)

				// Verify
				So(err, ShouldBeNil)
				expected := &pb.ListProjectsResponse{Projects: []*pb.Project{}}
				So(projectsResponse, ShouldResembleProto, expected)
			})
			Convey("Valid request", func() {
				// Run
				projectsResponse, err := server.List(ctx, request)

				// Verify
				So(err, ShouldBeNil)
				expected := &pb.ListProjectsResponse{Projects: []*pb.Project{
					{
						Name:        "projects/chrome",
						DisplayName: "Chrome",
						Project:     "chrome",
					},
					{
						Name:        "projects/chromium",
						DisplayName: "Chromium",
						Project:     "chromium",
					},
				}}
				So(projectsResponse, ShouldResembleProto, expected)
			})
		})
	})
}
