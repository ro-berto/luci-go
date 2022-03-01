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
// module to implement rate limiting for requests.
//
// Not intended to be run locally. A local demo can be found under
// server/quota/example.
package main

import (
	"net/http"

	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/quota"
	pb "go.chromium.org/luci/server/quota/proto"
	"go.chromium.org/luci/server/quota/quotaconfig"
	"go.chromium.org/luci/server/redisconn"
	"go.chromium.org/luci/server/router"
)

func main() {
	modules := []module.Module{
		cfgmodule.NewModuleFromFlags(),
		quota.NewModuleFromFlags(),
		redisconn.NewModuleFromFlags(),
	}

	server.Main(nil, modules, func(srv *server.Server) error {
		// Initialize a static, in-memory implementation of quotaconfig.Interface.
		m, err := quotaconfig.NewMemory(srv.Context, []*pb.Policy{
			// Policy governing a global rate limit of one request per minute to the
			// /global-rate-limit-endpoint handler. 60 resources are available and
			// the handler consumes 60 resources every time it's called (see below),
			// while the policy is configured to automatically replenish one resource
			// every second. This quota can be reset by sending a request to the
			// /global-rate-limit-reset handler. 60 resources are replenished every time
			// it's called (see below), and the default 60 resources also functions as a
			// cap.
			{
				Name:          "global-rate-limit",
				Resources:     60,
				Replenishment: 1,
			},
		})
		if err != nil {
			panic(err)
		}

		// Register the quotaconfig.Interface.
		srv.Context = quota.Use(srv.Context, m)

		// Set up a rate-limited endpoint by debiting 60 resources every time.
		// Returns an error if enough resources aren't available.
		srv.Routes.GET("/global-rate-limit-endpoint", nil, func(c *router.Context) {
			if err := quota.UpdateQuota(c.Context, map[string]int64{
				"global-rate-limit": -60,
			}, nil); err != nil {
				errors.Log(c.Context, errors.Annotate(err, "debit quota").Err())
				// TODO(crbug/1280055): Differentiate between errors.
				http.Error(c.Writer, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			if _, err := c.Writer.Write([]byte("OK\n")); err != nil {
				errors.Log(c.Context, errors.Annotate(err, "writer").Err())
			}
		})

		// Set up a quota reset endpoint by restoring 60 resources every time.
		// The total resources cap at 60, so repeated calls are fine.
		srv.Routes.GET("/global-rate-limit-reset", nil, func(c *router.Context) {
			if err := quota.UpdateQuota(c.Context, map[string]int64{
				"global-rate-limit": 60,
			}, nil); err != nil {
				errors.Log(c.Context, errors.Annotate(err, "credit quota").Err())
				http.Error(c.Writer, "internal server error", http.StatusInternalServerError)
			}
			if _, err := c.Writer.Write([]byte("OK\n")); err != nil {
				errors.Log(c.Context, errors.Annotate(err, "writer").Err())
			}
		})
		return nil
	})
}
