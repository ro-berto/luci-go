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

package triager

import (
	"fmt"
	"time"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/prjmanager/prjpb"
	"go.chromium.org/luci/cv/internal/run"
)

// triageCLs decides whether individual CLs ought to be acted upon.
func triageCLs(c *prjpb.Component, pm pmState) map[int64]*clInfo {
	cls := make(map[int64]*clInfo, len(c.GetClids()))
	for _, clid := range c.GetClids() {
		cls[clid] = &clInfo{
			pcl:       pm.MustPCL(clid),
			purgingCL: pm.PurgingCL(clid), // may be nil
		}
	}
	for index, r := range c.GetPruns() {
		for _, clid := range r.GetClids() {
			info := cls[clid]
			info.runIndexes = append(info.runIndexes, int32(index))
		}
	}
	for _, info := range cls {
		info.triage(c, pm)
	}
	return cls
}

// clInfo represents a CL in the PM component of CLs.
type clInfo struct {
	pcl *prjpb.PCL
	// runIndexes are indexes of Component.PRuns which references this CL.
	runIndexes []int32
	// purgingCL is set if CL is already being purged.
	purgingCL *prjpb.PurgingCL

	triagedCL
}

// lastCQVoteTriggered returns the last triggered time by CQ vote among this CL
// and its triggered deps. Can be zero time.Time if neither are triggered.
func (info *clInfo) lastCQVoteTriggered() time.Time {
	t := info.pcl.GetTriggers().GetCqVoteTrigger()
	thisPB := t.GetTime()
	switch {
	case thisPB == nil && info.deps == nil:
		return time.Time{}
	case thisPB == nil:
		return info.deps.lastCQVoteTriggered
	case info.deps == nil || info.deps.lastCQVoteTriggered.IsZero():
		return thisPB.AsTime()
	default:
		this := thisPB.AsTime()
		if info.deps.lastCQVoteTriggered.Before(this) {
			return this
		}
		return info.deps.lastCQVoteTriggered
	}
}

// triagedCL is the result of CL triage (see clInfo.triage()).
//
// Note: This doesn't take into account `combine_cls.stabilization_delay`,
// thus a CL may be ready or with purgeReason, but due to stabilization delay,
// it shouldn't be acted upon *yet*.
type triagedCL struct {
	// deps are triaged deps, set only if CL is watched by exactly 1 config group.
	// of the current project.
	deps *triagedDeps
	// purgeReasons is set if the CL ought to be purged.
	//
	// Not set if CL is .purgingCL is non-nil since CL is already being purged.
	purgeReasons []*prjpb.PurgeReason
	// cqReady is true if it can be used in creation of new CQ-Vote Runs.
	//
	// If true, purgeReason must be nil, and deps must be OK though they may contain
	// not-yet-loaded deps.
	cqReady bool

	// nprReady is true if it can be used in the creation of a new patchset
	// run.
	nprReady bool
}

func isCQVotePurging(purgingCL *prjpb.PurgingCL) bool {
	return purgingCL.GetTriggers().GetCqVoteTrigger() != nil || purgingCL.GetAllActiveTriggers()
}

func isNPRVotePurging(purgingCL *prjpb.PurgingCL) bool {
	return purgingCL.GetTriggers().GetNewPatchsetRunTrigger() != nil || purgingCL.GetAllActiveTriggers()
}

func (info *clInfo) prunCountByType(c *prjpb.Component) (int, int) {
	var nCQVoteRuns, nNewPatchsetRuns int
	for _, i := range info.runIndexes {
		switch mode := run.Mode(c.Pruns[i].GetMode()); mode {
		case run.NewPatchsetRun:
			nNewPatchsetRuns++
		case run.DryRun, run.FullRun, run.QuickDryRun, "":
			nCQVoteRuns++
		default:
			panic(fmt.Errorf("unsupported runmode: %s", mode))
		}
	}
	return nCQVoteRuns, nNewPatchsetRuns
}

