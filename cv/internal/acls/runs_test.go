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

package acls

import (
	"testing"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/common"
	"go.chromium.org/luci/cv/internal/configs/prjcfg/prjcfgtest"
	"go.chromium.org/luci/cv/internal/configs/validation"
	"go.chromium.org/luci/cv/internal/cvtesting"
	"go.chromium.org/luci/cv/internal/run"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestLoadRun(t *testing.T) {
	t.Parallel()

	Convey("Check NewRunReadChecker works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()

		const projectPublic = "infra"
		const projectInternal = "infra-internal"

		prjcfgtest.Create(ctx, projectPublic, &cfgpb.Config{
			// TODO(crbug/1233963): update this test to stop relying on legacy-based
			// ACL.
			CqStatusHost: validation.CQStatusHostPublic,
			ConfigGroups: []*cfgpb.ConfigGroup{{
				Name: "first",
			}},
		})
		prjcfgtest.Create(ctx, projectInternal, &cfgpb.Config{
			// TODO(crbug/1233963): update this test to stop relying on legacy-based
			// ACL.
			CqStatusHost: validation.CQStatusHostInternal,
			ConfigGroups: []*cfgpb.ConfigGroup{{
				Name: "first",
			}},
		})

		publicRun := &run.Run{
			ID:            common.RunID(projectPublic + "/123-visible"),
			EVersion:      5,
			ConfigGroupID: prjcfgtest.MustExist(ctx, projectPublic).ConfigGroupIDs[0],
		}
		internalRun := &run.Run{
			ID:            common.RunID(projectInternal + "/456-invisible"),
			EVersion:      5,
			ConfigGroupID: prjcfgtest.MustExist(ctx, projectInternal).ConfigGroupIDs[0],
		}
		So(datastore.Put(ctx, publicRun, internalRun), ShouldBeNil)

		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
		})

		Convey("Run doesn't exist", func() {
			_, err1 := run.LoadRun(ctx, common.RunID("foo/bar"), NewRunReadChecker())
			So(err1, ShouldHaveAppStatus, codes.NotFound)

			Convey("No access must be indistinguishable from not existing Run", func() {
				_, err2 := run.LoadRun(ctx, internalRun.ID, NewRunReadChecker())
				So(err2, ShouldHaveAppStatus, codes.NotFound)

				st1, _ := appstatus.Get(err1)
				st2, _ := appstatus.Get(err2)
				So(st1.Message(), ShouldResemble, st2.Message())
				So(st1.Details(), ShouldBeEmpty)
				So(st2.Details(), ShouldBeEmpty)
			})
		})

		Convey("OK public", func() {
			r, err := run.LoadRun(ctx, publicRun.ID, NewRunReadChecker())
			So(err, ShouldBeNil)
			So(r, cvtesting.SafeShouldResemble, publicRun)
		})

		Convey("OK internal", func() {
			// TODO(crbug/1233963): add a test once non-legacy ACLs are working.
			ctx = auth.WithState(ctx, &authtest.FakeState{
				Identity:       "user:googler@example.com",
				IdentityGroups: []string{"googlers"},
			})
			r, err := run.LoadRun(ctx, internalRun.ID, NewRunReadChecker())
			So(err, ShouldBeNil)
			So(r, cvtesting.SafeShouldResemble, internalRun)
		})

		Convey("PermissionDenied", func() {
			ctx = auth.WithState(ctx, &authtest.FakeState{
				Identity:       "user:public-user@example.com",
				IdentityGroups: []string{"insufficient"},
			})
			_, err := run.LoadRun(ctx, internalRun.ID, NewRunReadChecker())
			So(err, ShouldHaveAppStatus, codes.NotFound)
		})
	})
}
