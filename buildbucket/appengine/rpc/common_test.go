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

package rpc

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1beta2"

	// TODO(crbug/1242998): Remove once safe get becomes datastore default.
	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"

	"go.chromium.org/luci/server/bqlog"

	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"
	"google.golang.org/grpc/metadata"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/clock/testclock"
	. "go.chromium.org/luci/common/testing/assertions"
)

func init() {
	// By default cause Build.Proto.UpdateTime fields to be unset
	model.OverrideGlobalBuildUpdateTimeClock(nil)
}

func TestLogToBQ(t *testing.T) {
	t.Parallel()

	Convey("logToBQ", t, func(c C) {
		b := &bqlog.Bundler{
			CloudProject: "project",
			Dataset:      "dataset",
		}
		ctx := withBundler(context.Background(), b)
		b.RegisterSink(bqlog.Sink{
			Prototype: &pb.PRPCRequestLog{},
			Table:     "table",
		})
		b.Start(ctx, &bqlog.FakeBigQueryWriter{
			Send: func(req *storagepb.AppendRowsRequest) error {
				rows := req.GetProtoRows().GetRows().GetSerializedRows()
				// TODO(crbug/1250459): Check that rows being sent to BQ look correct.
				c.So(len(rows), ShouldEqual, 1)
				return nil
			},
		})
		defer b.Shutdown(ctx)

		logToBQ(ctx, "id", "parent", "method")
	})
}

