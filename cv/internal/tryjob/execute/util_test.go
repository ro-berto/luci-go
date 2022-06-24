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

package execute

import (
	"testing"

	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/tryjob"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestComposeReason(t *testing.T) {
	Convey("ComposeReason", t, func() {
		Convey("panics", func() {
			So(func() {
				_ = composeReason(nil)
			}, ShouldPanicLike, "called without tryjobs")
		})
		Convey("works", func() {
			r := composeReason([]*tryjob.Tryjob{
				// restricted.
				{
					ExternalID: tryjob.MustBuildbucketID("test.com", 123456790),
					Definition: &tryjob.Definition{
						ResultVisibility: cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED,
					},
					Result: &tryjob.Result{
						Backend: &tryjob.Result_Buildbucket_{
							Buildbucket: &tryjob.Result_Buildbucket{
								SummaryMarkdown: "A couple\nof lines\nwith secret details",
							},
						},
					},
				},
				// restricted but empty.
				{
					ExternalID: tryjob.MustBuildbucketID("test.com", 123456791),
					Definition: &tryjob.Definition{
						ResultVisibility: cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED,
					},
					Result: &tryjob.Result{},
				},
				// un-restricted.
				{
					ExternalID: tryjob.MustBuildbucketID("test.com", 123456792),
					Definition: &tryjob.Definition{},
					Result: &tryjob.Result{
						Backend: &tryjob.Result_Buildbucket_{
							Buildbucket: &tryjob.Result_Buildbucket{
								SummaryMarkdown: "A couple\nof lines\nwith public details",
							},
						},
					},
				},
			})
			So(r, ShouldEqual, "Failed Tryjobs:\n* https://test.com/build/123456790\n* https://test.com/build/123456791\n* https://test.com/build/123456792\n A couple\n of lines\n with public details")
		})
	})
}

func TestComposeLaunchFailureReason(t *testing.T) {
	Convey("Compose Launch Failure Reason", t, func() {
		defFoo := &tryjob.Definition{
			Backend: &tryjob.Definition_Buildbucket_{
				Buildbucket: &tryjob.Definition_Buildbucket{
					Host: "buildbucket.example.com",
					Builder: &buildbucketpb.BuilderID{
						Project: "ProjectFoo",
						Bucket:  "BucketFoo",
						Builder: "BuilderFoo",
					},
				},
			},
		}
		Convey("Single", func() {
			Convey("restricted", func() {
				defFoo.ResultVisibility = cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED
				reason := composeLaunchFailureReason(map[*tryjob.Definition]string{
					defFoo: "permission denied",
				})
				So(reason, ShouldEqual, "Failed to launch one tryjob. The tryjob name can't be shown due to configuration. Please contact your Project admin for help.")
			})
			Convey("public", func() {
				reason := composeLaunchFailureReason(map[*tryjob.Definition]string{
					defFoo: "permission denied",
				})
				So(reason, ShouldEqual, "Failed to launch tryjob `ProjectFoo/BucketFoo/BuilderFoo`. Reason: permission denied")
			})
		})
		defBar := &tryjob.Definition{
			Backend: &tryjob.Definition_Buildbucket_{
				Buildbucket: &tryjob.Definition_Buildbucket{
					Host: "buildbucket.example.com",
					Builder: &buildbucketpb.BuilderID{
						Project: "ProjectBar",
						Bucket:  "BucketBar",
						Builder: "BuilderBar",
					},
				},
			},
		}
		Convey("Multiple", func() {
			Convey("All restricted", func() {
				defFoo.ResultVisibility = cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED
				defBar.ResultVisibility = cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED
				reason := composeLaunchFailureReason(map[*tryjob.Definition]string{
					defFoo: "permission denied",
					defBar: "builder not found",
				})
				So(reason, ShouldEqual, "Failed to launch 2 tryjobs. The tryjob names can't be shown due to configuration. Please contact your Project admin for help.")
			})
			Convey("Partial restricted", func() {
				defBar.ResultVisibility = cfgpb.CommentLevel_COMMENT_LEVEL_RESTRICTED
				reason := composeLaunchFailureReason(map[*tryjob.Definition]string{
					defFoo: "permission denied",
					defBar: "builder not found",
				})
				So(reason, ShouldEqual, "Failed to launch the following tryjobs:\n* `ProjectFoo/BucketFoo/BuilderFoo`; Failure reason: permission denied\n\nIn addition to the tryjobs above, failed to launch 1 tryjob. But the tryjob names can't be shown due to configuration. Please contact your Project admin for help.")
			})
			Convey("All public", func() {
				reason := composeLaunchFailureReason(map[*tryjob.Definition]string{
					defFoo: "permission denied",
					defBar: "builder not found",
				})
				So(reason, ShouldEqual, "Failed to launch the following tryjobs:\n* `ProjectBar/BucketBar/BuilderBar`; Failure reason: builder not found\n* `ProjectFoo/BucketFoo/BuilderFoo`; Failure reason: permission denied")
			})
		})
	})
}
