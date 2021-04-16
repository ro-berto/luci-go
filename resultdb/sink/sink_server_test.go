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

package sink

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
	sinkpb "go.chromium.org/luci/resultdb/sink/proto/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestReportTestResults(t *testing.T) {
	t.Parallel()

	ctx := metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs(AuthTokenKey, authTokenValue("secret")))

	Convey("ReportTestResults", t, func() {
		// close and drain the server to enforce all the requests processed.
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		cfg := testServerConfig("", "secret")
		tr, cleanup := validTestResult()
		defer cleanup()

		var sentReq *pb.BatchCreateTestResultsRequest
		cfg.Recorder.(*mockRecorder).batchCreateTestResults = func(c context.Context, in *pb.BatchCreateTestResultsRequest) (*pb.BatchCreateTestResultsResponse, error) {
			sentReq = in
			return nil, nil
		}

		expected := &pb.TestResult{
			TestId:       tr.TestId,
			ResultId:     tr.ResultId,
			Expected:     tr.Expected,
			Status:       tr.Status,
			SummaryHtml:  tr.SummaryHtml,
			StartTime:    tr.StartTime,
			Duration:     tr.Duration,
			Tags:         tr.Tags,
			TestMetadata: tr.TestMetadata,
		}

		checkResults := func() {
			sink, err := newSinkServer(ctx, cfg)
			sink.(*sinkpb.DecoratedSink).Service.(*sinkServer).resultIDBase = "foo"
			sink.(*sinkpb.DecoratedSink).Service.(*sinkServer).resultCounter = 100
			So(err, ShouldBeNil)
			defer closeSinkServer(ctx, sink)

			req := &sinkpb.ReportTestResultsRequest{
				TestResults: []*sinkpb.TestResult{tr},
			}
			// Clone because the RPC impl mutates the request objects.
			req = proto.Clone(req).(*sinkpb.ReportTestResultsRequest)
			_, err = sink.ReportTestResults(ctx, req)
			So(err, ShouldBeNil)

			closeSinkServer(ctx, sink)
			So(sentReq, ShouldNotBeNil)
			So(sentReq.Requests, ShouldHaveLength, 1)
			So(sentReq.Requests[0].TestResult, ShouldResembleProto, expected)
		}

		Convey("works", func() {
			Convey("with ServerConfig.TestIDPrefix", func() {
				cfg.TestIDPrefix = "ninja://foo/bar/"
				tr.TestId = "HelloWorld.TestA"
				expected.TestId = "ninja://foo/bar/HelloWorld.TestA"
				checkResults()
			})

			Convey("with ServerConfig.BaseVariant", func() {
				base := []string{"bucket", "try", "builder", "linux-rel"}
				cfg.BaseVariant = pbutil.Variant(base...)
				expected.Variant = pbutil.Variant(base...)
				checkResults()
			})

			Convey("with ServerConfig.BaseTags", func() {
				t1, t2 := pbutil.StringPairs("t1", "v1"), pbutil.StringPairs("t2", "v2")
				// (nil, nil)
				cfg.BaseTags, tr.Tags, expected.Tags = nil, nil, nil
				checkResults()

				// (tag, nil)
				cfg.BaseTags, tr.Tags, expected.Tags = t1, nil, t1
				checkResults()

				// (nil, tag)
				cfg.BaseTags, tr.Tags, expected.Tags = nil, t1, t1
				checkResults()

				// (tag1, tag2)
				cfg.BaseTags, tr.Tags, expected.Tags = t1, t2, append(t1, t2...)
				checkResults()
			})
		})

		Convey("generates a random ResultID, if omitted", func() {
			tr.ResultId = ""
			expected.ResultId = "foo-00101"
			checkResults()
		})

		Convey("duration", func() {
			Convey("with CoerceNegativeDuration", func() {
				cfg.CoerceNegativeDuration = true

				// duration == nil
				tr.Duration, expected.Duration = nil, nil
				checkResults()

				// duration == 0
				tr.Duration, expected.Duration = ptypes.DurationProto(0), ptypes.DurationProto(0)
				checkResults()

				// duration > 0
				tr.Duration, expected.Duration = ptypes.DurationProto(8), ptypes.DurationProto(8)
				checkResults()

				// duration < 0
				tr.Duration = ptypes.DurationProto(-8)
				expected.Duration = ptypes.DurationProto(0)
				checkResults()
			})
			Convey("without CoerceNegativeDuration", func() {
				// duration < 0
				tr.Duration = ptypes.DurationProto(-8)
				sink, err := newSinkServer(ctx, cfg)
				So(err, ShouldBeNil)

				req := &sinkpb.ReportTestResultsRequest{TestResults: []*sinkpb.TestResult{tr}}
				_, err = sink.ReportTestResults(ctx, req)
				So(err, ShouldErrLike, "duration: is < 0")
			})
		})
		Convey("with ServerConfig.TestLocationBase", func() {
			cfg.TestLocationBase = "//base/"
			tr.TestMetadata.Location.FileName = "artifact_dir/a_test.cc"
			expected.TestMetadata = proto.Clone(expected.TestMetadata).(*pb.TestMetadata)
			expected.TestMetadata.Location.FileName = "//base/artifact_dir/a_test.cc"
			checkResults()
		})

		Convey("with ServerConfig.LocationTags", func() {
			rootTags := pbutil.StringPairs(
				"feature", "feature1",
				"monorail_component", "Monorail>Component",
				"teamEmail", "team_email@chromium.org",
				"os", "WINDOWS",
			)

			subTags := pbutil.StringPairs(
				"feature", "feature2",
				"feature", "feature3",
				"monorail_component", "Monorail>Component>Sub",
			)

			cfg.LocationTags = &sinkpb.LocationTags{
				Repos: map[string]*sinkpb.LocationTags_Repo{
					"https://chromium.googlesource.com/chromium/src": {
						Dirs: map[string]*sinkpb.LocationTags_Dir{
							".": {
								Tags: rootTags,
							},
							"artifact_dir": {
								Tags: subTags,
							},
						},
					},
				},
			}
			expected.Tags = append(expected.Tags, pbutil.StringPairs(
				"feature", "feature2",
				"feature", "feature3",
				"monorail_component", "Monorail>Component>Sub",
				"teamEmail", "team_email@chromium.org",
				"os", "WINDOWS",
			)...)
			pbutil.SortStringPairs(expected.Tags)
			checkResults()
		})

		Convey("returns an error if artifacts are invalid", func() {
			sink, err := newSinkServer(ctx, cfg)
			So(err, ShouldBeNil)
			defer closeSinkServer(ctx, sink)

			report := func(trs ...*sinkpb.TestResult) error {
				_, err := sink.ReportTestResults(ctx, &sinkpb.ReportTestResultsRequest{TestResults: trs})
				return err
			}

			tr.Artifacts["art2"] = &sinkpb.Artifact{}
			So(report(tr), ShouldHaveRPCCode, codes.InvalidArgument, "either file_path or contents must be provided")

			// "no such file or directory"
			tr.Artifacts["art2"] = &sinkpb.Artifact{Body: &sinkpb.Artifact_FilePath{FilePath: "not_exist"}}
			So(report(tr), ShouldHaveRPCCode, codes.FailedPrecondition, "querying file info")
		})
	})
}

func TestReportInvocationLevelArtifacts(t *testing.T) {
	t.Parallel()

	Convey("ReportInvocationLevelArtifacts", t, func() {
		ctx := metadata.NewIncomingContext(
			context.Background(),
			metadata.Pairs(AuthTokenKey, authTokenValue("secret")))
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		cfg := testServerConfig("", "secret")
		sink, err := newSinkServer(ctx, cfg)
		So(err, ShouldBeNil)
		defer closeSinkServer(ctx, sink)

		art := &sinkpb.Artifact{Body: &sinkpb.Artifact_Contents{Contents: []byte("123")}}
		req := &sinkpb.ReportInvocationLevelArtifactsRequest{
			Artifacts: map[string]*sinkpb.Artifact{"art1": art},
		}
		_, err = sink.ReportInvocationLevelArtifacts(ctx, req)
		So(err, ShouldBeNil)

		// Duplicated artifact will be rejected.
		_, err = sink.ReportInvocationLevelArtifacts(ctx, req)
		So(err, ShouldErrLike, `artifact "art1" has already been uploaded`)
	})
}
