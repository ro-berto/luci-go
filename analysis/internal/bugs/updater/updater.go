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

package updater

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	"go.chromium.org/luci/analysis/internal/analysis"
	"go.chromium.org/luci/analysis/internal/bugs"
	"go.chromium.org/luci/analysis/internal/clustering"
	"go.chromium.org/luci/analysis/internal/clustering/algorithms"
	"go.chromium.org/luci/analysis/internal/clustering/algorithms/rulesalgorithm"
	"go.chromium.org/luci/analysis/internal/clustering/rules"
	"go.chromium.org/luci/analysis/internal/clustering/rules/lang"
	"go.chromium.org/luci/analysis/internal/clustering/runs"
	"go.chromium.org/luci/analysis/internal/config/compiledcfg"
	pb "go.chromium.org/luci/analysis/proto/v1"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server/span"
)

// testnameThresholdInflationPercent is the percentage factor by which
// the bug filing threshold is inflated when applied to test-name clusters.
// This is to bias bug-filing towards failure reason clusters, which are
// seen as generally better scoped and more actionable (because they
// focus on one reason for the test failing.)
//
// The value of 34% was selected as it is sufficient to inflate any threshold
// values which are a '3' (e.g. CV runs rejected) to a '4'. Otherwise integer
// discretization of the statistics would cancel out any intended bias.
//
// If changing this value, please also update the comment in
// project_config.proto.
const testnameThresholdInflationPercent = 34

// mergeIntoCycleErr is the error returned if a cycle is detected in a bug's
// merged-into graph when handling a bug marked as duplicate.
var mergeIntoCycleErr = errors.New("a cycle was detected in the bug merged-into graph")

// ruleDefinitionTooLongErr is the error returned if merging two failure
// association rules results in a rule that is too long.
var ruleDefinitionTooLongErr = errors.New("the merged rule definition is too long")

// mergeIntoCycleMessage is the message posted on bugs when LUCI Analysis
// cannot deal with a bug marked as the duplicate of another because of
// a duplicate bug.
const mergeIntoCycleMessage = "LUCI Analysis cannot merge the failure" +
	" association rule for this bug into the rule for the merged-into bug," +
	" because a cycle was detected in the bug merged-into graph. Please" +
	" manually resolve the cycle, or update rules manually and archive the" +
	" rule for this bug."

// ruleDefinitionTooLongMessage is the message posted on bugs when
// LUCI Analysis cannot deal with a bug marked as the duplicate of another
// because the merged rule would be too long.
const ruleDefinitionTooLongMessage = "LUCI Analysis cannot merge the failure" +
	" association rule for this bug into the rule for the merged-into bug," +
	" because the merged failure association rule would be too long. Please" +
	" manually update the rule for the merged-into bug and archive the" +
	" rule for this bug."

// BugManager implements bug creation and bug updates for a bug-tracking
// system. The BugManager determines bug content and priority given a
// cluster.
type BugManager interface {
	// Create creates a new bug for the given request, returning its name,
	// or any encountered error.
	Create(ctx context.Context, cluster *bugs.CreateRequest) (string, error)
	// Update updates the specified list of bugs.
	Update(ctx context.Context, bugs []bugs.BugUpdateRequest) ([]bugs.BugUpdateResponse, error)
	// GetMergedInto reads the bug the given bug is merged into (if any).
	// This is to allow step-wise discovery of the canonical bug a bug
	// is merged into (if it exists and there is no cycle in the bug
	// merged-into graph).
	GetMergedInto(ctx context.Context, bug bugs.BugID) (*bugs.BugID, error)
	// UpdateDuplicateSource updates the source bug of a duplicate
	// bug relationship.
	// It normally posts a message advising the user LUCI Analysis
	// has merged the rule for the source bug to the destination
	// (merged-into) bug, and provides a new link to the failure
	// association rule.
	// If a cycle was detected, it instead posts a message that the
	// duplicate bug could not be handled and marks the bug no
	// longer a duplicate to break the cycle.
	UpdateDuplicateSource(ctx context.Context, request bugs.UpdateDuplicateSourceRequest) error
	// UpdateDuplicateDestination updates the destination bug of a duplicate
	// bug relationship.
	// It posts a message advising the user LUCI Analysis
	// has merged the rule for the source bug to the destination
	// (merged-into) bug, and provides a link to the failure
	// association rule.
	UpdateDuplicateDestination(ctx context.Context, destinationBug bugs.BugID) error
}

