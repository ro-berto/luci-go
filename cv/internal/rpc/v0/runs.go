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

package rpc

import (
	"context"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"

	apiv0pb "go.chromium.org/luci/cv/api/v0"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/rpc/versioning"
	"go.chromium.org/luci/cv/internal/run"
)

const allowGroup = "service-luci-change-verifier-v0-api-users"

// checkCanUseAPI ensures that calling user is granted permission to use
// unstable v0 API.
func checkCanUseAPI(ctx context.Context, name string) error {
	switch yes, err := auth.IsMember(ctx, allowGroup); {
	case err != nil:
		return appstatus.Errorf(codes.Internal, "failed to check ACL")
	case !yes:
		return appstatus.Errorf(codes.PermissionDenied, "not a member of %s", allowGroup)
	default:
		logging.Warningf(ctx, "%s is calling %s", auth.CurrentIdentity(ctx), name)
		return nil
	}
}

// RunsServer implements rpc v0 APIs.
type RunsServer struct {
	apiv0pb.UnimplementedRunsServer
}

// GetRun implements apiv0pb.RunsServer.
func (s *RunsServer) GetRun(ctx context.Context, req *apiv0pb.GetRunRequest) (resp *apiv0pb.Run, err error) {
	defer func() { err = appstatus.GRPCifyAndLog(ctx, err) }()
	if err = checkCanUseAPI(ctx, "Runs.GetRun"); err != nil {
		return
	}

	id, err := toInternalRunID(req.GetId())
	if err != nil {
		return nil, err
	}

	r := run.Run{ID: id}
	switch err := datastore.Get(ctx, &r); {
	case err == datastore.ErrNoSuchEntity:
		return nil, appstatus.Errorf(codes.NotFound, "run not found")
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch Run").Tag(transient.Tag).Err()
	}
	if err := checkCanReadRun(ctx, &r); err != nil {
		return nil, err
	}

	rcls, err := run.LoadRunCLs(ctx, r.ID, r.CLs)
	if err != nil {
		return nil, err
	}
	gcls := make([]*apiv0pb.GerritChange, len(rcls))
	sCLSet := common.MakeCLIDsSet(r.Submission.GetSubmittedCls()...)
	fCLSet := common.MakeCLIDsSet(r.Submission.GetFailedCls()...)
	sCLIndexes := make([]int32, 0, len(fCLSet))
	fCLIndexes := make([]int32, 0, len(sCLSet))

	for i, rcl := range rcls {
		host, change, err := rcl.ExternalID.ParseGobID()
		switch {
		case err != nil:
			// As of Sep 2, 2021, CV works only with Gerrit (GoB) CL.
			panic(errors.Annotate(err, "ParseGobID").Err())
		case sCLSet.Has(rcl.ID):
			sCLIndexes = append(sCLIndexes, int32(i))
		case fCLSet.Has(rcl.ID):
			fCLIndexes = append(fCLIndexes, int32(i))
		}
		gcls[i] = &apiv0pb.GerritChange{
			Host:     host,
			Change:   change,
			Patchset: rcl.Detail.GetPatchset(),
		}
	}

	tryjobs := make([]*apiv0pb.Tryjob, len(r.Tryjobs.GetTryjobs()))
	for i, tj := range r.Tryjobs.GetTryjobs() {
		tryjobs[i] = &apiv0pb.Tryjob{
			Status: versioning.TryjobStatusV0(tj.Status),
		}
		// result
		if result := tj.GetResult(); result != nil {
			tryjobs[i].Result = &apiv0pb.Tryjob_Result{
				Status: versioning.TryjobResultStatusV0(result.Status),
			}
			if bb := result.GetBuildbucket(); bb != nil {
				tryjobs[i].Result.Backend = &apiv0pb.Tryjob_Result_Buildbucket_{
					Buildbucket: &apiv0pb.Tryjob_Result_Buildbucket{
						Id:     bb.Id,
						Status: bb.Status,
					},
				}
			}
		}
		// definition
		if bb := tj.GetDefinition().GetBuildbucket(); bb != nil {
			tryjobs[i].Definition = &apiv0pb.Tryjob_Definition{
				Backend: &apiv0pb.Tryjob_Definition_Buildbucket_{
					Buildbucket: &apiv0pb.Tryjob_Definition_Buildbucket{
						Host:    bb.Host,
						Builder: bb.Builder,
					},
				},
			}
		}
	}

	var submission *apiv0pb.Run_Submission
	if len(sCLIndexes) > 0 || len(fCLIndexes) > 0 {
		submission = &apiv0pb.Run_Submission{
			SubmittedClIndexes: sCLIndexes,
			FailedClIndexes:    fCLIndexes,
		}
	}

	return &apiv0pb.Run{
		Id:         r.ID.PublicID(),
		Eversion:   int64(r.EVersion),
		Status:     versioning.RunStatusV0(r.Status),
		Mode:       string(r.Mode),
		CreateTime: common.TspbNillable(r.CreateTime),
		StartTime:  common.TspbNillable(r.StartTime),
		UpdateTime: common.TspbNillable(r.UpdateTime),
		EndTime:    common.TspbNillable(r.EndTime),
		Owner:      string(r.Owner),
		Cls:        gcls,
		Tryjobs:    tryjobs,
		Submission: submission,
	}, nil
}

func checkCanReadRun(ctx context.Context, r *run.Run) error {
	// TODO(crbug/1233963): implement.
	return nil
}

func toInternalRunID(id string) (common.RunID, error) {
	if id == "" {
		return "", appstatus.Errorf(codes.InvalidArgument, "Run ID is required")
	}
	internalID, err := common.FromPublicRunID(id)
	if err != nil {
		return "", appstatus.Errorf(codes.InvalidArgument, err.Error())
	}
	return internalID, nil
}
