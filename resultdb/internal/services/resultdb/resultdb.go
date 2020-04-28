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

package resultdb

import (
	"context"
	"time"

	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/artifactcontent"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

// resultDBServer implements pb.ResultDBServer.
//
// It does not return gRPC-native errors; use DecoratedResultDB with
// internal.CommonPostlude.
type resultDBServer struct {
	generateArtifactURL func(ctx context.Context, artifactName string) (url string, expiration time.Time, err error)
}

// Options is resultdb server configuration.
type Options struct {
	// Use http:// (not https://) for URLs pointing back to ResultDB
	InsecureSelfURLs bool
	// Host name for all user-content URLs.
	ContentHostname string
}

// InitServer initializes a resultdb server.
func InitServer(srv *server.Server, opts Options) error {
	contentServer, err := artifactcontent.NewServer(srv.Context, opts.InsecureSelfURLs, opts.ContentHostname)
	if err != nil {
		return err
	}

	contentServer.InstallHandlers(srv.VirtualHost(opts.ContentHostname))

	pb.RegisterResultDBServer(srv.PRPC, &pb.DecoratedResultDB{
		Service: &resultDBServer{
			generateArtifactURL: contentServer.GenerateSignedURL,
		},
		Prelude:  internal.CommonPrelude,
		Postlude: internal.CommonPostlude,
	})

	// Register an empty Recorder server only to make the discovery service
	// list it.
	// The actual traffic will be directed to another deployment, i.e. this
	// binary will never see Recorder RPCs.
	// TODO(nodir): replace this hack with a separate discovery Deployment that
	// dynamically fetches discovery documents from other deployments and
	// returns their union.
	pb.RegisterRecorderServer(srv.PRPC, nil)
	pb.RegisterDeriverServer(srv.PRPC, nil)

	srv.PRPC.AccessControl = prpc.AllowOriginAll
	return nil
}

// ListArtifacts implements pb.ResultDBServer.
func (s *resultDBServer) ListArtifacts(ctx context.Context, in *pb.ListArtifactsRequest) (*pb.ListArtifactsResponse, error) {
	// TODO(crbug.com/1071258): implement.
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