// BugUpdater performs updates to Monorail bugs and failure association
// rules to keep them in sync with clusters generated by analysis.
type BugUpdater struct {
	// project is the LUCI project to act on behalf of.
	project string
	// analysisClient provides access to cluster analysis.
	analysisClient AnalysisClient
	// managers stores the manager responsible for updating bugs for each
	// bug tracking system (monorail, buganizer, etc.).
	managers map[string]BugManager
	// projectCfg is the snapshot of project configuration to use for
	// the auto-bug filing run.
	projectCfg *compiledcfg.ProjectConfig
	// MaxBugsFiledPerRun is the maximum number of bugs to file each time
	// BugUpdater runs. This throttles the rate of changes to monorail.
	MaxBugsFiledPerRun int
}

// NewBugUpdater initialises a new BugUpdater. The specified impact thresholds are used
// when determining whether to a file a bug.
func NewBugUpdater(project string, mgrs map[string]BugManager, ac AnalysisClient, projectCfg *compiledcfg.ProjectConfig) *BugUpdater {
	return &BugUpdater{
		project:            project,
		managers:           mgrs,
		analysisClient:     ac,
		projectCfg:         projectCfg,
		MaxBugsFiledPerRun: 1, // Default value.
	}
}

// Run updates files/updates bugs to match high-impact clusters as
// identified by analysis. Each bug has a corresponding failure association
// rule.
// The passed progress should reflect the progress of re-clustering as captured
// in the latest analysis.
func (b *BugUpdater) Run(ctx context.Context, progress *runs.ReclusteringProgress) error {
	// Verify we are not currently reclustering to a new version of
	// algorithms or project configuration. If we are, we should
	// suspend bug creation, priority updates and auto-closure
	// as cluster impact is unreliable.
	impactValid := b.verifyClusterImpactValid(ctx, progress)

	rules, err := rules.ReadActive(span.Single(ctx), b.project)
	if err != nil {
		return errors.Annotate(err, "read active failure association rules").Err()
	}

	impactByRuleID := make(map[string]*bugs.ClusterImpact)
	if impactValid {
		// We want to read analysis for two categories of clusters:
		// - Bug Clusters: to update the priority of filed bugs.
		// - Impactful Suggested Clusters: if any suggested clusters may be
		//    near the threshold to file a new bug for, we want to
		//    read them, so we can file a bug. (Note: the thresholding applied
		//    here is weaker than the actual bug filing criteria which is
		//    implemented in this package, it exists mainly to avoid pulling
		//    back all suggested clusters).
		clusters, err := b.analysisClient.ReadImpactfulClusters(ctx, analysis.ImpactfulClusterReadOptions{
			Project:                  b.project,
			Thresholds:               b.projectCfg.Config.BugFilingThreshold,
			AlwaysIncludeBugClusters: true,
		})
		if err != nil {
			return errors.Annotate(err, "read impactful clusters").Err()
		}

		// blockedSourceClusterIDs is the set of source cluster IDs for which
		// filing new bugs should be suspended.
		blockedSourceClusterIDs := make(map[string]struct{})
		for _, r := range rules {
			if !progress.IncorporatesRulesVersion(r.CreationTime) {
				// If a bug cluster was recently filed for a source cluster, and
				// re-clustering and analysis is not yet complete (to move the
				// impact from the source cluster to the bug cluster), do not file
				// another bug for the source cluster.
				// (Of course, if a bug cluster was filed for a source cluster,
				// but the bug cluster's failure association rule was subsequently
				// modified (e.g. narrowed), it is allowed to file another bug
				// if the residual impact justifies it.)
				blockedSourceClusterIDs[r.SourceCluster.Key()] = struct{}{}
			}
		}

		if err := b.fileNewBugs(ctx, clusters, blockedSourceClusterIDs); err != nil {
			return err
		}

		for _, cluster := range clusters {
			if cluster.ClusterID.Algorithm == rulesalgorithm.AlgorithmName {
				// Use only impact from latest algorithm version.
				ruleID := cluster.ClusterID.ID
				impactByRuleID[ruleID] = ExtractResidualImpact(cluster)
			}
		}
	}

	// Prepare bug update requests.
	bugUpdatesBySystem := make(map[string][]bugs.BugUpdateRequest)
	for _, r := range rules {
		var impact *bugs.ClusterImpact

		// Impact is valid if re-clustering and analysis ran on the latest
		// version of this failure association rule. This avoids bugs getting
		// erroneous priority changes while impact information is incomplete.
		ruleImpactValid := impactValid &&
			progress.IncorporatesRulesVersion(r.PredicateLastUpdated)

		if ruleImpactValid {
			var ok bool
			impact, ok = impactByRuleID[r.RuleID]
			if !ok {
				// If there is no analysis, this means the cluster is
				// empty. Use empty impact.
				impact = &bugs.ClusterImpact{}
			}
		}
		// Else leave impact as nil. Bug-updating code takes this as an
		// indication valid impact is not available and will not attempt
		// priority updates/auto-closure.

		bugUpdates := bugUpdatesBySystem[r.BugID.System]
		bugUpdates = append(bugUpdates, bugs.BugUpdateRequest{
			Bug:           r.BugID,
			Impact:        impact,
			IsManagingBug: r.IsManagingBug,
			RuleID:        r.RuleID,
		})
		bugUpdatesBySystem[r.BugID.System] = bugUpdates
	}

	var duplicateBugs []bugs.BugID
	var ruleIDsToArchive []string

	// Perform bug updates.
	for system, bugsToUpdate := range bugUpdatesBySystem {
		if system == bugs.BuganizerSystem {
			// Updating buganizer bugs is currently not supported. This is a
			// known limitation.
			continue
		}
		manager, ok := b.managers[system]
		if !ok {
			logging.Warningf(ctx, "Encountered bug(s) with an unrecognised manager: %q", system)
			continue
		}
		responses, err := manager.Update(ctx, bugsToUpdate)
		if err != nil {
			return err
		}

		for i, rsp := range responses {
			if rsp.IsDuplicate {
				duplicateBugs = append(duplicateBugs, bugsToUpdate[i].Bug)
			} else if rsp.ShouldArchive {
				ruleIDsToArchive = append(ruleIDsToArchive, bugsToUpdate[i].RuleID)
			}
		}
	}

	// Handle rules which need to be archived because the bugs were:
	// - Verified for >30 days and managed by LUCI Analysis, OR
	// - In any closed state for > 30 days and not managed by LUCI Analysis, OR
	// -
	if err := b.archiveRules(ctx, ruleIDsToArchive); err != nil {
		return errors.Annotate(err, "archive rules").Err()
	}

	// Handle bugs marked as duplicate.
	for _, bug := range duplicateBugs {
		err := b.handleDuplicateBug(ctx, bug)
		if err == mergeIntoCycleErr {
			request := bugs.UpdateDuplicateSourceRequest{
				Bug:          bug,
				ErrorMessage: mergeIntoCycleMessage,
			}
			if err := b.updateDuplicateSource(ctx, request); err != nil {
				return errors.Annotate(err, "update source bug after a cycle was found").Err()
			}
		} else if err == ruleDefinitionTooLongErr {
			request := bugs.UpdateDuplicateSourceRequest{
				Bug:          bug,
				ErrorMessage: ruleDefinitionTooLongMessage,
			}
			if err := b.updateDuplicateSource(ctx, request); err != nil {
				return errors.Annotate(err, "update source bug after merging rule definition was found too long").Err()
			}
		} else if err != nil {
			return errors.Annotate(err, "handling bug (%s) marked as duplicate", bug).Err()
		}
	}

	return nil
}

