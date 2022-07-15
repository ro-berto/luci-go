// Copyright 2016 The LUCI Authors.
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

package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.chromium.org/luci/server/auth/signing"
	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/router"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func withSigner(s signing.Signer) router.MiddlewareChain {
	return router.NewMiddlewareChain(
		func(c *router.Context, next router.Handler) {
			c.Context = ModifyConfig(c.Context, func(cfg Config) Config {
				cfg.Signer = s
				return cfg
			})
			next(c)
		},
	)
}

func TestCertificatesHandler(t *testing.T) {
	t.Parallel()

	call := func(s signing.Signer) (*signing.PublicCertificates, error) {
		r := router.New()
		InstallHandlers(r, withSigner(s))
		ts := httptest.NewServer(r)
		defer ts.Close()
		// Note: there are two contexts. One for outer /certificates call
		// (this one), and another for /certificates request handler (it is setup
		// in the middleware chain above).
		ctx := caching.WithEmptyProcessCache(context.Background())
		ctx = ModifyConfig(ctx, func(cfg Config) Config {
			cfg.AnonymousTransport = func(context.Context) http.RoundTripper {
				return http.DefaultTransport
			}
			return cfg
		})
		return signing.FetchCertificates(ctx, ts.URL+"/auth/api/v1/server/certificates")
	}

	Convey("Works", t, func() {
		certs, err := call(&phonySigner{})
		So(err, ShouldBeNil)
		So(len(certs.Certificates), ShouldEqual, 1)
	})

	Convey("No signer", t, func() {
		_, err := call(nil)
		So(err, ShouldErrLike, "HTTP code (404)")
	})

	Convey("Error getting certs", t, func() {
		_, err := call(&phonySigner{errors.New("fail")})
		So(err, ShouldErrLike, "HTTP code (500)")
	})
}

func TestServiceInfoHandler(t *testing.T) {
	t.Parallel()

	Convey("Works", t, func() {
		r := router.New()
		signer := &phonySigner{}
		InstallHandlers(r, withSigner(signer))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/auth/api/v1/server/info", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		So(w.Body.String(), ShouldResemble,
			`{"app_id":"phony-app","app_runtime":"go",`+
				`"app_runtime_version":"go1.5.1",`+
				`"app_version":"1234-abcdef","service_account_name":`+
				`"phony-app-account@example.com"}`+"\n")

		signer.err = errors.New("fail")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/auth/api/v1/server/info", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 500)
		So(w.Body.String(), ShouldResemble, "{\"error\":\"Can't grab service info - fail\"}\n")
	})
}

func TestClientIDHandler(t *testing.T) {
	t.Parallel()

	Convey("Works", t, func() {
		var clientIDErr error
		modConfig := router.NewMiddlewareChain(
			func(c *router.Context, next router.Handler) {
				c.Context = ModifyConfig(c.Context, func(cfg Config) Config {
					cfg.FrontendClientID = func(context.Context) (string, error) {
						return "fake-client-id", clientIDErr
					}
					return cfg
				})
				next(c)
			},
		)

		r := router.New()
		InstallHandlers(r, modConfig)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/auth/api/v1/server/client_id", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 200)
		So(w.Body.String(), ShouldResemble, `{"client_id":"fake-client-id"}`+"\n")

		clientIDErr = errors.New("fail")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/auth/api/v1/server/client_id", nil)
		r.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 500)
		So(w.Body.String(), ShouldResemble, "{\"error\":\"Can't grab the client ID - fail\"}\n")
	})
}

///

type phonySigner struct {
	err error
}

func (s *phonySigner) SignBytes(ctx context.Context, blob []byte) (string, []byte, error) {
	if s.err != nil {
		return "", nil, s.err
	}
	return "phonyKey", []byte("signature"), nil
}

func (s *phonySigner) Certificates(ctx context.Context) (*signing.PublicCertificates, error) {
	if s.err != nil {
		return nil, s.err
	}
	return &signing.PublicCertificates{
		Certificates: []signing.Certificate{
			{
				KeyName:            "phonyKey",
				X509CertificatePEM: "phonyPEM",
			},
		},
	}, nil
}

func (s *phonySigner) ServiceInfo(ctx context.Context) (*signing.ServiceInfo, error) {
	if s.err != nil {
		return nil, s.err
	}
	return &signing.ServiceInfo{
		AppID:              "phony-app",
		AppRuntime:         "go",
		AppRuntimeVersion:  "go1.5.1",
		AppVersion:         "1234-abcdef",
		ServiceAccountName: "phony-app-account@example.com",
	}, nil
}
