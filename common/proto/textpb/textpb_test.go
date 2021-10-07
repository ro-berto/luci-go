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

package textpb

import (
	"testing"

	"go.chromium.org/luci/common/proto/textpb/internal"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestFormat(t *testing.T) {
	Convey("Works", t, func() {
		leaf := &internal.Leaf{
			Str:  "{\n\"not a json really\":\n\"zzz\"\n}",
			Json: "{\n\"this is json 1\":\n\"zzz\"\n}",
			JsonRep: []string{
				"{\n\"this is json 2\":\n\"zzz\"\n}",
				"{\n\"this is json 3\":\n\"zzz\"\n}",
			},
		}

		msg := internal.Container{
			Leaf:    leaf,
			LeafRep: []*internal.Leaf{leaf, leaf},
			Mapping: map[string]*internal.Leaf{"key": leaf},
		}

		out, err := Marshal(&msg)
		So(err, ShouldBeNil)
		So(string(out), ShouldEqual, `leaf {
  str: "{\n\"not a json really\":\n\"zzz\"\n}"
  json:
    '{'
    '  "this is json 1": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 2": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 3": "zzz"'
    '}'
}
leaf_rep {
  str: "{\n\"not a json really\":\n\"zzz\"\n}"
  json:
    '{'
    '  "this is json 1": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 2": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 3": "zzz"'
    '}'
}
leaf_rep {
  str: "{\n\"not a json really\":\n\"zzz\"\n}"
  json:
    '{'
    '  "this is json 1": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 2": "zzz"'
    '}'
  json_rep:
    '{'
    '  "this is json 3": "zzz"'
    '}'
}
mapping {
  key: "key"
  value {
    str: "{\n\"not a json really\":\n\"zzz\"\n}"
    json:
      '{'
      '  "this is json 1": "zzz"'
      '}'
    json_rep:
      '{'
      '  "this is json 2": "zzz"'
      '}'
    json_rep:
      '{'
      '  "this is json 3": "zzz"'
      '}'
  }
}
`)
	})

	Convey("Malformed JSON", t, func() {
		_, err := Format([]byte(`
			str: "nah don't care about that one"
			json: "not a json, boom"
		`), (&internal.Leaf{}).ProtoReflect().Descriptor())
		So(err, ShouldErrLike, "value for 'json' must be valid JSON, got value 'not a json, boom'")
	})
}
