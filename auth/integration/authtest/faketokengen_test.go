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

package authtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/integration/localauth"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/lucictx"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFakeTokenGenerator(t *testing.T) {
	t.Parallel()

	Convey("With fakes", t, func() {
		ctx, _ := testclock.UseTime(context.Background(), testclock.TestRecentTimeUTC)

		gen := FakeTokenGenerator{KeepRecord: true}

		srv := localauth.Server{
			TokenGenerators: map[string]localauth.TokenGenerator{
				"authtest": &gen,
			},
			DefaultAccountID: "authtest",
		}
		la, err := srv.Start(ctx)
		So(err, ShouldBeNil)
		Reset(func() { srv.Stop(ctx) })
		ctx = lucictx.SetLocalAuth(ctx, la)

		Convey("Access tokens", func() {
			for idx, scope := range []string{"A", "B"} {
				auth := auth.NewAuthenticator(ctx, auth.SilentLogin, auth.Options{
					Scopes: []string{scope, "zzz"},
				})

				email, err := auth.GetEmail()
				So(err, ShouldBeNil)
				So(email, ShouldEqual, DefaultFakeEmail)

				tok, err := auth.GetAccessToken(time.Minute)
				So(err, ShouldBeNil)
				So(tok.AccessToken, ShouldEqual, fmt.Sprintf("fake_token_%d", idx))

				// Expiry is rounded to integer number of seconds, since that's the
				// granularity of OAuth token expiration. Compare int unix timestamps to
				// account for that.
				So(tok.Expiry.Unix(), ShouldEqual,
					testclock.TestRecentTimeUTC.Add(DefaultFakeLifetime).Unix())
			}

			So(gen.TokenScopes("fake_token_0"), ShouldResemble, []string{"A", "zzz"})
			So(gen.TokenScopes("fake_token_1"), ShouldResemble, []string{"B", "zzz"})
		})

		Convey("ID tokens", func() {
			for idx, aud := range []string{"A", "B"} {
				auth := auth.NewAuthenticator(ctx, auth.SilentLogin, auth.Options{
					UseIDTokens: true,
					Audience:    aud,
				})

				email, err := auth.GetEmail()
				So(err, ShouldBeNil)
				So(email, ShouldEqual, DefaultFakeEmail)

				tok, err := auth.GetAccessToken(time.Minute)
				So(err, ShouldBeNil)
				So(tok.AccessToken, ShouldEqual, fmt.Sprintf("fake_token_%d", idx))

				// Expiry is rounded to integer number of seconds, since that's the
				// granularity of OAuth token expiration. Compare int unix timestamps to
				// account for that.
				So(tok.Expiry.Unix(), ShouldEqual,
					testclock.TestRecentTimeUTC.Add(DefaultFakeLifetime).Unix())
			}

			So(gen.TokenScopes("fake_token_0"), ShouldResemble, []string{"audience:A"})
			So(gen.TokenScopes("fake_token_1"), ShouldResemble, []string{"audience:B"})
		})
	})
}