// triage sets the triagedCL part of clInfo.
//
// Expects non-triagedCL part of clInfo to be already set.
// panics iff component is not in a valid state.
func (info *clInfo) triage(c *prjpb.Component, pm pmState) {
	nCQVoteRuns, nNewPatchsetRuns := info.prunCountByType(c)
	var triageCQTrigger, triageNPRTrigger bool
	switch {
	case nCQVoteRuns > 0:
		// Once CV supports API-based triggering, a CL may be both in purged
		// state and have an incomplete Run for the same type of trigger at the
		// same time. The presence in a Run is more important, so treat it as
		// such.
		info.triageInCQVoteRun(pm)
	case isCQVotePurging(info.purgingCL):
		info.triageInCQVotePurge(pm)
	case info.pcl.GetTriggers().GetCqVoteTrigger() != nil:
		triageCQTrigger = true
	}

	switch {
	case nNewPatchsetRuns > 0:
		info.triageInNewPatchsetRun(pm)
	case isNPRVotePurging(info.purgingCL):
		info.triageInNewPatchsetPurge(pm)
	case info.pcl.GetTriggers().GetNewPatchsetRunTrigger() != nil:
		triageNPRTrigger = true
	}
	info.triageNewTriggers(pm, triageCQTrigger, triageNPRTrigger)
}

func (info *clInfo) triageInCQVoteRun(pm pmState) {
	if !info.pcl.GetSubmitted() && info.pclStatusReadyForTriage() && info.pcl.GetTriggers().GetCqVoteTrigger() != nil {
		pcl := info.pcl
		if len(pcl.GetConfigGroupIndexes()) != 1 {
			// This is expected if project config has changed, but Run's reaction to it
			// via OnRunFinished event hasn't yet reached PM.
			return
		}
		cgIndex := pcl.GetConfigGroupIndexes()[0]
		info.deps = triageDeps(pcl, cgIndex, pm)
		// A purging CL must not be "ready" to avoid creating new Runs with them.
		if info.deps.OK() && !isCQVotePurging(info.purgingCL) {
			info.cqReady = true
		}
	}
}

func (info *clInfo) triageInNewPatchsetRun(pm pmState) {
	if len(info.pcl.GetConfigGroupIndexes()) != 1 {
		// This is expected if project config has changes, but Run's reation to
		// it via OnRunFinished event has not yet reached PM.
		return
	}
	if !info.pcl.GetSubmitted() && info.pclStatusReadyForTriage() && info.pcl.GetTriggers().GetNewPatchsetRunTrigger() != nil &&
		!isNPRVotePurging(info.purgingCL) {
		info.nprReady = true
	}
}

func (info *clInfo) pclStatusReadyForTriage() bool {
	switch s := info.pcl.GetStatus(); s {
	case prjpb.PCL_DELETED, prjpb.PCL_UNWATCHED, prjpb.PCL_UNKNOWN:
		return false
	case prjpb.PCL_OK:
		return true
	default:
		panic(fmt.Errorf("PCL has unrecognized status %s", s))
	}
}

func (info *clInfo) triageInCQVotePurge(pm pmState) {
	// The PM hasn't noticed yet the completion of the async purge.
	// The result of purging is modified CL, which may be observed by PM earlier
	// than completion of purge.
	//
	// Thus, consider these CLs in potential Run Creation, but don't mark them
	// ready in order to avoid creating new Runs.
	if !info.pcl.GetSubmitted() && info.pclStatusReadyForTriage() && info.pcl.Triggers.GetCqVoteTrigger() != nil {
		cgIndexes := info.pcl.GetConfigGroupIndexes()
		switch len(cgIndexes) {
		case 0:
			panic(fmt.Errorf("PCL %d without ConfigGroup index not possible for CL not referenced by any Runs (partitioning bug?)", info.pcl.GetClid()))
		case 1:
			info.deps = triageDeps(info.pcl, cgIndexes[0], pm)
			// info.deps.OK() may be true, for example if user has already corrected the
			// mistake that previously resulted in purging op. However, don't mark CL
			// ready until purging op completes or expires.
		}
	}
}

func (info *clInfo) triageInNewPatchsetPurge(pm pmState) {
	// The PM hasn't noticed yet the completion of the async purge.
	// The result of purging is modified CL, which may be observed by PM earlier
	// than completion of purge.
	//
	// Thus, consider these CLs in potential Run Creation, but don't mark them
	// ready in order to avoid creating new Runs.
	if !info.pcl.GetSubmitted() && info.pclStatusReadyForTriage() && info.pcl.Triggers.GetNewPatchsetRunTrigger() != nil {
		if len(info.pcl.GetConfigGroupIndexes()) == 0 {
			panic(fmt.Errorf("PCL %d without ConfigGroup index not possible for CL not referenced by any Runs (partitioning bug?)", info.pcl.GetClid()))
		}
	}
}

