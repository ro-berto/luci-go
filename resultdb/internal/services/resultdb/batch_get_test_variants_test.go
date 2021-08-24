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

package resultdb

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/resultdb/internal/testutil"
	"go.chromium.org/luci/resultdb/internal/testutil/insert"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func variantHash(pairs ...string) string {
	return pbutil.VariantHash(pbutil.Variant(pairs...))
}

func getTvStrings(tvs []*pb.TestVariant) []string {
	tvStrings := make([]string, len(tvs))
	for i, tv := range tvs {
		tvStrings[i] = fmt.Sprintf("%d/%s/%s", int32(tv.Status), tv.TestId, tv.VariantHash)
	}
	return tvStrings
}

func TestBatchGetTestVariants(t *testing.T) {
	Convey(`BatchGetTestVariants`, t, func() {
		ctx := auth.WithState(testutil.SpannerTestContext(t), &authtest.FakeState{
			Identity: "user:someone@example.com",
			IdentityPermissions: []authtest.RealmPermission{
				{Realm: "testproject:testrealm", Permission: permListTestResults},
			},
		})

		testutil.MustApply(ctx, insert.Invocation("i0", pb.Invocation_ACTIVE, map[string]interface{}{"Realm": "testproject:testrealm"}))
		testutil.MustApply(ctx, insert.Invocation("i1", pb.Invocation_ACTIVE, map[string]interface{}{"Realm": "testproject:testrealm"}))
		testutil.MustApply(ctx, testutil.CombineMutations(
			insert.TestResults("i0", "test1", pbutil.Variant("a", "b"), pb.TestStatus_PASS),
			insert.TestResults("i0", "test2", pbutil.Variant("c", "d"), pb.TestStatus_PASS),
			insert.TestResults("i0", "test3", pbutil.Variant("a", "b"), pb.TestStatus_FAIL),
			insert.TestResults("i0", "test4", pbutil.Variant("g", "h"), pb.TestStatus_SKIP),
			insert.TestResults("i1", "test1", pbutil.Variant("e", "f"), pb.TestStatus_PASS),
			insert.TestResults("i1", "test3", pbutil.Variant("c", "d"), pb.TestStatus_PASS),
		)...)

		srv := &resultDBServer{}

		Convey(`Valid request`, func() {
			res, err := srv.BatchGetTestVariants(ctx, &pb.BatchGetTestVariantsRequest{
				Invocation: "invocations/i0",
				TestVariants: []*pb.BatchGetTestVariantsRequest_TestVariantIdentifier{
					{TestId: "test1", VariantHash: variantHash("a", "b")},
					{TestId: "test3", VariantHash: variantHash("a", "b")},
					{TestId: "test4", VariantHash: variantHash("g", "h")},
				},
			})
			So(err, ShouldBeNil)

			// NOTE: The order isn't important here, we just don't have a
			// matcher that does an unordered comparison.
			So(getTvStrings(res.TestVariants), ShouldResemble, []string{
				fmt.Sprintf("10/test3/%s", variantHash("a", "b")),
				fmt.Sprintf("20/test4/%s", variantHash("g", "h")),
				fmt.Sprintf("50/test1/%s", variantHash("a", "b")),
			})
		})

		Convey(`Requesting > 500 variants fails`, func() {
			req := pb.BatchGetTestVariantsRequest{
				Invocation:   "invocations/i0",
				TestVariants: make([]*pb.BatchGetTestVariantsRequest_TestVariantIdentifier, 501),
			}
			for i := 0; i < 500; i += 1 {
				req.TestVariants[i] = &pb.BatchGetTestVariantsRequest_TestVariantIdentifier{
					TestId:      "test1",
					VariantHash: variantHash("a", "b"),
				}
			}

			_, err := srv.BatchGetTestVariants(ctx, &req)
			So(err, ShouldHaveAppStatus, codes.InvalidArgument)
		})

		Convey(`Request including missing variants omits said variants`, func() {
			res, err := srv.BatchGetTestVariants(ctx, &pb.BatchGetTestVariantsRequest{
				Invocation: "invocations/i0",
				TestVariants: []*pb.BatchGetTestVariantsRequest_TestVariantIdentifier{
					{TestId: "test1", VariantHash: variantHash("a", "b")},
					{TestId: "test1", VariantHash: variantHash("x", "y")},
				},
			})
			So(err, ShouldBeNil)

			So(getTvStrings(res.TestVariants), ShouldResemble, []string{
				fmt.Sprintf("50/test1/%s", variantHash("a", "b")),
			})
		})

		Convey(`Request doesn't return variants from other invocations`, func() {
			res, err := srv.BatchGetTestVariants(ctx, &pb.BatchGetTestVariantsRequest{
				Invocation: "invocations/i0",
				TestVariants: []*pb.BatchGetTestVariantsRequest_TestVariantIdentifier{
					{TestId: "test1", VariantHash: variantHash("e", "f")},
					{TestId: "test3", VariantHash: variantHash("c", "d")},
				},
			})
			So(err, ShouldBeNil)

			So(res.TestVariants, ShouldBeEmpty)
		})

		Convey(`Request combines test ID and variant hash correctly`, func() {
			res, err := srv.BatchGetTestVariants(ctx, &pb.BatchGetTestVariantsRequest{
				Invocation: "invocations/i0",
				TestVariants: []*pb.BatchGetTestVariantsRequest_TestVariantIdentifier{
					{TestId: "test1", VariantHash: variantHash("a", "b")},
					{TestId: "test3", VariantHash: variantHash("c", "d")},
				},
			})
			So(err, ShouldBeNil)

			// Testing that we don't match test3, a:b, even though we've
			// requested that test id and variant hash separately.
			So(getTvStrings(res.TestVariants), ShouldResemble, []string{
				fmt.Sprintf("50/test1/%s", variantHash("a", "b")),
			})
		})
	})
}
