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

package loggingutil

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/bisection/util/testutil"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/gae/impl/memory"
)

func TestLogging(t *testing.T) {
	t.Parallel()
	c := memory.Use(context.Background())

	Convey("Logging", t, func() {
		testutil.CreateCompileFailureAnalysisAnalysisChain(c, 123, 456)
		c, err := UpdateLoggingWithAnalysisID(c, 456)
		So(err, ShouldBeNil)

		// Check the logging fields
		entries := logging.GetFields(c).SortedEntries()
		So(checkEntries(entries, "analyzed_bbid", int64(123)), ShouldBeTrue)
		So(checkEntries(entries, "analysis_id", int64(456)), ShouldBeTrue)
	})
}

func checkEntries(entries []*logging.FieldEntry, k string, v interface{}) bool {
	for _, entry := range entries {
		if entry.Key == k && entry.Value == v {
			return true
		}
	}
	return false
}
