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

package buganizer

import (
	"github.com/golang/protobuf/proto"

	configpb "go.chromium.org/luci/analysis/proto/config"
)

func ChromeOSTestConfig() *configpb.BuganizerProject {
	return &configpb.BuganizerProject{
		DefaultComponent: &configpb.BuganizerComponent{
			Id: 1234567,
		},
		PriorityMappings:          createPriorityMappings(),
		PriorityHysteresisPercent: 10,
	}
}

func createPriorityMappings() []*configpb.BuganizerProject_PriorityMapping {
	return []*configpb.BuganizerProject_PriorityMapping{
		{
			Priority: configpb.BuganizerPriority_P0,
			Threshold: &configpb.ImpactThreshold{
				TestResultsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(1000),
				},
				TestRunsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(100),
				},
			},
		},
		{
			Priority: configpb.BuganizerPriority_P1,
			Threshold: &configpb.ImpactThreshold{
				TestResultsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(500),
				},
				TestRunsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(50),
				},
			},
		},
		{
			Priority: configpb.BuganizerPriority_P2,
			Threshold: &configpb.ImpactThreshold{
				TestResultsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(100),
				},
				TestRunsFailed: &configpb.MetricThreshold{
					OneDay: proto.Int64(10),
				},
			},
		},
		{
			Priority: configpb.BuganizerPriority_P3,
			// Should be less onerous than the bug-filing thresholds
			// used in BugUpdater tests, to avoid bugs that were filed
			// from being immediately closed.
			Threshold: &configpb.ImpactThreshold{
				TestResultsFailed: &configpb.MetricThreshold{
					OneDay:   proto.Int64(50),
					ThreeDay: proto.Int64(300),
					SevenDay: proto.Int64(1), // Set to 1 so that we check hysteresis never rounds down to 0 and prevents bugs from closing.
				},
			},
		},
	}
}
