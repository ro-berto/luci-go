// Copyright 2021 The LUCI Authors.
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

package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/auth/openid"
	"go.chromium.org/luci/server/encryptedcookies/session/sessionpb"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestTokenEndpoint(t *testing.T) {
	t.Parallel()

	Convey("With fake endpoint", t, func() {
		ctx := context.Background()
		ctx, _ = testclock.UseTime(ctx, testclock.TestTimeUTC)
		ctx = authtest.MockAuthConfig(ctx)

		type mockedResponse struct {
			status int
			body   []byte
		}
		resp := make(chan mockedResponse, 1)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "Not a POST", 400)
				return
			}
			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				http.Error(w, "Bad content type", 400)
				return
			}
			r.ParseForm()
			if r.PostForm.Get("k1") != "v1" || r.PostForm.Get("k2") != "v2" {
				http.Error(w, "Wrong POST body", 400)
				return
			}
			mockedResp := <-resp
			w.WriteHeader(mockedResp.status)
			w.Write(mockedResp.body)
		}))
		defer ts.Close()
		doc := openid.DiscoveryDoc{TokenEndpoint: ts.URL}

		mockResponse := func(status int, body interface{}) {
			var blob []byte
			if str, ok := body.(string); ok {
				blob = []byte(str)
			} else {
				blob, _ = json.Marshal(body)
			}
			resp <- mockedResponse{status, blob}
		}

		call := func() (*sessionpb.Private, time.Time, error) {
			return HitTokenEndpoint(ctx, &doc, map[string]string{
				"k1": "v1",
				"k2": "v2",
			})
		}

		Convey("Happy path", func() {
			mockResponse(200, map[string]interface{}{
				"access_token":  "access_token",
				"refresh_token": "refresh_token",
				"id_token":      "id_token",
				"expires_in":    3600,
			})
			priv, exp, err := call()
			So(err, ShouldBeNil)
			So(priv, ShouldResembleProto, &sessionpb.Private{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				IdToken:      "id_token",
			})
			So(exp.Equal(testclock.TestTimeUTC.Add(time.Hour)), ShouldBeTrue)
		})

		Convey("Fatal err", func() {
			mockResponse(400, "Boom")
			_, _, err := call()
			So(err, ShouldErrLike, `got HTTP 400`)
			So(err, ShouldErrLike, `with body "Boom"`)
			So(transient.Tag.In(err), ShouldBeFalse)
		})

		Convey("Transient err", func() {
			mockResponse(500, "Boom")
			_, _, err := call()
			So(err, ShouldErrLike, `got HTTP 500`)
			So(err, ShouldErrLike, `with body "Boom"`)
			So(transient.Tag.In(err), ShouldBeTrue)
		})
	})
}