func (b *BugUpdater) verifyClusterImpactValid(ctx context.Context, progress *runs.ReclusteringProgress) bool {
	if progress.IsReclusteringToNewAlgorithms() {
		logging.Warningf(ctx, "Auto-bug filing paused for project %s as re-clustering to new algorithms is in progress.", b.project)
		return false
	}
	if progress.IsReclusteringToNewConfig() {
		logging.Warningf(ctx, "Auto-bug filing paused for project %s as re-clustering to new configuration is in progress.", b.project)
		return false
	}
	if algorithms.AlgorithmsVersion != progress.Next.AlgorithmsVersion {
		logging.Warningf(ctx, "Auto-bug filing paused for project %s as bug-filing is running mismatched algorithms version %v (want %v).",
			b.project, algorithms.AlgorithmsVersion, progress.Next.AlgorithmsVersion)
		return false
	}
	if !b.projectCfg.LastUpdated.Equal(progress.Next.ConfigVersion) {
		logging.Warningf(ctx, "Auto-bug filing paused for project %s as bug-filing is running mismatched config version %v (want %v).",
			b.project, b.projectCfg.LastUpdated, progress.Next.ConfigVersion)
		return false
	}
	return true
}

// fileNewBugs files new bugs for suggested clusters whose residual impact
// exceed the configured bug-filing threshold. Clusters specified in
// blockedClusterIDs will not have a bug filed. This can be used to
// suppress bug-filing for suggested clusters that have recently had a
// bug filed for them and re-clustering is not yet complete.
func (b *BugUpdater) fileNewBugs(ctx context.Context, clusters []*analysis.Cluster, blockedClusterIDs map[string]struct{}) error {
	sortByBugFilingPreference(clusters)

	var toCreateBugsFor []*analysis.Cluster
	for _, cluster := range clusters {
		if cluster.ClusterID.IsBugCluster() {
			// Never file another bug for a bug cluster.
			continue
		}

		// Was a bug recently filed for this suggested cluster?
		// We want to avoid race conditions whereby we file multiple bug
		// clusters for the same suggested cluster, because re-clustering and
		// re-analysis has not yet run and moved residual impact from the
		// suggested cluster to the bug cluster.
		_, ok := blockedClusterIDs[cluster.ClusterID.Key()]
		if ok {
			// Do not file a bug.
			continue
		}

		// Were the failures are confined to only automation CLs
		// and/or 1-2 user CLs? In other words, are the failures in this
		// clusters unlikely to be present in the tree?
		if cluster.DistinctUserCLsWithFailures7d.Residual < 3 &&
			cluster.PostsubmitBuildsWithFailures7d.Residual == 0 {
			// Do not file a bug.
			continue
		}

		// Only file a bug if the residual impact exceeds the threshold.
		impact := ExtractResidualImpact(cluster)
		bugFilingThreshold := b.projectCfg.Config.BugFilingThreshold
		if cluster.ClusterID.IsTestNameCluster() {
			// Use an inflated threshold for test name clusters to bias
			// bug creation towards failure reason clusters.
			bugFilingThreshold =
				bugs.InflateThreshold(b.projectCfg.Config.BugFilingThreshold,
					testnameThresholdInflationPercent)
		}
		if !impact.MeetsThreshold(bugFilingThreshold) {
			continue
		}

		toCreateBugsFor = append(toCreateBugsFor, cluster)
	}

	// File new bugs.
	bugsFiled := 0
	for _, cluster := range toCreateBugsFor {
		// Throttle how many bugs may be filed each time.
		if bugsFiled >= b.MaxBugsFiledPerRun {
			break
		}
		created, err := b.createBug(ctx, cluster)
		if err != nil {
			return err
		}
		if created {
			bugsFiled++
		}
	}
	return nil
}

