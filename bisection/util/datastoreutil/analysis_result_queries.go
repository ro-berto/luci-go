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

// Package datastoreutil contains utility functions related to datastore entities
package datastoreutil

import (
	"context"
	"fmt"

	"go.chromium.org/luci/bisection/model"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/gae/service/datastore"
)

// GetBuild returns the failed build in the datastore with the given Buildbucket ID
// Note: if the build is not found, this will return (nil, nil)
func GetBuild(c context.Context, bbid int64) (*model.LuciFailedBuild, error) {
	build := &model.LuciFailedBuild{Id: bbid}
	switch err := datastore.Get(c, build); {
	case err == datastore.ErrNoSuchEntity:
		return nil, nil
	case err != nil:
		return nil, err
	}

	return build, nil
}

// GetAnalysisForBuild returns the failure analysis associated with the given Buildbucket ID
// Note: if the build or its analysis is not found, this will return (nil, nil)
func GetAnalysisForBuild(c context.Context, bbid int64) (*model.CompileFailureAnalysis, error) {
	buildModel, err := GetBuild(c, bbid)
	if (err != nil) || (buildModel == nil) {
		return nil, err
	}

	cfModel := &model.CompileFailure{
		Id:    bbid,
		Build: datastore.KeyForObj(c, buildModel),
	}
	switch err := datastore.Get(c, cfModel); {
	case err == datastore.ErrNoSuchEntity:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		//continue
	}

	// If the compile failure was "merged" into another compile failure,
	// use the merged one instead.
	cfKey := datastore.KeyForObj(c, cfModel)
	if cfModel.MergedFailureKey != nil {
		cfKey = cfModel.MergedFailureKey
	}

	// Get the analysis for the compile failure
	q := datastore.NewQuery("CompileFailureAnalysis").Eq("compile_failure", cfKey)
	analyses := []*model.CompileFailureAnalysis{}
	err = datastore.GetAll(c, q, &analyses)
	if err != nil {
		return nil, err
	}
	if len(analyses) == 0 {
		return nil, nil
	}
	if len(analyses) > 1 {
		logging.Warningf(c, "Found more than one analysis for build %d", bbid)
	}
	return analyses[0], nil
}

// GetHeuristicAnalysis returns the heuristic analysis associated with the given failure analysis
func GetHeuristicAnalysis(c context.Context, analysis *model.CompileFailureAnalysis) (*model.CompileHeuristicAnalysis, error) {
	// Gets heuristic analysis results.
	q := datastore.NewQuery("CompileHeuristicAnalysis").Ancestor(datastore.KeyForObj(c, analysis))
	heuristicAnalyses := []*model.CompileHeuristicAnalysis{}
	err := datastore.GetAll(c, q, &heuristicAnalyses)

	if err != nil {
		return nil, err
	}

	if len(heuristicAnalyses) == 0 {
		// No heuristic analysis
		return nil, nil
	}

	if len(heuristicAnalyses) > 1 {
		logging.Warningf(c, "Found multiple heuristic analysis for analysis %d", analysis.Id)
	}

	heuristicAnalysis := heuristicAnalyses[0]
	return heuristicAnalysis, nil
}

// GetSuspectsForHeuristicAnalysis returns the heuristic suspects identified by the given heuristic analysis
func GetSuspectsForHeuristicAnalysis(c context.Context, heuristicAnalysis *model.CompileHeuristicAnalysis) ([]*model.Suspect, error) {
	// Getting the suspects for heuristic analysis
	suspects := []*model.Suspect{}
	q := datastore.NewQuery("Suspect").Ancestor(datastore.KeyForObj(c, heuristicAnalysis)).Order("-score")
	err := datastore.GetAll(c, q, &suspects)
	if err != nil {
		return nil, err
	}

	return suspects, nil
}

