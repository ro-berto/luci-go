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

package ledcmd

import (
	"fmt"
	"net/http"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	swarmbucket "go.chromium.org/luci/common/api/buildbucket/swarmbucket/v1"
	swarming "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/grpc/prpc"
)

func newSwarmClient(authClient *http.Client, host string) *swarming.Service {
	swarm, err := swarming.New(authClient)
	if err != nil {
		panic(err)
	}
	swarm.BasePath = fmt.Sprintf("https://%s/_ah/api/swarming/v1/", host)
	return swarm
}

func newSwarmbucketClient(authClient *http.Client, host string) *swarmbucket.Service {
	// TODO(iannucci): Switch this to prpc endpoints
	sbucket, err := swarmbucket.New(authClient)
	if err != nil {
		panic(err)
	}
	sbucket.BasePath = fmt.Sprintf("https://%s/_ah/api/swarmbucket/v1/", host)
	return sbucket
}

func newBuildbucketClient(authClient *http.Client, host string) bbpb.BuildsClient {
	return bbpb.NewBuildsPRPCClient(&prpc.Client{
		C:    authClient,
		Host: host,
	})
}
