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

package submit

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/memlogger"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/run"
	"go.chromium.org/luci/cv/internal/run/eventpb"
	"go.chromium.org/luci/cv/internal/run/runtest"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubmitter(t *testing.T) {
	Convey("Submitter", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const (
			lProject = "test_proj"
			gHost1   = "gerrit-1.example.com"
			gHost2   = "gerrit-2.example.com"
			gHost3   = "gerrit-3.example.com"
		)
		ci1 := gf.CI(1, gf.PS(3), gf.AllRevs(), gf.CQ(2))
		ci2 := gf.CI(2, gf.PS(5), gf.AllRevs(), gf.CQ(2))
		ci3 := gf.CI(3, gf.PS(7), gf.AllRevs(), gf.CQ(2))
		ct.GFake.AddFrom(gf.WithCIs(gHost1, gf.ACLRestricted(lProject), ci1))
		ct.GFake.AddFrom(gf.WithCIs(gHost2, gf.ACLRestricted(lProject), ci2))
		ct.GFake.AddFrom(gf.WithCIs(gHost3, gf.ACLRestricted(lProject), ci3))

		now := ct.Clock.Now().UTC()
		s := RunCLsSubmitter{
			runID:    common.MakeRunID(lProject, now, 1, []byte("deadbeef")),
			deadline: now.Add(1 * time.Minute),
			clids:    common.CLIDs{1, 2, 3},
			rm:       run.NewNotifier(ct.TQDispatcher),
			gFactory: ct.GFactory(),
		}
		cl1 := &run.RunCL{
			ID:         1,
			ExternalID: changelist.MustGobID(gHost1, ci1.GetNumber()),
			Run:        datastore.MakeKey(ctx, run.RunKind, string(s.runID)),
			Detail: &changelist.Snapshot{
				Kind: &changelist.Snapshot_Gerrit{
					Gerrit: &changelist.Gerrit{
						Host: gHost1,
						Info: ci1,
					},
				},
			},
		}
		cl2 := &run.RunCL{
			ID:         2,
			ExternalID: changelist.MustGobID(gHost2, ci2.GetNumber()),
			Run:        datastore.MakeKey(ctx, run.RunKind, string(s.runID)),
			Detail: &changelist.Snapshot{
				Kind: &changelist.Snapshot_Gerrit{
					Gerrit: &changelist.Gerrit{
						Host: gHost2,
						Info: ci2,
					},
				},
			},
		}
		cl3 := &run.RunCL{
			ID:         3,
			ExternalID: changelist.MustGobID(gHost3, ci3.GetNumber()),
			Run:        datastore.MakeKey(ctx, run.RunKind, string(s.runID)),
			Detail: &changelist.Snapshot{
				Kind: &changelist.Snapshot_Gerrit{
					Gerrit: &changelist.Gerrit{
						Host: gHost3,
						Info: ci3,
					},
				},
			},
		}
		So(datastore.Put(ctx,
			&run.Run{
				ID:         s.runID,
				Status:     run.Status_RUNNING,
				CreateTime: now,
				StartTime:  now,
				CLs:        s.clids,
			},
			cl1, cl2, cl3,
		), ShouldBeNil)
		So(datastore.RunInTransaction(ctx, func(ctx context.Context) error {
			waitlisted, err := TryAcquire(ctx, s.rm.NotifyReadyForSubmission, s.runID, nil)
			So(err, ShouldBeNil)
			So(waitlisted, ShouldBeFalse)
			return err
		}, nil), ShouldBeNil)

		verifyRunReleased := func(runID common.RunID) {
			current, waitlist, err := LoadCurrentAndWaitlist(ctx, runID)
			So(err, ShouldBeNil)
			So(current, ShouldNotEqual, runID)
			So(waitlist.Index(runID), ShouldBeLessThan, 0) // doesn't exist
		}

		Convey("Submit successfully", func() {
			So(s.Submit(ctx), ShouldBeNil)
			verifyRunReleased(s.runID)
			runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 1)
			So(ct.GFake.GetChange(gHost1, 1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 2)
			So(ct.GFake.GetChange(gHost2, 2).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 3)
			So(ct.GFake.GetChange(gHost3, 3).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			So(ct.GFake.Requests(), ShouldHaveLength, len(s.clids)) // len(s.clids) SubmitRevision calls
			runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
				&eventpb.SubmissionCompleted{
					Result:                eventpb.SubmissionResult_SUCCEEDED,
					QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
				},
			)
		})

		// TODO(crbug/1199880): support flakiness for Gerrit fake to test submit
		// will retry individual CL on transient error and not release queue
		// for transient failure.
		// Also test that submission has exhausted the allocated time.

		Convey("Submit fails permanently when", func() {
			Convey("Not holding Submit Queue", func() {
				ctx = memlogger.Use(ctx)
				log := logging.Get(ctx).(*memlogger.MemLogger)
				So(datastore.RunInTransaction(ctx, func(ctx context.Context) error {
					return Release(ctx, s.rm.NotifyReadyForSubmission, s.runID)
				}, nil), ShouldBeNil)
				So(s.Submit(ctx), ShouldBeNil)
				runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
					&eventpb.SubmissionCompleted{
						Result:                eventpb.SubmissionResult_FAILED_PERMANENT,
						QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
					},
				)
				So(log, memlogger.ShouldHaveLog, logging.Error, "BUG: run no longer holds submit queue, currently held by")
			})

			Convey("No submit privilege", func() {
				// Submit gHost1/1 successfully but lack of submission right to
				// gHost2/2.
				ct.GFake.MutateChange(gHost2, 2, func(c *gf.Change) {
					c.ACLs = gf.ACLGrant(gf.OpSubmit, codes.PermissionDenied, "another_project")
				})
				So(s.Submit(ctx), ShouldBeNil)
				verifyRunReleased(s.runID)
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 1)
				So(ct.GFake.GetChange(gHost1, 1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 2)
				So(ct.GFake.GetChange(gHost2, 2).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW)
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 3)
				So(ct.GFake.GetChange(gHost3, 3).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW) // 3 not attempted
				runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
					&eventpb.SubmissionCompleted{
						Result: eventpb.SubmissionResult_FAILED_PERMANENT,
						FailureReason: &eventpb.SubmissionCompleted_ClFailures{
							ClFailures: &eventpb.SubmissionCompleted_CLSubmissionFailures{
								Failures: []*eventpb.SubmissionCompleted_CLSubmissionFailure{
									{Clid: 2, Message: permDeniedMsg},
								},
							},
						},
						QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
					},
				)
			})

			Convey("A new revision is uploaded ", func() {
				// gHost2/2 gets a new PS.
				ct.GFake.MutateChange(gHost2, 2, func(c *gf.Change) {
					c.Info = proto.Clone(ci2).(*gerritpb.ChangeInfo)
					gf.PS(6)(c.Info)
				})
				So(s.Submit(ctx), ShouldBeNil)
				verifyRunReleased(s.runID)
				So(ct.GFake.GetChange(gHost1, 1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 2)
				So(ct.GFake.GetChange(gHost2, 2).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW) // 3 not attempted
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 3)
				So(ct.GFake.GetChange(gHost3, 3).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW)
				runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
					&eventpb.SubmissionCompleted{
						Result: eventpb.SubmissionResult_FAILED_PERMANENT,
						FailureReason: &eventpb.SubmissionCompleted_ClFailures{
							ClFailures: &eventpb.SubmissionCompleted_CLSubmissionFailures{
								Failures: []*eventpb.SubmissionCompleted_CLSubmissionFailure{
									{
										Clid:    2,
										Message: fmt.Sprintf(failedPreconditionMsgFmt, fmt.Sprintf("revision %s is not current revision", ci2.GetCurrentRevision())),
									},
								},
							},
						},
						QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
					},
				)
			})
		})

		Convey("Change has already been merged", func() {
			ct.GFake.MutateChange(gHost1, 1, func(c *gf.Change) {
				c.Info = proto.Clone(ci1).(*gerritpb.ChangeInfo)
				gf.Status(gerritpb.ChangeStatus_MERGED)(c.Info)
			})
			// Submitter should receive FailedPrecondition failure from Gerrit
			// for Submit RPC. But the subsequent GetChange will figure out that
			// Change has been merged already and consider submission of gHost1/1
			// as a success.
			So(s.Submit(ctx), ShouldBeNil)
			verifyRunReleased(s.runID)
			runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 1)
			So(ct.GFake.GetChange(gHost1, 1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_MERGED)
			So(ct.GFake.Requests(), ShouldHaveLength, len(s.clids)+1) // 1 extra getChange call
			runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
				&eventpb.SubmissionCompleted{
					Result:                eventpb.SubmissionResult_SUCCEEDED,
					QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
				},
			)
		})

		Convey("CL with dependencies", func() {
			Convey("optmize submission of stack CLs", func() {
				cl2.Detail.Deps = append(cl2.Detail.Deps, &changelist.Dep{
					Clid: int64(cl1.ID),
					Kind: changelist.DepKind_HARD,
				})
				So(datastore.Put(ctx, cl2), ShouldBeNil)
				So(s.Submit(ctx), ShouldBeNil)
				verifyRunReleased(s.runID)
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 1, 2) // submitted together
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 3)
				So(ct.GFake.Requests(), ShouldHaveLength, len(s.clids)-1) // no submit call for 1
			})

			Convey("not optimize if cls don't form a chain", func() {
				cl2.Detail.Deps = append(cl2.Detail.Deps, &changelist.Dep{
					Clid: int64(cl1.ID),
					Kind: changelist.DepKind_SOFT,
				})
				cl3.Detail.Deps = append(cl3.Detail.Deps, &changelist.Dep{
					Clid: int64(cl1.ID),
					Kind: changelist.DepKind_HARD,
				})
				So(datastore.Put(ctx, cl2, cl3), ShouldBeNil)
				// 1 is parent of 3 and 2 is not. However, 2 need to be submitted before
				// 3 but after 1. Therefore, optimization for stack shouldn't take
				// effect. CLs have to be submitted one by one.
				So(s.Submit(ctx), ShouldBeNil)
				verifyRunReleased(s.runID)
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 1)
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 2)
				runtest.AssertReceivedCLsSubmitted(ctx, s.runID, 3)
				So(ct.GFake.Requests(), ShouldHaveLength, len(s.clids))
			})

			Convey("failed to submit top cl in the stack", func() {
				cl2.Detail.Deps = append(cl2.Detail.Deps, &changelist.Dep{
					Clid: int64(cl1.ID),
					Kind: changelist.DepKind_HARD,
				})
				So(datastore.Put(ctx, cl2), ShouldBeNil)
				ct.GFake.MutateChange(gHost2, 2, func(c *gf.Change) {
					c.ACLs = gf.ACLGrant(gf.OpSubmit, codes.PermissionDenied, "another_project")
				})
				So(s.Submit(ctx), ShouldBeNil)
				verifyRunReleased(s.runID)
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 1, 2)
				So(ct.GFake.GetChange(gHost1, 1).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW)
				So(ct.GFake.GetChange(gHost2, 2).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW)
				runtest.AssertNotReceivedCLsSubmitted(ctx, s.runID, 3)
				So(ct.GFake.GetChange(gHost3, 3).Info.GetStatus(), ShouldEqual, gerritpb.ChangeStatus_NEW)
				runtest.AssertReceivedSubmissionCompleted(ctx, s.runID,
					&eventpb.SubmissionCompleted{
						Result: eventpb.SubmissionResult_FAILED_PERMANENT,
						FailureReason: &eventpb.SubmissionCompleted_ClFailures{
							ClFailures: &eventpb.SubmissionCompleted_CLSubmissionFailures{
								Failures: []*eventpb.SubmissionCompleted_CLSubmissionFailure{
									{
										Clid:    1,
										Message: fmt.Sprintf(topOfStackFailureMsgFmt, cl2.ExternalID.MustURL()),
									},
									{
										Clid:    2,
										Message: permDeniedMsg,
									},
								},
							},
						},
						QueueReleaseTimestamp: timestamppb.New(clock.Now(ctx)),
					},
				)
			})
		})
	})
}
