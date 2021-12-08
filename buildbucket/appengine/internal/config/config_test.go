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

package config

import (
	"context"
	"testing"

	"go.chromium.org/luci/config/validation"
	"go.chromium.org/luci/gae/impl/memory"

	// TODO(crbug/1242998): Remove once safe get becomes datastore default.
	_ "go.chromium.org/luci/gae/service/datastore/crbug1242998safeget"

	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	Convey("get settings.cfg", t, func() {
		settingsCfg := &pb.SettingsCfg{Resultdb: &pb.ResultDBSettings{Hostname: "testing.results.api.cr.dev"}}
		ctx := memory.Use(context.Background())
		SetTestSettingsCfg(ctx, settingsCfg)
		cfg, err := GetSettingsCfg(ctx)
		So(err, ShouldBeNil)
		So(cfg, ShouldResembleProto, settingsCfg)
	})

	Convey("validate settings.cfg", t, func() {
		vctx := &validation.Context{
			Context: context.Background(),
		}
		configSet := "services/${appid}"
		path := "settings.cfg"

		Convey("bad proto", func() {
			content := []byte(` bad: "bad" `)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "invalid SettingsCfg proto message")
		})

		Convey("empty settings.cfg", func() {
			content := []byte(` `)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "logdog.hostname unspecified")
		})

		Convey("no swarming cfg", func() {
			content := []byte(`
				logdog {
					hostname: "logs.chromium.org"
				}
				resultdb {
					hostname: "results.api.cr.dev"
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize(), ShouldBeNil)
		})

		Convey("no milo", func() {
			content := []byte(`swarming{}`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "milo_hostname unspecified")
		})

		Convey("invalid user_packages", func() {
			content := []byte(`
				swarming {
					milo_hostname: "ci.chromium.org"
					user_packages {
						package_name: ""
						version: "git_revision:c84736ceb5ddcc3f6e6d1e6c4d602bb024ceb1b2"
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "(swarming / user_packages #0): package_name is required")
		})

		Convey("no /${platform} in bbagent", func() {
			content := []byte(`
				swarming {
					milo_hostname: "ci.chromium.org"
					bbagent_package {
						package_name: "infra/tools/luci/bbagent"
						version: "git_revision:60be805bf35a766cdf7d80bdf0a066dce30691e8"
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "(swarming / bbagent_package): package_name must end with '/${platform}'")
		})

		Convey("invalid builders.regex in experiments", func() {
			content := []byte(`
				experiment {
					experiments {
						name: "luci.buildbucket.bbagent_getbuild"
						builders {
							regex: "(no right parenthesis"
						}
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "(experiment.experiments #0): builders.regex \"(no right parenthesis\": invalid regex")
		})

		Convey("wrong default_value in experiments", func() {
			content := []byte(`
				experiment {
					experiments {
						name: "luci.buildbucket.bbagent_getbuild"
						default_value: 101
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "default_value must be in [0,100]")
		})

		Convey("wrong minimum_value in experiments", func() {
			content := []byte(`
				experiment {
					experiments {
						name: "luci.buildbucket.bbagent_getbuild"
						default_value: 10
						minimum_value: 5
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "minimum_value must be in [${default_value},100]")
		})

		Convey("invalid inactive in experiments", func() {
			content := []byte(`
				experiment {
					experiments {
						name: "luci.buildbucket.bbagent_getbuild"
						minimum_value: 5
						inactive: true
					}
				}
			`)
			So(validateSettingsCfg(vctx, configSet, path, content), ShouldBeNil)
			So(vctx.Finalize().Error(), ShouldContainSubstring, "default_value and minimum_value must both be 0 when inactive is true")
		})

		Convey("OK", func() {
			var okCfg = `
				swarming {
					milo_hostname: "ci.chromium.org"
					global_caches {
						path: "git"
					}
					bbagent_package {
						package_name: "infra/tools/luci/bbagent/${platform}"
						version: "git_revision:60be805bf35a766cdf7d80bdf0a066dce30691e8"
					}
					kitchen_package {
						package_name: "infra/tools/luci/kitchen/${platform}"
						version: "git_revision:63874080a20260642c8df82d4f4885ff30b33fb6"
					}
					user_packages {
						package_name: "infra/3pp/tools/git/${platform}"
						version: "version:2@2.33.0.chromium.6"
					}
				}
				logdog {
					hostname: "logs.chromium.org"
				}
				resultdb {
					hostname: "results.api.cr.dev"
				}
				known_public_gerrit_hosts: "swiftshader.googlesource.com"
				known_public_gerrit_hosts: "webrtc.googlesource.com"
				experiment {
					experiments {
						name: "luci.buildbucket.bbagent_getbuild"
					}
					# inactive experiments
					experiments { name: "luci.use_realms" inactive: true }
				}
			`
			So(validateSettingsCfg(vctx, configSet, path, []byte(okCfg)), ShouldBeNil)
			So(vctx.Finalize(), ShouldBeNil)
		})
	})
}
