// Copyright 2017 The LUCI Authors.
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
	"net/http"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/metric"
	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/cron"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/mailer"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/luci_notify/config"
	"go.chromium.org/luci/luci_notify/notify"
)

var buildbucketPubSub = metric.NewCounter(
	"luci/notify/buildbucket-pubsub",
	"Number of received Buildbucket PubSub messages",
	nil,
	// "success", "transient-failure" or "permanent-failure"
	field.String("status"),
)

func main() {
	modules := []module.Module{
		cfgmodule.NewModuleFromFlags(),
		cron.NewModuleFromFlags(),
		gaeemulation.NewModuleFromFlags(),
		mailer.NewModuleFromFlags(),
		tq.NewModuleFromFlags(),
	}

	notify.InitDispatcher(&tq.Default)

	server.Main(nil, modules, func(srv *server.Server) error {
		// Cron endpoints.
		cron.RegisterHandler("update-config", config.UpdateHandler)
		cron.RegisterHandler("update-tree-status", notify.UpdateTreeStatus)

		// Buildbucket Pub/Sub endpoint.
		srv.Routes.POST("/_ah/push-handlers/buildbucket", nil,
			func(c *router.Context) {
				ctx, cancel := context.WithTimeout(c.Context, notify.PUBSUB_POST_REQUEST_TIMEOUT)
				defer cancel()
				c.Context = ctx

				status := ""
				switch err := notify.BuildbucketPubSubHandler(c); {
				case transient.Tag.In(err):
					status = "transient-failure"
					logging.Errorf(ctx, "transient failure: %s", err)
					// Retry the message.
					c.Writer.WriteHeader(http.StatusInternalServerError)

				case err != nil:
					status = "permanent-failure"
					logging.Errorf(ctx, "permanent failure: %s", err)

				default:
					status = "success"
				}

				buildbucketPubSub.Add(ctx, 1, status)
			})

		return nil
	})
}
