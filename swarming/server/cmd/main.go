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
	"flag"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/cron"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/secrets"
	"go.chromium.org/luci/server/tq"

	"go.chromium.org/luci/swarming/server/botsrv"
	"go.chromium.org/luci/swarming/server/hmactoken"
	"go.chromium.org/luci/swarming/server/rbe"
	"go.chromium.org/luci/swarming/server/testing/integrationmocks"
)

func main() {
	modules := []module.Module{
		gaeemulation.NewModuleFromFlags(),
		cron.NewModuleFromFlags(),
		tq.NewModuleFromFlags(),
		secrets.NewModuleFromFlags(),
	}

	hmacSecret := flag.String(
		"shared-hmac-secret",
		"sm://shared-hmac",
		"A name of a secret with an HMAC key to use to produce various tokens.",
	)

	exposeIntegrationMocks := flag.Bool(
		"expose-integration-mocks",
		false,
		"If set, expose endpoints for running integration tests. Must be used locally only.",
	)

	server.Main(nil, modules, func(srv *server.Server) error {
		tokenSecret, err := hmactoken.NewRotatingSecret(srv.Context, *hmacSecret)
		if err != nil {
			return err
		}

		rbeSessions, err := rbe.NewSessionServer(srv.Context, tokenSecret)
		if err != nil {
			return err
		}

		botSrv := botsrv.New(srv.Context, srv.Routes, tokenSecret)

		botsrv.InstallHandler(botSrv, "/swarming/api/v1/bot/rbe/ping", pingHandler)
		botsrv.InstallHandler(botSrv, "/swarming/api/v1/bot/rbe/session/create", rbeSessions.CreateBotSession)
		botsrv.InstallHandler(botSrv, "/swarming/api/v1/bot/rbe/session/update", rbeSessions.UpdateBotSession)

		// Helpers for running local integration tests. They fake some of Swarming
		// Python server behavior.
		if *exposeIntegrationMocks {
			if srv.Options.Prod {
				return errors.Reason("-expose-integration-mocks should not be used with -prod").Err()
			}
			integrationmocks.RegisterIntegrationMocksServer(srv.PRPC, integrationmocks.New(
				srv.Context,
				tokenSecret,
			))
		}

		return nil
	})
}

////////////////////////////////////////////////////////////////////////////////

// pingRequest is a JSON structure of the ping request payload.
type pingRequest struct {
	// Dimensions is dimensions reported by the bot.
	Dimensions map[string][]string `json:"dimensions"`
	// State is the state reported by the bot.
	State map[string]interface{} `json:"state"`
	// Version is the bot version.
	Version string `json:"version"`
	// RBEState is RBE-related state reported by the bot.
	RBEState struct {
		// Instance if the full RBE instance name to use.
		Instance string `json:"instance"`
		// PollToken is base64-encoded HMAC-tagged internalspb.PollState.
		PollToken []byte `json:"poll_token"`
	} `json:"rbe_state"`
}

func (r *pingRequest) ExtractPollToken() []byte               { return r.RBEState.PollToken }
func (r *pingRequest) ExtractSessionToken() []byte            { return nil }
func (r *pingRequest) ExtractDimensions() map[string][]string { return r.Dimensions }

func pingHandler(ctx context.Context, body *pingRequest, r *botsrv.Request) (botsrv.Response, error) {
	logging.Infof(ctx, "Dimensions: %v", r.Dimensions)
	logging.Infof(ctx, "PollState: %v", r.PollState)
	logging.Infof(ctx, "Bot version: %s", body.Version)
	if body.RBEState.Instance != r.PollState.RbeInstance {
		logging.Errorf(ctx, "RBE instance mismatch: reported %q, expecting %q",
			body.RBEState.Instance, r.PollState.RbeInstance,
		)
	}
	return nil, nil
}