func (info *clInfo) addPurgeReason(t *run.Trigger, clError *changelist.CLError) {
	switch {
	case t == nil:
		info.purgeReasons = append(info.purgeReasons, &prjpb.PurgeReason{
			ClError: clError,
			ApplyTo: &prjpb.PurgeReason_AllActiveTriggers{
				AllActiveTriggers: true,
			},
		})
	case run.Mode(t.Mode) == run.NewPatchsetRun:
		info.purgeReasons = append(info.purgeReasons, &prjpb.PurgeReason{
			ClError: clError,
			ApplyTo: &prjpb.PurgeReason_Triggers{
				Triggers: &run.Triggers{
					NewPatchsetRunTrigger: t,
				},
			},
		})
	case (run.Mode(t.Mode) == run.DryRun ||
		run.Mode(t.Mode) == run.FullRun ||
		run.Mode(t.Mode) == run.QuickDryRun):
		info.purgeReasons = append(info.purgeReasons, &prjpb.PurgeReason{
			ClError: clError,
			ApplyTo: &prjpb.PurgeReason_Triggers{
				Triggers: &run.Triggers{
					CqVoteTrigger: t,
				},
			},
		})
	default:
		panic(fmt.Sprintf("impossible Run mode %q", t.GetMode()))
	}
}

func (info *clInfo) triageNewTriggers(pm pmState, triageCQTrigger, triageNPRTrigger bool) {
	pcl := info.pcl
	for _, r := range pcl.GetPurgeReasons() {
		switch {
		case r.GetAllActiveTriggers():
			triageCQTrigger, triageNPRTrigger = false, false
		case r.GetTriggers().GetNewPatchsetRunTrigger() != nil:
			triageNPRTrigger = false
		case r.GetTriggers().GetCqVoteTrigger() != nil:
			triageCQTrigger = false
		}
	}
	info.purgeReasons = append(info.purgeReasons, pcl.GetPurgeReasons()...)
	if !triageCQTrigger && !triageNPRTrigger {
		return
	}
	clid := pcl.GetClid()
	assumption := "not possible for CL not referenced by any Runs (partitioning bug?)"
	switch s := pcl.GetStatus(); s {
	case prjpb.PCL_DELETED, prjpb.PCL_UNWATCHED, prjpb.PCL_UNKNOWN:
		panic(fmt.Errorf("PCL %d status %s %s", clid, s, assumption))
	case prjpb.PCL_OK:
		// OK.
	default:
		panic(fmt.Errorf("PCL has unrecognized status %s", s))
	}

	if pcl.GetSubmitted() {
		panic(fmt.Errorf("PCL %d submitted %s", clid, assumption))
	}

	cgIndexes := pcl.GetConfigGroupIndexes()
	switch len(cgIndexes) {
	case 0:
		panic(fmt.Errorf("PCL %d without ConfigGroup index %s", clid, assumption))
	case 1:
		// if either trigger is being purged, do not mark it as ready.
		if triageCQTrigger {
			if info.deps = triageDeps(pcl, cgIndexes[0], pm); info.deps.OK() {
				info.cqReady = true
			} else {
				info.addPurgeReason(info.pcl.Triggers.GetCqVoteTrigger(), info.deps.makePurgeReason())
			}
		}
		if triageNPRTrigger {
			info.nprReady = true
		}
	default:
		cgNames := make([]string, len(cgIndexes))
		for i, idx := range cgIndexes {
			cgNames[i] = pm.ConfigGroup(idx).ID.Name()
		}
		var purgeTrigger *run.Trigger
		switch {
		case triageCQTrigger && triageNPRTrigger:
			purgeTrigger = nil // purge whole CL
		case triageCQTrigger:
			purgeTrigger = pcl.GetTriggers().GetCqVoteTrigger()
		case triageNPRTrigger:
			purgeTrigger = pcl.GetTriggers().GetNewPatchsetRunTrigger()
		}
		info.addPurgeReason(purgeTrigger, &changelist.CLError{
			Kind: &changelist.CLError_WatchedByManyConfigGroups_{
				WatchedByManyConfigGroups: &changelist.CLError_WatchedByManyConfigGroups{
					ConfigGroups: cgNames,
				},
			},
		})
	}
}
