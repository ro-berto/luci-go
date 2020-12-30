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

package common

import (
	"fmt"
	"time"
)

// RunID is an unique RunID to identify a Run in CV.
//
// RunID is string like `luciProject/timeComponent-hexHashDigest` consisting of
// 5 parts:
//   1. The LUCI Project that this Run belongs to.
//      Purpose: separates load on Datastore from different projects.
//   2. `/` separator.
//   3. (`endOfTheWorld` - CreateTime) in ms precision, left-padded with zeros
//      to 13 digits. See `Run.CreateTime` Doc.
//      Purpose: ensures queries by default orders runs of the same project by
//      most recent first.
//   4. `-` separator.
//   5. A hex digest string uniquely identifying the set of CLs involved in
//      this Run.
//      Purpose: ensures two simultaneously started Runs in the same project
//      won't have the same RunID.
type RunID string

// CV will be dead on 2336-10-19T17:46:40Z (10^10s after 2020-01-01T00:00:00Z).
var endOfTheWorld = time.Date(2336, 10, 19, 17, 46, 40, 0, time.UTC)

// LUCIProject this Run belongs to.
func (id RunID) LUCIProject() string {
	for i, c := range id {
		if c == '/' {
			return string(id[:i])
		}
	}
	panic(fmt.Errorf("invalid run ID %q", id))
}

// RunIDs is a convenience type to faciliate handling of run RunIDs.
type RunIDs []RunID

// sort.Interface copy-pasta.
func (ids RunIDs) Less(i, j int) bool { return ids[i] < ids[j] }
func (ids RunIDs) Len() int           { return len(ids) }
func (ids RunIDs) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }

// WithoutSorted returns a subsequence of IDs without excluded IDs.
//
// Both this and the excluded slices must be sorted.
//
// If this and excluded IDs are disjoint, return this slice.
// Otherwise, returns a copy without excluded IDs.
func (ids RunIDs) WithoutSorted(exclude RunIDs) RunIDs {
	remaining := ids
	ret := ids
	mutated := false
	for {
		switch {
		case len(remaining) == 0:
			return ret
		case len(exclude) == 0:
			if mutated {
				ret = append(ret, remaining...)
			}
			return ret
		case remaining[0] < exclude[0]:
			if mutated {
				ret = append(ret, remaining[0])
			}
			remaining = remaining[1:]
		case remaining[0] > exclude[0]:
			exclude = exclude[1:]
		default:
			if !mutated {
				// Must copy all IDs that were skipped.
				mutated = true
				n := len(ids) - len(remaining)
				ret = make(RunIDs, n, len(ids)-1)
				copy(ret, ids) // copies len(ret) == n elements.
			}
			remaining = remaining[1:]
			exclude = exclude[1:]
		}
	}
}

// Equal checks if two assumed-to-be-sorted slices are equal.
func (ids RunIDs) Equal(other RunIDs) bool {
	if len(ids) != len(other) {
		return false
	}
	for i, id := range ids {
		if id != other[i] {
			return false
		}
	}
	return true
}

// MakeRunIDs returns RunIDs from list of strings.
func MakeRunIDs(ids ...string) RunIDs {
	ret := make(RunIDs, len(ids))
	for i, id := range ids {
		ret[i] = RunID(id)
	}
	return ret
}
