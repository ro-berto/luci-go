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

// Package main contains a binary demonstrating how to use the server/quota
// module to implement rate limiting for requests. To try it locally:
//	$ go run .&
//	<observe logs, wait for the server to start>
//	$ curl http://localhost:8800/global-rate-limit-endpoint
//	<observe logs, expected 200 OK>
//	$ curl http://localhost:8800/global-rate-limit-endpoint
//	<observe logs, expected 429 Rate Limit Exceeded>
//	<wait 60 seconds>
//	$ curl http://localhost:8800/global-rate-limit-endpoint
//	<observe logs, expected 200 OK>
package main

import (
	"net/http"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"

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
		quota.NewModuleFromFlags(),
		redisconn.NewModuleFromFlags(),
	}

	// Configure an in-memory redis database.
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	server.Main(nil, modules, func(srv *server.Server) error {
		// Initialize a static, in-memory implementation of quotaconfig.Interface.
		m, err := quotaconfig.NewMemory(srv.Context, []*pb.Policy{
			// Policy governing a global rate limit of one request per minute to the
			// /global-rate-limit-endpoint handler. 60 resources are available and
			// the handler consumes 60 resources every time it's called (see below),
			// while the policy is configured to automatically replenish one resource
			// every second.
			{
				Name:          "global-rate-limit",
				Resources:     60,
				Replenishment: 1,
			},
		})
		if err != nil {
			panic(err)
		}

		// Register the quotaconfig.Interface and &redis.Pool.
		srv.Context = redisconn.UsePool(quota.Use(srv.Context, m), &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", s.Addr())
			},
		})

		// Set up a rate-limited endpoint.
		srv.Routes.GET("/global-rate-limit-endpoint", nil, func(c *router.Context) {
			if err := quota.DebitQuota(c.Context, map[string]int64{
				"global-rate-limit": 60,
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
		return nil
	})
}
