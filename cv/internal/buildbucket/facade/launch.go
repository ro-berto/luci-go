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

package bbfacade

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/sync/parallel"

	"go.chromium.org/luci/cv/api/recipe/v1"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/tryjob"
)

const propertyKey = "$recipe_engine/cq"

// Launch schedules requested Tryjobs in Buildbucket.
//
// The Tryjobs will include relevant info from the Run (e.g. Run mode) and
// involves all provided CLs.
//
// Updates the Tryjobs that are scheduled successfully in Buildbucket in place.
// The following fields will be updated:
//  * ExternalID
//  * Status
//  * Result
//
// Returns nil if all tryjobs have been successfully launched. Otherwise,
// returns `errors.MultiError` where each element is the launch error of
// the corresponding Tryjob.
//
// Uses Tryjob ID as the request key for deduplication. This ensures only one
// Buildbucket build will be scheduled for one Tryjob within the deduplication
// window (currently 1 min. in Buildbucket).
func (f *Facade) Launch(ctx context.Context, tryjobs []*tryjob.Tryjob, r *run.Run, cls []*run.RunCL) error {
	tryjobsByHost := splitTryjobsByHost(tryjobs)
	tryjobToIndex := make(map[*tryjob.Tryjob]int, len(tryjobs))
	for i, tj := range tryjobs {
		tryjobToIndex[tj] = i
	}
	launchErrs := errors.NewLazyMultiError(len(tryjobs))
	poolErr := parallel.WorkPool(min(len(tryjobsByHost), 8), func(work chan<- func() error) {
		for host, tryjobs := range tryjobsByHost {
			host, tryjobs := host, tryjobs
			work <- func() error {
				err := f.schedule(ctx, host, r, cls, tryjobs)
				switch merrs, ok := err.(errors.MultiError); {
				case err == nil:
				case !ok:
					// assign singular error to all tryjobs.
					for _, tj := range tryjobs {
						launchErrs.Assign(tryjobToIndex[tj], err)
					}
				default:
					for i, tj := range tryjobs {
						launchErrs.Assign(tryjobToIndex[tj], merrs[i])
					}
				}
				return nil
			}
		}
	})
	if poolErr != nil {
		panic(fmt.Errorf("impossible"))
	}
	return launchErrs.Get()
}

func splitTryjobsByHost(tryjobs []*tryjob.Tryjob) map[string][]*tryjob.Tryjob {
	ret := make(map[string][]*tryjob.Tryjob, 2) // normally, at most 2 host.
	for _, tj := range tryjobs {
		bbDef := tj.Definition.GetBuildbucket()
		if bbDef == nil {
			panic(fmt.Errorf("launch non-Buildbucket Tryjob (%T) with Buildbucket backend", tj.Definition.GetBackend()))
		}
		ret[bbDef.GetHost()] = append(ret[bbDef.GetHost()], tj)
	}
	return ret
}

func (f *Facade) schedule(ctx context.Context, host string, r *run.Run, cls []*run.RunCL, tryjobs []*tryjob.Tryjob) error {
	bbClient, err := f.ClientFactory.MakeClient(ctx, host, r.ID.LUCIProject())
	if err != nil {
		return errors.Annotate(err, "failed to create Buildbucket client").Err()
	}
	batchReq, err := prepareBatchRequest(tryjobs, r, cls)
	if err != nil {
		return errors.Annotate(err, "failed to create batch schedule build request").Err()
	}
	batchRes, err := bbClient.Batch(ctx, batchReq)
	if err != nil {
		return err
	}
	ret := errors.NewLazyMultiError(len(tryjobs))
	for i, res := range batchRes.GetResponses() {
		switch res.GetResponse().(type) {
		case *bbpb.BatchResponse_Response_ScheduleBuild:
			build := res.GetScheduleBuild()
			status, result, err := toTryjobStatusAndResult(ctx, build)
			if err != nil {
				ret.Assign(i, err)
			}
			tj := tryjobs[i]
			tj.ExternalID = tryjob.MustBuildbucketID(host, build.Id)
			tj.Status = status
			tj.Result = result
		case *bbpb.BatchResponse_Response_Error:
			ret.Assign(i, status.ErrorProto(res.GetError()))
		default:
			panic(fmt.Errorf("unexpected response type: %T", res.GetResponse()))
		}
	}
	return ret.Get()
}

