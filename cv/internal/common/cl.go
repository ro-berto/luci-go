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

import "sort"

// CLID is a unique ID of a CL used internally in CV.
//
// It's just 8 bytes long and is thus much shorter than ExternalID,
// which reduces CPU & RAM & storage costs of CL graphs for multi-CL Runs.
type CLID int64

// CLIDsAsInt64s returns proto representation of CLIDs.
func CLIDsAsInt64s(ids []CLID) []int64 {
	r := make([]int64, len(ids))
	for i, id := range ids {
		r[i] = int64(id)
	}
	return r
}

// CLIDs is a convenience type to facilitate handling of a slice of CLID.
type CLIDs []CLID

// Dedupe removes duplicates in place.
//
// Note: Does not perserve original order.
func (p *CLIDs) Dedupe() {
	clids := *p
	if len(clids) <= 1 {
		return
	}
	sort.Slice(clids, func(i, j int) bool { return clids[i] < clids[j] })
	n := 0
	for i := 1; i < len(clids); i++ {
		if clids[i] == clids[n] {
			continue
		}
		n++
		clids[n] = clids[i]
	}
	*p = clids[:n+1]
}
