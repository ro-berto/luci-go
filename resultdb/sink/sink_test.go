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

package sink

import (
	"context"
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/lucictx"

	"go.chromium.org/luci/resultdb/pbutil"
	sinkpb "go.chromium.org/luci/resultdb/sink/proto/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestNewServer(t *testing.T) {
	t.Parallel()

	Convey("NewServer", t, func() {
		ctx := context.Background()
		cfg := testServerConfig(":42", "my_token")

		Convey("succeeds", func() {
			srv, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(srv, ShouldNotBeNil)
		})
		Convey("generates a random auth token, if missing", func() {
			cfg.AuthToken = ""
			srv, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(srv.cfg.AuthToken, ShouldNotEqual, "")
		})
		Convey("uses the default max leases, if missing or 0", func() {
			srv, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(srv.cfg.ArtChannelMaxLeases, ShouldEqual, DefaultArtChannelMaxLeases)
			So(srv.cfg.TestResultChannelMaxLeases, ShouldEqual, DefaultTestResultChannelMaxLeases)
		})
		Convey("use the custom max leases, if specified", func() {
			cfg.ArtChannelMaxLeases, cfg.TestResultChannelMaxLeases = 123, 456
			srv, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(srv.cfg.ArtChannelMaxLeases, ShouldEqual, 123)
			So(srv.cfg.TestResultChannelMaxLeases, ShouldEqual, 456)
			testServerConfig("", "my_token")
		})
		Convey("with TestLocationBase", func() {
			// empty
			cfg.TestLocationBase = ""
			_, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// valid
			cfg.TestLocationBase = "//base"
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// invalid - not starting with double slahes
			cfg.TestLocationBase = "base"
			_, err = NewServer(ctx, cfg)
			So(err, ShouldErrLike, "TestLocationBase: doesn't start with //")
		})
		Convey("with BaseTags", func() {
			// empty
			cfg.BaseTags = nil
			_, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// valid - unique keys
			cfg.BaseTags = pbutil.StringPairs("k1", "v1", "k2", "v2")
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// valid - duplicate keys
			cfg.BaseTags = pbutil.StringPairs("k1", "v1", "k1", "v2")
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// valid - empty value
			cfg.BaseTags = pbutil.StringPairs("k1", "")
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)

			// invalid - empty key
			cfg.BaseTags = pbutil.StringPairs("", "v1")
			_, err = NewServer(ctx, cfg)
			So(err, ShouldErrLike, "key: unspecified")
		})
		Convey("with MaxBatchableArtifactSize", func() {
			// default
			cfg.MaxBatchableArtifactSize = 0
			s, err := NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(s.cfg.MaxBatchableArtifactSize, ShouldNotEqual, 0)

			// valid
			cfg.MaxBatchableArtifactSize = 512 * 1024
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(cfg.MaxBatchableArtifactSize, ShouldNotEqual, 0)

			cfg.MaxBatchableArtifactSize = 10 * 1024 * 1024
			_, err = NewServer(ctx, cfg)
			So(err, ShouldBeNil)
			So(cfg.MaxBatchableArtifactSize, ShouldNotEqual, 0)

			// invalid - too big
			cfg.MaxBatchableArtifactSize = 10*1024*1024 + 1
			_, err = NewServer(ctx, cfg)
			So(err, ShouldErrLike, "is greater than 10MiB")
		})
	})
}

func TestServer(t *testing.T) {
	t.Parallel()

	Convey("Server", t, func() {
		req := &sinkpb.ReportTestResultsRequest{}
		ctx := context.Background()

		srvCfg := testServerConfig("", "secret")
		srv, err := NewServer(ctx, srvCfg)
		So(err, ShouldBeNil)

		Convey("Start assigns a random port, if missing cfg.Address", func() {
			So(srv.Config().Address, ShouldBeEmpty)
			So(srv.Start(ctx), ShouldBeNil)
			So(srv.Config().Address, ShouldNotBeEmpty)
		})
		Convey("Start fails", func() {
			So(srv.Start(ctx), ShouldBeNil)

			Convey("if called twice", func() {
				So(srv.Start(ctx), ShouldErrLike, "cannot call Start twice")
			})

			Convey("after being closed", func() {
				So(srv.Close(ctx), ShouldBeNil)
				So(srv.Start(ctx), ShouldErrLike, "cannot call Start twice")
			})
		})

		Convey("Close closes the HTTP server", func() {
			So(srv.Start(ctx), ShouldBeNil)
			So(srv.Close(ctx), ShouldBeNil)

			_, err := reportTestResults(ctx, srv.Config().Address, "secret", req)
			// The error could be a connection error or write-error.
			// e.g.,
			// "No connection could be made", "connection refused", "write: broken pipe"
			//
			// The error messages could be different by OS, and this test simply checks
			// whether err != nil.
			So(err, ShouldNotBeNil)
		})

		Convey("Close fails before Start being called", func() {
			So(srv.Close(ctx), ShouldErrLike, ErrCloseBeforeStart)
		})

		Convey("Shutdown closes Done", func() {
			isClosed := func() bool {
				select {
				case <-srv.Done():
					return true
				default:
					return false
				}
			}
			So(srv.Start(ctx), ShouldBeNil)

			// wait until the server is up.
			_, err := reportTestResults(ctx, srv.Config().Address, "secret", req)
			So(err, ShouldBeNil)

			So(isClosed(), ShouldBeFalse)
			So(srv.Shutdown(ctx), ShouldBeNil)
			So(isClosed(), ShouldBeTrue)
		})

		Convey("Run", func() {
			handlerErr := make(chan error, 1)
			runErr := make(chan error)
			expected := errors.New("an error-1")

			Convey("succeeds", func() {
				// launch a go routine with Run
				go func() {
					runErr <- Run(ctx, srvCfg, func(ctx context.Context, cfg ServerConfig) error {
						return <-handlerErr
					})
				}()

				// finish the callback and verify that srv.Run returned what the callback
				// returned.
				handlerErr <- expected
				So(<-runErr, ShouldEqual, expected)
			})

			Convey("serves requests", func() {
				So(srv.Start(ctx), ShouldBeNil)

				Convey("with 200 OK", func() {
					res, err := reportTestResults(ctx, srv.Config().Address, "secret", req)
					So(err, ShouldBeNil)
					So(res, ShouldNotBeNil)
				})

				Convey("with 401 Unauthorized if the auth_token missing", func() {
					_, err := reportTestResults(ctx, srv.Config().Address, "", req)
					So(status.Code(err), ShouldEqual, codes.Unauthenticated)
				})

				Convey("with 403 Forbidden if auth_token mismatched", func() {
					_, err := reportTestResults(ctx, srv.Config().Address, "not-a-secret", req)
					So(status.Code(err), ShouldEqual, codes.PermissionDenied)
				})
			})
		})
	})
}

func TestServerExport(t *testing.T) {
	t.Parallel()

	Convey("Export returns the configured address and auth_token", t, func() {
		ctx := context.Background()
		srv, err := NewServer(ctx, testServerConfig(":42", "hello"))
		So(err, ShouldBeNil)

		ctx = srv.Export(ctx)
		sink := lucictx.GetResultSink(ctx)
		So(sink, ShouldNotBeNil)
		So(sink, ShouldNotBeNil)
		So(sink.Address, ShouldEqual, ":42")
		So(sink.AuthToken, ShouldEqual, "hello")
	})
}
