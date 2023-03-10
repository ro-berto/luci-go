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

package tryjob

import (
	"testing"

	"go.chromium.org/luci/cv/internal/common"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestCLPatchset(t *testing.T) {
	t.Parallel()

	Convey("CLPatchset works", t, func() {
		var cl common.CLID
		var ps int32
		Convey("with max values", func() {
			cl, ps = 0x7fff_ffff_ffff_ffff, 0x7fff_ffff
		})
		Convey("with min values", func() {
			cl, ps = 0, 0
		})
		clps := MakeCLPatchset(cl, ps)
		parsedCl, parsedPs, err := clps.Parse()
		So(err, ShouldBeNil)
		So(parsedCl, ShouldEqual, cl)
		So(parsedPs, ShouldEqual, ps)
	})
	Convey("CLPatchset fails", t, func() {
		Convey("with bad number of values", func() {
			_, _, err := CLPatchset("1/1").Parse()
			So(err, ShouldErrLike, "CLPatchset in unexpected format")
		})
		Convey("with bad version", func() {
			_, _, err := CLPatchset("8/8/8").Parse()
			So(err, ShouldErrLike, "unsupported version")
		})
		Convey("with bad CLID", func() {
			_, _, err := CLPatchset("1/4d35683b24371b75c5f3fda0d48796638dc0d695/7").Parse()
			So(err, ShouldErrLike, "clid segment in unexpected format")
			_, _, err = CLPatchset("1/gerrit/chromium-review.googlesource.com/3530834/7").Parse()
			So(err, ShouldErrLike, "unexpected format")
		})
		Convey("with bad patchset", func() {
			_, _, err := CLPatchset("1/1/ps1").Parse()
			So(err, ShouldErrLike, "patchset segment in unexpected format")
		})
	})
}
