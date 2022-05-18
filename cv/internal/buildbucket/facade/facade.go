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

package bbfacade

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/proto/structmask"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/cv/internal/buildbucket"
	"go.chromium.org/luci/cv/internal/tryjob"
)

// Facade provides APIs that LUCI CV can use to interact with buildbucket
// tryjobs.
type Facade struct {
	ClientFactory buildbucket.ClientFactory
}

func (f *Facade) Kind() string {
	return "buildbucket"
}

var defaultMask *bbpb.BuildMask

func init() {
	defaultMask = &bbpb.BuildMask{
		Fields: &fieldmaskpb.FieldMask{
			Paths: []string{
				"create_time",
				"id",
				"output.properties",
				"status",
				"status_details",
				"update_time",
			},
		},
	}
	for _, key := range outputPropKeys {
		defaultMask.OutputProperties = append(defaultMask.OutputProperties, &structmask.StructMask{
			Path: []string{key},
		})
	}
}

// Update retrieves the Buildbucket build corresponding to the given Tryjob,
// parses its output and returns its current Status and Result.
//
// It does not modify the given Tryjob.
func (f *Facade) Update(ctx context.Context, saved *tryjob.Tryjob) (tryjob.Status, *tryjob.Result, error) {
	host, buildID, err := saved.ExternalID.ParseBuildbucketID()
	if err != nil {
		return 0, nil, err
	}

	bbClient, err := f.ClientFactory.MakeClient(ctx, host, saved.LUCIProject())
	if err != nil {
		return 0, nil, err
	}

	build, err := bbClient.GetBuild(ctx, &bbpb.GetBuildRequest{Id: buildID, Mask: defaultMask})
	switch code := status.Code(err); {
	case code == codes.OK:
		return parseStatusAndResult(ctx, build)
	case grpcutil.IsTransientCode(code) || code == codes.DeadlineExceeded:
		return 0, nil, transient.Tag.Apply(err)
	default:
		return 0, nil, err
	}
}

// CancelTryjob asks buildbucket to cancel a running tryjob.
//
// It returns nil error if the buildbucket build is ended.
func (f *Facade) CancelTryjob(ctx context.Context, tj *tryjob.Tryjob) error {
	host, buildID, err := tj.ExternalID.ParseBuildbucketID()
	if err != nil {
		return err
	}

	bbClient, err := f.ClientFactory.MakeClient(ctx, host, tj.LUCIProject())
	if err != nil {
		return err
	}

	_, err = bbClient.CancelBuild(ctx, &bbpb.CancelBuildRequest{
		Id: buildID,
	})
	return err
}
