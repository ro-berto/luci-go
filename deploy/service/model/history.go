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

package model

import (
	"context"

	"go.chromium.org/luci/deploy/api/modelpb"
	"go.chromium.org/luci/gae/service/datastore"
)

// shouldRecordHistory returns true if the new history entry should be recorded.
//
// It is skipped if it is not sufficiently interesting compared to the last
// committed entry.
func shouldRecordHistory(next, prev *modelpb.AssetHistory) bool {
	nextDecision := next.Decision.Decision
	prevDecision := prev.Decision.Decision
	switch {
	case nextDecision != prevDecision:
		return true // changes are always interesting
	case isActuateDecision(nextDecision):
		return true // active actuations are also always interesting
	case nextDecision == modelpb.ActuationDecision_SKIP_UPTODATE:
		return false // repeating UPTODATE decisions are boring, it is steady state
	case nextDecision == modelpb.ActuationDecision_SKIP_DISABLED:
		return false // repeating DISABLED decisions are also boring
	case nextDecision == modelpb.ActuationDecision_SKIP_LOCKED:
		return !sameLocks(next.Decision.Locks, prev.Decision.Locks)
	case nextDecision == modelpb.ActuationDecision_SKIP_BROKEN:
		return true // errors are always interesting (for retries and alerts)
	default:
		panic("unreachable")
	}
}

func sameLocks(a, b []*modelpb.ActuationLock) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Id != b[i].Id {
			return false
		}
	}
	return true
}

// commitHistory prepares the history entries for commit and emits TQ tasks.
//
// Must be called inside a transaction. Returns a list of entities to
// transactionally store.
func commitHistory(ctx context.Context, entries []*modelpb.AssetHistory) ([]interface{}, error) {
	// TODO: Emit TQ tasks.
	toPut := make([]interface{}, len(entries))
	for idx, entry := range entries {
		toPut[idx] = &AssetHistory{
			ID:      entry.HistoryId,
			Parent:  datastore.NewKey(ctx, "Asset", entry.AssetId, 0, nil),
			Entry:   entry,
			Created: asTime(entry.Actuation.Created),
		}
	}
	return toPut, nil
}