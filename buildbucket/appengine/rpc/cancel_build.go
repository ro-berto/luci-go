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

package rpc

import (
	"context"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/buildbucket/appengine/internal/perm"
	"go.chromium.org/luci/buildbucket/appengine/model"
	"go.chromium.org/luci/buildbucket/appengine/tasks"
	pb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/buildbucket/protoutil"
)

// validateCancel validates the given request.
func validateCancel(req *pb.CancelBuildRequest) error {
	var err error
	switch {
	case req.GetId() == 0:
		return errors.Reason("id is required").Err()
	case req.SummaryMarkdown == "":
		return errors.Reason("summary_markdown is required").Err()
	case teeErr(validateSummaryMarkdown(req.SummaryMarkdown), &err) != nil:
		return errors.Annotate(err, "summary_markdown").Err()
	}
	return nil
}

// CancelBuild handles a request to cancel a build. Implements pb.BuildsServer.
func (*Builds) CancelBuild(ctx context.Context, req *pb.CancelBuildRequest) (*pb.Build, error) {
	if err := validateCancel(req); err != nil {
		return nil, appstatus.BadRequest(err)
	}
	m, err := model.NewBuildMask("", req.Fields, req.Mask)
	if err != nil {
		return nil, appstatus.Errorf(codes.InvalidArgument, "invalid mask")
	}

	bld, err := getBuild(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if err := perm.HasInBuilder(ctx, perm.BuildsCancel, bld.Proto.Builder); err != nil {
		return nil, err
	}

	if protoutil.IsEnded(bld.Proto.Status) {
		return bld.ToProto(ctx, m)
	}

	bld, err = tasks.StartCancel(ctx, req.Id, req.SummaryMarkdown)
	if err != nil {
		return nil, err
	}
	return bld.ToProto(ctx, m)
}
