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

package nthsection

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/smartystreets/goconvey/convey"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/bisection/internal/gitiles"
	lbm "go.chromium.org/luci/bisection/model"
	lbpb "go.chromium.org/luci/bisection/proto"
)

func TestAnalyze(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())
	cl := testclock.New(testclock.TestTimeUTC)
	c = clock.Set(c, cl)

	gitilesResponse := lbm.ChangeLogResponse{
		Log: []*lbm.ChangeLog{
			{
				Commit:  "3424",
				Message: "Use TestActivationManager for all page activations\n\nblah blah\n\nChange-Id: blah\nBug: blah\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/3472129\nReviewed-by: blah blah\n",
			},
			{
				Commit:  "3425",
				Message: "Second Commit\n\nblah blah\n\nChange-Id: blah\nBug: blah\nReviewed-on: https://chromium-review.googlesource.com/c/chromium/src/+/3472130\nReviewed-by: blah blah\n",
			},
		},
	}
	gitilesResponseStr, _ := json.Marshal(gitilesResponse)

	c = gitiles.MockedGitilesClientContext(c, map[string]string{
		"https://chromium.googlesource.com/chromium/src/+log/12345..23456": string(gitilesResponseStr),
	})

	Convey("CheckBlameList", t, func() {
		cfa := &lbm.CompileFailureAnalysis{
			Id: 123,
		}
		So(datastore.Put(c, cfa), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		rr := &lbpb.RegressionRange{
			LastPassed: &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "chromium/src",
				Id:      "12345",
			},
			FirstFailed: &buildbucketpb.GitilesCommit{
				Host:    "chromium.googlesource.com",
				Project: "chromium/src",
				Id:      "23456",
			},
		}

		nsa, err := Analyze(c, cfa, rr)
		So(err, ShouldBeNil)
		So(nsa, ShouldNotBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		// Fetch the nth section analysis
		q := datastore.NewQuery("CompileNthSectionAnalysis")
		nthsectionAnalyses := []*lbm.CompileNthSectionAnalysis{}
		err = datastore.GetAll(c, q, &nthsectionAnalyses)
		So(err, ShouldBeNil)
		So(len(nthsectionAnalyses), ShouldEqual, 1)
		nsa = nthsectionAnalyses[0]

		diff := cmp.Diff(nsa.BlameList, &lbpb.BlameList{
			Commits: []*lbpb.BlameListSingleCommit{
				{
					Commit:      "3424",
					ReviewTitle: "Use TestActivationManager for all page activations",
					ReviewUrl:   "https://chromium-review.googlesource.com/c/chromium/src/+/3472129",
				},
				{
					Commit:      "3425",
					ReviewTitle: "Second Commit",
					ReviewUrl:   "https://chromium-review.googlesource.com/c/chromium/src/+/3472130",
				},
			},
		}, cmp.Comparer(proto.Equal))
		So(diff, ShouldEqual, "")
	})
}
