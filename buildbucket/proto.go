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

package buildbucket

import (
	"go.chromium.org/luci/common/data/stringset"

	pb "go.chromium.org/luci/buildbucket/proto"
)

// This file contains helper functions for pb package.
// TODO(nodir): move existing helpers from pb to this file.

// BuildTokenHeader is the name of gRPC metadata header indicating the build
// token (see BuildSecrets.BuildToken).
// It is required in UpdateBuild RPC.
// Defined in
// https://chromium.googlesource.com/infra/infra/+/c189064/appengine/cr-buildbucket/v2/api.py#35
const BuildTokenHeader = "x-build-token"

// Well-known experiment strings.
//
// See the Builder.experiments field documentation.
const (
	ExperimentBBCanarySoftware = "luci.buildbucket.canary_software"
	ExperimentNonProduction    = "luci.non_production"

	ExperimentBBAgent   = "luci.buildbucket.use_bbagent"
	ExperimentRecipePY3 = "luci.recipes.use_python3"
	ExperimentUseRealms = "luci.use_realms"
)

// WellKnownExperiments is the list of all well-known experiments.
var WellKnownExperiments = stringset.NewFromSlice(
	ExperimentBBCanarySoftware,
	ExperimentNonProduction,
	ExperimentBBAgent,
	ExperimentRecipePY3,
	ExperimentUseRealms,
)

var (
	// DisallowedAppendTagKeys is the set of tag keys which cannot be set via
	// UpdateBuild. Clients calling UpdateBuild must strip these before making
	// the request.
	DisallowedAppendTagKeys = stringset.NewFromSlice("build_address", "buildset", "builder")
)

// StripDisallowedTagKeys modifies `tags` in-place to remove tags with
// DisallowedAppendTagKeys keys.
//
// This does not preserve the order of `tags`.
func StripDisallowedTagKeys(tags *[]*pb.StringPair) {
	if tags == nil {
		return
	}

	ts := *tags

	for i := len(ts) - 1; i >= 0; i-- {
		if DisallowedAppendTagKeys.Has(ts[i].Key) {
			ts[i] = ts[len(ts)-1]
			ts[len(ts)-1] = nil // allow pruned StringPair to be gc'd.
			ts = ts[:len(ts)-1]
		}
	}

	*tags = ts
}
