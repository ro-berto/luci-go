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

package gobmap

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/gae/filter/featureBreaker"
	"go.chromium.org/luci/gae/filter/featureBreaker/flaky"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"golang.org/x/sync/errgroup"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	"go.chromium.org/luci/cv/internal/cvtesting"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGobMapUpdateAndLookup(t *testing.T) {
	t.Parallel()

	ctx := memory.Use(context.Background())
	datastore.GetTestable(ctx).AutoIndex(true)
	datastore.GetTestable(ctx).Consistent(true)

	if testing.Verbose() {
		ctx = logging.SetLevel(gologger.StdConfig.Use(ctx), logging.Debug)
	}

	// First set up an example project with two config groups to show basic
	// regular usage; there is a "main" group which matches a main ref, and
	// another fallback group that matches many other refs, but not all.
	prjcfgtest.Create(ctx, "chromium", &cfgpb.Config{
		ConfigGroups: []*cfgpb.ConfigGroup{
			{
				Name: "group_main",
				Gerrit: []*cfgpb.ConfigGroup_Gerrit{
					{
						Url: "https://cr-review.gs.com/",
						Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
							{
								Name:      "cr/src",
								RefRegexp: []string{"refs/heads/main"},
							},
						},
					},
				},
			},
			{
				// This is the fallback group, so "refs/heads/main" should be
				// handled by the main group but not this one, even though it
				// matches the include regexp list.
				Name:     "group_other",
				Fallback: cfgpb.Toggle_YES,
				Gerrit: []*cfgpb.ConfigGroup_Gerrit{
					{
						Url: "https://cr-review.gs.com/",
						Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
							{
								Name:             "cr/src",
								RefRegexp:        []string{"refs/heads/.*"},
								RefRegexpExclude: []string{"refs/heads/123"},
							},
						},
					},
				},
			},
		},
	})

	update := func(lProject string) error {
		meta := prjcfgtest.MustExist(ctx, lProject)
		cgs, err := meta.GetConfigGroups(ctx)
		if err != nil {
			panic(err)
		}
		return Update(ctx, &meta, cgs)
	}

	Convey("Update with nonexistent project stores nothing", t, func() {
		So(Update(ctx, &prjcfg.Meta{Project: "bogus", Status: prjcfg.StatusNotExists}, nil), ShouldBeNil)
		mps := []*mapPart{}
		q := datastore.NewQuery(mapKind)
		So(datastore.GetAll(ctx, q, &mps), ShouldBeNil)
		So(mps, ShouldBeEmpty)
	})

	Convey("Lookup nonexistent project returns empty result", t, func() {
		So(
			lookup(ctx, "foo-review.gs.com", "repo", "refs/heads/main"),
			ShouldBeEmpty)
	})

	Convey("Basic behavior with one project", t, func() {
		So(update("chromium"), ShouldBeNil)

		Convey("Lookup with main ref returns main group", func() {
			// Note that even though the other config group also matches,
			// only the main config group is applicable since the other one
			// is the fallback config group.
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"),
				ShouldResemble,
				map[string][]string{
					"chromium": {"group_main"},
				})
		})

		Convey("Lookup with other ref returns other group", func() {
			// refs/heads/something matches other group, but not main group.
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/something"),
				ShouldResemble,
				map[string][]string{
					"chromium": {"group_other"},
				})
		})

		Convey("Lookup excluded ref returns nothing", func() {
			// refs/heads/123 is specifically excluded from the "other" group,
			// and also not included in main group.
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/123"),
				ShouldBeEmpty)
		})

		Convey("For a ref with no matching groups the result is empty", func() {
			// If a ref doesn't match any include patterns then no groups
			// match.
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/branch-heads/beta"),
				ShouldBeEmpty)
		})
	})

	Convey("Lookup again returns nothing for disabled project", t, func() {
		// Simulate deleting project. Projects that are deleted are first disabled
		// in practice.
		prjcfgtest.Disable(ctx, "chromium")
		So(update("chromium"), ShouldBeNil)
		So(lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"), ShouldBeEmpty)
	})

	Convey("With two matches and no fallback...", t, func() {
		// Simulate the project being updated so that the "other" group is no
		// longer a fallback group. Now some refs will match both groups.
		prjcfgtest.Enable(ctx, "chromium")
		prjcfgtest.Update(ctx, "chromium", &cfgpb.Config{
			ConfigGroups: []*cfgpb.ConfigGroup{
				{
					Name: "group_main",
					Gerrit: []*cfgpb.ConfigGroup_Gerrit{
						{
							Url: "https://cr-review.gs.com/",
							Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
								{
									Name:      "cr/src",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
					},
				},
				{
					Name: "group_other",
					Gerrit: []*cfgpb.ConfigGroup_Gerrit{
						{
							Url: "https://cr-review.gs.com/",
							Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
								{
									Name:             "cr/src",
									RefRegexp:        []string{"refs/heads/.*"},
									RefRegexpExclude: []string{"refs/heads/123"},
								},
							},
						},
					},
					Fallback: cfgpb.Toggle_NO,
				},
			},
		})

		Convey("Lookup main ref matching two refs", func() {
			// This adds coverage for matching two groups.
			So(update("chromium"), ShouldBeNil)
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"),
				ShouldResemble,
				map[string][]string{"chromium": {"group_main", "group_other"}})
		})
	})

	Convey("With two repos in main group and no other group...", t, func() {
		// This update includes both additions and removals,
		// and also tests multiple hosts.
		prjcfgtest.Update(ctx, "chromium", &cfgpb.Config{
			ConfigGroups: []*cfgpb.ConfigGroup{
				{
					Name: "group_main",
					Gerrit: []*cfgpb.ConfigGroup_Gerrit{
						{
							Url: "https://cr-review.gs.com/",
							Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
								{
									Name:      "cr/src",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
						{
							Url: "https://cr2-review.gs.com/",
							Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
								{
									Name:      "cr2/src",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
					},
				},
			},
		})
		So(update("chromium"), ShouldBeNil)

		Convey("main group matches two different hosts", func() {

			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"),
				ShouldResemble,
				map[string][]string{"chromium": {"group_main"}})
			So(
				lookup(ctx, "cr2-review.gs.com", "cr2/src", "refs/heads/main"),
				ShouldResemble,
				map[string][]string{"chromium": {"group_main"}})
		})

		Convey("other group no longer exists", func() {
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/something"),
				ShouldBeEmpty)
		})
	})

	Convey("With another project matching the same ref...", t, func() {
		// Below another project is created that watches the same repo and ref.
		// This tests multiple projects matching for one Lookup.
		prjcfgtest.Create(ctx, "foo", &cfgpb.Config{
			ConfigGroups: []*cfgpb.ConfigGroup{
				{
					Name: "group_foo",
					Gerrit: []*cfgpb.ConfigGroup_Gerrit{
						{
							Url: "https://cr-review.gs.com/",
							Projects: []*cfgpb.ConfigGroup_Gerrit_Project{
								{
									Name:      "cr/src",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
					},
				},
			},
		})
		So(update("foo"), ShouldBeNil)

		Convey("main group matches two different projects", func() {
			So(
				lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"),
				ShouldResemble,
				map[string][]string{
					"chromium": {"group_main"},
					"foo":      {"group_foo"},
				})
		})
	})

	Convey("Lookup again after correcting the config mistake by deleting the second project", t, func() {
		prjcfgtest.Delete(ctx, "foo")
		meta, err := prjcfg.GetLatestMeta(ctx, "foo")
		So(err, ShouldBeNil)
		So(Update(ctx, &meta, nil), ShouldBeNil)
		So(
			lookup(ctx, "cr-review.gs.com", "cr/src", "refs/heads/main"),
			ShouldResemble,
			map[string][]string{
				"chromium": {"group_main"},
			})
	})
}

func TestGobMapConcurrentUpdates(t *testing.T) {
	t.Parallel()

	Convey("Update() works under flaky Datastore and lots of concurrent tries", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const (
			projects         = 2
			versions         = 20
			repos            = 20
			repoPresenceProb = 0.05
			workers          = 10
			taskRedundancy   = 3 // # of workers doing the same Update() task.
		)

		const (
			gHost = "cr-review.gs.com"
			gRef  = "refs/heads/main"
		)
		// Each LUCI projects gets the same number of config versions.
		// Each version has a random non-empty subset of repos (Gerrit projects).
		var tasks []struct {
			meta prjcfg.Meta
			cgs  []*prjcfg.ConfigGroup
		}
		for v := 1; v <= versions; v++ {
			for lp := 1; lp <= projects; lp++ {
				lProject := fmt.Sprintf("project-%d", lp)
				var gerritProjects []*cfgpb.ConfigGroup_Gerrit_Project
				for i := 1; i <= repos; i++ {
					if mathrand.Float32(ctx) <= repoPresenceProb || (len(gerritProjects) == 0 && i == repos) {
						gerritProjects = append(gerritProjects, &cfgpb.ConfigGroup_Gerrit_Project{
							Name:      fmt.Sprintf("repo-%d", i),
							RefRegexp: []string{gRef},
						})
					}
				}
				cfg := &cfgpb.Config{ConfigGroups: []*cfgpb.ConfigGroup{{
					Name:   fmt.Sprintf("%d-%d", lp, v),
					Gerrit: []*cfgpb.ConfigGroup_Gerrit{{Url: "https://" + gHost, Projects: gerritProjects}},
				}}}
				if v == 1 {
					prjcfgtest.Create(ctx, lProject, cfg)
				} else {
					prjcfgtest.Update(ctx, lProject, cfg)
				}

				task := struct {
					meta prjcfg.Meta
					cgs  []*prjcfg.ConfigGroup
				}{meta: prjcfgtest.MustExist(ctx, lProject)}
				var err error
				if task.cgs, err = task.meta.GetConfigGroups(ctx); err != nil {
					panic(err)
				}
				for t := 1; t <= taskRedundancy; t++ {
					tasks = append(tasks, task)
				}
			}
		}

		ctx, fb := featureBreaker.FilterRDS(ctx, nil)
		// Use a single random source for all flaky.Errors(...) instances. Otherwise
		// they repeat the same random pattern each time withBrokenDS is called.
		rnd := rand.NewSource(0)
		// Make datastore a bit faulty.
		fb.BreakFeaturesWithCallback(
			flaky.Errors(flaky.Params{
				Rand:                             rnd,
				DeadlineProbability:              0.01,
				ConcurrentTransactionProbability: 0.01,
			}),
			featureBreaker.DatastoreFeatures...,
		)

		// Run workers. Each worker process Update tasks in order.
		// Each task is retried until it succeeds.
		eg, egCtx := errgroup.WithContext(ctx)
		retries := make([]int, workers)
		for w := 0; w < workers; w++ {
			w := w
			eg.Go(func() error {
				for i := w; i < len(tasks); i += workers {
				retryLoop:
					for {
						// Simulate passage of time but slow enough that some updates
						// succeed before the lease expiry.
						ct.Clock.Add(maxUpdateDuration / workers)
						switch err := Update(egCtx, &tasks[i].meta, tasks[i].cgs); {
						case err == nil:
							break retryLoop
						case ctx.Err() != nil:
							// This test should be fast. If test context expired, fail
							// quickly.
							return err
						default:
							retries[w]++
						}
					}
				}
				return nil
			})
		}
		So(eg.Wait(), ShouldBeNil)

		// If individual retries exceed 1K, it's probably a good idea to tweak
		// parameters s.t. test runs faster.
		t.Logf("Retries per each worker: %v", retries)

		// "Fix" datastore, letting us examine it.
		fb.BreakFeaturesWithCallback(
			func(context.Context, string) error { return nil },
			featureBreaker.DatastoreFeatures...,
		)
		for p := 1; p <= projects; p++ {
			project := fmt.Sprintf("project-%d", p)

			// Compute which repos we expect to see.
			expectedRepos := stringset.Set{}
			meta := prjcfgtest.MustExist(ctx, project)
			cgs, err := meta.GetConfigGroups(ctx)
			So(err, ShouldBeNil)
			for _, pr := range cgs[0].Content.GetGerrit()[0].GetProjects() {
				expectedRepos.Add(pr.GetName())
			}

			// Ensure the map contains these repos and only them.
			// NOTE: this test reproducibly fails because gobmap.Update is not really
			// safe to call concurrently, so asserted are marked with SkipSo.
			// TODO(crbug/1179286): fix the code and the test.
			var mps []*mapPart
			So(datastore.GetAll(ctx, datastore.NewQuery(mapKind).Eq("Project", project), &mps), ShouldBeNil)
			for _, mp := range mps {
				SkipSo(mp.ConfigHash, ShouldResemble, meta.Hash())
				hostAndRepo := strings.SplitN(mp.Parent.StringID(), "/", 2)
				So(hostAndRepo[0], ShouldResemble, gHost)
				SkipSo(expectedRepos.Del(hostAndRepo[1]), ShouldBeTrue)
			}
			SkipSo(expectedRepos, ShouldBeEmpty)
		}
	})
}

// lookup is a test helper function to return just the projects and config
// group names returned by Lookup.
func lookup(ctx context.Context, host, repo, ref string) map[string][]string {
	ret := map[string][]string{}
	ac, err := Lookup(ctx, host, repo, ref)
	So(err, ShouldBeNil)
	for _, p := range ac.Projects {
		var names []string
		for _, id := range p.ConfigGroupIds {
			parts := strings.Split(id, "/")
			So(len(parts), ShouldEqual, 2)
			names = append(names, parts[1])
		}
		ret[p.Name] = names
	}
	return ret
}
