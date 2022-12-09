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

package openid

import (
	"context"
	"net/http"
	"testing"
	"time"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/signing/signingtest"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestGoogleComputeAuthMethod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx, _ = testclock.UseTime(ctx, time.Unix(1442540000, 0))

	signer := signingtest.NewSigner(nil)
	certs, _ := signer.Certificates(ctx)
	keyID := signer.KeyNameForTest()

	mintVMToken := func(tok IDToken) string {
		return idTokenForTest(ctx, &tok, keyID, signer)
	}

	const fakeHost = "fake-host.example.com"

	method := &GoogleComputeAuthMethod{
		Header:        "X-Token-Header",
		AudienceCheck: AudienceMatchesHost,
		certs:         certs,
	}
	call := func(authHeader string) (*auth.User, error) {
		u, _, err := method.Authenticate(ctx, &http.Request{
			Host:   fakeHost,
			Header: http.Header{"X-Token-Header": {authHeader}},
		})
		return u, err
	}

	Convey("Skipped if no header", t, func() {
		user, err := call("")
		So(err, ShouldBeNil)
		So(user, ShouldBeNil)
	})

	Convey("Valid token", t, func() {
		tok := IDToken{
			Iss:           "https://accounts.google.com",
			Sub:           "example@example.gserviceaccount.com",
			Email:         "example@example.gserviceaccount.com",
			EmailVerified: true,
			Aud:           "https://" + fakeHost,
			Iat:           clock.Now(ctx).Unix(),
			Exp:           clock.Now(ctx).Add(time.Hour).Unix(),
		}
		tok.Google.ComputeEngine.ProjectID = "example.com:project-id"
		tok.Google.ComputeEngine.Zone = "zone-id"
		tok.Google.ComputeEngine.InstanceName = "instance-id"

		user, err := call(mintVMToken(tok))
		So(err, ShouldBeNil)
		So(user, ShouldResemble, &auth.User{
			Identity: "bot:instance-id@gce.project-id.example.com",
			Extra: &GoogleComputeTokenInfo{
				Audience:       "https://" + fakeHost,
				ServiceAccount: "example@example.gserviceaccount.com",
				Instance:       "instance-id",
				Zone:           "zone-id",
				Project:        "example.com:project-id",
			},
		})
	})

	Convey("No GCE info", t, func() {
		tok := IDToken{
			Iss:           "https://accounts.google.com",
			Sub:           "example@example.gserviceaccount.com",
			Email:         "example@example.gserviceaccount.com",
			EmailVerified: true,
			Aud:           "https://" + fakeHost,
			Iat:           clock.Now(ctx).Unix(),
			Exp:           clock.Now(ctx).Add(time.Hour).Unix(),
		}

		_, err := call(mintVMToken(tok))
		So(err, ShouldErrLike, "no google.compute_engine in the GCE VM token")
	})

	Convey("Bad audience info", t, func() {
		tok := IDToken{
			Iss:           "https://accounts.google.com",
			Sub:           "example@example.gserviceaccount.com",
			Email:         "example@example.gserviceaccount.com",
			EmailVerified: true,
			Aud:           "https://WRONG",
			Iat:           clock.Now(ctx).Unix(),
			Exp:           clock.Now(ctx).Add(time.Hour).Unix(),
		}
		tok.Google.ComputeEngine.ProjectID = "example.com:project-id"
		tok.Google.ComputeEngine.Zone = "zone-id"
		tok.Google.ComputeEngine.InstanceName = "instance-id"

		_, err := call(mintVMToken(tok))
		So(err, ShouldEqual, auth.ErrBadAudience)
	})
}
