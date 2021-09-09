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

// Package main is the main entry point for the app.
package main

import (
	"go.chromium.org/luci/common/proto/access"
	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/cron"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/gerritauth"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/tq"

	// TODO(crbug/1242998): Remove once safe get becomes datastore default.
	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"
	// Enable datastore transactional tasks support.
	_ "go.chromium.org/luci/server/tq/txn/datastore"

	"go.chromium.org/luci/buildbucket/appengine/internal/config"
	"go.chromium.org/luci/buildbucket/appengine/rpc"
	pb "go.chromium.org/luci/buildbucket/proto"
)

func main() {
	mods := []module.Module{
		cfgmodule.NewModuleFromFlags(),
		cron.NewModuleFromFlags(),
		gaeemulation.NewModuleFromFlags(),
		gerritauth.NewModuleFromFlags(),
		tq.NewModuleFromFlags(),
	}

	server.Main(nil, mods, func(srv *server.Server) error {
		srv.PRPC.AccessControl = prpc.AllowOriginAll
		srv.PRPC.Authenticator = &auth.Authenticator{
			Methods: []auth.Method{
				// The default method used by majority of clients.
				&auth.GoogleOAuth2Method{
					Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
				},
				// For authenticating calls from Gerrit plugins.
				&gerritauth.Method,
			},
		}

		access.RegisterAccessServer(srv.PRPC, &access.UnimplementedAccessServer{})
		pb.RegisterBuildsServer(srv.PRPC, rpc.NewBuilds())
		pb.RegisterBuildersServer(srv.PRPC, rpc.NewBuilders())
		// TODO(crbug/1082369): Remove this workaround once field masks can be decoded.
		srv.PRPC.HackFixFieldMasksForJSON = true
		cron.RegisterHandler("update_config", config.UpdateSettingsCfg)
		return nil
	})
}
