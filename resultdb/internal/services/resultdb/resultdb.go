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

// Package resultdb provides implementation for luci.resultdb.v1.ResultDB
// service.
package resultdb

import (
	"context"
	"time"

	"google.golang.org/genproto/googleapis/bytestream"
	sppb "google.golang.org/genproto/googleapis/spanner/v1"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/gerritauth"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/artifactcontent"
	uipb "go.chromium.org/luci/resultdb/internal/proto/ui"
	"go.chromium.org/luci/resultdb/internal/rpcutil"
	"go.chromium.org/luci/resultdb/internal/spanutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// resultDBServer implements pb.ResultDBServer.
//
// It does not return gRPC-native errors; use DecoratedResultDB with
// internal.CommonPostlude.
type resultDBServer struct {
	generateArtifactURL func(ctx context.Context, requestHost, artifactName string) (url string, expiration time.Time, err error)
}

// Options is resultdb server configuration.
type Options struct {
	// InsecureSelfURLs is set to true to use http:// (not https://) for URLs
	// pointing back to ResultDB.
	InsecureSelfURLs bool

	// ContentHostnameMap maps a Host header of GetArtifact request to a host name
	// to use for all user-content URLs.
	//
	// Special key "*" indicates a fallback.
	ContentHostnameMap map[string]string

	// ArtifactRBEInstance is the name of the RBE instance to use for artifact
	// storage. Example: "projects/luci-resultdb/instances/artifacts".
	ArtifactRBEInstance string
}

// InitServer initializes a resultdb server.
func InitServer(srv *server.Server, opts Options) error {
	contentServer, err := newArtifactContentServer(srv.Context, opts)
	if err != nil {
		return errors.Annotate(err, "failed to create an artifact content server").Err()
	}

	// Serve all possible content hostnames.
	hosts := stringset.New(len(opts.ContentHostnameMap))
	for _, v := range opts.ContentHostnameMap {
		hosts.Add(v)
	}
	for _, host := range hosts.ToSortedSlice() {
		contentServer.InstallHandlers(srv.VirtualHost(host))
	}

	rdbSvr := &resultDBServer{
		generateArtifactURL: contentServer.GenerateSignedURL,
	}
	pb.RegisterResultDBServer(srv.PRPC, &pb.DecoratedResultDB{
		Service:  rdbSvr,
		Postlude: internal.CommonPostlude,
	})

	uipb.RegisterUIServer(srv.PRPC, &uipb.DecoratedUI{
		Service:  rdbSvr,
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

	// Allow cross-origin calls, in particular calls using Gerrit auth headers.
	srv.PRPC.AccessControl = func(context.Context, string) prpc.AccessControlDecision {
		return prpc.AccessControlDecision{
			AllowCrossOriginRequests: true,
			AllowCredentials:         true,
			AllowHeaders:             []string{gerritauth.Method.Header},
		}
	}

	// TODO(crbug/1082369): Remove this workaround once field masks can be decoded.
	srv.PRPC.HackFixFieldMasksForJSON = true

	srv.RegisterUnaryServerInterceptor(spanutil.SpannerDefaultsInterceptor(sppb.RequestOptions_PRIORITY_MEDIUM))
	srv.RegisterUnaryServerInterceptor(rpcutil.IdentityKindCountingInterceptor())
	return nil
}

func newArtifactContentServer(ctx context.Context, opts Options) (*artifactcontent.Server, error) {
	if opts.ArtifactRBEInstance == "" {
		return nil, errors.Reason("opts.ArtifactRBEInstance is required").Err()
	}

	conn, err := artifactcontent.RBEConn(ctx)
	if err != nil {
		return nil, err
	}
	bs := bytestream.NewByteStreamClient(conn)

	return &artifactcontent.Server{
		InsecureURLs: opts.InsecureSelfURLs,
		HostnameProvider: func(requestHost string) string {
			if host, ok := opts.ContentHostnameMap[requestHost]; ok {
				return host
			}
			return opts.ContentHostnameMap["*"]
		},

		ReadCASBlob: func(ctx context.Context, req *bytestream.ReadRequest) (bytestream.ByteStream_ReadClient, error) {
			return bs.Read(ctx, req)
		},
		RBECASInstanceName: opts.ArtifactRBEInstance,
	}, nil
}
