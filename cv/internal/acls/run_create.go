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
	"context"
	"fmt"
	"strings"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/server/auth"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg"
	"go.chromium.org/luci/cv/internal/run"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
)

const (
	okButDueToOthers     = "CV cannot continue this run due to errors on the other CL(s) included in this run."
	ownerNotCommitter    = "CV cannot trigger the Run for %q because %q is not a committer."
	ownerNotDryRunner    = "CV cannot trigger the Run for %q because %q is not a dry-runner."
	notOwnerNotCommitter = "CV cannot trigger the Run for %q because %q is neither the CL owner nor a committer."
	noLGTM               = "This CL needs to be approved first to trigger a Run."
	suspiciouslyNoLGTM   = noLGTM + " " +
		"However, all requirements appear to be satisfied. " +
		"It's likely caused by an issue in Gerrit or Gerrit configuration. " +
		"Please contact your Git admin."
	untrustedDeps = "" +
		"CV cannot trigger the Run because of the following dependencies. " +
		"They must be approved because their owners are not committers. " +
		"Alternatively, you can ask the owner of this CL to trigger a dry-run."
	suspiciouslyUntrustedDeps = "" +
		"The above list contains unapproved CLs that satisfy all the submit requirements. " +
		"It's likely caused by an issue in Gerrit or Gerrit configuration. " +
		"Please contact your Git admin."
)

// runCreateChecker holds the evaluation results of a CL Run, and checks
// if the Run can be created.
type runCreateChecker struct {
	cl                      *changelist.CL
	runMode                 run.Mode
	allowOwnerIfSubmittable cfgpb.Verifiers_GerritCQAbility_CQAction
	commGroups              []string // committer groups
	dryGroups               []string // dry-runner groups

	owner         identity.Identity // the CL owner
	triggerer     identity.Identity // the Run triggerer
	isApproved    bool              // if the CL has been approved (LGTMed) in Gerrit
	depsToExamine common.CLIDs      // deps that are possibly untrusted.
	trustedDeps   common.CLIDsSet   // deps that have been proven to be trustable.
}

func (ck runCreateChecker) canTrustDeps(ctx context.Context) (evalResult, error) {
	if len(ck.depsToExamine) == 0 {
		return yes, nil
	}
	deps := make([]*changelist.CL, 0, len(ck.depsToExamine))
	for _, id := range ck.depsToExamine {
		if !ck.trustedDeps.Has(id) {
			deps = append(deps, &changelist.CL{ID: id})
		}
	}
	if len(deps) == 0 {
		return yes, nil
	}

	// Fetch the CL entity of the deps and examine if they are trustable.
	// CV never removes CL entities. Hence, this handles transient and
	// datastore.ErrNoSuchEntity in the same way.
	if err := changelist.LoadCLs(ctx, deps); err != nil {
		return no, err
	}
	untrusted := deps[:0]
	for _, d := range deps {
		// Dep is trusted, if
		// - it has been approved in Gerrit, OR
		// - the owner is a committer
		switch isApproved, err := checkApproval(d.Snapshot); {
		case err != nil:
			return no, errors.Annotate(err, "dep-CL(%d)", d.ID).Err()
		case isApproved:
			ck.trustedDeps.Add(d.ID)
			continue
		}

		depOwner, err := d.Snapshot.OwnerIdentity()
		if err != nil {
			return no, errors.Annotate(err, "dep-CL(%d)", d.ID).Err()
		}
		switch isCommitter, err := ck.isCommitter(ctx, depOwner); {
		case err != nil:
			return no, errors.Annotate(err,
				"dep-CL(%d): checking if owner %q is a committer", d.ID, depOwner).Err()
		case isCommitter:
			ck.trustedDeps.Add(d.ID)
			continue
		}
		untrusted = append(untrusted, d)
	}
	if len(untrusted) == 0 {
		return yes, nil
	}
	return noWithReason(untrustedDepsReason(ctx, untrusted)), nil
}

func (ck runCreateChecker) canCreateRun(ctx context.Context) (evalResult, error) {
	switch ck.runMode {
	case run.FullRun:
		return ck.canCreateFullRun(ctx)
	case run.DryRun, run.QuickDryRun:
		return ck.canCreateDryRun(ctx)
	default:
		panic(fmt.Errorf("unknown mode %q", ck.runMode))
	}
}