func TestValidateTags(t *testing.T) {
	t.Parallel()

	Convey("validate build set", t, func() {
		Convey("valid", func() {
			// test gitiles format
			gitiles := fmt.Sprintf("commit/gitiles/chromium.googlesource.com/chromium/src/+/%s", strings.Repeat("a", 40))
			So(validateBuildSet(gitiles), ShouldBeNil)
			// test gerrit format
			So(validateBuildSet("patch/gerrit/chromium-review.googlesource.com/123/456"), ShouldBeNil)
			// test user format
			So(validateBuildSet("myformat/x"), ShouldBeNil)
		})
		Convey("invalid", func() {
			So(validateBuildSet("commit/gitiles/chromium.googlesource.com/chromium/src.git/+/aaa"), ShouldErrLike, `does not match regex "^commit/gitiles/([^/]+)/(.+?)/\+/([a-f0-9]{40})$"`)
			gitiles := fmt.Sprintf("commit/gitiles/chromium.googlesource.com/a/chromium/src/+/%s", strings.Repeat("a", 40))
			So(validateBuildSet(gitiles), ShouldErrLike, `gitiles project must not start with "a/"`)
			gitiles = fmt.Sprintf("commit/gitiles/chromium.googlesource.com/chromium/src.git/+/%s", strings.Repeat("a", 40))
			So(validateBuildSet(gitiles), ShouldErrLike, `gitiles project must not end with ".git"`)

			So(validateBuildSet("patch/gerrit/chromium-review.googlesource.com/aa/bb"), ShouldErrLike, `does not match regex "^patch/gerrit/([^/]+)/(\d+)/(\d+)$"`)
			So(validateBuildSet(strings.Repeat("a", 2000)), ShouldErrLike, "buildset tag is too long")
		})
	})

	Convey("validate tags", t, func() {
		Convey("invalid", func() {
			// in general
			So(validateTags([]*pb.StringPair{{Key: "k:1", Value: "v"}}, TagNew), ShouldErrLike, "cannot have a colon")

			// build address
			So(validateTags([]*pb.StringPair{{Key: "build_address", Value: "v"}}, TagNew), ShouldErrLike, `tag "build_address" is reserved`)
			So(validateTags([]*pb.StringPair{{Key: "build_address", Value: "v"}}, TagAppend), ShouldErrLike, `cannot be added to an existing build`)

			// buildset
			So(validateTags([]*pb.StringPair{{Key: "buildset", Value: "patch/gerrit/foo"}}, TagNew), ShouldErrLike, `does not match regex "^patch/gerrit/([^/]+)/(\d+)/(\d+)$"`)

			gitiles1 := fmt.Sprintf("commit/gitiles/chromium.googlesource.com/chromium/src/+/%s", strings.Repeat("a", 40))
			gitiles2 := fmt.Sprintf("commit/gitiles/chromium.googlesource.com/chromium/src/+/%s", strings.Repeat("b", 40))
			So(validateTags([]*pb.StringPair{
				{Key: "buildset", Value: gitiles1},
				{Key: "buildset", Value: gitiles2},
			}, TagNew),
				ShouldErrLike,
				fmt.Sprintf(`tag "buildset:%s" conflicts with tag "buildset:%s"`, gitiles2, gitiles1))
			So(validateTags([]*pb.StringPair{
				{Key: "buildset", Value: gitiles1},
				{Key: "buildset", Value: gitiles1},
			}, TagNew),
				ShouldBeNil)

			// builder
			So(validateTags([]*pb.StringPair{
				{Key: "builder", Value: "1"},
				{Key: "builder", Value: "2"},
			}, TagNew),
				ShouldErrLike,
				`tag "builder:2" conflicts with tag "builder:1"`)
			So(validateTags([]*pb.StringPair{
				{Key: "builder", Value: "1"},
				{Key: "builder", Value: "1"},
			}, TagNew),
				ShouldBeNil)
			So(validateTags([]*pb.StringPair{{Key: "builder", Value: "v"}}, TagAppend), ShouldErrLike, "cannot be added to an existing build")
		})
	})

	Convey("validate summary_markdown", t, func() {
		Convey("valid", func() {
			So(validateSummaryMarkdown("[this](http://example.org) is a link"), ShouldBeNil)
		})

		Convey("too big", func() {
			So(validateSummaryMarkdown(strings.Repeat("☕", summaryMarkdownMaxLength)), ShouldErrLike, "too big to accept")
		})
	})

	Convey("validateCommit", t, func() {
		Convey("nil", func() {
			err := validateCommit(nil)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("empty", func() {
			cm := &pb.GitilesCommit{}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("project", func() {
			cm := &pb.GitilesCommit{
				Host: "host",
			}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "project is required")
		})

		Convey("id", func() {
			Convey("invalid id", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Id:      "id",
				}
				err := validateCommit(cm)
				// sha1
				So(err, ShouldErrLike, "id must match")
			})

			Convey("position", func() {
				cm := &pb.GitilesCommit{
					Host:     "host",
					Project:  "project",
					Id:       "id",
					Position: 1,
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "position requires ref")
			})
		})

		Convey("ref", func() {
			Convey("invalid ref", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Ref:     "ref",
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "ref must match")
			})

			Convey("valid, but w/ invalid id", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Ref:     "refs/r1",
					Id:      "id",
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "id must match")
			})
		})

		Convey("neither id nor ref", func() {
			cm := &pb.GitilesCommit{
				Host:    "host",
				Project: "project",
			}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "one of")
		})

		Convey("valid", func() {
			Convey("id", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Id:      "1234567890123456789012345678901234567890",
				}
				err := validateCommit(cm)
				So(err, ShouldBeNil)
			})

			Convey("ref", func() {
				cm := &pb.GitilesCommit{
					Host:     "host",
					Project:  "project",
					Ref:      "refs/ref",
					Position: 1,
				}
				err := validateCommit(cm)
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("validateCommitWithRef", t, func() {
		Convey("nil", func() {
			So(validateCommitWithRef(nil), ShouldErrLike, "ref is required")
		})

		Convey("empty", func() {
			So(validateCommitWithRef(&pb.GitilesCommit{}), ShouldErrLike, "ref is required")
		})

		Convey("with id", func() {
			cm := &pb.GitilesCommit{
				Host:    "host",
				Project: "project",
				Id:      "id",
			}
			So(validateCommitWithRef(cm), ShouldErrLike, "ref is required")
		})

		Convey("with ref", func() {
			cm := &pb.GitilesCommit{
				Host:     "host",
				Project:  "project",
				Ref:      "refs/",
				Position: 1,
			}
			So(validateCommitWithRef(cm), ShouldBeNil)
		})
	})
}

func TestValidateBuildToken(t *testing.T) {
	t.Parallel()

	Convey("validateBuildToken", t, func() {
		ctx := context.Background()
		ctx, _ = testclock.UseTime(ctx, time.Unix(1444945245, 0))
		b := &model.Build{}
		tk1 := "a token"
		tk2 := "b token"

		Convey("Works", func() {
			b.UpdateToken = tk1
			ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(BuildTokenKey, tk1))
			So(validateBuildToken(ctx, b), ShouldBeNil)
		})

		Convey("Fails", func() {
			Convey("if unmatched", func() {
				b.UpdateToken = tk1
				ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(BuildTokenKey, tk2))
				So(validateBuildToken(ctx, b), ShouldNotBeNil)
			})
			Convey("if missing", func() {
				b.UpdateToken = tk1
				So(validateBuildToken(ctx, b), ShouldNotBeNil)
			})
		})
	})
}