// archiveRules archives the given list of rules.
func (b *BugUpdater) archiveRules(ctx context.Context, ruleIDs []string) error {
	if len(ruleIDs) == 0 {
		return nil
	}
	// Limit the number of rules that can be archived at once to stay
	// well within Spanner mutation limits. The rest will be handled
	// in the next bug-filing run.
	if len(ruleIDs) > 100 {
		ruleIDs = ruleIDs[:100]
	}
	f := func(ctx context.Context) error {
		// Perform atomic read-update of rule.
		rs, err := rules.ReadMany(ctx, b.project, ruleIDs)
		if err != nil {
			return errors.Annotate(err, "read rules to archive").Err()
		}
		for _, r := range rs {
			r.IsActive = false
			if err := rules.Update(ctx, r, rules.UpdateOptions{
				PredicateUpdated: true,
			}, rules.LUCIAnalysisSystem); err != nil {
				// Validation error. Actual save happens upon transaction
				// commit.
				return errors.Annotate(err, "update rules").Err()
			}
		}
		return nil
	}
	_, err := span.ReadWriteTransaction(ctx, f)
	return err
}

// handleDuplicateBug handles a duplicate bug, merging its failure association
// rule with the bug it is ultimately merged into (creating the rule if it does
// not exist). The original rule is archived.
func (b *BugUpdater) handleDuplicateBug(ctx context.Context, bug bugs.BugID) error {
	// Chase the bug merged-into graph until we find the sink of the graph.
	// (The canonical bug of the chain of duplicate bugs.)
	destBug, err := b.resolveMergedIntoBug(ctx, bug)
	if err != nil {
		// May return mergeIntoCycleErr.
		return err
	}

	var destinationBugRuleID string

	f := func(ctx context.Context) error {
		sourceRule, _, err := readRuleForBugAndProject(ctx, bug, b.project)
		if err != nil {
			return errors.Annotate(err, "reading rule for source bug").Err()
		}
		if !sourceRule.IsActive {
			// The source rule is no longer active. This is a race condition
			// as we only do bug updates for rules that exist at the time
			// we start bug updates.
			// An inactive rule does not match any failures so merging the
			// it into another rule should have no effect anyway.
			return nil
		}
		// Try and read the rule for the bug we are merging into.
		destinationRule, anyRuleManagingDestBug, err :=
			readRuleForBugAndProject(ctx, destBug, b.project)
		if err != nil {
			return errors.Annotate(err, "reading rule for destination bug").Err()
		}
		if destinationRule == nil {
			// Simply update the source rule to point to the new bug.
			sourceRule.BugID = destBug

			// Only one rule can manage a bug at a given time.
			// Even if there is no rule in this project which manages
			// the destination bug, there could a rule in a different project.
			if anyRuleManagingDestBug {
				sourceRule.IsManagingBug = false
			}

			err = rules.Update(ctx, sourceRule, rules.UpdateOptions{
				PredicateUpdated: false,
			}, rules.LUCIAnalysisSystem)
			if err != nil {
				// Indicates validation error. Should never happen.
				return err
			}

			destinationBugRuleID = sourceRule.RuleID
			return nil
		} else {
			// The bug we are a duplicate of already has a rule.
			if destinationRule.IsActive {
				// Merge the source and destination rules with an "OR".
				mergedRule, err := lang.Merge(destinationRule.RuleDefinition, sourceRule.RuleDefinition)
				if err != nil {
					return errors.Annotate(err, "merging rules").Err()
				}
				if len(mergedRule) > rules.MaxRuleDefinitionLength {
					// The merged rule is too long to store.
					return ruleDefinitionTooLongErr
				}
				destinationRule.RuleDefinition = mergedRule
			} else {
				// Else: an inactive rule does not match any failures, so we should
				// use only the rule from the source bug.
				destinationRule.RuleDefinition = sourceRule.RuleDefinition
			}

			// Disable the source rule.
			sourceRule.IsActive = false
			err = rules.Update(ctx, sourceRule, rules.UpdateOptions{
				PredicateUpdated: true,
			}, rules.LUCIAnalysisSystem)
			if err != nil {
				// Indicates validation error. Should never happen.
				return err
			}

			// Update the rule on the destination rule.
			destinationRule.IsActive = true
			err = rules.Update(ctx, destinationRule, rules.UpdateOptions{
				PredicateUpdated: true,
			}, rules.LUCIAnalysisSystem)
			if err != nil {
				return err
			}

			destinationBugRuleID = destinationRule.RuleID
			return nil
		}
	}
	// Update source and destination rules in one transaction, to ensure
	// consistency.
	_, err = span.ReadWriteTransaction(ctx, f)
	if err != nil {
		return err
	}

	request := bugs.UpdateDuplicateSourceRequest{
		Bug:               bug,
		DestinationRuleID: destinationBugRuleID,
	}
	if err := b.updateDuplicateSource(ctx, request); err != nil {
		return errors.Annotate(err, "updating source bug").Err()
	}
	if err := b.updateDuplicateDestination(ctx, destBug); err != nil {
		return errors.Annotate(err, "updating destination bug %s", destBug).Err()
	}

	return err
}

