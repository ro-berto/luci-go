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

package main

import (
	"context"
	"fmt"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/authctx"
	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/lucictx"
)

func setResultDBContext(ctx context.Context, buildProto *bbpb.Build, secrets *bbpb.BuildSecrets) context.Context {
	if invocation := buildProto.GetInfra().GetResultdb().GetInvocation(); invocation != "" {
		// For buildbucket builds, buildbucket creates the invocations and saves the
		// info in build proto.
		// Then bbagent uses the info from build proto to set resultdb
		// parameters in the luci context.
		return lucictx.SetResultDB(ctx, &lucictx.ResultDB{
			Hostname: buildProto.Infra.Resultdb.Hostname,
			CurrentInvocation: &lucictx.ResultDBInvocation{
				Name:        invocation,
				UpdateToken: secrets.ResultdbInvocationUpdateToken,
			},
		})
	} else if resultDBCtx := lucictx.GetResultDB(ctx); resultDBCtx != nil {
		// For led builds, swarming creates the invocations and sets resultdb
		// parameters in luci context.
		// Then bbagent gets the parameters from luci context and updates build proto.
		buildProto.Infra.Resultdb = &bbpb.BuildInfra_ResultDB{
			Hostname:   resultDBCtx.Hostname,
			Invocation: resultDBCtx.CurrentInvocation.Name,
		}
		return ctx
	}
	return ctx
}

func setBuildbucketContext(ctx context.Context, hostname *string, secrets *bbpb.BuildSecrets) context.Context {
	// Set `buildbucket` in the context.
	bbCtx := lucictx.GetBuildbucket(ctx)
	if bbCtx == nil || bbCtx.Hostname != *hostname || bbCtx.ScheduleBuildToken != secrets.BuildToken {
		ctx = lucictx.SetBuildbucket(ctx, &lucictx.Buildbucket{
			Hostname:           *hostname,
			ScheduleBuildToken: secrets.BuildToken,
		})
		if bbCtx != nil {
			logging.Warningf(ctx, "buildbucket context is overwritten.")
		}
	}
	return ctx
}

func setRealmContext(ctx context.Context, input *bbpb.BBAgentArgs) context.Context {
	// Populate `realm` in the context based on the build's bucket if there's no
	// realm there already.
	if lucictx.GetRealm(ctx).GetName() == "" {
		project := input.Build.Builder.Project
		bucket := input.Build.Builder.Bucket
		if project != "" && bucket != "" {
			ctx = lucictx.SetRealm(ctx, &lucictx.Realm{
				Name: fmt.Sprintf("%s:%s", project, bucket),
			})
		} else {
			logging.Warningf(ctx, "Bad BuilderID in the build proto: %s", input.Build.Builder)
		}
	}
	return ctx
}

func setLocalAuth(ctx context.Context) context.Context {
	// If asked to use the GCE account, create a new local auth context so it
	// can be properly picked through out the rest of bbagent process tree. Use
	// it as the default task account and as a "system" account (so it is used
	// for things like Logdog PubSub calls).
	authCtx := authctx.Context{
		ID:                  "bbagent",
		Options:             auth.Options{Method: auth.GCEMetadataMethod},
		ExposeSystemAccount: true,
	}
	err := authCtx.Launch(ctx, "")
	check(ctx, errors.Annotate(err, "failed launch the local LUCI auth context").Err())
	defer authCtx.Close(ctx)

	// Switch the default auth in the context to the one we just setup.
	return authCtx.SetLocalAuth(ctx)
}
