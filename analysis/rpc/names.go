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

package rpc

import (
	"fmt"
	"regexp"

	"go.chromium.org/luci/common/errors"

	"go.chromium.org/luci/analysis/internal/clustering"
	"go.chromium.org/luci/analysis/internal/clustering/rules"
	"go.chromium.org/luci/analysis/internal/config"
)

// Regular expressions for matching resource names used in APIs.
var (
	GenericKeyPattern = "[a-z0-9\\-]+"
	RuleNameRe        = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)/rules/(` + rules.RuleIDRePattern + `)$`)
	// ClusterNameRe performs partial validation of a cluster resource name.
	// Cluster algorithm and ID must be further validated by
	// ClusterID.Validate().
	ClusterNameRe = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)/clusters/(` + GenericKeyPattern + `)/(` + GenericKeyPattern + `)$`)
	// ClusterFailuresNameRe performs a partial validation of the resource
	// name for a cluster's failures.
	// Cluster algorithm and ID must be further validated by
	// ClusterID.Validate().
	ClusterFailuresNameRe      = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)/clusters/(` + GenericKeyPattern + `)/(` + GenericKeyPattern + `)/failures$`)
	ProjectNameRe              = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)$`)
	ProjectConfigNameRe        = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)/config$`)
	ReclusteringProgressNameRe = regexp.MustCompile(`^projects/(` + config.ProjectRePattern + `)/reclusteringProgress$`)
)

// parseRuleName parses a rule resource name into its constituent ID parts.
func parseRuleName(name string) (project, ruleID string, err error) {
	match := RuleNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", "", errors.New("invalid rule name, expected format: projects/{project}/rules/{rule_id}")
	}
	return match[1], match[2], nil
}

// parseProjectName parses a project resource name into a project ID.
func parseProjectName(name string) (project string, err error) {
	match := ProjectNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", errors.New("invalid project name, expected format: projects/{project}")
	}
	return match[1], nil
}

// parseProjectConfigName parses a project config resource name into a project ID.
func parseProjectConfigName(name string) (project string, err error) {
	match := ProjectConfigNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", errors.New("invalid project config name, expected format: projects/{project}/config")
	}
	return match[1], nil
}

// parseReclusteringProgressName parses a reclustering progress resource name
// into its constituent project ID part.
func parseReclusteringProgressName(name string) (project string, err error) {
	match := ReclusteringProgressNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", errors.New("invalid reclustering progress name, expected format: projects/{project}/reclusteringProgress")
	}
	return match[1], nil
}

// parseClusterName parses a cluster resource name into its constituent ID
// parts. Algorithm aliases are resolved to concrete algorithm names.
func parseClusterName(name string) (project string, clusterID clustering.ClusterID, err error) {
	match := ClusterNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", clustering.ClusterID{}, errors.New("invalid cluster name, expected format: projects/{project}/clusters/{cluster_alg}/{cluster_id}")
	}
	algorithm := resolveAlgorithm(match[2])
	id := match[3]
	cID := clustering.ClusterID{Algorithm: algorithm, ID: id}
	if err := cID.Validate(); err != nil {
		return "", clustering.ClusterID{}, errors.Annotate(err, "invalid cluster identity").Err()
	}
	return match[1], cID, nil
}

// parseClusterFailuresName parses the resource name for a cluster's failures
// into its constituent ID parts. Algorithm aliases are resolved to
// concrete algorithm names.
func parseClusterFailuresName(name string) (project string, clusterID clustering.ClusterID, err error) {
	match := ClusterFailuresNameRe.FindStringSubmatch(name)
	if match == nil {
		return "", clustering.ClusterID{}, errors.New("invalid cluster failures name, expected format: projects/{project}/clusters/{cluster_alg}/{cluster_id}/failures")
	}
	algorithm := resolveAlgorithm(match[2])
	id := match[3]
	cID := clustering.ClusterID{Algorithm: algorithm, ID: id}
	if err := cID.Validate(); err != nil {
		return "", clustering.ClusterID{}, errors.Annotate(err, "invalid cluster identity").Err()
	}
	return match[1], cID, nil
}

func ruleName(project, ruleID string) string {
	return fmt.Sprintf("projects/%s/rules/%s", project, ruleID)
}
