// Copyright 2019 The LUCI Authors.
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

package authctx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGitConfig(t *testing.T) {
	t.Parallel()

	Convey("Works", t, func() {
		tmp, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)
		defer os.RemoveAll(tmp)

		gen := func(name string, cfg gitConfig) string {
			So(cfg.Write(filepath.Join(tmp, name)), ShouldBeNil)
			body, err := ioutil.ReadFile(filepath.Join(tmp, name))
			So(err, ShouldBeNil)
			return string(body)
		}

		cfg := gitConfig{
			IsWindows:           false,
			UserEmail:           "email@example.com",
			UserName:            "name",
			UseCredentialHelper: true,
			KnownGerritHosts:    []string{"host-a", "host-b"},
		}

		So(gen("unix", cfg), ShouldEqual, `# Autogenerated.

[user]
  email = email@example.com
  name = name

[core]
  deltaBaseCacheLimit = 2g

[pack]
  packSizeLimit = 2g

[http]
  version = HTTP/1.1
  # Request the GFE return debug headers as an encrypted blob in
  # X-Encrypted-Debug-Headers.
  extraheader = X-Return-Encrypted-Headers: all

[gc]
  autodetach = false

[credential]
  helper = luci

[url "https://host-a/a/"]
  insteadOf = https://host-a/a/
  insteadOf = https://host-a/

[url "https://host-b/a/"]
  insteadOf = https://host-b/a/
  insteadOf = https://host-b/
`)

		cfg.IsWindows = true
		So(gen("win", cfg), ShouldEqual, `# Autogenerated.

[user]
  email = email@example.com
  name = name

[core]
  deltaBaseCacheLimit = 2g
  fscache = true
  symlinks = false
  autocrlf = false
  filemode = false

[pack]
  packSizeLimit = 2g

[http]
  version = HTTP/1.1
  # Request the GFE return debug headers as an encrypted blob in
  # X-Encrypted-Debug-Headers.
  extraheader = X-Return-Encrypted-Headers: all

  # This is path inside mingw64 installation, bundled with Git For Windows.
  sslCAinfo = /ssl/certs/ca-bundle.crt

[diff "astextplain"]
  textconv = astextplain

[gc]
  autodetach = false

[credential]
  helper = luci

[url "https://host-a/a/"]
  insteadOf = https://host-a/a/
  insteadOf = https://host-a/

[url "https://host-b/a/"]
  insteadOf = https://host-b/a/
  insteadOf = https://host-b/
`)
	})
}