// GetSuspectForNthSectionAnalysis returns the heuristic suspects identified by the given heuristic analysis
func GetSuspectForNthSectionAnalysis(c context.Context, nthsectionAnalysis *model.CompileNthSectionAnalysis) (*model.Suspect, error) {
	// Getting the suspects for nthsection analysis
	suspects := []*model.Suspect{}
	q := datastore.NewQuery("Suspect").Ancestor(datastore.KeyForObj(c, nthsectionAnalysis))
	err := datastore.GetAll(c, q, &suspects)
	if err != nil {
		return nil, err
	}
	if len(suspects) == 0 {
		return nil, nil
	}
	if len(suspects) > 0 {
		logging.Warningf(c, "nthsectionAnalysis has more than 1 suspect %d", len(suspects))
	}
	return suspects[0], nil
}

// GetCompileFailureForAnalysisID gets CompileFailure for analysisID.
func GetCompileFailureForAnalysisID(c context.Context, analysisID int64) (*model.CompileFailure, error) {
	cfa, err := GetCompileFailureAnalysis(c, analysisID)
	if err != nil {
		return nil, err
	}
	return GetCompileFailureForAnalysis(c, cfa)
}

// GetCompileFailureForAnalysis gets CompileFailure for analysis
func GetCompileFailureForAnalysis(c context.Context, cfa *model.CompileFailureAnalysis) (*model.CompileFailure, error) {
	compileFailure := &model.CompileFailure{
		Id: cfa.CompileFailure.IntID(),
		// We need to specify the parent here because this is a multi-part key.
		Build: cfa.CompileFailure.Parent(),
	}
	err := datastore.Get(c, compileFailure)
	if err != nil {
		return nil, errors.Annotate(err, "getting compile failure for analysis %d", cfa.Id).Err()
	}
	return compileFailure, nil
}

// GetFailedBuildForAnalysis gets LuciFailedBuild for analysis.
func GetFailedBuildForAnalysis(c context.Context, cfa *model.CompileFailureAnalysis) (*model.LuciFailedBuild, error) {
	cf, err := GetCompileFailureForAnalysis(c, cfa)
	if err != nil {
		return nil, errors.Annotate(err, "getting compile failure for analysis %d", cfa.Id).Err()
	}
	build := &model.LuciFailedBuild{Id: cf.Build.IntID()}
	err = datastore.Get(c, build)
	if err != nil {
		return nil, errors.Annotate(err, "getting failed build for analysis %d", cfa.Id).Err()
	}
	return build, nil
}

// GetRerunsForRerunBuild returns all SingleRerun for a rerunBuild
func GetRerunsForRerunBuild(c context.Context, rerunBuild *model.CompileRerunBuild) ([]*model.SingleRerun, error) {
	q := datastore.NewQuery("SingleRerun").Eq("rerun_build", datastore.KeyForObj(c, rerunBuild)).Order("start_time")
	singleReruns := []*model.SingleRerun{}
	err := datastore.GetAll(c, q, &singleReruns)
	return singleReruns, errors.Annotate(err, "get reruns for rerun build %d", rerunBuild.Id).Err()
}

// GetLastRerunForRerunBuild returns the last SingleRerun for a rerunBuild (based on start_time)
func GetLastRerunForRerunBuild(c context.Context, rerunBuild *model.CompileRerunBuild) (*model.SingleRerun, error) {
	reruns, err := GetRerunsForRerunBuild(c, rerunBuild)
	if err != nil {
		return nil, err
	}
	if len(reruns) == 0 {
		return nil, fmt.Errorf("got no SingleRerun for build %d", rerunBuild.Id)
	}
	return reruns[len(reruns)-1], nil
}

// GetNthSectionAnalysis returns the nthsection analysis associated with the given failure analysis
func GetNthSectionAnalysis(c context.Context, analysis *model.CompileFailureAnalysis) (*model.CompileNthSectionAnalysis, error) {
	q := datastore.NewQuery("CompileNthSectionAnalysis").Ancestor(datastore.KeyForObj(c, analysis))
	nthSectionAnalyses := []*model.CompileNthSectionAnalysis{}
	err := datastore.GetAll(c, q, &nthSectionAnalyses)

	if err != nil {
		return nil, errors.Annotate(err, "couldn't get nthsection analysis for analysis %d", analysis.Id).Err()
	}

	if len(nthSectionAnalyses) == 0 {
		return nil, nil
	}

	if len(nthSectionAnalyses) > 1 {
		return nil, fmt.Errorf("found more than 1 nth section analysis for analysis %d", analysis.Id)
	}

	return nthSectionAnalyses[0], nil
}

