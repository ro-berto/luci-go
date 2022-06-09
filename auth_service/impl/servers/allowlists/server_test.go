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

package allowlists

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/auth_service/api/rpcpb"
	"go.chromium.org/luci/auth_service/impl/model"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestAllowlistsServer(t *testing.T) {
	t.Parallel()
	srv := Server{}
	createdTime := time.Date(2021, time.September, 16, 15, 20, 0, 0, time.UTC)

	Convey("GetAllowlist RPC call", t, func() {
		ctx := memory.Use(context.Background())

		request := &rpcpb.GetAllowlistRequest{
			Name: "test-allowlist",
		}

		_, err := srv.GetAllowlist(ctx, request)
		So(err, ShouldHaveGRPCStatus, codes.NotFound)

		// Allowlist built from model.AuthIPAllowlist definition.
		So(datastore.Put(ctx,
			&model.AuthIPAllowlist{
				AuthVersionedEntityMixin: model.AuthVersionedEntityMixin{},
				Parent:                   model.RootKey(ctx),
				ID:                       "test-allowlist",
				Subnets: []string{
					"127.0.0.1/24",
					"127.0.0.127/24",
				},
				Description: "This is a test allowlist.",
				CreatedTS:   createdTime,
				CreatedBy:   "user:test-user-1",
			}), ShouldBeNil)

		expectedResponse := &rpcpb.Allowlist{
			Name: "test-allowlist",
			Subnets: []string{
				"127.0.0.1/24",
				"127.0.0.127/24",
			},
			Description: "This is a test allowlist.",
			CreatedTs:   timestamppb.New(createdTime),
			CreatedBy:   "user:test-user-1",
		}

		actualResponse, err := srv.GetAllowlist(ctx, request)
		So(err, ShouldBeNil)
		So(actualResponse, ShouldResembleProto, expectedResponse)
	})

	Convey("ListAllowlists RPC call", t, func() {
		ctx := memory.Use(context.Background())

		So(datastore.Put(ctx,
			&model.AuthIPAllowlist{
				AuthVersionedEntityMixin: model.AuthVersionedEntityMixin{},
				ID:                       "z-test-allowlist",
				Parent:                   model.RootKey(ctx),
				Subnets: []string{
					"127.0.0.1/24",
					"127.0.0.127/24",
				},
				Description: "This is a test allowlist, should show up last.",
				CreatedTS:   createdTime,
				CreatedBy:   "user:test-user-2",
			},
			&model.AuthIPAllowlist{
				AuthVersionedEntityMixin: model.AuthVersionedEntityMixin{},
				ID:                       "a-test-allowlist",
				Parent:                   model.RootKey(ctx),
				Subnets: []string{
					"0.0.0.0/0",
				},
				Description: "This is a test allowlist, should show up first.",
				CreatedTS:   createdTime,
				CreatedBy:   "user:test-user-1",
			},
			&model.AuthIPAllowlist{
				AuthVersionedEntityMixin: model.AuthVersionedEntityMixin{},
				ID:                       "test-allowlist",
				Parent:                   model.RootKey(ctx),
				Subnets:                  []string{},
				Description:              "This is a test allowlist, should show up second.",
				CreatedTS:                createdTime,
				CreatedBy:                "user:test-user-3",
			}), ShouldBeNil)

		// Expected response, build with pb.
		expectedAllowlists := &rpcpb.ListAllowlistsResponse{
			Allowlists: []*rpcpb.Allowlist{
				{
					Name: "a-test-allowlist",
					Subnets: []string{
						"0.0.0.0/0",
					},
					Description: "This is a test allowlist, should show up first.",
					CreatedTs:   timestamppb.New(createdTime),
					CreatedBy:   "user:test-user-1",
				},
				{
					Name:        "test-allowlist",
					Subnets:     []string{},
					Description: "This is a test allowlist, should show up second.",
					CreatedTs:   timestamppb.New(createdTime),
					CreatedBy:   "user:test-user-3",
				},
				{
					Name: "z-test-allowlist",
					Subnets: []string{
						"127.0.0.1/24",
						"127.0.0.127/24",
					},
					Description: "This is a test allowlist, should show up last.",
					CreatedTs:   timestamppb.New(createdTime),
					CreatedBy:   "user:test-user-2",
				},
			},
		}

		actualResponse, err := srv.ListAllowlists(ctx, &emptypb.Empty{})
		So(err, ShouldBeNil)
		So(expectedAllowlists.Allowlists, ShouldResembleProto, actualResponse.Allowlists)
	})
}