func prepareBatchRequest(tryjobs []*tryjob.Tryjob, r *run.Run, cls []*run.RunCL) (*bbpb.BatchRequest, error) {
	gcs := makeGerritChanges(cls)
	nonExpProp, expProp, err := makeProperties(r.Mode)
	if err != nil {
		return nil, errors.Annotate(err, "failed to make input properties").Err()
	}
	nonExpTags, expTags, err := makeTags(r, cls)
	if err != nil {
		return nil, errors.Annotate(err, "failed to make tags").Err()
	}
	batchReq := &bbpb.BatchRequest{
		Requests: make([]*bbpb.BatchRequest_Request, len(tryjobs)),
	}
	for i, tj := range tryjobs {
		def := tj.Definition
		req := &bbpb.ScheduleBuildRequest{
			RequestId:     strconv.Itoa(int(tj.ID)),
			Builder:       def.GetBuildbucket().GetBuilder(),
			Properties:    nonExpProp,
			GerritChanges: gcs,
			Tags:          nonExpTags,
			Mask:          TryjobBuildMask,
		}
		if def.GetExperimental() {
			req.Properties = expProp
			req.Tags = expTags
		}
		batchReq.Requests[i] = &bbpb.BatchRequest_Request{
			Request: &bbpb.BatchRequest_Request_ScheduleBuild{
				ScheduleBuild: req,
			},
		}
	}
	return batchReq, nil
}

func makeProperties(mode run.Mode) (nonexp, exp *structpb.Struct, err error) {
	in := &recipe.Input{
		Active:   true,
		DryRun:   mode == run.DryRun,
		RunMode:  string(mode),
		TopLevel: true,
	}
	if nonexp, err = makeCVProperties(in); err != nil {
		return nil, nil, err
	}
	in.Experimental = true
	if exp, err = makeCVProperties(in); err != nil {
		return nil, nil, err
	}
	return nonexp, exp, nil
}

func makeCVProperties(in *recipe.Input) (*structpb.Struct, error) {
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}
	return structpb.NewStruct(map[string]interface{}{propertyKey: raw})
}

func makeGerritChanges(cls []*run.RunCL) []*bbpb.GerritChange {
	ret := make([]*bbpb.GerritChange, len(cls))
	for i, cl := range cls {
		g := cl.Detail.GetGerrit()
		if g == nil {
			panic(fmt.Errorf("change backend (%T) is not supported", cl.Detail.GetKind()))
		}
		ret[i] = &bbpb.GerritChange{
			Host:     g.GetHost(),
			Project:  g.GetInfo().GetProject(),
			Change:   g.GetInfo().GetNumber(),
			Patchset: int64(cl.Detail.GetPatchset()),
		}
	}
	return ret
}

func makeTags(r *run.Run, cls []*run.RunCL) (nonExp, exp []*bbpb.StringPair, err error) {
	var commonTags []*bbpb.StringPair
	addTag := func(key string, values ...string) {
		for _, v := range values {
			commonTags = append(commonTags, &bbpb.StringPair{Key: key, Value: v})
		}
	}
	addTag("user_agent", "cq")
	addTag("cq_attempt_key", r.ID.AttemptKey())
	addTag("cq_cl_group_key", run.ComputeCLGroupKey(cls, false))
	addTag("cq_equivalent_cl_group_key", run.ComputeCLGroupKey(cls, true))
	owners := stringset.New(1) // normally 1 owner 1 triggerer
	triggerers := stringset.New(1)
	for _, cl := range cls {
		ownerID, err := cl.Detail.OwnerIdentity()
		if err != nil {
			return nil, nil, err
		}
		owners.Add(ownerID.Email())
		triggerers.Add(cl.Trigger.GetEmail())
	}
	addTag("cq_cl_owner", owners.ToSlice()...)
	addTag("cq_triggerer", triggerers.ToSlice()...)
	// TODO(crbug/1323978): support custom tags
	nonExp = append([]*bbpb.StringPair{{Key: "cq_experimental", Value: "false"}}, commonTags...)
	exp = append([]*bbpb.StringPair{{Key: "cq_experimental", Value: "true"}}, commonTags...)
	sortTags(nonExp)
	sortTags(exp)
	return nonExp, exp, nil
}

func sortTags(tags []*bbpb.StringPair) {
	sort.Slice(tags, func(i, j int) bool {
		switch strings.Compare(tags[i].GetKey(), tags[j].GetKey()) {
		case 0:
			return strings.Compare(tags[i].GetValue(), tags[j].GetValue()) < 0
		case -1:
			return true
		case 1:
			return false
		default:
			panic(fmt.Errorf("unreachable"))
		}
	})
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