func (ck runCreateChecker) canCreateFullRun(ctx context.Context) (evalResult, error) {
	// A committer can run a full run, as long as the CL has been approved.
	isCommitter, err := ck.isCommitter(ctx, ck.triggerer)
	if err != nil {
		return no, err
	}
	if isCommitter {
		if ck.isApproved {
			return yes, nil
		}
		return noWithReason(noLGTMReason(ctx, ck.cl)), nil
	}

	// A non-committer can trigger a full-run,
	// if all of the following conditions are met.
	//
	// 1) triggerer == owner
	// 2) triggerer is a dry-runner OR cg.AllowOwnerIfSubmittable == COMMIT
	// 3) the CL has been approved in Gerrit.
	//
	// That is, a dry-runner can trigger a full-run for own CLs w/ an approval.
	// For more context, crbug.com/692611 and go/cq-after-lgtm.
	if ck.triggerer != ck.owner {
		return noWithReason(fmt.Sprintf(notOwnerNotCommitter, ck.triggerer, ck.triggerer)), nil
	}
	isDryRunner, err := ck.isDryRunner(ctx, ck.triggerer)
	if err != nil {
		return no, err
	}
	if !isDryRunner && ck.allowOwnerIfSubmittable != cfgpb.Verifiers_GerritCQAbility_COMMIT {
		return noWithReason(fmt.Sprintf(ownerNotCommitter, ck.triggerer, ck.triggerer)), nil
	}
	if !ck.isApproved {
		return noWithReason(noLGTMReason(ctx, ck.cl)), nil
	}
	return yes, nil
}

func (ck runCreateChecker) canCreateDryRun(ctx context.Context) (evalResult, error) {
	// A committer can trigger a [Quick]DryRun w/o approval for own CLs.
	isCommitter, err := ck.isCommitter(ctx, ck.triggerer)
	if err != nil {
		return no, err
	}
	if isCommitter {
		if ck.triggerer == ck.owner {
			return yes, nil
		}
		// In order for a committer to trigger a dry-run for
		// someone else' CL, all the dependencies, of which owner
		// is not a committer, must be approved in Gerrit.
		return ck.canTrustDeps(ctx)
	}

	// A non-committer can trigger a dry-run,
	// if all of the following conditions are met.
	//
	// 1) triggerer == owner
	// 2) triggerer is a dry-runner
	//    OR
	//    cg.AllowOwnerIfSubmittable in [COMMIT, DRY_RUN] AND
	//      the CL has been approved in Gerrit.
	// 3) all the deps are trusted.
	//
	// A dep is trusted, if at least one of the following conditions are met.
	// - the dep is one of the CLs included in the Run
	// - the owner of the dep is a committer
	// - the dep has been approved in Gerrit
	//
	// For more context, crbug.com/692611 and go/cq-after-lgtm.
	if ck.triggerer != ck.owner {
		return noWithReason(fmt.Sprintf(notOwnerNotCommitter, ck.triggerer, ck.triggerer)), nil
	}
	isDryRunner, err := ck.isDryRunner(ctx, ck.triggerer)
	if err != nil {
		return no, err
	}
	if !isDryRunner {
		switch ck.allowOwnerIfSubmittable {
		case cfgpb.Verifiers_GerritCQAbility_DRY_RUN:
		case cfgpb.Verifiers_GerritCQAbility_COMMIT:
		default:
			return noWithReason(fmt.Sprintf(ownerNotDryRunner, ck.triggerer, ck.triggerer)), nil
		}
		if !ck.isApproved {
			return noWithReason(noLGTMReason(ctx, ck.cl)), nil
		}
		return ck.canTrustDeps(ctx)
	}
	return yes, nil
}

func (ck runCreateChecker) isDryRunner(ctx context.Context, id identity.Identity) (bool, error) {
	if len(ck.dryGroups) == 0 {
		return false, nil
	}
	return auth.GetState(ctx).DB().IsMember(ctx, id, ck.dryGroups)
}

