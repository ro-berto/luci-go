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

package datastoreutil

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/bisection/model"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
)

func TestGetBuild(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("No build found", t, func() {
		buildModel, err := GetBuild(c, 100)
		So(err, ShouldBeNil)
		So(buildModel, ShouldBeNil)
	})

	Convey("Build found", t, func() {
		// Prepare datastore
		failed_build := &model.LuciFailedBuild{
			Id: 101,
		}
		So(datastore.Put(c, failed_build), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		buildModel, err := GetBuild(c, 101)
		So(err, ShouldBeNil)
		So(buildModel, ShouldNotBeNil)
		So(buildModel.Id, ShouldEqual, 101)
	})
}

func TestGetAnalysisForBuild(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("No build found", t, func() {
		analysis, err := GetAnalysisForBuild(c, 100)
		So(err, ShouldBeNil)
		So(analysis, ShouldBeNil)
	})

	Convey("No analysis found", t, func() {
		// Prepare datastore
		failedBuild := &model.LuciFailedBuild{
			Id: 101,
		}
		So(datastore.Put(c, failedBuild), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		analysis, err := GetAnalysisForBuild(c, 101)
		So(err, ShouldBeNil)
		So(analysis, ShouldBeNil)
	})

	Convey("Analysis found", t, func() {
		// Prepare datastore
		failedBuild := &model.LuciFailedBuild{
			Id: 101,
		}
		So(datastore.Put(c, failedBuild), ShouldBeNil)

		compileFailure := &model.CompileFailure{
			Id:    101,
			Build: datastore.KeyForObj(c, failedBuild),
		}
		So(datastore.Put(c, compileFailure), ShouldBeNil)

		compileFailureAnalysis := &model.CompileFailureAnalysis{
			Id:                 1230001,
			CompileFailure:     datastore.KeyForObj(c, compileFailure),
			FirstFailedBuildId: 101,
		}
		So(datastore.Put(c, compileFailureAnalysis), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		analysis, err := GetAnalysisForBuild(c, 101)
		So(err, ShouldBeNil)
		So(analysis, ShouldNotBeNil)
		So(analysis.Id, ShouldEqual, 1230001)
		So(analysis.FirstFailedBuildId, ShouldEqual, 101)
	})

	Convey("Related analysis found", t, func() {
		// Prepare datastore
		firstFailedBuild := &model.LuciFailedBuild{
			Id: 200,
		}
		So(datastore.Put(c, firstFailedBuild), ShouldBeNil)

		firstCompileFailure := &model.CompileFailure{
			Id:    200,
			Build: datastore.KeyForObj(c, firstFailedBuild),
		}
		So(datastore.Put(c, firstCompileFailure), ShouldBeNil)

		failedBuild := &model.LuciFailedBuild{
			Id: 201,
		}
		So(datastore.Put(c, failedBuild), ShouldBeNil)

		compileFailure := &model.CompileFailure{
			Id:               201,
			Build:            datastore.KeyForObj(c, failedBuild),
			MergedFailureKey: datastore.KeyForObj(c, firstCompileFailure),
		}
		So(datastore.Put(c, compileFailure), ShouldBeNil)

		compileFailureAnalysis := &model.CompileFailureAnalysis{
			Id:             1230002,
			CompileFailure: datastore.KeyForObj(c, firstCompileFailure),
		}
		So(datastore.Put(c, compileFailureAnalysis), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		analysis, err := GetAnalysisForBuild(c, 201)
		So(err, ShouldBeNil)
		So(analysis, ShouldNotBeNil)
		So(analysis.Id, ShouldEqual, 1230002)
	})
}

func TestGetHeuristicAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("No heuristic analysis found", t, func() {
		compileFailureAnalysis := &model.CompileFailureAnalysis{
			Id: 1230003,
		}
		heuristicAnalysis, err := GetHeuristicAnalysis(c, compileFailureAnalysis)
		So(err, ShouldBeNil)
		So(heuristicAnalysis, ShouldBeNil)
	})

	Convey("Heuristic analysis found", t, func() {
		// Prepare datastore
		compileFailureAnalysis := &model.CompileFailureAnalysis{
			Id: 1230003,
		}
		So(datastore.Put(c, compileFailureAnalysis), ShouldBeNil)

		compileHeuristicAnalysis := &model.CompileHeuristicAnalysis{
			Id:             4560001,
			ParentAnalysis: datastore.KeyForObj(c, compileFailureAnalysis),
		}
		So(datastore.Put(c, compileHeuristicAnalysis), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		heuristicAnalysis, err := GetHeuristicAnalysis(c, compileFailureAnalysis)
		So(err, ShouldBeNil)
		So(heuristicAnalysis, ShouldNotBeNil)
		So(heuristicAnalysis.Id, ShouldEqual, 4560001)
	})
}

func TestGetSuspects(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	datastore.GetTestable(c).AutoIndex(true)

	Convey("No suspects found", t, func() {
		// Prepare datastore
		heuristicAnalysis := &model.CompileHeuristicAnalysis{
			Id: 700,
		}
		So(datastore.Put(c, heuristicAnalysis), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		suspects, err := GetSuspects(c, heuristicAnalysis)
		So(err, ShouldBeNil)
		So(len(suspects), ShouldEqual, 0)
	})

	Convey("All suspects found", t, func() {
		// Prepare datastore
		heuristicAnalysis := &model.CompileHeuristicAnalysis{
			Id: 701,
		}
		So(datastore.Put(c, heuristicAnalysis), ShouldBeNil)

		suspect1 := &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, heuristicAnalysis),
			Score:          1,
		}
		suspect2 := &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, heuristicAnalysis),
			Score:          3,
		}
		suspect3 := &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, heuristicAnalysis),
			Score:          4,
		}
		suspect4 := &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, heuristicAnalysis),
			Score:          2,
		}
		So(datastore.Put(c, suspect1), ShouldBeNil)
		So(datastore.Put(c, suspect2), ShouldBeNil)
		So(datastore.Put(c, suspect3), ShouldBeNil)
		So(datastore.Put(c, suspect4), ShouldBeNil)

		// Add a different heuristic analysis with its own suspect
		otherHeuristicAnalysis := &model.CompileHeuristicAnalysis{
			Id: 702,
		}
		So(datastore.Put(c, heuristicAnalysis), ShouldBeNil)
		otherSuspect := &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, otherHeuristicAnalysis),
			Score:          5,
		}
		So(datastore.Put(c, otherSuspect), ShouldBeNil)

		datastore.GetTestable(c).CatchupIndexes()

		suspects, err := GetSuspects(c, heuristicAnalysis)
		So(err, ShouldBeNil)
		So(len(suspects), ShouldEqual, 4)
		So(suspects[0].Score, ShouldEqual, 4)
		So(suspects[1].Score, ShouldEqual, 3)
		So(suspects[2].Score, ShouldEqual, 2)
		So(suspects[3].Score, ShouldEqual, 1)
	})
}

func TestGetCompileFailureForAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	datastore.GetTestable(c).AutoIndex(true)

	Convey("No analysis found", t, func() {
		_, err := GetCompileFailureForAnalysis(c, 100)
		So(err, ShouldNotBeNil)
	})

	Convey("Have analysis for compile failure", t, func() {
		build := &model.LuciFailedBuild{
			Id: 111,
		}
		So(datastore.Put(c, build), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		compileFailure := &model.CompileFailure{
			Id:    123,
			Build: datastore.KeyForObj(c, build),
		}
		So(datastore.Put(c, compileFailure), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		analysis := &model.CompileFailureAnalysis{
			Id:             456,
			CompileFailure: datastore.KeyForObj(c, compileFailure),
		}
		So(datastore.Put(c, analysis), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		cf, err := GetCompileFailureForAnalysis(c, 456)
		So(err, ShouldBeNil)
		So(cf.Id, ShouldEqual, 123)
	})
}

func TestGetRerunsForRerunBuild(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	datastore.GetTestable(c).AddIndexes(&datastore.IndexDefinition{
		Kind: "SingleRerun",
		SortBy: []datastore.IndexColumn{
			{
				Property: "rerun_build",
			},
			{
				Property: "start_time",
			},
		},
	})
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	Convey("GetRerunsForRerunBuild", t, func() {
		// Set up reruns
		rerunBuildModel := &model.CompileRerunBuild{
			Id: 8800,
		}
		So(datastore.Put(c, rerunBuildModel), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		reruns, err := GetRerunsForRerunBuild(c, rerunBuildModel)
		So(err, ShouldBeNil)
		So(len(reruns), ShouldEqual, 0)
		_, err = GetLastRerunForRerunBuild(c, rerunBuildModel)
		So(err, ShouldNotBeNil)

		singleRerun1 := &model.SingleRerun{
			Id:         1,
			RerunBuild: datastore.KeyForObj(c, rerunBuildModel),
			StartTime:  clock.Now(c),
		}

		singleRerun2 := &model.SingleRerun{
			Id:         2,
			RerunBuild: datastore.KeyForObj(c, rerunBuildModel),
			StartTime:  clock.Now(c).Add(time.Hour),
		}

		singleRerun3 := &model.SingleRerun{
			Id:         3,
			RerunBuild: datastore.KeyForObj(c, rerunBuildModel),
			StartTime:  clock.Now(c).Add(time.Minute),
		}

		So(datastore.Put(c, singleRerun1), ShouldBeNil)
		So(datastore.Put(c, singleRerun2), ShouldBeNil)
		So(datastore.Put(c, singleRerun3), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		reruns, err = GetRerunsForRerunBuild(c, rerunBuildModel)
		So(err, ShouldBeNil)
		So(len(reruns), ShouldEqual, 3)
		r, err := GetLastRerunForRerunBuild(c, rerunBuildModel)
		So(err, ShouldBeNil)
		So(r.Id, ShouldEqual, 2)
	})
}

func TestGetNthSectionAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("No nthsection analysis found", t, func() {
		cfa := &model.CompileFailureAnalysis{
			Id: 123,
		}
		nsa, err := GetNthSectionAnalysis(c, cfa)
		So(err, ShouldBeNil)
		So(nsa, ShouldBeNil)
	})

	Convey("More than one nthsection analysis found", t, func() {
		cfa := &model.CompileFailureAnalysis{
			Id: 456,
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		nsa1 := &model.CompileNthSectionAnalysis{
			ParentAnalysis: datastore.KeyForObj(c, cfa),
		}
		nsa2 := &model.CompileNthSectionAnalysis{
			ParentAnalysis: datastore.KeyForObj(c, cfa),
		}
		So(datastore.Put(c, nsa1), ShouldBeNil)
		So(datastore.Put(c, nsa2), ShouldBeNil)

		_, err := GetNthSectionAnalysis(c, cfa)
		So(err, ShouldNotBeNil)
	})

	Convey("Nthsection analysis found", t, func() {
		cfa := &model.CompileFailureAnalysis{
			Id: 789,
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		nsa1 := &model.CompileNthSectionAnalysis{
			ParentAnalysis: datastore.KeyForObj(c, cfa),
			Id:             345,
		}
		So(datastore.Put(c, nsa1), ShouldBeNil)
		nsa, err := GetNthSectionAnalysis(c, cfa)
		So(err, ShouldBeNil)
		So(nsa.Id, ShouldEqual, 345)
	})
}

func TestGetCompileFailureAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("GetCompileFailureAnalysis", t, func() {
		_, err := GetCompileFailureAnalysis(c, 123)
		So(err, ShouldNotBeNil)
		cfa := &model.CompileFailureAnalysis{
			Id: 123,
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		analysis, err := GetCompileFailureAnalysis(c, 123)
		So(err, ShouldBeNil)
		So(analysis.Id, ShouldEqual, 123)
	})
}

func TestGetOtherSuspectsWithSameCL(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("GetOtherSuspectsWithSameCL", t, func() {
		suspect := &model.Suspect{
			ReviewUrl: "https://this/is/review/url",
			Id:        123,
		}
		So(datastore.Put(c, suspect), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()
		suspects, err := GetOtherSuspectsWithSameCL(c, suspect)
		So(err, ShouldBeNil)
		So(len(suspects), ShouldEqual, 0)

		suspect1 := &model.Suspect{
			ReviewUrl: "https://this/is/review/url",
			Id:        124,
		}
		So(datastore.Put(c, suspect1), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()
		suspects, err = GetOtherSuspectsWithSameCL(c, suspect)
		So(err, ShouldBeNil)
		So(len(suspects), ShouldEqual, 1)
		So(suspects[0].Id, ShouldEqual, 124)
	})
}

func TestGetLatestBuildFailureAndAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	datastore.GetTestable(c).AddIndexes(&datastore.IndexDefinition{
		Kind: "LuciFailedBuild",
		SortBy: []datastore.IndexColumn{
			{
				Property: "project",
			},
			{
				Property: "bucket",
			},
			{
				Property: "builder",
			},
			{
				Property:   "end_time",
				Descending: true,
			},
		},
	})
	datastore.GetTestable(c).CatchupIndexes()
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	Convey("GetLatestBuildFailureAndAnalysis", t, func() {
		build, err := GetLatestBuildFailureForBuilder(c, "project", "bucket", "builder")
		So(err, ShouldBeNil)
		So(build, ShouldBeNil)
		analysis, err := GetLatestAnalysisForBuilder(c, "project", "bucket", "builder")
		So(err, ShouldBeNil)
		So(analysis, ShouldBeNil)

		bf1 := &model.LuciFailedBuild{
			Id: 123,
			LuciBuild: model.LuciBuild{
				Project: "project",
				Bucket:  "bucket",
				Builder: "builder",
				EndTime: clock.Now(c),
			},
		}

		bf2 := &model.LuciFailedBuild{
			Id: 456,
			LuciBuild: model.LuciBuild{
				Project: "project",
				Bucket:  "bucket",
				Builder: "builder",
				EndTime: clock.Now(c).Add(time.Hour),
			},
		}

		So(datastore.Put(c, bf1), ShouldBeNil)
		So(datastore.Put(c, bf2), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()
		build, err = GetLatestBuildFailureForBuilder(c, "project", "bucket", "builder")
		So(err, ShouldBeNil)
		So(build.Id, ShouldEqual, 456)

		cf1 := &model.CompileFailure{
			Id:    123,
			Build: datastore.KeyForObj(c, bf1),
		}
		So(datastore.Put(c, cf1), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		cf2 := &model.CompileFailure{
			Id:               456,
			Build:            datastore.KeyForObj(c, bf2),
			MergedFailureKey: datastore.KeyForObj(c, cf1),
		}
		So(datastore.Put(c, cf2), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		cfa := &model.CompileFailureAnalysis{
			Id:             789,
			CompileFailure: datastore.KeyForObj(c, cf1),
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		analysis, err = GetLatestAnalysisForBuilder(c, "project", "bucket", "builder")
		So(err, ShouldBeNil)
		So(analysis.Id, ShouldEqual, 789)
	})
}

func TestGetRerunsForAnalysis(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	Convey("GetRerunsForAnalysis", t, func() {
		cfa := &model.CompileFailureAnalysis{
			Id: 123,
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		nsa := &model.CompileNthSectionAnalysis{
			ParentAnalysis: datastore.KeyForObj(c, cfa),
		}
		So(datastore.Put(c, nsa), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		rr1 := &model.SingleRerun{
			Type:     model.RerunBuildType_CulpritVerification,
			Analysis: datastore.KeyForObj(c, cfa),
		}

		rr2 := &model.SingleRerun{
			Id:       333,
			Type:     model.RerunBuildType_NthSection,
			Analysis: datastore.KeyForObj(c, cfa),
		}

		So(datastore.Put(c, rr1), ShouldBeNil)
		So(datastore.Put(c, rr2), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		reruns, err := GetRerunsForAnalysis(c, cfa)
		So(err, ShouldBeNil)
		So(len(reruns), ShouldEqual, 2)

		nsectionReruns, err := GetRerunsForNthSectionAnalysis(c, nsa)
		So(err, ShouldBeNil)
		So(len(nsectionReruns), ShouldEqual, 1)
		So(nsectionReruns[0].Id, ShouldEqual, 333)
	})
}
