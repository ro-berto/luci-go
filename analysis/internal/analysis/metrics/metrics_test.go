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

package metrics

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetrics(t *testing.T) {
	Convey(`Metrics`, t, func() {
		Convey(`Metrics have unique sort orders`, func() {
			usedSortOrders := make(map[int]bool)
			for _, m := range ComputedMetrics {
				unique := !usedSortOrders[m.SortPriority]
				So(unique, ShouldBeTrue)
				usedSortOrders[m.SortPriority] = true
			}
		})
		Convey(`Metrics have unique IDs`, func() {
			usedIDs := make(map[string]bool)
			for _, m := range ComputedMetrics {
				unique := !usedIDs[m.ID.String()]
				So(unique, ShouldBeTrue)
				usedIDs[m.ID.String()] = true
			}
		})
	})
}
