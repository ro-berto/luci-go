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
