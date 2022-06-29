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

package util

import (
	"fmt"
	"testing"
	"time"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsActionTakenOnGerritCL(t *testing.T) {
	t.Parallel()

	Convey("IsActionTakenOnGerritCL works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const (
			lProject = "infra"
			gHost    = "g-review.example.com"
			gChange  = 1111
		)
		var runID = common.MakeRunID(lProject, ct.Clock.Now(), 1, []byte("deadbeef"))

		// CL is on PS#2 in Gerrit while PS#1 in datastore
		ciInCV := gf.CI(gChange, gf.PS(1))
		revisionInCV := ciInCV.GetCurrentRevision()
		cl := changelist.MustGobID(gHost, gChange).MustCreateIfNotExists(ctx)
		cl.Snapshot = &changelist.Snapshot{
			Kind: &changelist.Snapshot_Gerrit{Gerrit: &changelist.Gerrit{
				Host: gHost,
				Info: ciInCV,
			}},
			ExternalUpdateTime: timestamppb.New(ct.Clock.Now().Add(-1 * time.Minute)),
		}
		cl.EVersion++
		rcl := &run.RunCL{
			ID:         cl.ID,
			ExternalID: cl.ExternalID,
			IndexedID:  cl.ID,
			Run:        datastore.MakeKey(ctx, common.RunKind, string(runID)),
			Detail:     cl.Snapshot,
		}
		So(datastore.Put(ctx, cl, rcl), ShouldBeNil)

		ciInGerrit := proto.Clone(ciInCV).(*gerritpb.ChangeInfo)
		gf.PS(2)(ciInGerrit)
		revisionInGerrit := ciInGerrit.GetCurrentRevision()
		ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLRestricted(lProject), ciInGerrit))

		Convey("Action already taken for CL in CV", func() {
			expectedActionTime := ct.Clock.Now().Add(-3 * time.Minute)

			actionTime, err := IsActionTakenOnGerritCL(ctx, ct.GFake, rcl,
				[]gerritpb.QueryOption{gerritpb.QueryOption_CURRENT_REVISION},
				func(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
					So(ci.GetCurrentRevision(), ShouldEqual, revisionInCV)
					return expectedActionTime
				})
			So(err, ShouldBeNil)
			So(actionTime, ShouldEqual, expectedActionTime)
		})

		Convey("Action taken for CL in Gerrit", func() {
			expectedActionTime := ct.Clock.Now().Add(-1 * time.Minute)

			actionTime, err := IsActionTakenOnGerritCL(ctx, ct.GFake, rcl,
				[]gerritpb.QueryOption{gerritpb.QueryOption_CURRENT_REVISION},
				func(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
					switch ci.GetCurrentRevision() {
					case revisionInCV:
						return time.Time{}
					case revisionInGerrit:
						return expectedActionTime
					default:
						panic(fmt.Errorf("impossible revision %s", ci.GetCurrentRevision()))
					}
				})
			So(err, ShouldBeNil)
			So(actionTime, ShouldEqual, expectedActionTime)
		})

		Convey("Action not even taken for CL in Gerrit", func() {
			actionTime, err := IsActionTakenOnGerritCL(ctx, ct.GFake, rcl,
				[]gerritpb.QueryOption{gerritpb.QueryOption_CURRENT_REVISION},
				func(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
					switch ci.GetCurrentRevision() {
					case revisionInCV:
						return time.Time{}
					case revisionInGerrit:
						return time.Time{}
					default:
						panic(fmt.Errorf("impossible revision %s", ci.GetCurrentRevision()))
					}
				})
			So(err, ShouldBeNil)
			So(actionTime.IsZero(), ShouldBeTrue)
		})

		Convey("Don't go to Gerrit if CL is refreshed recently", func() {
			cl.Snapshot.ExternalUpdateTime = timestamppb.New(ct.Clock.Now().Add(-StaleCLAgeThreshold / 2))
			So(datastore.Put(ctx, cl), ShouldBeNil)

			actionTime, err := IsActionTakenOnGerritCL(ctx, ct.GFake, rcl,
				[]gerritpb.QueryOption{gerritpb.QueryOption_CURRENT_REVISION},
				func(rcl *run.RunCL, ci *gerritpb.ChangeInfo) time.Time {
					So(ci.GetCurrentRevision(), ShouldEqual, revisionInCV)
					return time.Time{}
				})
			So(err, ShouldBeNil)
			So(actionTime.IsZero(), ShouldBeTrue)
		})

	})
}
