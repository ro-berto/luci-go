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
	"go.chromium.org/luci/common/tsmon/field"
	"go.chromium.org/luci/common/tsmon/metric"
	"go.chromium.org/luci/common/tsmon/types"
)

var (
	// V2 is a collection of metric objects for V2 metrics.
	V2 = struct {
		BuildCount      metric.Int
		BuilderPresence metric.Bool
		MaxAgeScheduled metric.Float
	}{
		BuildCount: metric.NewIntWithTargetType(
			"buildbucket/v2/builds/count",
			(&Builder{}).Type(),
			"Number of pending/running prod builds",
			nil,
			field.String("status"),
		),
		BuilderPresence: metric.NewBoolWithTargetType(
			"buildbucket/v2/builder/presence",
			(&Builder{}).Type(),
			"A constant, always-true metric that indicates the presence of LUCI Builder",
			nil,
		),
		MaxAgeScheduled: metric.NewFloatWithTargetType(
			"buildbucket/v2/builds/max_age_scheduled",
			(&Builder{}).Type(),
			"Age of the oldest SCHEDULED build",
			&types.MetricMetadata{Units: types.Seconds},
		),
	}
)
