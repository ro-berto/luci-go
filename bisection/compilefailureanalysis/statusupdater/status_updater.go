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

// Package statusupdater updates analysis status based on the data in datastore.
package statusupdater

import (
	"context"

	"go.chromium.org/luci/bisection/model"
	pb "go.chromium.org/luci/bisection/proto"
	"go.chromium.org/luci/bisection/util/datastoreutil"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/gae/service/datastore"
)

func UpdateAnalysisStatus(c context.Context, cfa *model.CompileFailureAnalysis) error {
	// If there are confirmed culprit
	if len(cfa.VerifiedCulprits) > 0 {
		return updateStatus(c, cfa, pb.AnalysisStatus_FOUND, pb.AnalysisRunStatus_ENDED)
	}

	// Fetch heuristic and nthsection analysis
	ha, err := datastoreutil.GetHeuristicAnalysis(c, cfa)
	if err != nil {
		return errors.Annotate(err, "couldn't fetch heuristic analysis of analysis %d", cfa.Id).Err()
	}

	nsa, err := datastoreutil.GetNthSectionAnalysis(c, cfa)
	if err != nil {
		return errors.Annotate(err, "couldn't fetch nthsection analysis of analysis %d", cfa.Id).Err()
	}

	haveUnfinishedReruns, err := analysisStillHaveUnfinishedReruns(c, cfa)
	if err != nil {
		return errors.Annotate(err, "couldn't decide if analysis %d has unfinished rerun", cfa.Id).Err()
	}

	// No nth-section run. Just consider the heuristic analysis.
	if nsa == nil || nsa.Status == pb.AnalysisStatus_ERROR {
		if ha == nil || ha.Status == pb.AnalysisStatus_ERROR {
			return updateStatus(c, cfa, pb.AnalysisStatus_ERROR, pb.AnalysisRunStatus_ENDED)
		}
		if ha.Status != pb.AnalysisStatus_SUSPECTFOUND {
			return updateStatus(c, cfa, ha.Status, ha.RunStatus)
		}
		// Heuristic found suspect. So analysis could be in progress or ended
		// depend on if there is any rerun in progress
		if haveUnfinishedReruns {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_STARTED)
		} else {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_ENDED)
		}
	}

	// No heuristic analysis (for some reasons). Just consider nth section
	if ha == nil || ha.Status == pb.AnalysisStatus_ERROR {
		if nsa == nil || nsa.Status == pb.AnalysisStatus_ERROR {
			return updateStatus(c, cfa, pb.AnalysisStatus_ERROR, pb.AnalysisRunStatus_ENDED)
		}

		if nsa.Status != pb.AnalysisStatus_SUSPECTFOUND {
			return updateStatus(c, cfa, nsa.Status, nsa.RunStatus)
		}
		// nsa found suspect. So analysis could be in progress or ended
		// depend on if there is any rerun in progress
		if haveUnfinishedReruns {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_STARTED)
		} else {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_ENDED)
		}
	}

	// Both heuristic and nthsection analysis present
	gotSuspect := (ha.Status == pb.AnalysisStatus_SUSPECTFOUND || nsa.Status == pb.AnalysisStatus_SUSPECTFOUND)
	if gotSuspect {
		inProgress := (ha.Status == pb.AnalysisStatus_RUNNING || nsa.Status == pb.AnalysisStatus_RUNNING)
		if haveUnfinishedReruns || inProgress {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_STARTED)
		} else {
			return updateStatus(c, cfa, pb.AnalysisStatus_SUSPECTFOUND, pb.AnalysisRunStatus_ENDED)
		}
	}

	// No suspect -> either in progress or notfound
	if ha.Status == pb.AnalysisStatus_NOTFOUND && nsa.Status == pb.AnalysisStatus_NOTFOUND {
		return updateStatus(c, cfa, pb.AnalysisStatus_NOTFOUND, pb.AnalysisRunStatus_ENDED)
	}
	return updateStatus(c, cfa, pb.AnalysisStatus_RUNNING, pb.AnalysisRunStatus_STARTED)
}

func updateStatus(c context.Context, cfa *model.CompileFailureAnalysis, status pb.AnalysisStatus, runStatus pb.AnalysisRunStatus) error {
	return datastore.RunInTransaction(c, func(c context.Context) error {
		e := datastore.Get(c, cfa)
		if e != nil {
			return e
		}

		// If the run has ended or canceled, we don't want to do anything
		if cfa.RunStatus == pb.AnalysisRunStatus_ENDED || cfa.RunStatus == pb.AnalysisRunStatus_CANCELED {
			return nil
		}

		// All the same, no need to update
		if cfa.RunStatus == runStatus && cfa.Status == status {
			return nil
		}

		cfa.Status = status
		cfa.RunStatus = runStatus
		if runStatus == pb.AnalysisRunStatus_ENDED || runStatus == pb.AnalysisRunStatus_CANCELED {
			cfa.EndTime = clock.Now(c)
		}
		return datastore.Put(c, cfa)
	}, nil)
}

func analysisStillHaveUnfinishedReruns(c context.Context, cfa *model.CompileFailureAnalysis) (bool, error) {
	reruns, err := datastoreutil.GetRerunsForAnalysis(c, cfa)
	if err != nil {
		return false, err
	}
	for _, rerun := range reruns {
		if rerun.Status == pb.RerunStatus_IN_PROGRESS {
			return true, nil
		}
	}
	return false, nil
}
