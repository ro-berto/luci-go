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

package poller

import (
	"sort"
	"testing"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/cvtesting"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScheduleRefreshTasks(t *testing.T) {
	t.Parallel()

	Convey("scheduleRefreshTasks works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const lProject = "chromium"
		const gHost = "chromium-review.example.com"
		const gRepo = "infra/infra"

		pm := pmMock{}
		clUpdater := clUpdaterMock{}
		p := New(ct.TQDispatcher, ct.GFake.Factory(), &clUpdater, &pm)

		changes := []int64{1, 2, 3, 4, 5}
		const notYetSaved = 4

		var knownIDs common.CLIDs
		for _, i := range changes {
			if i == notYetSaved {
				continue
			}
			cl, err := changelist.MustGobID(gHost, i).GetOrInsert(ctx, func(cl *changelist.CL) {
				// In practice, cl.Snapshot would be populated, but for this test it
				// doesn't matter.
			})
			So(err, ShouldBeNil)
			knownIDs = append(knownIDs, cl.ID)
		}
		sort.Sort(knownIDs)

		err := p.scheduleRefreshTasks(ctx, lProject, gHost, changes)
		So(err, ShouldBeNil)

		// PM must be notified immediately on CLs already saved.
		ids := pm.projects[lProject]
		sort.Sort(ids)
		So(ids, ShouldResemble, knownIDs)

		// CL Updater must have scheduled tasks.
		etas := clUpdater.peekETAs()
		payloads := clUpdater.popPayloadsByETA()
		So(payloads, ShouldHaveLength, len(changes))
		// Tasks must be somewhat distributed in time.
		mid := ct.Clock.Now().Add(fullPollInterval / 2)
		So(etas[1], ShouldHappenBefore, mid)
		So(etas[3], ShouldHappenAfter, mid)
		// For not yet saved CL, PM must be forcefully notified.
		var forced []int64
		for _, p := range payloads {
			if p.GetForceNotify() {
				forced = append(forced, p.GetChange())
			}
		}
		So(forced, ShouldResemble, []int64{notYetSaved})
	})
}
