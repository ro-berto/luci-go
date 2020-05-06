// Copyright 2019 The LUCI Authors.
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

package resultdb

import (
	"testing"
	"time"

	"go.chromium.org/luci/common/clock/testclock"

	"go.chromium.org/luci/resultdb/internal/testutil"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateGetInvocationRequest(t *testing.T) {
	t.Parallel()
	Convey(`ValidateGetInvocationRequest`, t, func() {
		Convey(`Valid`, func() {
			req := &pb.GetInvocationRequest{Name: "invocations/valid_id_0"}
			So(validateGetInvocationRequest(req), ShouldBeNil)
		})

		Convey(`Invalid name`, func() {
			Convey(`, missing`, func() {
				req := &pb.GetInvocationRequest{}
				So(validateGetInvocationRequest(req), ShouldErrLike, "name missing")
			})

			Convey(`, invalid format`, func() {
				req := &pb.GetInvocationRequest{Name: "bad_name"}
				So(validateGetInvocationRequest(req), ShouldErrLike, "does not match")
			})
		})
	})
}

func TestGetInvocation(t *testing.T) {
	Convey(`GetInvocation`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		ct := testclock.TestRecentTimeUTC
		deadline := ct.Add(time.Hour)

		// Insert some Invocations.
		testutil.MustApply(ctx,
			testutil.InsertInvocation("including", pb.Invocation_ACTIVE, map[string]interface{}{
				"CreateTime": ct,
				"Deadline":   deadline,
			}),
			testutil.InsertInvocation("included0", pb.Invocation_FINALIZED, nil),
			testutil.InsertInvocation("included1", pb.Invocation_FINALIZED, nil),
			testutil.InsertInclusion("including", "included0"),
			testutil.InsertInclusion("including", "included1"),
		)

		// Fetch back the top-level Invocation.
		srv := newTestResultDBService()
		req := &pb.GetInvocationRequest{Name: "invocations/including"}
		inv, err := srv.GetInvocation(ctx, req)
		So(err, ShouldBeNil)
		So(inv, ShouldResembleProto, &pb.Invocation{
			Name:                "invocations/including",
			State:               pb.Invocation_ACTIVE,
			CreateTime:          pbutil.MustTimestampProto(ct),
			Deadline:            pbutil.MustTimestampProto(deadline),
			IncludedInvocations: []string{"invocations/included0", "invocations/included1"},
		})
	})
}
