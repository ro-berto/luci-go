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

// Package runtest implements tests for working with Run Manager.
package runtest

import (
	"sort"

	"go.chromium.org/luci/server/tq/tqtesting"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/run/eventpb"
)

// Runs returns list of runs from tasks for Run Manager.
func Runs(in tqtesting.TaskList) (runs common.RunIDs) {
	for _, t := range in.SortByETA() {
		switch v := t.Payload.(type) {
		case *eventpb.PokeRunTask:
			runs = append(runs, common.RunID(v.GetRunId()))
		case *eventpb.KickPokeRunTask:
			runs = append(runs, common.RunID(v.GetRunId()))
		}
	}
	return runs
}

// SortedRuns returns sorted list of runs from tasks for Run Manager.
func SortedRuns(in tqtesting.TaskList) common.RunIDs {
	runs := Runs(in)
	sort.Sort(runs)
	return runs
}
