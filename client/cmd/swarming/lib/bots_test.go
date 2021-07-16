// Copyright 2018 The LUCI Authors.
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

package lib

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func botsExpectErr(argv []string, errLike string) {
	b := botsRun{}
	b.Init(&testAuthFlags{})
	fullArgv := append([]string{"-server", "http://localhost:9050"}, argv...)
	err := b.GetFlags().Parse(fullArgv)
	So(err, ShouldBeNil)
	So(b.Parse(), ShouldErrLike, errLike)
}

func TestBotsParse(t *testing.T) {
	Convey(`Make sure that Parse fails with -quiet without -json.`, t, func() {
		botsExpectErr([]string{"-quiet"}, "specify -json")
	})

	Convey(`Make sure that Parse fails with -count and -field.`, t, func() {
		botsExpectErr([]string{"-count", "-field", "myField"}, "-field cannot")
	})
}
