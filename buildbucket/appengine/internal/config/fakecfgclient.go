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

package config

import (
	"context"

	"go.chromium.org/luci/common/errors"
	luciconfig "go.chromium.org/luci/config"
)

var chromiumRevision = "deadbeef"
var chromiumBuildbucketCfg = `
	buckets {
		name: "try"
		swarming {
			task_template_canary_percentage { value: 10 }
			builders {
				name: "linux"
				dimensions: "os:Linux"
				exe {
					cipd_version: "refs/heads/main"
					cipd_package: "infra/recipe_bundle"
					cmd: ["luciexe"]
				}
				swarming_host: "swarming.example.com"
				task_template_canary_percentage {
					value: 10
				}
			}
		}
	}
	buckets {
		name: "master.tryserver.chromium.linux"
	}
	buckets {
		name: "master.tryserver.chromium.win"
	}
`
var dartRevision = "deadbeef"
var dartBuildbucketCfg = `
	buckets {
		name: "try"
		swarming {
			builders {
				name: "linux"
				dimensions: "pool:Dart.LUCI"
				exe {
					cipd_version: "refs/heads/main"
					cipd_package: "infra/recipe_bundle"
					cmd: ["luciexe"]
				}
			}
		}
	}
`
var v8Revision = ""
var v8BuildbucketCfg = `
	buckets {
		name: "master.tryserver.v8"
	}
`

// fakeCfgClient mocks the luciconfig.Interface.
type fakeCfgClient struct {
	luciconfig.Interface
}

func (*fakeCfgClient) GetProjectConfigs(ctx context.Context, path string, metaOnly bool) ([]luciconfig.Config, error) {
	if path == "${appid}.cfg" {
		chromiumCfg := luciconfig.Config{
			Meta: luciconfig.Meta{
				ConfigSet:"projects/chromium",
				Path: "fake-cr-buildbucket.cfg",
				Revision: chromiumRevision,
			},
			Content: chromiumBuildbucketCfg,
		}
		dartCfg := luciconfig.Config{
			Meta: luciconfig.Meta{
				ConfigSet:"projects/dart",
				Path: "fake-cr-buildbucket.cfg",
				Revision: dartRevision,
			},
			Content: dartBuildbucketCfg,
		}
		v8Cfg := luciconfig.Config{
			Meta: luciconfig.Meta{
				ConfigSet:"projects/v8",
				Path: "fake-cr-buildbucket.cfg",
				Revision: v8Revision,
			},
			Content: v8BuildbucketCfg,
		}
		configs := []luciconfig.Config{chromiumCfg, dartCfg, v8Cfg}
		configsToReturn := []luciconfig.Config{}
		for _, cfg := range configs {
			if cfg.Content != "" {
				configsToReturn = append(configsToReturn, cfg)
			}
		}
		return configsToReturn, nil
	}
	return nil, errors.New("not found")
}
