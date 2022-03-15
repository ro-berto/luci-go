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

package acls

import (
	"fmt"
	"testing"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/cvtesting"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckRunCLs(t *testing.T) {
	t.Parallel()

	const (
		lProject   = "chromium"
		gerritHost = "chromium-review.googlesource.com"
		committers = "committer-group"
		dryRunners = "dry-runner-group"
	)

	Convey("CheckRunCreate", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		cg := prjcfg.ConfigGroup{
			Content: &cfgpb.ConfigGroup{
				Verifiers: &cfgpb.Verifiers{
					GerritCqAbility: &cfgpb.Verifiers_GerritCQAbility{
						CommitterList:    []string{committers},
						DryRunAccessList: []string{dryRunners},
					},
				},
			},
		}

		authState := &authtest.FakeState{FakeDB: authtest.NewFakeDB()}
		ctx = auth.WithState(ctx, authState)
		addMember := func(email, grp string) {
			id, err := identity.MakeIdentity("user:" + email)
			So(err, ShouldBeNil)
			authState.FakeDB.(*authtest.FakeDB).AddMocks(authtest.MockMembership(id, grp))
		}
		addCommitter := func(email string) {
			addMember(email, committers)
		}
		addDryRunner := func(email string) {
			addMember(email, dryRunners)
		}

		// test helpers
		var cls []*changelist.CL
		var trs []*run.Trigger
		addCL := func(triggerer, owner string, m run.Mode) *changelist.CL {
			id := len(cls) + 1
			cls = append(cls, &changelist.CL{
				ID:         common.CLID(id),
				ExternalID: changelist.MustGobID(gerritHost, int64(id)),
				Snapshot: &changelist.Snapshot{
					Kind: &changelist.Snapshot_Gerrit{
						Gerrit: &changelist.Gerrit{
							Host: gerritHost,
							Info: &gerritpb.ChangeInfo{
								Owner: &gerritpb.AccountInfo{
									Email: owner,
								},
							},
						},
					},
				},
			})
			trs = append(trs, &run.Trigger{
				Email: triggerer,
				Mode:  string(m),
			})
			return cls[len(cls)-1]
		}

		mustOK := func() {
			res, err := CheckRunCreate(ctx, &cg, trs, cls)
			So(err, ShouldBeNil)
			So(res.OK(), ShouldBeTrue)
		}
		mustFailWith := func(cl *changelist.CL, format string, args ...interface{}) {
			res, err := CheckRunCreate(ctx, &cg, trs, cls)
			So(err, ShouldBeNil)
			So(res.OK(), ShouldBeFalse)
			So(res.Failure(cl), ShouldContainSubstring, fmt.Sprintf(format, args...))
		}
		approveCL := func(cl *changelist.CL) {
			cl.Snapshot.GetGerrit().GetInfo().Submittable = true
		}
		setAllowOwner := func(action cfgpb.Verifiers_GerritCQAbility_CQAction) {
			cg.Content.Verifiers.GerritCqAbility.AllowOwnerIfSubmittable = action
		}

		Convey("mode == FullRun", func() {
			m := run.FullRun

			Convey("triggerer == owner", func() {
				tr, owner := "t@example.org", "t@example.org"
				cl := addCL(tr, owner, m)

				Convey("triggerer is a committer", func() {
					addCommitter(tr)

					// Should succeed w/ approval.
					mustFailWith(cl, noLGTM)
					approveCL(cl)
					mustOK()
				})
				Convey("triggerer is a dry-runner", func() {
					addDryRunner(tr)

					// Dry-runner can trigger a full-run for own CL w/ approval.
					mustFailWith(cl, noLGTM)
					approveCL(cl)
					mustOK()
				})
				Convey("triggerer is neither dry-runner nor committer", func() {
					Convey("CL approved", func() {
						// Should fail, even if it was approved.
						approveCL(cl)
						mustFailWith(cl, "%q is not a committer", "user:"+tr)
						// unless AllowOwnerIfSubmittable == COMMIT
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustOK()
					})
					Convey("CL not approved", func() {
						// Should fail always.
						mustFailWith(cl, "%q is not a committer", "user:"+tr)
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustFailWith(cl, noLGTM)
					})
				})
			})

			Convey("triggerer != owner", func() {
				tr, owner := "t@example.org", "o@example.org"
				cl := addCL(tr, owner, m)

				Convey("triggerer is a committer", func() {
					addCommitter(tr)

					// Should succeed w/ approval.
					mustFailWith(cl, noLGTM)
					approveCL(cl)
					mustOK()
				})
				Convey("triggerer is a dry-runner", func() {
					addDryRunner(tr)

					// Dry-runner cannot trigger a full-run for someone else' CL,
					// w/ or w/o approval.
					mustFailWith(cl, "neither the CL owner nor a committer")
					approveCL(cl)
					mustFailWith(cl, "neither the CL owner nor a committer")

					// AllowOwnerIfSubmittable doesn't change the decision, either.
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
					mustFailWith(cl, "neither the CL owner nor a committer")
				})

				Convey("triggerer is neither dry-runner nor committer", func() {
					// Should fail always.
					mustFailWith(cl, "neither the CL owner nor a committer")
					approveCL(cl)
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
					mustFailWith(cl, "neither the CL owner nor a committer")
				})
			})
		})

		Convey("mode == DryRun", func() {
			m := run.DryRun

			Convey("triggerer == owner", func() {
				tr, owner := "t@example.org", "t@example.org"
				cl := addCL(tr, owner, m)

				Convey("triggerer is a committer", func() {
					// Committers can trigger a dry-run for someone else' CL
					// w/o approval.
					addCommitter(tr)
					mustOK()
				})
				Convey("triggerer is a dry-runner", func() {
					// Should succeed w/o approval.
					addDryRunner(tr)
					mustOK()
				})
				Convey("triggerer is neither dry-runner nor committer", func() {
					Convey("CL approved", func() {
						// Should fail, even if it was approved.
						approveCL(cl)
						mustFailWith(cl, "%q is not a dry-runner", "user:"+owner)
						// Unless AllowOwnerIfSubmittable == DRY_RUN
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_DRY_RUN)
						mustOK()
						// Or, COMMIT
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustOK()
					})
					Convey("CL not approved", func() {
						// Should fail always.
						mustFailWith(cl, "%q is not a dry-runner", "user:"+owner)
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustFailWith(cl, noLGTM)
					})
				})
			})

			Convey("triggerer != owner", func() {
				tr, owner := "t@example.org", "o@example.org"
				cl := addCL(tr, owner, m)

				Convey("triggerer is a committer", func() {
					// Should suceed w/ or w/o approval.
					addCommitter(tr)
					mustOK()
					approveCL(cl)
					mustOK()
				})
				Convey("triggerer is a dry-runner", func() {
					// Only committers can trigger a dry-run for someone else' CL.
					addDryRunner(tr)
					mustFailWith(cl, "neither the CL owner nor a committer")
					approveCL(cl)
					mustFailWith(cl, "neither the CL owner nor a committer")
					// AllowOwnerIfSubmittable doesn't change the decision, either.
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
					mustFailWith(cl, "neither the CL owner nor a committer")
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_DRY_RUN)
					mustFailWith(cl, "neither the CL owner nor a committer")
				})
				Convey("triggerer is neither dry-runner nor committer", func() {
					// Only committers can trigger a dry-run for someone else' CL.
					mustFailWith(cl, "neither the CL owner nor a committer")
					approveCL(cl)
					mustFailWith(cl, "neither the CL owner nor a committer")
					// AllowOwnerIfSubmittable doesn't change the decision, either.
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
					mustFailWith(cl, "neither the CL owner nor a committer")
					setAllowOwner(cfgpb.Verifiers_GerritCQAbility_DRY_RUN)
					mustFailWith(cl, "neither the CL owner nor a committer")
				})
			})
		})

		Convey("multiple CLs", func() {
			m := run.DryRun
			tr, owner := "t@example.org", "t@example.org"
			cl1 := addCL(tr, owner, m)
			cl2 := addCL(tr, owner, m)
			setAllowOwner(cfgpb.Verifiers_GerritCQAbility_DRY_RUN)

			Convey("all CLs passed", func() {
				approveCL(cl1)
				approveCL(cl2)
				mustOK()
			})
			Convey("all CLs failed", func() {
				mustFailWith(cl1, noLGTM)
				mustFailWith(cl2, noLGTM)
			})
			Convey("Some CLs failed", func() {
				approveCL(cl1)
				mustFailWith(cl1, "CV cannot continue this run due to errors on the other CL(s)")
				mustFailWith(cl2, noLGTM)
			})
		})
	})
}