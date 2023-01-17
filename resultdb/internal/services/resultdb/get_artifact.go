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
	"strings"
	"time"

	"cloud.google.com/go/storage"

	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/resultdb/internal/artifacts"
	"go.chromium.org/luci/resultdb/internal/config"
	"go.chromium.org/luci/resultdb/internal/gsutil"
	"go.chromium.org/luci/resultdb/internal/invocations"
	"go.chromium.org/luci/resultdb/internal/permissions"
	"go.chromium.org/luci/resultdb/pbutil"
	configpb "go.chromium.org/luci/resultdb/proto/config"
	pb "go.chromium.org/luci/resultdb/proto/v1"
	"go.chromium.org/luci/resultdb/rdbperms"
	"go.chromium.org/luci/server/auth/realms"
	"go.chromium.org/luci/server/span"
)

func verifyReadArtifactPermission(ctx context.Context, name string) error {
	invIDStr, _, _, _, inputErr := pbutil.ParseArtifactName(name)
	if inputErr != nil {
		return appstatus.BadRequest(inputErr)
	}

	return permissions.VerifyInvocation(ctx, invocations.ID(invIDStr), rdbperms.PermGetArtifact)
}

func validateGetArtifactRequest(req *pb.GetArtifactRequest) error {
	if err := pbutil.ValidateArtifactName(req.Name); err != nil {
		return errors.Annotate(err, "name").Err()
	}

	return nil
}

// isAllowedBucketPrefix returns true if the requested object in bucket is allowed to be accessed by all the given
// globalRealms.
func isAllowedBucketPrefix(ctx context.Context, bucket string, object string, globalRealms []string) (isAllowed bool, err error) {

	// A single realm is passed in most cases (e.g. GetArtifact and
	// ListArtifacts). When multiple realms are passed (e.g. QueryArtifacts),
	// check that the access is allowed in all realms.
	for _, globalRealm := range globalRealms {
		project, realm := realms.Split(globalRealm)
		cfg, err := config.Project(ctx, project)
		if err != nil {
			if errors.Is(err, config.ErrNotFoundProjectConfig) {
				return false, nil
			}
			return false, err
		}

		var allowList *configpb.RealmGcsAllowList
		for _, list := range cfg.RealmGcsAllowlist {
			if list.Realm == realm {
				allowList = list
				break
			}
		}
		if allowList == nil {
			return false, nil
		}

		var allowedPrefixes *configpb.GcsBucketPrefixes
		for _, prefixes := range allowList.GcsBucketPrefixes {
			if prefixes.Bucket == bucket {
				allowedPrefixes = prefixes
				break
			}
		}
		if allowedPrefixes == nil {
			return false, nil
		}

		var isAllowed bool
		for _, prefix := range allowedPrefixes.AllowedPrefixes {
			if prefix == "*" {
				isAllowed = true
				break
			}
			if ok := strings.HasPrefix(object, prefix); ok {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return false, nil
		}
	}
	return true, nil
}

// GetArtifact implements pb.ResultDBServer.
func (s *resultDBServer) GetArtifact(ctx context.Context, in *pb.GetArtifactRequest) (*pb.Artifact, error) {
	if err := verifyReadArtifactPermission(ctx, in.Name); err != nil {
		return nil, err
	}

	if err := validateGetArtifactRequest(in); err != nil {
		return nil, appstatus.BadRequest(err)
	}

	ctx, cancel := span.ReadOnlyTransaction(ctx)
	defer cancel()

	art, err := artifacts.Read(ctx, in.Name)
	if err != nil {
		return nil, err
	}

	invIDStr, _, _, _, _ := pbutil.ParseArtifactName(in.Name)
	realm, err := invocations.ReadRealm(ctx, invocations.ID(invIDStr))
	if err != nil {
		return nil, err
	}

	if err := s.populateFetchURLs(ctx, []string{realm}, art); err != nil {
		return nil, err
	}

	return art, nil
}

// populateFetchURLs populates FetchUrl and FetchUrlExpiration fields
// of the artifacts. Uses queriedRealms for GCS Artifacts ACL checking.
//
// Must be called from within some gRPC request handler.
func (s *resultDBServer) populateFetchURLs(ctx context.Context, queriedRealms []string, artifacts ...*pb.Artifact) error {
	// Extract Host header (may be empty) from the request to use it as a basis
	// for generating artifact URLs.
	requestHost := ""
	md, _ := metadata.FromIncomingContext(ctx)
	if val := md.Get("host"); len(val) > 0 {
		requestHost = val[0]
	}

	// Client to fetch from Google Storage
	var gsClient *storage.Client

	for _, a := range artifacts {

		if a.GcsUri != "" {
			now := clock.Now(ctx).UTC()

			if gsClient == nil {
				client, err := storage.NewClient(ctx)
				if err != nil {
					return err
				}
				gsClient = client
			}

			bucket, object := gsutil.Split(a.GcsUri)

			isAllowed, err := isAllowedBucketPrefix(ctx, bucket, object, queriedRealms)
			if err != nil {
				return err
			}

			exp := now.Add(7 * 24 * time.Hour)
			if !isAllowed {
				logging.Infof(ctx,
					"realms: %v are not allowed to access object: %s in bucket: %s",
					queriedRealms, object, bucket)
				a.FetchUrl = ""
				a.FetchUrlExpiration = pbutil.MustTimestampProto(exp)
				continue
			}

			var opts *storage.SignedURLOptions
			ctxOpts := ctx.Value(gsutil.Key("signedURLOpts"))
			if ctxOpts != nil {
				opts = ctxOpts.(*storage.SignedURLOptions)
			}
			url, err := gsutil.GenerateSignedURL(ctx, gsClient, bucket, object, exp, opts)

			if err != nil {
				return err
			}

			a.FetchUrl = url
			a.FetchUrlExpiration = pbutil.MustTimestampProto(exp)
		} else {
			url, exp, err := s.generateArtifactURL(ctx, requestHost, a.Name)
			if err != nil {
				return err
			}

			a.FetchUrl = url
			a.FetchUrlExpiration = pbutil.MustTimestampProto(exp)
		}
	}
	return nil
}
