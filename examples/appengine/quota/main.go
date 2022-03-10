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

// Package main contains a GAE app demonstrating how to use the server/quota
// module to implement rate limiting for requests. Navigate to /rpcexplorer
// to try out quota operations.
//
// Not intended to be run locally. A local demo can be found under
// server/quota/example.
package main

import (
	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/quota"
	quotapb "go.chromium.org/luci/server/quota/proto"
	"go.chromium.org/luci/server/quota/quotaconfig"
	"go.chromium.org/luci/server/redisconn"

	pb "go.chromium.org/luci/examples/appengine/quota/proto"
	"go.chromium.org/luci/examples/appengine/quota/rpc"
)

func main() {
	modules := []module.Module{
		cfgmodule.NewModuleFromFlags(),
		quota.NewModuleFromFlags(),
		redisconn.NewModuleFromFlags(),
	}

	server.Main(nil, modules, func(srv *server.Server) error {
		// Initialize a static, in-memory implementation of quotaconfig.Interface.
		// See the quota/rpc package for how these quotas are used.
		// TODO(crbug/1280055): Fetch from the config service.
		m, err := quotaconfig.NewMemory(srv.Context, []*quotapb.Policy{
			{
				Name:          "global-rate-limit",
				Resources:     60,
				Replenishment: 1,
			},
			{
				Name:          "per-user-rate-limit/${user}",
				Resources:     60,
				Replenishment: 1,
			},
		})
		if err != nil {
			panic(err)
		}

		// Register the quotaconfig.Interface.
		srv.Context = quota.Use(srv.Context, m)

		// Support authentication for per-user rate limit demo.
		srv.PRPC.Authenticator = &auth.Authenticator{
			Methods: []auth.Method{
				&auth.GoogleOAuth2Method{
					Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
				},
			},
		}

		// Register the demo pRPC service.
		pb.RegisterDemoServer(srv.PRPC, rpc.New())
		return nil
	})
}
