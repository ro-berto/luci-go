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

	"github.com/golang/protobuf/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"google.golang.org/grpc/codes"
)

const allowGroup = "luci-analysis-access"

// Checks if this call is allowed, returns an error if it is.
func checkAllowedPrelude(ctx context.Context, methodName string, req proto.Message) (context.Context, error) {
	if err := checkAllowed(ctx); err != nil {
		return ctx, err
	}
	return ctx, nil
}

// Logs and converts the errors to GRPC type errors.
func gRPCifyAndLogPostlude(ctx context.Context, methodName string, rsp proto.Message, err error) error {
	return appstatus.GRPCifyAndLog(ctx, err)
}

func checkAllowed(ctx context.Context) error {
	switch yes, err := auth.IsMember(ctx, allowGroup); {
	case err != nil:
		return errors.Annotate(err, "failed to check ACL").Err()
	case !yes:
		return appstatus.Errorf(codes.PermissionDenied, "not a member of %s", allowGroup)
	default:
		return nil
	}
}
