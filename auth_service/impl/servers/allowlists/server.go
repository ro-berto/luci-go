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

// Package allowlists contains Allowlists server implementation.
package allowlists

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/auth_service/api/rpcpb"
	"go.chromium.org/luci/auth_service/impl/model"
	"go.chromium.org/luci/gae/service/datastore"
)

// Server implements Allowlists server.
type Server struct {
	rpcpb.UnimplementedAllowlistsServer
}

// GetAllowlist implements the corresponding RPC method.
func (*Server) GetAllowlist(ctx context.Context, request *rpcpb.GetAllowlistRequest) (*rpcpb.Allowlist, error) {
	switch allowlist, err := model.GetAuthIPAllowlist(ctx, request.Name); {
	case err == nil:
		return allowlist.ToProto(), nil
	case err == datastore.ErrNoSuchEntity:
		return nil, status.Errorf(codes.NotFound, "no such allowlist %q", request.Name)
	default:
		return nil, status.Errorf(codes.Internal, "failed to fetch allowlist %q: %s", request.Name, err)
	}
}