func (ck runCreateChecker) isCommitter(ctx context.Context, id identity.Identity) (bool, error) {
	if len(ck.commGroups) == 0 {
		return false, nil
	}
	return auth.GetState(ctx).DB().IsMember(ctx, id, ck.commGroups)
}

// CheckRunCreate verifies that the user(s) who triggered Run are authorized
// to create the Run for the CLs.
func CheckRunCreate(ctx context.Context, cg *prjcfg.ConfigGroup, trs []*run.Trigger, cls []*changelist.CL) (CheckResult, error) {
	res := make(CheckResult, len(cls))
	cks, err := evaluateCLs(ctx, cg, trs, cls)
	if err != nil {
		return nil, err
	}
	for _, ck := range cks {
		switch result, err := ck.canCreateRun(ctx); {
		case err != nil:
			return nil, err
		case !result.ok:
			res[ck.cl] = result.reason
		}
	}
	return res, nil
}

func checkApproval(snap *changelist.Snapshot) (bool, error) {
	switch isSubmitted, err := snap.IsSubmitted(); {
	case err != nil:
		return false, err
	case isSubmitted:
		return true, nil
	}
	isSubmittable, err := snap.IsSubmittable()
	if err != nil {
		return false, err
	}
	return isSubmittable, nil
}

func evaluateCLs(ctx context.Context, cg *prjcfg.ConfigGroup, trs []*run.Trigger, cls []*changelist.CL) ([]*runCreateChecker, error) {
	gVerifier := cg.Content.Verifiers.GetGerritCqAbility()

	cks := make([]*runCreateChecker, len(cls))
	trustedDeps := make(common.CLIDsSet, len(cls))
	for i, cl := range cls {
		tr := trs[i]
		triggerer, err := identity.MakeIdentity("user:" + tr.Email)
		if err != nil {
			return nil, errors.Annotate(err, "CL(%d): triggerer %q", cl.ID, tr.Email).Err()
		}
		owner, err := cl.Snapshot.OwnerIdentity()
		if err != nil {
			return nil, errors.Annotate(err, "CL(%d)", cl.ID).Err()
		}
		isApproved, err := checkApproval(cl.Snapshot)
		if err != nil {
			return nil, errors.Annotate(err, "CL(%d)", cl.ID).Err()
		}
		// by default, all deps are untrusted, unless they are part of the Run.
		var depsToExamine common.CLIDs
		if len(cl.Snapshot.Deps) > 0 {
			depsToExamine = make(common.CLIDs, len(cl.Snapshot.Deps))
			for i, d := range cl.Snapshot.Deps {
				depsToExamine[i] = common.CLID(d.Clid)
			}
		}
		trustedDeps.Add(cl.ID)
		cks[i] = &runCreateChecker{
			cl:                      cl,
			runMode:                 run.Mode(tr.Mode),
			allowOwnerIfSubmittable: gVerifier.GetAllowOwnerIfSubmittable(),
			commGroups:              gVerifier.GetCommitterList(),
			dryGroups:               gVerifier.GetDryRunAccessList(),

			owner:         owner,
			triggerer:     triggerer,
			isApproved:    isApproved,
			depsToExamine: depsToExamine,
			trustedDeps:   trustedDeps,
		}
	}
	return cks, nil
}

// untrustedDepsReason generates a RunCreate rejection comment for untrusted deps.
func untrustedDepsReason(ctx context.Context, udeps []*changelist.CL) string {
	var sb strings.Builder
	anySuspicious := false
	sb.WriteString(untrustedDeps)
	for _, d := range udeps {
		fmt.Fprintf(&sb, "\n- %s", d.ExternalID.MustURL())
		if allSatisfied, msg := strSubmitReqsForUnapprovedCL(ctx, d); len(msg) > 0 {
			fmt.Fprintf(&sb, " %s", msg)
			anySuspicious = anySuspicious || allSatisfied
		}
	}
	if anySuspicious {
		fmt.Fprintf(&sb, "\n\n%s", suspiciouslyUntrustedDeps)
	}
	return sb.String()
}

