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

package heuristic

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/bisection/model"
)

func TestChangeLogAnalyzer(t *testing.T) {
	t.Parallel()

	Convey("AreRelelatedExtensions", t, func() {
		So(AreRelelatedExtensions("c", "cpp"), ShouldBeTrue)
		So(AreRelelatedExtensions("py", "pyc"), ShouldBeTrue)
		So(AreRelelatedExtensions("gyp", "gypi"), ShouldBeTrue)
		So(AreRelelatedExtensions("c", "py"), ShouldBeFalse)
		So(AreRelelatedExtensions("abc", "xyz"), ShouldBeFalse)
	})

	Convey("NormalizeObjectFilePath", t, func() {
		data := map[string]string{
			"obj/a/T.x.o":   "a/x.o",
			"obj/a/T.x.y.o": "a/x.y.o",
			"x.o":           "x.o",
			"obj/a/x.obj":   "a/x.obj",
			"a.cc.obj":      "a.cc.obj",
			"T.a.c.o":       "a.c.o",
			"T.a.o":         "a.o",
			"T.a.b.c":       "T.a.b.c",
		}
		for k, v := range data {
			So(NormalizeObjectFilePath(k), ShouldEqual, v)
		}
	})

	Convey("AnalyzeOneChangeLog", t, func() {
		c := context.Background()
		signal := &model.CompileFailureSignal{
			Files: map[string][]int{
				"src/a/b/x.cc":       {27},
				"obj/content/util.o": {},
			},
			Edges: []*model.CompileFailureEdge{
				{
					Dependencies: []string{
						"x/y/aa_impl_mac.cc",
						"y/z/bb_impl.cc",
					},
				},
			},
		}
		signal.CalculateDependencyMap(c)
		Convey("Changelog from a non-blamable email", func() {
			cl := &model.ChangeLog{
				Author: model.ChangeLogActor{
					Email: "chrome-release-bot@chromium.org",
				},
			}

			justification, err := AnalyzeOneChangeLog(c, signal, cl)
			So(err, ShouldBeNil)
			So(justification, ShouldResemble, &model.SuspectJustification{IsNonBlamable: true})
		})

		Convey("Changelog did not touch any file", func() {
			cl := &model.ChangeLog{
				ChangeLogDiffs: []model.ChangeLogDiff{
					{
						Type:    model.ChangeType_ADD,
						NewPath: "some_file.cc",
					},
				},
			}
			justification, err := AnalyzeOneChangeLog(c, signal, cl)
			So(err, ShouldBeNil)
			So(justification, ShouldResemble, &model.SuspectJustification{})
		})

		Convey("Changelog touched relevant files", func() {
			cl := &model.ChangeLog{
				ChangeLogDiffs: []model.ChangeLogDiff{
					{
						Type:    model.ChangeType_MODIFY,
						OldPath: "content/util.c",
						NewPath: "content/util.c",
					},
					{
						Type:    model.ChangeType_ADD,
						NewPath: "dir/a/b/x.cc",
					},
					{
						Type:    model.ChangeType_RENAME,
						OldPath: "unrelated_file_1.cc",
						NewPath: "unrelated_file_2.cc",
					},
					{
						Type:    model.ChangeType_DELETE,
						OldPath: "x/y/aa.h",
					},
					{
						Type:    model.ChangeType_MODIFY,
						OldPath: "y/z/bb.cc",
						NewPath: "y/z/bb.cc",
					},
				},
			}
			justification, err := AnalyzeOneChangeLog(c, signal, cl)
			So(err, ShouldBeNil)
			So(justification, ShouldResemble, &model.SuspectJustification{
				Items: []*model.SuspectJustificationItem{
					{
						Score:    10,
						FilePath: "dir/a/b/x.cc",
						Reason:   `The file "dir/a/b/x.cc" was added and it was in the failure log.`,
						Type:     model.JustificationType_FAILURELOG,
					},
					{
						Score:    2,
						FilePath: "content/util.c",
						Reason:   "The file \"content/util.c\" was modified. It was related to the file obj/content/util.o which was in the failure log.",
						Type:     model.JustificationType_FAILURELOG,
					},
					{
						Score:    1,
						FilePath: "x/y/aa.h",
						Reason:   "The file \"x/y/aa.h\" was deleted. It was related to the dependency x/y/aa_impl_mac.cc.",
						Type:     model.JustificationType_DEPENDENCY,
					},
					{
						Score:    1,
						FilePath: "y/z/bb.cc",
						Reason:   "The file \"y/z/bb.cc\" was modified. It was related to the dependency y/z/bb_impl.cc.",
						Type:     model.JustificationType_DEPENDENCY,
					},
				},
			})
		})
	})

	Convey("AnalyzeChangeLogs", t, func() {
		c := context.Background()
		signal := &model.CompileFailureSignal{
			Files: map[string][]int{
				"src/a/b/x.cc":       {27},
				"obj/content/util.o": {},
			},
		}

		Convey("Results should be sorted", func() {
			cls := []*model.ChangeLog{
				{
					Commit:  "abcd",
					Message: "First blah blah\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/123\n bla",
					ChangeLogDiffs: []model.ChangeLogDiff{
						{
							Type:    model.ChangeType_MODIFY,
							NewPath: "content/util.c",
						},
					},
				},
				{
					Commit:  "efgh",
					Message: "Second blah blah\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/456\n bla",
					ChangeLogDiffs: []model.ChangeLogDiff{
						{
							Type:    model.ChangeType_RENAME,
							OldPath: "unrelated_file_1.cc",
							NewPath: "unrelated_file_2.cc",
						},
					},
				},
				{
					Commit:  "wxyz",
					Message: "Third blah blah\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/789\n bla",
					ChangeLogDiffs: []model.ChangeLogDiff{
						{
							Type:    model.ChangeType_ADD,
							NewPath: "dir/a/b/x.cc",
						},
					},
				},
			}

			analysisResult, err := AnalyzeChangeLogs(c, signal, cls)
			So(err, ShouldBeNil)
			So(analysisResult, ShouldResemble, &model.HeuristicAnalysisResult{
				Items: []*model.HeuristicAnalysisResultItem{
					{
						Commit:      "wxyz",
						ReviewUrl:   "https://chromium-review.googlesource.com/c/chromium/src/+/789",
						ReviewTitle: "Third blah blah",
						Justification: &model.SuspectJustification{
							Items: []*model.SuspectJustificationItem{
								{
									Score:    10,
									FilePath: "dir/a/b/x.cc",
									Reason:   `The file "dir/a/b/x.cc" was added and it was in the failure log.`,
									Type:     model.JustificationType_FAILURELOG,
								},
							},
						},
					},
					{
						Commit:      "abcd",
						ReviewUrl:   "https://chromium-review.googlesource.com/c/chromium/src/+/123",
						ReviewTitle: "First blah blah",
						Justification: &model.SuspectJustification{
							Items: []*model.SuspectJustificationItem{
								{
									Score:    2,
									FilePath: "content/util.c",
									Reason:   "The file \"content/util.c\" was modified. It was related to the file obj/content/util.o which was in the failure log.",
									Type:     model.JustificationType_FAILURELOG,
								},
							},
						},
					},
				},
			})
		})
	})

}
