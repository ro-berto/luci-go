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

package changelist

import (
	"context"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/cvtesting"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestUpdaterSchedule(t *testing.T) {
	t.Parallel()

	Convey("Correctly generate dedup keys for Updater TQ tasks", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		Convey("Correctly generate dedup keys for Updater TQ tasks", func() {

			Convey("Diff CLIDs have diff dedup keys", func() {
				t := &UpdateCLTask{LuciProject: "proj", Id: 7}
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				t.Id = 8
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldNotResemble, k2)
			})

			Convey("Diff ExternalID have diff dedup keys", func() {
				t := &UpdateCLTask{LuciProject: "proj"}
				t.ExternalId = "kind1/foo/23"
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				t.ExternalId = "kind4/foo/56"
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldNotResemble, k2)
			})

			Convey("Even if ExternalID and internal ID refer to the same CL, they have diff dedup keys", func() {
				t1 := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				t2 := &UpdateCLTask{LuciProject: "proj", Id: 2}
				k1 := makeTaskDeduplicationKey(ctx, t1, 0)
				k2 := makeTaskDeduplicationKey(ctx, t2, 0)
				So(k1, ShouldNotResemble, k2)
			})

			Convey("Diff updatedHint have diff dedup keys", func() {
				t := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				t.UpdatedHint = timestamppb.New(ct.Clock.Now())
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				t.UpdatedHint = timestamppb.New(ct.Clock.Now().Add(time.Second))
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldNotResemble, k2)
			})

			Convey("Same CLs but diff LUCI projects have diff dedup keys", func() {
				t := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				t.LuciProject += "-diff"
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldNotResemble, k2)
			})

			Convey("Same CL at the same time is de-duped", func() {
				t := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldResemble, k2)

				Convey("Internal ID doesn't affect dedup based on ExternalID", func() {
					t.Id = 123
					k3 := makeTaskDeduplicationKey(ctx, t, 0)
					So(k3, ShouldResemble, k1)
				})
			})

			Convey("Same CL with a delay or after the same delay is de-duped", func() {
				t := &UpdateCLTask{LuciProject: "proj", Id: 123}
				k1 := makeTaskDeduplicationKey(ctx, t, time.Second)
				ct.Clock.Add(time.Second)
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldResemble, k2)
			})

			Convey("Same CL at mostly same time is also de-duped", func() {
				t := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				// NOTE: this check may fail if common.DistributeOffset is changed,
				// making new timestamp in the next epoch. If so, adjust the increment.
				ct.Clock.Add(time.Second)
				k2 := makeTaskDeduplicationKey(ctx, t, 0)
				So(k1, ShouldResemble, k2)
			})

			Convey("Same CL after sufficient time is no longer de-duped", func() {
				t := &UpdateCLTask{LuciProject: "proj", ExternalId: "kind1/foo/23"}
				k1 := makeTaskDeduplicationKey(ctx, t, 0)
				k2 := makeTaskDeduplicationKey(ctx, t, blindRefreshInterval)
				So(k1, ShouldNotResemble, k2)
			})
		})

		Convey("makeTQTitleForHumans works", func() {
			So(makeTQTitleForHumans(&UpdateCLTask{
				LuciProject: "proj",
				Id:          123,
			}), ShouldResemble, "proj/123")
			So(makeTQTitleForHumans(&UpdateCLTask{
				LuciProject: "proj",
				ExternalId:  "kind/xyz/44",
				Id:          123,
			}), ShouldResemble, "proj/123/kind/xyz/44")
			So(makeTQTitleForHumans(&UpdateCLTask{
				LuciProject: "proj",
				ExternalId:  "gerrit/chromium-review.googlesource.com/1111111",
				Id:          123,
			}), ShouldResemble, "proj/123/gerrit/chromium/1111111")
			So(makeTQTitleForHumans(&UpdateCLTask{
				LuciProject: "proj",
				ExternalId:  "gerrit/chromium-review.googlesource.com/1111111",
				UpdatedHint: timestamppb.New(testclock.TestRecentTimeUTC),
			}), ShouldResemble, "proj/gerrit/chromium/1111111/u2016-02-03T04:05:06Z")
		})

		Convey("Works overall", func() {
			u := NewUpdater(ct.TQDispatcher, nil)
			t := &UpdateCLTask{
				LuciProject: "proj",
				Id:          123,
				UpdatedHint: timestamppb.New(ct.Clock.Now().Add(-time.Second)),
			}
			delay := time.Minute
			So(u.ScheduleDelayed(ctx, t, delay), ShouldBeNil)
			So(ct.TQ.Tasks().Payloads(), ShouldResembleProto, []proto.Message{t})

			_, _ = Println("Dedup works")
			ct.Clock.Add(delay)
			So(u.Schedule(ctx, t), ShouldBeNil)
			So(ct.TQ.Tasks().Payloads(), ShouldHaveLength, 1)

			_, _ = Println("But not within the transaction")
			err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				return u.Schedule(ctx, t)
			}, nil)
			So(err, ShouldBeNil)
			So(ct.TQ.Tasks().Payloads(), ShouldResembleProto, []proto.Message{t, t})

			_, _ = Println("Once out of dedup window, schedules a new task")
			ct.Clock.Add(knownRefreshInterval)
			So(u.Schedule(ctx, t), ShouldBeNil)
			So(ct.TQ.Tasks().Payloads(), ShouldResembleProto, []proto.Message{t, t, t})
		})
	})
}
