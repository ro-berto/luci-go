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

package common

import (
	"sort"
)

// TryjobID is a unique ID of a Tryjob used internally in CV.
//
// This ID is not a Buildbucket Build ID.
// See also tryjob.Tryjob type.
type TryjobID int64

// TryjobIDs is a convenience type to facilitate handling of a slice of
// TryjobID.
type TryjobIDs []TryjobID

// Dedupe removes duplicates in place and sorts the slice.
//
// Note: Does not preserve original order.
func (p *TryjobIDs) Dedupe() {
	ids := *p
	if len(ids) <= 1 {
		return
	}
	sort.Sort(ids)
	n, prev, skipped := 0, ids[0], false
	for _, id := range ids[1:] {
		if id == prev {
			skipped = true
			continue
		}
		n++
		if skipped {
			ids[n] = id
		}
		prev = id
	}
	*p = ids[:n+1]
}

// Len is the number of elements in the collection.
func (ids TryjobIDs) Len() int {
	return len(ids)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (ids TryjobIDs) Less(i int, j int) bool {
	return ids[i] < ids[j]
}

// Swap swaps the elements with indexes i and j.
func (ids TryjobIDs) Swap(i int, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

// Set returns a new set of TryjobIDs.
func (ids TryjobIDs) Set() map[TryjobID]struct{} {
	if ids == nil {
		return nil
	}
	ret := make(map[TryjobID]struct{}, len(ids))
	for _, id := range ids {
		ret[id] = struct{}{}
	}
	return ret
}

// Contains returns true if TryjobID is inside these TryjobIDs.
func (ids TryjobIDs) Contains(id TryjobID) bool {
	for _, x := range ids {
		if x == id {
			return true
		}
	}
	return false
}

// ToInt64 returns a slice that contains all Tryjobs in int64 type.
func (ids TryjobIDs) ToInt64() []int64 {
	if ids == nil {
		return nil
	}
	ret := make([]int64, len(ids))
	for i, id := range ids {
		ret[i] = int64(id)
	}
	return ret
}

// MakeTryjobIDs returns TryjobIDs from list of TryjobID in int64.
func MakeTryjobIDs(ids ...int64) TryjobIDs {
	if ids == nil {
		return nil
	}
	ret := make(TryjobIDs, len(ids))
	for i, id := range ids {
		ret[i] = TryjobID(id)
	}
	return ret
}

// TryjobIDSet is convenience type to reduce the boilerplate.
type TryjobIDSet map[TryjobID]struct{}

// Add adds the provided Tryjob ID to the set.
func (s TryjobIDSet) Add(tjID TryjobID) {
	s[tjID] = struct{}{}
}

// Has returns true if the provided Tryjob ID is in the set.
//
// Otherwise, returns false.
func (s TryjobIDSet) Has(tjID TryjobID) bool {
	_, exists := s[tjID]
	return exists
}