// noLGTMReason generates a RunCreate rejection comment for unapproved CL.
func noLGTMReason(ctx context.Context, cl *changelist.CL) string {
	switch allSatisfied, msg := strSubmitReqsForUnapprovedCL(ctx, cl); {
	case allSatisfied:
		return suspiciouslyNoLGTM
	case len(msg) > 0:
		return fmt.Sprintf("%s This CL is not approved because requirement %s", noLGTM, msg)
	}
	return noLGTM
}

func strSubmitReqsForUnapprovedCL(ctx context.Context, cl *changelist.CL) (allSatisfied bool, msg string) {
	reqs := cl.Snapshot.GetGerrit().GetInfo().GetSubmitRequirements()
	if len(reqs) == 0 {
		return
	}
	join := func(ss []string) string {
		var sb strings.Builder
		sb.Grow(len(ss) * 16) // typically, len(submit_requirement.name) is < 16.
		comma := ""
		for _, s := range ss {
			fmt.Fprintf(&sb, "%s`%s`", comma, s)
			comma = ", "
		}
		return sb.String()
	}

	switch satisfied, unsatisfied := groupSubmitReqs(ctx, reqs); {
	// all were NOT_APPLICABLE?
	// just log the occurrence, but consider that
	// submit requirements agreed with Submittable.
	case len(satisfied) == 0 && len(unsatisfied) == 0:
		logging.Errorf(ctx, "CL(%d): all submit reqs(%d) are NOT_APPLICABLE", cl.ID, len(reqs))

	// all satisfied?
	// These are the cases where submit reqs and submittable DISAGREE with each other.
	case len(unsatisfied) == 0 && len(satisfied) == 1:
		msg = fmt.Sprintf("`%s` is satisfied, but the CL is not approved", satisfied[0])
		allSatisfied = true
	case len(unsatisfied) == 0 && len(satisfied) > 1:
		msg = fmt.Sprintf("%s are satisfied, but the CL is not approved", join(satisfied))
		allSatisfied = true

	case len(unsatisfied) == 1:
		msg = fmt.Sprintf("`%s` is not satisfied", unsatisfied[0])
	default:
		msg = fmt.Sprintf("%s are not satisfied", join(unsatisfied))
	}

	if allSatisfied {
		logging.Errorf(ctx, "CL(%d): all submit reqs satisfied; but CL not submittable", cl.ID)
	}
	return
}

func groupSubmitReqs(ctx context.Context, reqs []*gerritpb.SubmitRequirementResultInfo) (satisfied, unsatisfied []string) {
	if len(reqs) == 0 {
		return
	}
	satisfied = make([]string, 0, len(reqs))
	unsatisfied = make([]string, 0, len(reqs))
	for _, req := range reqs {
		switch req.Status {
		case gerritpb.SubmitRequirementResultInfo_SUBMIT_REQUIREMENT_STATUS_UNSPECIFIED:
			panic(errors.New("Unspecified SubmitRequirement.Status; this should never happen"))
		case gerritpb.SubmitRequirementResultInfo_NOT_APPLICABLE:

		// satisfied stauses
		case gerritpb.SubmitRequirementResultInfo_SATISFIED,
			gerritpb.SubmitRequirementResultInfo_OVERRIDDEN,
			gerritpb.SubmitRequirementResultInfo_FORCED:
			satisfied = append(satisfied, req.Name)

		// unsatified statuses
		case gerritpb.SubmitRequirementResultInfo_ERROR:
			// log the error. It may be helpful for diagnosing the reason of a Run rejection.
			logging.Warningf(ctx, "Gerrit reported SubmissionRequirement error %s", req)
			fallthrough
		case gerritpb.SubmitRequirementResultInfo_UNSATISFIED:
			unsatisfied = append(unsatisfied, req.Name)

		default:
			// This must be a bug in CV.
			//
			// common/api/gerrit returns an error if it receives a Status of which enum
			// doesn't exist in common/proto/gerrit. Hence, if a Status is unknown here,
			// this switch is missing the status, enumerated in common/proto/gerrit.
			logging.Errorf(ctx, "Unknown SubmitRequirementStatus %q", req.GetStatus())
			// Unknown enums are considered as a not-satisfied status.
			unsatisfied = append(unsatisfied, req.Name)
		}
	}
	return
}