// GetCompileFailureAnalysis gets compile failure analysis by its id
func GetCompileFailureAnalysis(c context.Context, analysisID int64) (*model.CompileFailureAnalysis, error) {
	analysis := &model.CompileFailureAnalysis{
		Id: analysisID,
	}
	err := datastore.Get(c, analysis)
	if err != nil {
		return nil, errors.Annotate(err, "couldn't get CompileFailureAnalysis %d", analysis.Id).Err()
	}
	return analysis, err
}

// GetOtherSuspectsWithSameCL returns the list of Suspect(from different analyses)
// that has the same reviewURL as this suspect.
// It is meant to check if the same CL is the suspects for multiple failures.
func GetOtherSuspectsWithSameCL(c context.Context, suspect *model.Suspect) ([]*model.Suspect, error) {
	suspects := []*model.Suspect{}
	q := datastore.NewQuery("Suspect").Eq("review_url", suspect.ReviewUrl)
	err := datastore.GetAll(c, q, &suspects)
	if err != nil {
		return nil, errors.Annotate(err, "failed GetSameSuspects").Err()
	}

	// Remove this suspect
	for i, s := range suspects {
		if s.Id == suspect.Id {
			return append(suspects[:i], suspects[i+1:]...), nil
		}
	}
	return suspects, nil
}

// GetLatestBuildFailureForBuilder returns the latest LuciFailedBuild model for a builderID
// If there is no build failure, return (nil, nil)
func GetLatestBuildFailureForBuilder(c context.Context, project string, bucket string, builder string) (*model.LuciFailedBuild, error) {
	builds := []*model.LuciFailedBuild{}
	q := datastore.NewQuery("LuciFailedBuild").Eq("project", project).Eq("bucket", bucket).Eq("builder", builder).Order("-end_time").Limit(1)
	err := datastore.GetAll(c, q, &builds)
	if err != nil {
		return nil, errors.Annotate(err, "failed querying LuciFailedBuild").Err()
	}

	if len(builds) == 0 {
		return nil, nil
	}
	return builds[0], nil
}

// GetLatestAnalysisForBuilder returns the latest CompileFailureAnalysis for a builderID
// If there is no analysis, return (nil, nil)
func GetLatestAnalysisForBuilder(c context.Context, project string, bucket string, builder string) (*model.CompileFailureAnalysis, error) {
	build, err := GetLatestBuildFailureForBuilder(c, project, bucket, builder)
	if err != nil {
		return nil, errors.Annotate(err, "cannot GetLatestBuildFailureForBuilder").Err()
	}
	if build == nil {
		return nil, nil
	}
	return GetAnalysisForBuild(c, build.Id)
}

// GetRerunsForAnalysis returns all reruns for an analysis
// The result is sorted by start_time
func GetRerunsForAnalysis(c context.Context, cfa *model.CompileFailureAnalysis) ([]*model.SingleRerun, error) {
	q := datastore.NewQuery("SingleRerun").Eq("analysis", datastore.KeyForObj(c, cfa)).Order("start_time")
	reruns := []*model.SingleRerun{}
	err := datastore.GetAll(c, q, &reruns)
	if err != nil {
		return nil, errors.Annotate(err, "getting reruns for analysis %d", cfa.Id).Err()
	}
	return reruns, nil
}

func GetRerunsForNthSectionAnalysis(c context.Context, nsa *model.CompileNthSectionAnalysis) ([]*model.SingleRerun, error) {
	q := datastore.NewQuery("SingleRerun").Eq("analysis", nsa.ParentAnalysis).Eq("rerun_type", model.RerunBuildType_NthSection)
	reruns := []*model.SingleRerun{}
	err := datastore.GetAll(c, q, &reruns)
	if err != nil {
		return nil, errors.Annotate(err, "getting reruns for analysis %d", nsa.ParentAnalysis.IntID()).Err()
	}
	return reruns, nil
}
