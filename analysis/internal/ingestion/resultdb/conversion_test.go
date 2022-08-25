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

package resultdb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestInvocationFromTestResultName(t *testing.T) {
	Convey("Valid input", t, func() {
		result, err := InvocationFromTestResultName("invocations/build-1234/tests/a/results/b")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "build-1234")
	})
	Convey("Invalid input", t, func() {
		_, err := InvocationFromTestResultName("")
		So(err, ShouldErrLike, "invalid test result name")

		_, err = InvocationFromTestResultName("projects/chromium/resource/b")
		So(err, ShouldErrLike, "invalid test result name")

		_, err = InvocationFromTestResultName("invocations/build-1234")
		So(err, ShouldErrLike, "invalid test result name")

		_, err = InvocationFromTestResultName("invocations//")
		So(err, ShouldErrLike, "invalid test result name")

		_, err = InvocationFromTestResultName("invocations/")
		So(err, ShouldErrLike, "invalid test result name")

		_, err = InvocationFromTestResultName("invocations")
		So(err, ShouldErrLike, "invalid test result name")
	})
}
