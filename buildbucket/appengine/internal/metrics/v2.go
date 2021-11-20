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

package metrics

import (
	"go.chromium.org/luci/common/tsmon/metric"
)

var (
	// V2 is a collection of metric objects for V2 metrics.
	V2 = struct {
		BuilderPresence metric.Bool
	}{
		BuilderPresence: metric.NewBoolWithTargetType(
			"v2/builder/presence",
			(&Builder{}).Type(),
			"A constant, always-true metric that indicates the presence of LUCI Builder",
			nil,
		),
	}
)