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

package metrics

import (
	"context"

	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/metric"
)

// TestResultRowStatus is a status of a test result row.
// Used in test result counter metric.
type TestResultRowStatus string

// Values of TestResultRowStatus type.
const (
	Inserted       TestResultRowStatus = "INSERTED"
	PurgeScheduled TestResultRowStatus = "PURGE_SCHEDULED"
	Purged         TestResultRowStatus = "PURGED"
)

var testResultCounter = metric.NewCounter(
	"resultdb/test_results/count",
	"Number of test results",
	nil,
	field.String("status")) // See TestResultRowStatus type.

// IncTestResultCount increments the test result counter.
func IncTestResultCount(ctx context.Context, count int, rowStatus TestResultRowStatus) {
	testResultCounter.Add(ctx, int64(count), string(rowStatus))
}
