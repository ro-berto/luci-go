// Copyright 2019 The LUCI Authors.
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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/trace"

	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/gerritauth"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/redisconn"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/secrets"

	"go.chromium.org/luci/examples/k8s/helloworld/apipb"
)

func main() {
	// Additional modules that extend the server functionality.
	modules := []module.Module{
		gaeemulation.NewModuleFromFlags(),
		redisconn.NewModuleFromFlags(),
		secrets.NewModuleFromFlags(),
		gerritauth.NewModuleFromFlags(),
	}

	server.Main(nil, modules, func(srv *server.Server) error {
		// pRPC example.
		apipb.RegisterGreeterServer(srv.PRPC, &greeterServer{})

		// Logging and tracing example.
		srv.Routes.GET("/", nil, func(c *router.Context) {
			logging.Debugf(c.Context, "Hello debug world")

			ctx, span := trace.StartSpan(c.Context, "Testing")
			logging.Infof(ctx, "Hello info world")
			time.Sleep(100 * time.Millisecond)
			span.End(nil)

			logging.Warningf(c.Context, "Hello warning world")
			c.Writer.Write([]byte("Hello, world"))

			logging.WithError(fmt.Errorf("boom")).Errorf(c.Context, "Hello error world")
		})

		// Authentication example.
		mw := router.NewMiddlewareChain(auth.Authenticate(
			&auth.GoogleOAuth2Method{
				Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
			},
			&gerritauth.Method,
		))
		srv.Routes.GET("/who", mw, func(c *router.Context) {
			logging.Infof(c.Context, "Authenticated as %s", auth.CurrentIdentity(c.Context))
			fmt.Fprintf(c.Writer, "Authenticated as %s\n", auth.CurrentIdentity(c.Context))
			if info := gerritauth.GetAssertedInfo(c.Context); info != nil {
				fmt.Fprintf(c.Writer, "Gerrit user: %v\n", info.User)
				fmt.Fprintf(c.Writer, "Gerrit CL: %v\n", info.Change)
			}
		})

		// Redis example.
		//
		// To run Redis for tests locally (in particular on OSX):
		//   docker run --name redis -p 6379:6379 --restart always --detach redis
		//
		// Then launch the example with "... -redis-addr :6379".
		//
		// Note that it makes Redis port available on 0.0.0.0. This is a necessity
		// when using Docker-for-Mac. Don't put any sensitive stuff there (or make
		// sure your firewall is configured to block external connections).
		srv.Routes.GET("/redis", nil, func(c *router.Context) {
			conn, err := redisconn.Get(c.Context)
			if err != nil {
				http.Error(c.Writer, err.Error(), 500)
				return
			}
			defer conn.Close()
			n, err := redis.Int(conn.Do("INCR", "testKey"))
			if err != nil {
				http.Error(c.Writer, err.Error(), 500)
				return
			}
			fmt.Fprintf(c.Writer, "%d\n", n)
		})

		srv.Routes.GET("/inc", nil, func(c *router.Context) {
			ctx := c.Context

			ent := TestEntity{ID: "test"}
			if err := datastore.Get(ctx, &ent); err != nil && err != datastore.ErrNoSuchEntity {
				http.Error(c.Writer, err.Error(), 500)
				return
			}
			ent.Value += 1
			if err := datastore.Put(ctx, &ent); err != nil {
				http.Error(c.Writer, err.Error(), 500)
				return
			}

			fmt.Fprintf(c.Writer, "%d\n", ent.Value)
		})

		srv.Routes.GET("/sign", nil, func(c *router.Context) {
			key, sig, err := auth.GetSigner(c.Context).SignBytes(c.Context, []byte("test"))
			if err != nil {
				http.Error(c.Writer, err.Error(), 500)
				return
			}
			fmt.Fprintf(c.Writer, "Key: %s\nSig: %s\n", key, base64.RawStdEncoding.EncodeToString(sig))
		})

		// Shared http.Client with monitoring instrumentation.
		tr, err := auth.GetRPCTransport(srv.Context, auth.NoAuth)
		if err != nil {
			return err
		}
		client := &http.Client{Transport: tr}

		srv.Routes.GET("/example.com", nil, func(c *router.Context) {
			var blob []byte
			resp, err := client.Get("http://example.com")
			if err == nil {
				defer resp.Body.Close()
				blob, err = ioutil.ReadAll(resp.Body)
			}
			if err != nil {
				http.Error(c.Writer, err.Error(), 500)
				return
			}
			c.Writer.Write(blob)
		})

		return nil
	})
}

type TestEntity struct {
	ID    string `gae:"$id"`
	Value int64  `gae:",noindex"`

	_ datastore.PropertyMap `gae:"-,extra"`
}

type greeterServer struct{}

func (*greeterServer) SayHi(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	logging.Infof(ctx, "Hi")
	time.Sleep(100 * time.Millisecond)
	return &emptypb.Empty{}, nil
}
