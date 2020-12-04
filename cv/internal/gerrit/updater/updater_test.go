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

package updater

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/gobmap"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRelatedChangeProcessing(t *testing.T) {
	t.Parallel()

	Convey("setGitDeps works", t, func() {
		ctx := context.Background()
		f := fetcher{
			change: 111,
			host:   "host",
			toUpdate: changelist.UpdateFields{
				Snapshot: &changelist.Snapshot{Kind: &changelist.Snapshot_Gerrit{Gerrit: &changelist.Gerrit{}}},
			},
		}

		Convey("No related changes", func() {
			err := f.setGitDeps(ctx, nil)
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Just itself", func() {
			// This isn't happening today, but CV shouldn't choke if Gerrit changes.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3), // No parents.
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3, "107_2"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but no deps", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3, "107_2"),
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but lacking this change", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})
		Convey("Has related, and several times itself", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 2, 2, "107_2"),
				gf.RelatedChange(111, 3, 3, "107_2"),
				gf.RelatedChange(114, 1, 3, "111_3"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})

		Convey("1 parent", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(107, 1, 3, "104_2"),
				gf.RelatedChange(111, 3, 3, "107_1"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
			})
		})

		Convey("Diamond", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(103, 2, 2),
				gf.RelatedChange(104, 2, 2, "103_2"),
				gf.RelatedChange(107, 1, 3, "104_2"),
				gf.RelatedChange(108, 1, 3, "104_2"),
				gf.RelatedChange(111, 3, 3, "107_1", "108_1"),
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
				{Change: 108, Immediate: true},
				{Change: 104, Immediate: false},
				{Change: 103, Immediate: false},
			})
		})

		Convey("Same revision, different changes", func() {
			c104 := gf.RelatedChange(104, 1, 1, "103_2")
			c105 := gf.RelatedChange(105, 1, 1, "103_2")
			c105.GetCommit().Id = c104.GetCommit().GetId()
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(103, 2, 2),
				c104,
				c105, // should be ignored, somewhat arbitrarily.
				gf.RelatedChange(111, 3, 3, "104_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 103, Immediate: false},
			})
		})

		Convey("2 parents which are the same change at different revisions", func() {
			// Actually happened, see https://crbug.com/988309.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(104, 1, 2, "long-ago-merged1"),
				gf.RelatedChange(107, 1, 1, "long-ago-merged2"),
				gf.RelatedChange(104, 2, 2, "107_1"),
				gf.RelatedChange(111, 3, 3, "104_1", "104_2"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 107, Immediate: false},
			})
		})
	})
}

func TestUpdateCLWorks(t *testing.T) {
	t.Parallel()

	Convey("Updating CL works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		const lProject = "infra"
		const gHost = "chromium"
		const gRepo = "depot_tools"

		ct.Cfg.Create(ctx, lProject, singleRepoConfig(gHost, gRepo))
		gobmap.Update(ctx, lProject)

		Convey("No access or permission denied", func() {
			assertDependentMetaOnly := func(change int) {
				cl := getCL(ctx, gHost, change)
				So(cl.Snapshot, ShouldBeNil)
				So(cl.ApplicableConfig, ShouldBeNil)
				So(cl.DependentMeta.GetByProject()[lProject].GetUpdateTime().AsTime(),
					ShouldResemble, ct.Clock.Now().UTC())
			}
			So(refreshExternal(ctx, lProject, gHost, 404, time.Time{}, 0), ShouldBeNil)
			assertDependentMetaOnly(404)
			So(refreshExternal(ctx, lProject, gHost, 403, time.Time{}, 0), ShouldBeNil)
			assertDependentMetaOnly(403)
		})

		Convey("Unhandled Gerrit error results in no CL update", func() {
			ci500 := gf.CI(500, gf.Project(gRepo), gf.Ref("refs/heads/main"))
			err5xx := func(gf.Operation, string) *status.Status {
				return status.New(codes.Internal, "boo")
			}
			Convey("fail to fetch change details", func() {
				ct.GFake.AddFrom(gf.WithCIs(gHost, err5xx, ci500))
				So(refreshExternal(ctx, lProject, gHost, 500, time.Time{}, 0), ShouldErrLike, "boo")
				cl := getCL(ctx, gHost, 500)
				So(cl, ShouldBeNil)
			})

			Convey("fail to get filelist", func() {
				calls := int32(0)
				okThenErr5xx := func(o gf.Operation, p string) *status.Status {
					if atomic.AddInt32(&calls, 1) == 1 {
						return status.New(codes.OK, "")
					} else {
						return err5xx(o, p)
					}
				}
				ct.GFake.AddFrom(gf.WithCIs(gHost, okThenErr5xx, ci500))
				So(refreshExternal(ctx, lProject, gHost, 500, time.Time{}, 0), ShouldErrLike, "boo")
				cl := getCL(ctx, gHost, 500)
				So(cl, ShouldBeNil)
			})
		})
	})
}

func getCL(ctx context.Context, host string, change int) *changelist.CL {
	eid, err := changelist.GobID(host, int64(change))
	So(err, ShouldBeNil)
	cl, err := eid.Get(ctx)
	if err == datastore.ErrNoSuchEntity {
		return nil
	}
	So(err, ShouldBeNil)
	return cl
}

func singleRepoConfig(gHost string, gRepos ...string) *cfgpb.Config {
	projects := make([]*cfgpb.ConfigGroup_Gerrit_Project, len(gRepos))
	for i, gRepo := range gRepos {
		projects[i] = &cfgpb.ConfigGroup_Gerrit_Project{
			Name:      gRepo,
			RefRegexp: []string{"refs/heads/main"},
		}
	}
	return &cfgpb.Config{
		ConfigGroups: []*cfgpb.ConfigGroup{
			{
				Name: "main",
				Gerrit: []*cfgpb.ConfigGroup_Gerrit{
					{
						Url:      "https://" + gHost + "/",
						Projects: projects,
					},
				},
			},
		},
	}
}
