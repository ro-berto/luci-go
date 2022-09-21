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
	"go.chromium.org/luci/gae/service/datastore"
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
		npRunners  = "new-patchset-runner-group"
	)

	Convey("CheckRunCreate", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		cg := prjcfg.ConfigGroup{
			Content: &cfgpb.ConfigGroup{
				Verifiers: &cfgpb.Verifiers{
					GerritCqAbility: &cfgpb.Verifiers_GerritCQAbility{
						CommitterList:            []string{committers},
						DryRunAccessList:         []string{dryRunners},
						NewPatchsetRunAccessList: []string{npRunners},
					},
				},
			},
		}

		authState := &authtest.FakeState{FakeDB: authtest.NewFakeDB()}
		ctx = auth.WithState(ctx, authState)
		addMember := func(email, grp string) {
			id, err := identity.MakeIdentity(fmt.Sprintf("%s:%s", identity.User, email))
			So(err, ShouldBeNil)
			authState.FakeDB.(*authtest.FakeDB).AddMocks(authtest.MockMembership(id, grp))
		}
		addCommitter := func(email string) {
			addMember(email, committers)
		}
		addDryRunner := func(email string) {
			addMember(email, dryRunners)
		}
		addNPRunner := func(email string) {
			addMember(email, npRunners)
		}

		// test helpers
		var cls []*changelist.CL
		var trs []*run.Trigger
		var clid int64
		addCL := func(triggerer, owner string, m run.Mode) *changelist.CL {
			clid++
			cl := &changelist.CL{
				ID:         common.CLID(clid),
				ExternalID: changelist.MustGobID(gerritHost, clid),
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
			}
			So(datastore.Put(ctx, cl), ShouldBeNil)
			cls = append(cls, cl)
			trs = append(trs, &run.Trigger{
				Email: triggerer,
				Mode:  string(m),
			})
			return cl
		}
		addDep := func(base *changelist.CL, owner string) *changelist.CL {
			clid++
			dep := &changelist.CL{
				ID:         common.CLID(clid),
				ExternalID: changelist.MustGobID(gerritHost, clid),
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
			}
			So(datastore.Put(ctx, dep), ShouldBeNil)
			base.Snapshot.Deps = append(base.Snapshot.Deps, &changelist.Dep{Clid: clid})
			return dep
		}

		mustOK := func() {
			res, err := CheckRunCreate(ctx, &cg, trs, cls)
			So(err, ShouldBeNil)
			So(res.FailuresSummary(), ShouldBeEmpty)
			So(res.OK(), ShouldBeTrue)
		}
		mustFailWith := func(cl *changelist.CL, format string, args ...interface{}) CheckResult {
			res, err := CheckRunCreate(ctx, &cg, trs, cls)
			So(err, ShouldBeNil)
			So(res.OK(), ShouldBeFalse)
			So(res.Failure(cl), ShouldContainSubstring, fmt.Sprintf(format, args...))
			return res
		}
		approveCL := func(cl *changelist.CL) {
			cl.Snapshot.GetGerrit().GetInfo().Submittable = true
			So(datastore.Put(ctx, cl), ShouldBeNil)
		}
		submitCL := func(cl *changelist.CL) {
			cl.Snapshot.GetGerrit().GetInfo().Status = gerritpb.ChangeStatus_MERGED
			So(datastore.Put(ctx, cl), ShouldBeNil)
		}
		setAllowOwner := func(action cfgpb.Verifiers_GerritCQAbility_CQAction) {
			cg.Content.Verifiers.GerritCqAbility.AllowOwnerIfSubmittable = action
		}

		addSubmitReq := func(cl *changelist.CL, name string, st gerritpb.SubmitRequirementResultInfo_Status) {
			ci := cl.Snapshot.Kind.(*changelist.Snapshot_Gerrit).Gerrit.Info
			ci.SubmitRequirements = append(ci.SubmitRequirements,
				&gerritpb.SubmitRequirementResultInfo{Name: name, Status: st})
			So(datastore.Put(ctx, cl), ShouldBeNil)
		}
		satisfyReq := func(cl *changelist.CL, name string) {
			addSubmitReq(cl, name, gerritpb.SubmitRequirementResultInfo_SATISFIED)
		}
		unsatisfyReq := func(cl *changelist.CL, name string) {
			addSubmitReq(cl, name, gerritpb.SubmitRequirementResultInfo_UNSATISFIED)
		}
		naReq := func(cl *changelist.CL, name string) {
			addSubmitReq(cl, name, gerritpb.SubmitRequirementResultInfo_NOT_APPLICABLE)
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
					unsatisfyReq(cl, "Code-Review")
					mustFailWith(cl, fmt.Sprintf(noLGTMWithReqs, "missing `Code-Review`"))
					approveCL(cl)
					mustOK()
				})
				Convey("triggerer is neither dry-runner nor committer", func() {
					Convey("CL approved", func() {
						// Should fail, even if it was approved.
						approveCL(cl)
						mustFailWith(cl, "CV cannot start a Run for `%s` because the user is not a committer", tr)
						// unless AllowOwnerIfSubmittable == COMMIT
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustOK()
					})
					Convey("CL not approved", func() {
						// Should fail always.
						mustFailWith(cl, "CV cannot start a Run for `%s` because the user is not a committer", tr)
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustFailWith(cl, noLGTM)
					})
				})
				Convey("suspiciously noLGTM", func() {
					addDryRunner(tr)
					addSubmitReq(cl, "Code-Review", gerritpb.SubmitRequirementResultInfo_SATISFIED)
					mustFailWith(cl, noLGTMSuspicious)
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
				Convey("suspiciously noLGTM", func() {
					addCommitter(tr)
					addSubmitReq(cl, "Code-Review", gerritpb.SubmitRequirementResultInfo_SATISFIED)
					mustFailWith(cl, noLGTMSuspicious)
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
						mustFailWith(cl, "CV cannot start a Run for `%s` because the user is not a dry-runner", owner)
						// Unless AllowOwnerIfSubmittable == DRY_RUN
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_DRY_RUN)
						mustOK()
						// Or, COMMIT
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustOK()
					})
					Convey("CL not approved", func() {
						// Should fail always.
						mustFailWith(cl, "CV cannot start a Run for `%s` because the user is not a dry-runner", owner)
						setAllowOwner(cfgpb.Verifiers_GerritCQAbility_COMMIT)
						mustFailWith(cl, noLGTM)
					})
				})
			})

			Convey("triggerer != owner", func() {
				tr, owner := "t@example.org", "o@example.org"
				cl := addCL(tr, owner, m)

				Convey("triggerer is a committer", func() {
					// Should succeed w/ or w/o approval.
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

			Convey("w/ dependencies", func() {
				// if triggerer is not the owner, but a committer, then
				// untrusted deps should be checked.
				tr, owner := "t@example.org", "o@example.org"
				cl := addCL(tr, owner, m)
				addCommitter(tr)

				dep1 := addDep(cl, "dep_owner1@example.org")
				dep2 := addDep(cl, "dep_owner2@example.org")
				dep1URL := dep1.ExternalID.MustURL()
				dep2URL := dep2.ExternalID.MustURL()

				Convey("untrusted", func() {
					res := mustFailWith(cl, untrustedDeps)
					So(res.Failure(cl), ShouldContainSubstring, dep1URL)
					So(res.Failure(cl), ShouldContainSubstring, dep2URL)
					// if the deps have no submit requirements, the rejection message
					// shouldn't contain a warning for suspicious CLs.
					So(res.Failure(cl), ShouldNotContainSubstring, untrustedDepsSuspicious)

					Convey("but dep2 satisfies all the SubmitRequirements", func() {
						naReq(dep1, "Code-Review")
						unsatisfyReq(dep1, "Code-Owner")
						satisfyReq(dep2, "Code-Review")
						satisfyReq(dep2, "Code-Owner")
						res := mustFailWith(cl, untrustedDeps)
						So(res.Failure(cl), ShouldContainSubstring, fmt.Sprintf(""+
							"- %s missing approval, although `Code-Review` and `Code-Owner` are satisfied",
							dep2URL,
						))
						So(res.Failure(cl), ShouldContainSubstring, untrustedDepsSuspicious)
					})

					Convey("because all the deps have unsatisfied requirements", func() {
						dep3 := addDep(cl, "dep_owner3@example.org")
						dep3URL := dep3.ExternalID.MustURL()

						unsatisfyReq(dep1, "Code-Review")
						unsatisfyReq(dep2, "Code-Review")
						unsatisfyReq(dep2, "Code-Owner")
						unsatisfyReq(dep3, "Code-Review")
						unsatisfyReq(dep3, "Code-Owner")
						unsatisfyReq(dep3, "Code-Quiz")

						res := mustFailWith(cl, untrustedDeps)
						So(res.Failure(cl), ShouldNotContainSubstring, untrustedDepsSuspicious)
						So(res.Failure(cl), ShouldContainSubstring, fmt.Sprintf(""+
							"- %s missing `Code-Review`\n"+
							"- %s missing `Code-Review` and `Code-Owner`\n"+
							"- %s missing `Code-Review`, `Code-Owner`, and `Code-Quiz`",
							dep1URL, dep2URL, dep3URL,
						))
					})
				})
				Convey("trusted because it's apart of the Run", func() {
					cls = append(cls, dep1, dep2)
					trs = append(trs, &run.Trigger{Email: tr, Mode: string(m)})
					trs = append(trs, &run.Trigger{Email: tr, Mode: string(m)})
					mustOK()
				})
				Convey("trusted because of an approval", func() {
					approveCL(dep1)
					approveCL(dep2)
					mustOK()
				})
				Convey("trusterd because they have been merged already", func() {
					submitCL(dep1)
					submitCL(dep2)
					mustOK()
				})
				Convey("trusted because the owner is a committer", func() {
					addCommitter("dep_owner1@example.org")
					addCommitter("dep_owner2@example.org")
					mustOK()
				})
				Convey("a mix of untrusted and trusted deps", func() {
					addCommitter("dep_owner1@example.org")
					res := mustFailWith(cl, untrustedDeps)
					So(res.Failure(cl), ShouldNotContainSubstring, dep1URL)
					So(res.Failure(cl), ShouldContainSubstring, dep2URL)
				})
			})
		})

		Convey("mode == NewPatchsetRun", func() {
			tr, owner := "t@example.org", "t@example.org"
			cl := addCL(tr, owner, run.NewPatchsetRun)
			Convey("owner is disallowed", func() {
				mustFailWith(cl, "CL owner is not in the allowlist.")
			})
			Convey("owner is allowed", func() {
				addNPRunner(owner)
				mustOK()
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
				mustFailWith(cl1, "CV cannot start a Run due to errors in the following CL(s)")
				mustFailWith(cl2, noLGTM)
			})
		})
	})
}
