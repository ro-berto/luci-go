// Copyright 2015 The LUCI Authors.
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

package pattern

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPattern(t *testing.T) {
	t.Parallel()

	Convey("Pattern", t, func() {
		Convey("Exact", func() {
			p := Exact("a")
			So(p.Match("a"), ShouldBeTrue)
			So(p.Match("b"), ShouldBeFalse)
			So(p.String(), ShouldEqual, "exact:a")
		})

		Convey("Regex", func() {
			p := Regexp(regexp.MustCompile("^ab?$"))
			So(p.Match("a"), ShouldBeTrue)
			So(p.Match("ab"), ShouldBeTrue)
			So(p.Match("b"), ShouldBeFalse)
			So(p.String(), ShouldEqual, "regex:^ab?$")
		})

		Convey("Any", func() {
			So(Any.Match("a"), ShouldBeTrue)
			So(Any.Match("b"), ShouldBeTrue)
			So(Any.String(), ShouldEqual, "*")
		})

		Convey("None", func() {
			So(None.Match("a"), ShouldBeFalse)
			So(None.Match("b"), ShouldBeFalse)
			So(None.String(), ShouldEqual, "")
		})

		Convey("Parse", func() {
			Convey("Good", func() {
				patterns := []Pattern{
					Exact("a"),
					Regexp(regexp.MustCompile("^ab$")),
					Any,
					None,
				}
				for _, p := range patterns {
					p2, err := Parse(p.String())
					So(err, ShouldBeNil)
					So(p2.String(), ShouldEqual, p.String())
				}
			})

			Convey("Bad", func() {
				bad := []string{
					":",
					"a:",
					"a:b",
					"regex:)",
				}
				for _, s := range bad {
					_, err := Parse(s)
					So(err, ShouldNotBeNil)
				}
			})
		})
	})
}
