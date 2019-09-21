// Copyright 2019 The LUCI Authors.
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

package formats

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	rdb "go.chromium.org/luci/results"
	resultspb "go.chromium.org/luci/results/proto/v1"
	"go.chromium.org/luci/results/util"
)

// GTestResults represents the structure as described to be generated in
// https://cs.chromium.org/chromium/src/base/test/launcher/test_results_tracker.h?l=83&rcl=96020cfd447cb285acfa1a96c37a67ed22fa2499
// (base::TestResultsTracker::SaveSummaryAsJSON)
//
// Fields not used by Test Results are omitted.
type GTestResults struct {
	AllTests   []string `json:"all_tests"`
	GlobalTags []string `json:"global_tags"`

	// PerIterationData is a vector of run iterations, each mapping test names to a list of test data.
	PerIterationData []map[string][]*GTestRunResult `json:"per_iteration_data"`

	// TestLocations maps test names to their location in code.
	TestLocations map[string]*Location `json:"test_locations"`
}

// GTestRunResult represents the per_iteration_data as described in
// https://cs.chromium.org/chromium/src/base/test/launcher/test_results_tracker.h?l=83&rcl=96020cfd447cb285acfa1a96c37a67ed22fa2499
// (base::TestResultsTracker::SaveSummaryAsJSON)
//
// Fields not used by Test Results are omitted.
type GTestRunResult struct {
	Status        string `json:"status"`
	ElapsedTimeMs int    `json:"elapsed_time_ms"`

	LosslessSnippet     bool   `json:"losless_snippet"`
	OutputSnippetBase64 string `json:"output_snippet_base64"`
}

// Location describes a code location.
type Location struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

// ConvertFromJSON reads the provided reader into the receiver.
func (r *GTestResults) ConvertFromJSON(ctx context.Context, reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(r); err != nil {
		return err
	}

	if len(r.AllTests) == 0 {
		return errors.Reason(`missing "all_tests" field in JSON`).Err()
	}

	return nil
}

// ToInvocation converts a GTestResults to a (partial) resultspb.Invocation.
func (r *GTestResults) ToInvocation(ctx context.Context, req *resultspb.DeriveInvocationFromSwarmingRequest) (*resultspb.Invocation, error) {
	// testsToVariants maps GTest base name to the variant IDs present to Results.
	// We store these maps to handle the same test and variants appearing in multiple iterations.
	testsToVariants := map[string]map[string][]*resultspb.TestResult{}
	variantDefs := map[string]*resultspb.VariantDef{}

	// In theory, we can have multiple iterations. This seems rare in practice, so log if we do see
	// more than one to confirm and track.
	if len(r.PerIterationData) > 1 {
		logging.Infof(ctx,
			"Got %d iterations for task %s on %s", len(r.PerIterationData), req.Task.Id, req.Task.Hostname)
	}

	// Assume the invocation is complete for now; if any results are NOTRUN, we'll mark as otherwise.
	incomplete := false

	for i, data := range r.PerIterationData {
		for name, results := range data {
			baseName, params := extractGTestParameters(name)

			// Derive the variant definition from base variant and params.
			variantDefMap := rdb.MergeTestVariantMaps(req.BaseTestVariant.Def, params)
			if err := variantDefMap.Validate(); err != nil {
				return nil, errors.Annotate(err,
					"iteration %d of test %s failed to validate variant def", i, name).Err()
			}
			variantDef := variantDefMap.Proto()
			varID := variantDef.Digest
			if _, ok := variantDefs[varID]; !ok {
				variantDefs[varID] = variantDef
			}

			// Generate individual result protos.
			rpbs := make([]*resultspb.TestResult, len(results))
			for i, result := range results {
				if rpb, err := r.convertRunResult(ctx, name, result); err != nil {
					return nil, errors.Annotate(err,
						"iteration %d of test %s failed to convert run result", i, name).Err()

				} else {
					rpbs[i] = rpb
				}

				// TODO(jchinlee): Verify that it's indeed the case that getting NOTRUN results in the final
				// results indicates the task was incomplete.
				// TODO(jchinlee): Check how unexpected SKIPPED tests should be handled.
				if result.Status == "NOTRUN" {
					incomplete = true
				}
			}

			// Store the processed test result into the correct part of the overall map.
			testPath := req.TestPathPrefix + baseName
			if _, ok := testsToVariants[testPath]; !ok {
				testsToVariants[testPath] = map[string][]*resultspb.TestResult{}
			}
			testsToVariants[testPath][varID] = append(testsToVariants[testPath][varID], rpbs...)
		}
	}

	// Now that we have all the results mapped by test path and variant, populate the invocation.
	inv := &resultspb.Invocation{
		Incomplete:  incomplete,
		VariantDefs: variantDefs,
		Tags:        []*resultspb.StringPair{util.StringPair("test_framework", "gtest")},
	}

	for testPath, variants := range testsToVariants {
		testpb := &resultspb.Invocation_Test{Path: testPath}

		for id, results := range variants {
			testpb.Variants = append(testpb.Variants, &resultspb.Invocation_TestVariant{
				VariantId: id,
				Results:   results,
			})
		}

		inv.Tests = append(inv.Tests, testpb)
	}
	sort.Slice(inv.Tests, func(i, j int) bool { return inv.Tests[i].Path < inv.Tests[j].Path })

	// Populate the tags.
	for _, tag := range r.GlobalTags {
		inv.Tags = append(inv.Tags, util.StringPair("gtest_global_tag", tag))
	}
	SortTagsInPlace(inv.Tags)

	return inv, nil
}

