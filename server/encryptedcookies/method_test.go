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

package encryptedcookies

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNormalizeURL(t *testing.T) {
	t.Parallel()

	Convey("Normalizes good URLs", t, func(ctx C) {
		cases := []struct {
			in  string
			out string
		}{
			{"", "/"},
			{"/", "/"},
			{"/?asd=def#blah", "/?asd=def#blah"},
			{"/abc/def", "/abc/def"},
			{"/blah//abc///def/", "/blah/abc/def/"},
			{"/blah/..//./abc/", "/abc/"},
			{"/abc/%2F/def", "/abc/def"},
		}
		for _, c := range cases {
			out, err := normalizeURL(c.in)
			if err != nil {
				ctx.Printf("Failed while checking %q\n", c.in)
				So(err, ShouldBeNil)
			}
			So(out, ShouldEqual, c.out)
		}
	})

	Convey("Rejects bad URLs", t, func(ctx C) {
		cases := []string{
			"//",
			"///",
			"://",
			":",
			"http://another/abc/def",
			"abc/def",
			"//host.example.com",
		}
		for _, c := range cases {
			_, err := normalizeURL(c)
			if err == nil {
				ctx.Printf("Didn't fail while testing %q\n", c)
			}
			So(err, ShouldNotBeNil)
		}
	})
}
