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

package job

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestEditCIPDPkgs(t *testing.T) {
	t.Parallel()

	runCases(t, "EditCIPDPkgs", []testCase{
		{
			name: "nil",
			fn: func(jd *Definition) {
				SoEdit(jd, func(je Editor) {
					je.CIPDPkgs(nil)
				})
				So(must(jd.Info().CIPDPkgs()), ShouldBeEmpty)
			},
		},

		{
			name: "add",
			fn: func(jd *Definition) {
				SoEdit(jd, func(je Editor) {
					je.CIPDPkgs(CIPDPkgs{
						"subdir:some/pkg":       "version",
						"other_subdir:some/pkg": "different_version",
						"some/other/pkg":        "latest",
						":more/other/pkg":       "whatever",
					})
				})
				if sw := jd.GetSwarming(); sw != nil && len(sw.GetTask().GetTaskSlices()) == 0 {
					So(must(jd.Info().CIPDPkgs()), ShouldBeEmpty)
				} else {
					So(must(jd.Info().CIPDPkgs()), ShouldResemble, CIPDPkgs{
						"subdir:some/pkg":       "version",
						"other_subdir:some/pkg": "different_version",
						"some/other/pkg":        "latest",
						"more/other/pkg":        "whatever",
					})
				}
			},
		},

		{
			name: "delete",
			fn: func(jd *Definition) {
				SoEdit(jd, func(je Editor) {
					je.CIPDPkgs(CIPDPkgs{
						"subdir:some/pkg":       "version",
						"other_subdir:some/pkg": "different_version",
						"some/other/pkg":        "latest",
					})
				})
				SoEdit(jd, func(je Editor) {
					je.CIPDPkgs(CIPDPkgs{
						"subdir:some/pkg": "",
						"some/other/pkg":  "",
					})
				})

				if sw := jd.GetSwarming(); sw != nil && len(sw.GetTask().GetTaskSlices()) == 0 {
					So(must(jd.Info().CIPDPkgs()), ShouldBeEmpty)
				} else {
					So(must(jd.Info().CIPDPkgs()), ShouldResemble, CIPDPkgs{
						"other_subdir:some/pkg": "different_version",
					})
				}
			},
		},

		{
			name: "edit v2 build",
			fn: func(jd *Definition) {
				if b := jd.GetBuildbucket(); b != nil && b.BbagentDownloadCIPDPkgs() {
					err := jd.Edit(func(je Editor) {
						je.CIPDPkgs(CIPDPkgs{
							"subdir:some/pkg": "version",
						})
					})
					So(err, ShouldErrLike, "not supported for Buildbucket v2 builds")
					return
				}
			},
			v2Build: true,
		},
	})
}
