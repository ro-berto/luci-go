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

package artifactcontent

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/secrets"
	"go.chromium.org/luci/server/secrets/testsecrets"

	artifactcontenttest "go.chromium.org/luci/resultdb/internal/artifactcontent/testutil"
	"go.chromium.org/luci/resultdb/internal/testutil"
	"go.chromium.org/luci/resultdb/internal/testutil/insert"
	pb "go.chromium.org/luci/resultdb/proto/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateSignedURL(t *testing.T) {
	Convey(`TestGenerateSignedURL`, t, func(c C) {
		ctx := testutil.TestingContext()

		ctx, _ = testclock.UseTime(ctx, testclock.TestRecentTimeUTC)
		ctx = secrets.Use(ctx, &testsecrets.Store{})
		ctx = authtest.MockAuthConfig(ctx)

		s := &Server{
			HostnameProvider: func(string) string {
				return "results.usercontent.example.com"
			},
		}
		So(s.Init(ctx), ShouldBeNil)

		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
		})

		Convey(`Basic case`, func() {
			url, exp, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/artifacts/a")
			So(err, ShouldBeNil)
			So(url, ShouldStartWith, "https://results.usercontent.example.com/invocations/inv/artifacts/a?token=")
			So(exp, ShouldResemble, clock.Now(ctx).UTC().Add(time.Hour))
		})

		Convey(`Escaped test id`, func() {
			url, exp, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/tests/t%2Ft/results/r/artifacts/a")
			So(err, ShouldBeNil)
			So(url, ShouldStartWith, "https://results.usercontent.example.com/invocations/inv/tests/t%2Ft/results/r/artifacts/a?token=")
			So(exp, ShouldResemble, clock.Now(ctx).UTC().Add(time.Hour))
		})
	})
}

func TestServeContent(t *testing.T) {
	Convey(`TestServeContent`, t, func(c C) {
		ctx := testutil.SpannerTestContext(t)

		ctx, _ = testclock.UseTime(ctx, testclock.TestRecentTimeUTC)
		ctx = secrets.Use(ctx, &testsecrets.Store{})
		ctx = authtest.MockAuthConfig(ctx)

		casReader := &artifactcontenttest.FakeCASReader{
			Res: []*bytestream.ReadResponse{
				{Data: []byte("contents")},
			},
		}
		var casReadErr error
		s := &Server{
			HostnameProvider: func(string) string {
				return "example.com"
			},
			RBECASInstanceName: "projects/example/instances/artifacts",
			ReadCASBlob: func(ctx context.Context, req *bytestream.ReadRequest) (bytestream.ByteStream_ReadClient, error) {
				return casReader, casReadErr
			},
		}
		So(s.Init(ctx), ShouldBeNil)
		s.testFetchIsolate = func(ctx context.Context, isolateURL string, w io.Writer) error {
			return fmt.Errorf("unexpected")
		}

		r := router.NewWithRootContext(ctx)
		s.InstallHandlers(r)

		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
		})

		fetch := func(rawurl string) (res *http.Response, contents string) {
			req, err := http.NewRequest("GET", rawurl, nil)
			So(err, ShouldBeNil)
			rec := httptest.NewRecorder()
			s.handleGET(&router.Context{
				Context: ctx,
				Request: req,
				Writer:  rec,
			})
			res = rec.Result()
			rawContents, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)
			defer res.Body.Close()
			return res, string(rawContents)
		}

		testutil.MustApply(ctx,
			insert.Invocation("inv", pb.Invocation_FINALIZED, nil),
			insert.Artifact("inv", "", "a", map[string]interface{}{
				"ContentType": "text/plain",
				"Size":        64,
				"IsolateURL":  "isolate://isolate.example.com/default-gzip/deadbeef",
			}),
			insert.Artifact("inv", "", "rbe", map[string]interface{}{
				"ContentType": "text/plain",
				"Size":        64,
				"RBECASHash":  "sha256:deadbeef",
			}),
			insert.Artifact("inv", "tr/t/t/r", "a", map[string]interface{}{
				"ContentType": "text/plain",
				"Size":        64,
				"RBECASHash":  "sha256:deadbeef",
			}),
		)

		s.testFetchIsolate = func(ctx context.Context, isolateURL string, w io.Writer) error {
			_, err := w.Write([]byte("contents"))
			return err
		}

		Convey(`Invalid resource name`, func() {
			res, _ := fetch("https://results.usercontent.example.com/invocations/inv")
			So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
		})

		Convey(`Invalid token`, func() {
			res, _ := fetch("https://results.usercontent.example.com/invocations/inv/artifacts/a?token=bad")
			So(res.StatusCode, ShouldEqual, http.StatusForbidden)
		})

		Convey(`No token`, func() {
			res, _ := fetch("https://results.usercontent.example.com/invocations/inv/artifacts/a")
			So(res.StatusCode, ShouldEqual, http.StatusUnauthorized)
		})

		Convey(`Escaped test id`, func() {
			u, _, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/tests/t%2Ft/results/r/artifacts/a")
			So(err, ShouldBeNil)
			res, actualContents := fetch(u)
			So(res.StatusCode, ShouldEqual, http.StatusOK)
			So(actualContents, ShouldEqual, "contents")
		})

		Convey(`limit`, func() {
			u, _, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/tests/t%2Ft/results/r/artifacts/a")
			So(err, ShouldBeNil)

			Convey(`empty`, func() {
				u += "&n="
				res, _ := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusOK)
			})

			Convey(`multiple`, func() {
				u += "&n=10&n=50"
				res, _ := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusOK)
			})

			Convey(`invalide`, func() {
				u += "&n=limit"
				res, _ := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey(`RBE-CAS`, func() {
			u, _, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/artifacts/rbe")
			So(err, ShouldBeNil)

			Convey(`Not found`, func() {
				casReadErr = status.Errorf(codes.NotFound, "not found")
				res, _ := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			Convey(`Recv error`, func() {
				casReader.ResErr = status.Errorf(codes.Internal, "internal error")
				res, _ := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

		})

		Convey(`E2E`, func() {
			Convey(`RBE-CAS`, func() {
				casReader.Res = []*bytestream.ReadResponse{
					{Data: []byte("first ")},
					{Data: []byte("second")},
				}

				u, _, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/artifacts/rbe")
				So(err, ShouldBeNil)
				res, actualContents := fetch(u)
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				So(actualContents, ShouldEqual, "first second")
				So(res.Header.Get("Content-Type"), ShouldEqual, "text/plain")
				So(res.Header.Get("Content-Length"), ShouldEqual, "64")
			})

			//Convey(`RBE-CAS with limit`, func() {
			//	casReader.Res = []*bytestream.ReadResponse{
			//		{Data: []byte("first ")},
			//		{Data: []byte("second")},
			//	}
			//
			//	u, _, err := s.GenerateSignedURL(ctx, "request.example.com", "invocations/inv/artifacts/rbe")
			//	So(err, ShouldBeNil)
			//	u += "&n=10"
			//	res, actualContents := fetch(u)
			//	So(res.StatusCode, ShouldEqual, http.StatusOK)
			//	So(actualContents, ShouldEqual, "first s...")
			//	So(res.Header.Get("Content-Type"), ShouldEqual, "text/plain")
			//	So(res.Header.Get("Content-Length"), ShouldEqual, "64")
			//})
		})
	})
}