// resolveMergedIntoBug resolves the bug the given bug is ultimately merged
// into.
func (b *BugUpdater) resolveMergedIntoBug(ctx context.Context, bug bugs.BugID) (bugs.BugID, error) {
	isResolved := false
	mergedIntoBug := bug
	const maxResolutionSteps = 5
	for i := 0; i < maxResolutionSteps; i++ {
		system := mergedIntoBug.System
		if system == bugs.BuganizerSystem {
			// Resolving the canonical "merged into" bug for bugs in
			// buganizer is not supported. We'll merge into the first
			// buganizer bug we see.
			isResolved = true
			break
		}
		manager, ok := b.managers[system]
		if !ok {
			return bugs.BugID{}, fmt.Errorf("encountered unknown bug system: %q", system)
		}
		mergedInto, err := manager.GetMergedInto(ctx, mergedIntoBug)
		if err != nil {
			return bugs.BugID{}, err
		}
		if mergedInto == nil {
			isResolved = true
			break
		} else {
			mergedIntoBug = *mergedInto
		}
	}
	if !isResolved {
		return bugs.BugID{}, mergeIntoCycleErr
	}
	if mergedIntoBug == bug {
		// This would normally never occur, but is possible in some
		// exceptional scenarios like race conditions where a cycle
		// is broken during the graph traversal, or a bug which
		// was marked as duplicate but is no longer marked as duplicate
		// now.
		return bugs.BugID{}, fmt.Errorf("cannot deduplicate a bug into itself")
	}
	return mergedIntoBug, nil
}

