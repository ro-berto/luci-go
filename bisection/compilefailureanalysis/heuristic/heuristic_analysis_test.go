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

package heuristic

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/bisection/model"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
)

func TestSaveResultsToDatastore(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("SaveResultsToDatastore", t, func() {
		heuristicAnalysis := &model.CompileHeuristicAnalysis{}
		err := datastore.Put(c, heuristicAnalysis)
		So(err, ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		result := &model.HeuristicAnalysisResult{
			Items: []*model.HeuristicAnalysisResultItem{
				{
					Commit:      "12345",
					ReviewUrl:   "this/is/review/url",
					ReviewTitle: "title",
					Justification: &model.SuspectJustification{
						Items: []*model.SuspectJustificationItem{
							{
								Score:  10,
								Reason: "failure reason",
								Type:   model.JustificationType_FAILURELOG,
							},
						},
					},
				},
			},
		}

		err = saveResultsToDatastore(c, heuristicAnalysis, result, "host", "proj", "ref")
		So(err, ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		suspects := []*model.Suspect{}
		q := datastore.NewQuery("Suspect")
		err = datastore.GetAll(c, q, &suspects)
		So(err, ShouldBeNil)
		So(len(suspects), ShouldEqual, 1)
		So(suspects[0], ShouldResemble, &model.Suspect{
			ParentAnalysis: datastore.KeyForObj(c, heuristicAnalysis),
			Id:             suspects[0].Id,
			ReviewUrl:      "this/is/review/url",
			ReviewTitle:    "title",
			Justification:  "failure reason",
			Score:          10,
			GitilesCommit: buildbucketpb.GitilesCommit{
				Host:    "host",
				Project: "proj",
				Ref:     "ref",
				Id:      "12345",
			},
			VerificationStatus: model.SuspectVerificationStatus_Unverified,
			Type:               model.SuspectType_Heuristic,
		})
	})
}