func toGTestStatus(s string) (resultspb.Status, error) {
	switch s {
	case "SUCCESS":
		return resultspb.Status_PASS, nil
	case "FAILURE":
		return resultspb.Status_FAIL, nil
	case "FAILURE_ON_EXIT":
		return resultspb.Status_FAIL, nil
	case "TIMEOUT":
		return resultspb.Status_ABORT, nil
	case "CRASH":
		return resultspb.Status_CRASH, nil
	case "SKIPPED":
		return resultspb.Status_SKIP, nil
	case "EXCESSIVE_OUTPUT":
		return resultspb.Status_FAIL, nil
	case "NOTRUN":
		return resultspb.Status_SKIP, nil
	default:
		// This would only happen if the set of possible GTest result statuses change and resultsdb has
		// not been updated to match.
		return resultspb.Status_STATUS_UNSPECIFIED, errors.Reason("unknown GTest status %q", s).Err()
	}
}

// extractGTestParameters extracts parameters from a test path as a mapping with "param/" keys.
func extractGTestParameters(testPath string) (basePath string, params rdb.VariantDefMap) {
	// TODO(jchinlee): Implement.
	basePath = testPath
	params = rdb.VariantDefMap{}
	return
}

func (r *GTestResults) convertRunResult(ctx context.Context, name string, result *GTestRunResult) (*resultspb.TestResult, error) {
	status, err := toGTestStatus(result.Status)
	if err != nil {
		return nil, err
	}

	rpb := &resultspb.TestResult{
		Status:   status,
		Tags:     []*resultspb.StringPair{util.StringPair("gtest_status", result.Status)},
		Duration: secondsToDuration(1e-6 * float64(result.ElapsedTimeMs)),
	}

	// Write the summary.
	if result.OutputSnippetBase64 != "" {
		outputBytes, err := base64.StdEncoding.DecodeString(result.OutputSnippetBase64)
		if err != nil {
			// Log the error, but we shouldn't fail to convert an entire invocation just because we can't
			// convert a summary.
			logging.Errorf(ctx, "Failed to convert OutputSnippetBase64 %q", result.OutputSnippetBase64)
		} else {
			// TODO(jchinlee): Escape Markdown.
			rpb.SummaryMarkdown = strings.ToValidUTF8(string(outputBytes), string(utf8.RuneError))
		}
	}

	// Store the correct output snippet.
	rpb.Tags = append(
		rpb.Tags,
		util.StringPair("lossless_snippet", strconv.FormatBool(result.LosslessSnippet)),
	)

	// Store the test code location.
	if loc, ok := r.TestLocations[name]; ok {
		rpb.Tags = append(rpb.Tags,
			util.StringPair("file", loc.File),
			util.StringPair("line", strconv.Itoa(loc.Line)),
		)
	}

	SortTagsInPlace(rpb.Tags)

	return rpb, nil
}
