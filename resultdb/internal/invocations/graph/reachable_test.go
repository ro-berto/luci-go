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

package graph

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/resultdb/internal/invocations"
)

func TestReachableInvocations(t *testing.T) {
	Convey(`ReachableInvocations`, t, func() {
		invs := make(ReachableInvocations)
		invs["0"] = ReachableInvocation{HasTestResults: true, HasTestExonerations: true, Realm: "testproject:testrealmA"}
		invs["1"] = ReachableInvocation{HasTestResults: true, HasTestExonerations: false, Realm: "testproject:testrealmB"}
		invs["2"] = ReachableInvocation{HasTestResults: true, HasTestExonerations: true, Realm: "testproject:testrealmC"}
		invs["3"] = ReachableInvocation{HasTestResults: false, HasTestExonerations: false, Realm: "testproject:testrealmC"}
		invs["4"] = ReachableInvocation{HasTestResults: false, HasTestExonerations: true, Realm: "testproject:testrealmB"}
		invs["5"] = ReachableInvocation{HasTestResults: false, HasTestExonerations: false, Realm: "testproject:testrealmA"}

		Convey(`Batches`, func() {
			So(invs.batches(2), ShouldResemble, []ReachableInvocations{
				{
					"3": ReachableInvocation{HasTestResults: false, HasTestExonerations: false, Realm: "testproject:testrealmC"},
					"4": ReachableInvocation{HasTestResults: false, HasTestExonerations: true, Realm: "testproject:testrealmB"},
				},
				{
					"0": ReachableInvocation{HasTestResults: true, HasTestExonerations: true, Realm: "testproject:testrealmA"},
					"1": ReachableInvocation{HasTestResults: true, HasTestExonerations: false, Realm: "testproject:testrealmB"},
				},
				{
					"2": ReachableInvocation{HasTestResults: true, HasTestExonerations: true, Realm: "testproject:testrealmC"},
					"5": ReachableInvocation{HasTestResults: false, HasTestExonerations: false, Realm: "testproject:testrealmA"},
				},
			})
		})
		Convey(`Marshal and unmarshal`, func() {
			b, err := invs.marshal()
			So(err, ShouldBeNil)

			result, err := unmarshalReachableInvocations(b)
			So(err, ShouldBeNil)

			So(result, ShouldResemble, invs)
		})
		Convey(`IDSet`, func() {
			invIDs := invs.IDSet()
			So(invIDs, ShouldResemble, invocations.NewIDSet("0", "1", "2", "3", "4", "5"))
		})
		Convey(`WithTestResultsIDSet`, func() {
			invIDs := invs.WithTestResultsIDSet()
			So(invIDs, ShouldResemble, invocations.NewIDSet("0", "1", "2"))
		})
		Convey(`WithExonerationsIDSet`, func() {
			invIDs := invs.WithExonerationsIDSet()
			So(invIDs, ShouldResemble, invocations.NewIDSet("0", "2", "4"))
		})
	})
}
