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

package rpcs

import (
	"context"
	"sort"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/deploy/api/modelpb"
	"go.chromium.org/luci/deploy/api/rpcpb"
	"go.chromium.org/luci/deploy/service/model"
)

// Assets is the implementation of deploy.service.Assets service.
type Assets struct {
	rpcpb.UnimplementedAssetsServer
}

// GetAsset implements the corresponding RPC method.
func (*Assets) GetAsset(ctx context.Context, req *rpcpb.GetAssetRequest) (resp *modelpb.Asset, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(ctx, err) }()

	entity := model.Asset{ID: req.AssetId}
	switch err := datastore.Get(ctx, &entity); {
	case err == datastore.ErrNoSuchEntity:
		return nil, status.Errorf(codes.NotFound, "no such asset")
	case err != nil:
		return nil, status.Errorf(codes.Internal, "datastore error when fetching the asset: %s", err)
	default:
		return checkAssetEntity(&entity)
	}
}

// ListAssets implements the corresponding RPC method.
func (*Assets) ListAssets(ctx context.Context, req *rpcpb.ListAssetsRequest) (resp *rpcpb.ListAssetsResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(ctx, err) }()

	q := datastore.NewQuery("Asset")

	var entities []*model.Asset
	if err = datastore.GetAll(ctx, q, &entities); err != nil {
		return nil, status.Errorf(codes.Internal, "datastore query to list assets failed: %s", err)
	}

	assets, err := sortedProtoList(entities)
	if err != nil {
		return nil, err
	}

	return &rpcpb.ListAssetsResponse{Assets: assets}, nil
}

// checkAssetEntity checks the proto payload of the asset entity is correct.
//
// Returns gRPC errors.
func checkAssetEntity(e *model.Asset) (*modelpb.Asset, error) {
	if e.Asset.GetId() != e.ID {
		return nil, status.Errorf(codes.Internal, "asset entity %q has bad proto payload %v", e.ID, e.Asset)
	}
	return e.Asset, nil
}

// sortedProtoList extracts Asset protos and sorts them by ID.
//
// Returns gRPC errors.
func sortedProtoList(entities []*model.Asset) ([]*modelpb.Asset, error) {
	out := make([]*modelpb.Asset, len(entities))
	for i, ent := range entities {
		var err error
		if out[i], err = checkAssetEntity(ent); err != nil {
			return nil, err
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Id < out[j].Id
	})
	return out, nil
}
