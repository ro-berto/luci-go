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

package run

import (
	"testing"
	"time"

	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/cvtesting"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCLQueryBuilder(t *testing.T) {
	t.Parallel()

	Convey("CLQueryBuilder works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		getAll := func(qb CLQueryBuilder) (out common.RunIDs) {
			keys, err := qb.GetAllRunKeys(ctx)
			So(err, ShouldBeNil)
			for _, k := range keys {
				out = append(out, common.RunID(k.StringID()))
			}
			return out
		}

		makeRun := func(proj string, delay time.Duration, clids ...common.CLID) common.RunID {
			createdAt := ct.Clock.Now().Add(delay)
			runID := common.MakeRunID(proj, createdAt, 1, []byte{0, byte(delay / time.Millisecond)})
			So(datastore.Put(ctx, &Run{ID: runID, CLs: clids}), ShouldBeNil)
			for _, clid := range clids {
				So(datastore.Put(ctx, &RunCL{
					Run:       datastore.MakeKey(ctx, RunKind, string(runID)),
					ID:        clid,
					IndexedID: clid,
				}), ShouldBeNil)
			}
			return runID
		}

		clA, clB, clZ := common.CLID(1), common.CLID(2), common.CLID(3)

		// RunID below are ordered lexicographically.
		bond9 := makeRun("bond", 9*time.Millisecond, clA)
		bond4 := makeRun("bond", 4*time.Millisecond, clA, clB)
		bond2 := makeRun("bond", 2*time.Millisecond, clA)
		dart5 := makeRun("dart", 5*time.Millisecond, clA)
		dart3 := makeRun("dart", 3*time.Millisecond, clA)
		rust1 := makeRun("rust", 1*time.Millisecond, clA, clB)
		xero7 := makeRun("xero", 7*time.Millisecond, clA)

		Convey("CL without Runs", func() {
			qb := CLQueryBuilder{CLID: clZ}
			So(getAll(qb), ShouldResemble, common.RunIDs(nil))
		})

		Convey("CL with some Runs", func() {
			qb := CLQueryBuilder{CLID: clB}
			So(getAll(qb), ShouldResemble, common.RunIDs{bond4, rust1})
		})

		Convey("CL with all Runs", func() {
			qb := CLQueryBuilder{CLID: clA}
			So(getAll(qb), ShouldResemble, common.RunIDs{bond9, bond4, bond2, dart5, dart3, rust1, xero7})
		})

		Convey("Reversed", func() {
			qb := CLQueryBuilder{CLID: clA, Descending: true}
			So(getAll(qb), ShouldResemble, common.RunIDs{xero7, rust1, dart3, dart5, bond2, bond4, bond9})
		})

		Convey("Filter by Project", func() {
			qb := CLQueryBuilder{CLID: clA, Project: "bond"}
			So(getAll(qb), ShouldResemble, common.RunIDs{bond9, bond4, bond2})
		})

		Convey("Filtering by Project and Min with diff project", func() {
			qb := CLQueryBuilder{CLID: clA, Project: "dart", Min: bond4}
			So(getAll(qb), ShouldResemble, common.RunIDs{dart5, dart3})

			qb = CLQueryBuilder{CLID: clA, Project: "dart", Min: rust1}
			_, err := qb.BuildKeysOnly(ctx).Finalize()
			So(err, ShouldEqual, datastore.ErrNullQuery)
		})

		Convey("Filtering by Project and Max with diff project", func() {
			qb := CLQueryBuilder{CLID: clA, Project: "dart", Max: xero7}
			So(getAll(qb), ShouldResemble, common.RunIDs{dart5, dart3})

			qb = CLQueryBuilder{CLID: clA, Project: "dart", Max: bond4}
			_, err := qb.BuildKeysOnly(ctx).Finalize()
			So(err, ShouldEqual, datastore.ErrNullQuery)
		})

		Convey("Before", func() {
			qb := CLQueryBuilder{CLID: clA}.BeforeInProject(bond9)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond4, bond2})

			qb = CLQueryBuilder{CLID: clA}.BeforeInProject(bond4)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond2})

			qb = CLQueryBuilder{CLID: clA}.BeforeInProject(bond2)
			So(getAll(qb), ShouldResemble, common.RunIDs(nil))
		})

		Convey("After", func() {
			qb := CLQueryBuilder{CLID: clA}.AfterInProject(bond2)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond9, bond4})

			qb = CLQueryBuilder{CLID: clA}.AfterInProject(bond4)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond9})

			qb = CLQueryBuilder{CLID: clA}.AfterInProject(bond9)
			So(getAll(qb), ShouldResemble, common.RunIDs(nil))
		})

		Convey("Next Run = After + Reversed + Limit 1", func() {
			qb := CLQueryBuilder{CLID: clA, Descending: true, Limit: 1}.AfterInProject(bond2)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond4})
		})

		Convey("After and Before", func() {
			qb := CLQueryBuilder{CLID: clA}.AfterInProject(bond2).BeforeInProject(bond9)
			So(getAll(qb), ShouldResemble, common.RunIDs{bond4})
		})

		Convey("Invalid usage panics", func() {
			So(func() { CLQueryBuilder{CLID: clA, Project: "dart"}.BeforeInProject(bond2) }, ShouldPanic)
			So(func() { CLQueryBuilder{CLID: clA, Project: "dart"}.AfterInProject(bond2) }, ShouldPanic)
			So(func() { CLQueryBuilder{CLID: clA}.AfterInProject(dart3).BeforeInProject(xero7) }, ShouldPanic)
		})
	})
}
