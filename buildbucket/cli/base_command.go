// Copyright 2016 The LUCI Authors.
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

package cli

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/maruel/subcommands"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/buildbucket/protoutil"
	"go.chromium.org/luci/cipd/version"
	"go.chromium.org/luci/common/lhttp"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/prpc"

	pb "go.chromium.org/luci/buildbucket/proto"
)

type baseCommandRun struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags
	host      string
	json      bool
	noColor   bool

	httpClient *http.Client
	client     pb.BuildsClient
}

func (r *baseCommandRun) RegisterGlobalFlags(p Params) {
	r.Flags.StringVar(
		&r.host,
		"host",
		p.DefaultBuildbucketHost,
		"Host for the buildbucket service instance.")
	r.Flags.BoolVar(
		&r.json,
		"json",
		false,
		"Print information in JSON format.")
	r.Flags.BoolVar(
		&r.noColor,
		"nocolor",
		false,
		"Disable coloration.")
	r.authFlags.Register(&r.Flags, p.Auth)
}

// initClients validates -host flag and initializes r.httpClient and r.client.
func (r *baseCommandRun) initClients(ctx context.Context) error {
	// Create HTTP Client.
	authOpts, err := r.authFlags.Options()
	if err != nil {
		return err
	}
	r.httpClient, err = auth.NewAuthenticator(ctx, auth.SilentLogin, authOpts).Client()
	if err != nil {
		return err
	}

	// Validate -host
	if r.host == "" {
		return fmt.Errorf("a host for the buildbucket service must be provided")
	}
	if strings.ContainsRune(r.host, '/') {
		return fmt.Errorf("invalid host %q", r.host)
	}

	// Create Buildbucket client.
	rpcOpts := prpc.DefaultOptions()
	rpcOpts.Insecure = lhttp.IsLocalHost(r.host)
	info, err := version.GetCurrentVersion()
	if err != nil {
		return err
	}
	rpcOpts.UserAgent = fmt.Sprintf("buildbucket CLI, instanceID=%q", info.InstanceID)
	r.client = pb.NewBuildsPRPCClient(&prpc.Client{
		C:       r.httpClient,
		Host:    r.host,
		Options: rpcOpts,
	})
	return nil
}

// batchAndDone executes req and prints the response.
func (r *baseCommandRun) batchAndDone(ctx context.Context, req *pb.BatchRequest) int {
	res, err := r.client.Batch(ctx, req)
	if err != nil {
		return r.done(ctx, err)
	}

	stderr := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, format, args...)
	}

	hasErr := false
	p := newStdoutPrinter(r.noColor)
	for i, res := range res.Responses {
		var build *pb.Build
		switch res := res.Response.(type) {

		case *pb.BatchResponse_Response_Error:
			hasErr = true

			// If we have multiple requests, print a request title.
			if len(req.Requests) > 1 {
				switch req := req.Requests[i].Request.(type) {
				case *pb.BatchRequest_Request_GetBuild:
					r := req.GetBuild
					if r.Id != 0 {
						stderr("build %d", r.Id)
					} else {
						stderr(`build "%s/%d"`, protoutil.FormatBuilderID(r.Builder), r.BuildNumber)
					}

				case *pb.BatchRequest_Request_CancelBuild:
					stderr("build %d", req.CancelBuild.Id)

				default:
					stderr("request #%d", i)
				}
				stderr(": ")
			}

			stderr("%s\n", res.Error.Message)
			continue

		case *pb.BatchResponse_Response_GetBuild:
			build = res.GetBuild
		case *pb.BatchResponse_Response_CancelBuild:
			build = res.CancelBuild
		case *pb.BatchResponse_Response_ScheduleBuild:
			build = res.ScheduleBuild
		default:
			panic("forgot to update batchAndDone()?")
		}

		if r.json {
			p.JSONPB(build)
		} else {
			if i > 0 {
				p.f("\n")
			}
			p.Build(build)
		}
	}
	if hasErr {
		return 1
	}
	return 0
}

func (r *baseCommandRun) done(ctx context.Context, err error) int {
	if err != nil {
		logging.Errorf(ctx, "%s", err)
		return 1
	}
	return 0
}

// retrieveBuildIDs converts build arguments to int64 build ids,
// where a build argument can be an int64 build or a
// "<project>/<bucket>/<builder>/<build_number>" string.
func (r *baseCommandRun) retrieveBuildIDs(ctx context.Context, builds []string) (buildIDs []int64, err error) {
	return retrieveBuildIDs(builds, func(req *pb.BatchRequest) (*pb.BatchResponse, error) {
		return r.client.Batch(ctx, req)
	})
}

func retrieveBuildIDs(builds []string, callBatch func(*pb.BatchRequest) (*pb.BatchResponse, error)) (buildIDs []int64, err error) {
	buildIDs = make([]int64, len(builds))
	batchReq := &pb.BatchRequest{
		Requests: make([]*pb.BatchRequest_Request, 0, len(builds)),
	}
	indexes := make([]int, 0, len(builds))
	idFieldMask := &field_mask.FieldMask{Paths: []string{"id"}}
	for i, b := range builds {
		getBuild, err := protoutil.ParseGetBuildRequest(b)
		if err != nil {
			return nil, fmt.Errorf("invalid build %q: %s", b, err)
		}
		if getBuild.Builder == nil {
			buildIDs[i] = getBuild.Id
		} else {
			getBuild.Fields = idFieldMask
			batchReq.Requests = append(batchReq.Requests, &pb.BatchRequest_Request{
				Request: &pb.BatchRequest_Request_GetBuild{GetBuild: getBuild},
			})
			indexes = append(indexes, i)
		}
	}

	if len(batchReq.Requests) == 0 {
		return buildIDs, nil
	}

	res, err := callBatch(batchReq)
	for i, res := range res.Responses {
		j := indexes[i]
		switch codes.Code(res.GetError().GetCode()) {
		case codes.OK:
			buildIDs[j] = res.GetGetBuild().Id
		case codes.NotFound:
			return nil, fmt.Errorf("build %q not found", builds[j])
		default:
			return nil, err
		}
	}
	return buildIDs, nil
}
