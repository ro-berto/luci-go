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

package clpurger

import (
	"testing"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/text"
	"google.golang.org/protobuf/types/known/timestamppb"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/trigger"
	"go.chromium.org/luci/cv/internal/prjmanager/prjpb"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPurgeCLFormatMessage(t *testing.T) {
	t.Parallel()

	Convey("PurgeCL formatMessage works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const gHost = "x-review.googlesource.com"
		ci := gf.CI(
			43, gf.PS(2), gf.Project("re/po"), gf.Ref("refs/heads/main"),
			gf.CQ(+2, testclock.TestRecentTimeUTC, gf.U("user-1")),
			gf.Updated(testclock.TestRecentTimeUTC),
		)
		cl := &changelist.CL{
			Snapshot: &changelist.Snapshot{
				Kind: &changelist.Snapshot_Gerrit{
					Gerrit: &changelist.Gerrit{
						Host: gHost,
						Info: ci,
					},
				},
			},
		}
		task := &prjpb.PurgeCLTask{
			LuciProject: "luci-prj",
			PurgingCl:   nil, // not relevant to this test.
			Trigger:     trigger.Find(ci, &cfgpb.ConfigGroup{}),
			Reasons: []*prjpb.CLError{
				{Kind: nil}, // Set below.
			},
		}

		mustFormat := func() string {
			s, err := formatMessage(ctx, task, cl)
			So(err, ShouldBeNil)
			So(s, ShouldNotContainSubstring, "<no value>")
			return s
		}

		Convey("Lacks owner email", func() {
			task.Reasons[0].Kind = &prjpb.CLError_OwnerLacksEmail{
				OwnerLacksEmail: true,
			}
			So(mustFormat(), ShouldContainSubstring, "set preferred email at https://x-review.googlesource.com/settings/#EmailAddresses")
		})
		Convey("Not yet supported mode", func() {
			task.Reasons[0].Kind = &prjpb.CLError_UnsupportedMode{
				UnsupportedMode: "CUSTOM_RUN",
			}
			So(mustFormat(), ShouldContainSubstring, `its mode "CUSTOM_RUN" is not supported`)
		})
		Convey("Depends on itself", func() {
			task.Reasons[0].Kind = &prjpb.CLError_SelfCqDepend{SelfCqDepend: true}
			So(mustFormat(), ShouldContainSubstring, `because it depends on itself`)
		})

		Convey("Watched by many config groups", func() {
			task.Reasons[0].Kind = &prjpb.CLError_WatchedByManyConfigGroups_{
				WatchedByManyConfigGroups: &prjpb.CLError_WatchedByManyConfigGroups{
					ConfigGroups: []string{"first", "second"},
				},
			}
			s := mustFormat()
			So(s, ShouldContainSubstring, text.Doc(`
				it is watched by more than 1 config group:
				  * first
				  * second

				Please
			`))
			So(s, ShouldContainSubstring, `current CL target ref is "refs/heads/main"`)
		})

		Convey("Invalid deps", func() {
			// Save a CL snapshot for each dep.
			deps := make(map[int]*changelist.Dep, 3)
			for i := 101; i <= 102; i++ {
				depCL, err := changelist.MustGobID(gHost, int64(i)).GetOrInsert(ctx, func(cl *changelist.CL) {
					cl.Snapshot = &changelist.Snapshot{
						LuciProject:           "whatever",
						MinEquivalentPatchset: 1,
						Patchset:              2,
						ExternalUpdateTime:    timestamppb.New(ct.Clock.Now()),
						Kind: &changelist.Snapshot_Gerrit{
							Gerrit: &changelist.Gerrit{
								Host: gHost,
								Info: gf.CI(i),
							},
						},
					}
				})
				So(err, ShouldBeNil)
				deps[i] = &changelist.Dep{Clid: int64(depCL.ID)}
			}
			invalidDeps := &prjpb.CLError_InvalidDeps{ /*set below*/ }
			task.Reasons[0].Kind = &prjpb.CLError_InvalidDeps_{InvalidDeps: invalidDeps}

			Convey("Unwatched", func() {
				invalidDeps.Unwatched = []*changelist.Dep{deps[101]}
				s := mustFormat()
				So(s, ShouldContainSubstring, text.Doc(`
				are not watched by the same LUCI project:
				  * https://x-review.googlesource.com/101

				Please check Cq-Depend
			`))
			})
			Convey("WrongConfigGroup", func() {
				invalidDeps.WrongConfigGroup = []*changelist.Dep{deps[101], deps[102]}
				s := mustFormat()
				So(s, ShouldContainSubstring, text.Doc(`
				its deps do not belong to the same config group:
				  * https://x-review.googlesource.com/101
				  * https://x-review.googlesource.com/102
			`))
			})
			Convey("IncompatMode", func() {
				invalidDeps.IncompatMode = []*changelist.Dep{deps[102], deps[101]}
				s := mustFormat()
				So(s, ShouldContainSubstring, text.Doc(`
				its mode "FULL_RUN" does not match mode on its dependencies:
				  * https://x-review.googlesource.com/101
				  * https://x-review.googlesource.com/102
			`))
			})
		})

		Convey("Several reasons", func() {
			task.Reasons = []*prjpb.CLError{
				{Kind: &prjpb.CLError_OwnerLacksEmail{OwnerLacksEmail: true}},
				{
					Kind: &prjpb.CLError_WatchedByManyConfigGroups_{
						WatchedByManyConfigGroups: &prjpb.CLError_WatchedByManyConfigGroups{
							ConfigGroups: []string{"first", "second"},
						},
					},
				},
			}
			res := mustFormat()
			So(res, ShouldContainSubstring, "set preferred email")
			So(res, ShouldContainSubstring, "more than 1 config group")
		})
	})
}