// updateDuplicateSource updates the source bug of a duplicate
// bug pair (source bug, destination bug).
// It either posts a message notifying the user the rule was successfully
// merged to the destination, or notifies the user of the error and
// marks the bug no longer a duplicate (to avoid repeated attempts to
// handle the problematic duplicate bug).
func (b *BugUpdater) updateDuplicateSource(ctx context.Context, request bugs.UpdateDuplicateSourceRequest) error {
	manager, ok := b.managers[request.Bug.System]
	if !ok {
		return fmt.Errorf("encountered unknown bug system: %q", request.Bug.System)
	}
	err := manager.UpdateDuplicateSource(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

// updateDuplicateDestination updates the destination bug of a duplicate
// bug pair (source bug, destination bug).
// It posts a message notifying the user the rule was successfully
// merged to the destination.
func (b *BugUpdater) updateDuplicateDestination(ctx context.Context, destinationBug bugs.BugID) error {
	manager, ok := b.managers[destinationBug.System]
	if !ok {
		// Not all destination bug systems need to be supported
		// in order to be able to merge a bug there. We simply
		// won't be able to post an update there.
		return nil
	}
	err := manager.UpdateDuplicateDestination(ctx, destinationBug)
	if err != nil {
		return err
	}
	return nil
}

// readRuleForBugAndProject reads the failure association rule for the given
// bug in the given project, if it exists. It additionally returns whether
// there is any rule in the system that manages the given bug, even if in
// a different project.
// If the rule cannot be read, it returns nil.
func readRuleForBugAndProject(ctx context.Context, bug bugs.BugID, project string) (rule *rules.FailureAssociationRule, anyRuleManaging bool, err error) {
	rules, err := rules.ReadByBug(ctx, bug)
	if err != nil {
		return nil, false, err
	}
	rule = nil
	anyRuleManaging = false
	for _, r := range rules {
		if r.IsManagingBug {
			anyRuleManaging = true
		}
		if r.Project == project {
			rule = r
		}
	}
	return rule, anyRuleManaging, nil
}

// sortByBugFilingPreference sorts clusters based on our preference
// to file bugs for these clusters.
func sortByBugFilingPreference(cs []*analysis.Cluster) {
	// The current ranking approach prefers filing bugs for clusters with more
	// impact, with a bias towards reason clusters.
	//
	// The order of this ranking is only important where there are
	// multiple competing clusters which meet the bug-filing threshold.
	// As bug filing runs relatively often, except in cases of contention,
	// the first bug to meet the threshold will be filed.
	sort.Slice(cs, func(i, j int) bool {
		presubmitRejects := func(cs *analysis.Cluster) analysis.Counts { return cs.PresubmitRejects7d }
		criticalFailuresExonerated := func(cs *analysis.Cluster) analysis.Counts { return cs.CriticalFailuresExonerated7d }
		failures := func(cs *analysis.Cluster) analysis.Counts { return cs.Failures7d }

		if equal, less := rankByMetric(cs[i], cs[j], presubmitRejects); !equal {
			return less
		}
		if equal, less := rankByMetric(cs[i], cs[j], criticalFailuresExonerated); !equal {
			return less
		}
		if equal, less := rankByMetric(cs[i], cs[j], failures); !equal {
			return less
		}
		// If all else fails, sort by cluster ID. This is mostly to ensure
		// the code behaves deterministically when under unit testing.
		if cs[i].ClusterID.Algorithm != cs[j].ClusterID.Algorithm {
			return cs[i].ClusterID.Algorithm < cs[j].ClusterID.Algorithm
		}
		return cs[i].ClusterID.ID < cs[j].ClusterID.ID
	})
}

func rankByMetric(a, b *analysis.Cluster, accessor func(*analysis.Cluster) analysis.Counts) (equal bool, less bool) {
	valueA := accessor(a).Residual
	valueB := accessor(b).Residual
	// If one cluster we are comparing with is a test name cluster,
	// give the other cluster an impact boost in the comparison, so
	// that we bias towards filing it (instead of the test name cluster).
	if b.ClusterID.IsTestNameCluster() {
		valueA = (valueA * (100 + testnameThresholdInflationPercent)) / 100
	}
	if a.ClusterID.IsTestNameCluster() {
		valueB = (valueB * (100 + testnameThresholdInflationPercent)) / 100
	}
	equal = (valueA == valueB)
	// a less than b in the sort order is defined as a having more impact
	// than b, so that clusters are sorted in descending impact order.
	less = (valueA > valueB)
	return equal, less
}

// createBug files a new bug for the given suggested cluster,
// and stores the association from bug to failures through a new
// failure association rule.
func (b *BugUpdater) createBug(ctx context.Context, cs *analysis.Cluster) (created bool, err error) {
	alg, err := algorithms.SuggestingAlgorithm(cs.ClusterID.Algorithm)
	if err == algorithms.ErrAlgorithmNotExist {
		// The cluster is for an old algorithm that no longer exists, or
		// for a new algorithm that is not known by us yet.
		// Do not file a bug. This is not an error, it is expected during
		// algorithm version changes.
		return false, nil
	}

	summary := clusterSummaryFromAnalysis(cs)

	// Double-check the failure matches the cluster. Generating a
	// failure association rule that does not match the suggested cluster
	// could result in indefinite creation of new bugs, as the system
	// will repeatedly create new failure association rules for the
	// same suggested cluster.
	// Mismatches should usually be transient as re-clustering will fix
	// up any incorrect clustering.
	if hex.EncodeToString(alg.Cluster(b.projectCfg, &summary.Example)) != cs.ClusterID.ID {
		return false, errors.New("example failure did not match cluster ID")
	}
	rule, err := b.generateFailureAssociationRule(alg, &summary.Example)
	if err != nil {
		return false, errors.Annotate(err, "obtain failure association rule").Err()
	}

	ruleID, err := rules.GenerateID()
	if err != nil {
		return false, errors.Annotate(err, "generating rule ID").Err()
	}

	description, err := alg.ClusterDescription(b.projectCfg, summary)
	if err != nil {
		return false, errors.Annotate(err, "prepare bug description").Err()
	}

	var monorailComponents []string
	for _, tc := range cs.TopMonorailComponents {
		// Any monorail component is associated for more than 30% of the
		// failures in the cluster should be on the filed bug.
		if tc.Count > ((cs.Failures7d.Nominal * 3) / 10) {
			monorailComponents = append(monorailComponents, tc.Value)
		}
	}
	request := &bugs.CreateRequest{
		Description:        description,
		Impact:             ExtractResidualImpact(cs),
		MonorailComponents: monorailComponents,
	}

	// For now, the only issue system supported is monorail.
	system := bugs.MonorailSystem
	mgr := b.managers[system]
	name, err := mgr.Create(ctx, request)
	if err == bugs.ErrCreateSimulated {
		// Create did not do anything because it is in simulation mode.
		// This is expected.
		return true, nil
	}
	if err != nil {
		return false, errors.Annotate(err, "create issue in %v", system).Err()
	}

	// Create a failure association rule associating the failures with a bug.
	r := &rules.FailureAssociationRule{
		Project:        b.project,
		RuleID:         ruleID,
		RuleDefinition: rule,
		BugID:          bugs.BugID{System: system, ID: name},
		IsActive:       true,
		IsManagingBug:  true,
		SourceCluster:  cs.ClusterID,
	}
	create := func(ctx context.Context) error {
		user := rules.LUCIAnalysisSystem
		return rules.Create(ctx, r, user)
	}
	if _, err := span.ReadWriteTransaction(ctx, create); err != nil {
		return false, errors.Annotate(err, "create bug cluster").Err()
	}

	return true, nil
}

func clusterSummaryFromAnalysis(c *analysis.Cluster) *clustering.ClusterSummary {
	example := clustering.Failure{
		TestID: c.ExampleTestID(),
	}
	if c.ExampleFailureReason.Valid {
		example.Reason = &pb.FailureReason{PrimaryErrorMessage: c.ExampleFailureReason.StringVal}
	}
	// A list of 5 commonly occuring tests are included in bugs created
	// for failure reason clusters, to improve searchability by test name.
	var topTests []string
	for _, tt := range c.TopTestIDs {
		topTests = append(topTests, tt.Value)
	}
	return &clustering.ClusterSummary{
		Example:  example,
		TopTests: topTests,
	}
}

func (b *BugUpdater) generateFailureAssociationRule(alg algorithms.Algorithm, failure *clustering.Failure) (string, error) {
	rule := alg.FailureAssociationRule(b.projectCfg, failure)

	// Check the generated rule is valid and matches the failure.
	// An improperly generated failure association rule could result
	// in uncontrolled creation of new bugs.
	expr, err := lang.Parse(rule)
	if err != nil {
		return "", errors.Annotate(err, "rule generated by %s did not parse", alg.Name()).Err()
	}
	match := expr.Evaluate(failure)
	if !match {
		reason := ""
		if failure.Reason != nil {
			reason = failure.Reason.PrimaryErrorMessage
		}
		return "", fmt.Errorf("rule generated by %s did not match example failure (testID: %q, failureReason: %q)",
			alg.Name(), failure.TestID, reason)
	}
	return rule, nil
}
