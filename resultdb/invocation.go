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

package resultdb

import (
	"sort"

	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// NormalizeInvocation converts inv to the canonical form.
func NormalizeInvocation(inv *pb.Invocation) {
	sortStringPairs(inv.Tags)
}

// NormalizeTestResult converts inv to the canonical form.
func NormalizeTestResult(tr *pb.TestResult) {
	sortStringPairs(tr.Tags)
}

// NormalizeTestResultSlice converts trs to the canonical form.
func NormalizeTestResultSlice(trs []*pb.TestResult) {
	for _, tr := range trs {
		NormalizeTestResult(tr)
	}
	sort.Slice(trs, func(i, j int) bool {
		a := trs[i]
		b := trs[j]
		if a.TestPath != b.TestPath {
			return a.TestPath < b.TestPath
		}
		return a.Name < b.Name
	})
}

// sortStringPairs sorts in-place the tags slice lexicographically by key, then value.
func sortStringPairs(tags []*pb.StringPair) {
	sort.Slice(tags, func(i, j int) bool {
		if tags[i].Key != tags[j].Key {
			return tags[i].Key < tags[j].Key
		}
		return tags[i].Value < tags[j].Value
	})
}
