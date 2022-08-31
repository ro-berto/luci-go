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

package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExtractorUtil(t *testing.T) {
	Convey("NormalizeFilePath", t, func() {
		data := map[string]string{
			"../a/b/c.cc":    "a/b/c.cc",
			"a/b/./c.cc":     "a/b/c.cc",
			"a/b/../c.cc":    "a/c.cc",
			"a\\b\\.\\c.cc":  "a/b/c.cc",
			"a\\\\b\\\\c.cc": "a/b/c.cc",
			"//a/b/c.cc":     "a/b/c.cc",
		}
		for fp, nfp := range data {
			So(NormalizeFilePath(fp), ShouldEqual, nfp)
		}
	})

	Convey("GetCanonicalFileName", t, func() {
		data := map[string]string{
			"../a/b/c.cc":   "c",
			"a/b/./d.dd":    "d",
			"a/b/c.xx":      "c",
			"a/b/c_impl.xx": "c",
		}
		for fp, name := range data {
			So(GetCanonicalFileName(fp), ShouldEqual, name)
		}
	})

	Convey("StripExtensionAndCommonSuffixFromFileName", t, func() {
		data := map[string]string{
			"a_file_impl_mac_test.cc": "a_file",
			"src/b_file_x11_ozone.h":  "src/b_file",
			"c_file.cc":               "c_file",
		}
		for k, v := range data {
			So(StripExtensionAndCommonSuffixFromFileName(k), ShouldEqual, v)
		}
	})

}
