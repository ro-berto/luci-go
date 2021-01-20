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
	"os"
	"text/template"
)

// gitConfigTempl is a template for .gitconfig file generated by Kitchen.
//
// Note that we have to do two insteadOf rewrites: / => /a/ and /a/ => /a/.
// The later is needed to allow https://.../a/ to be passed to git explicitly.
// Without it, git will rewrite the url into https://.../a/a/ which doesn't
// work.
//
// Windows-specific settings are taken from <git-install>/mingw64/etc/gitconfig.
// We can't pick up this default config because it specifies Git-For-Windows
// built-in credential.helper that interferes with git-credential-luci. Settings
// useless on bots (like fancy terminal colors) are omitted.
var gitConfigTempl = template.Must(template.New(".gitconfig").Parse(`# Autogenerated.

[user]
  email = {{.UserEmail}}
  name = {{.UserName}}

[core]
  deltaBaseCacheLimit = 2g
{{- if .IsWindows}}
  fscache = true
  symlinks = false
  autocrlf = false
  filemode = false{{end}}

[pack]
  packSizeLimit = 2g

[http]
  version = HTTP/1.1
  # Request the GFE return debug headers as an encrypted blob in
  # X-Encrypted-Debug-Headers.
  extraheader = X-Return-Encrypted-Headers: all
{{if .IsWindows}}
  # This is path inside mingw64 installation, bundled with Git For Windows.
  sslCAinfo = /ssl/certs/ca-bundle.crt

[diff "astextplain"]
  textconv = astextplain
{{end}}
[gc]
  autodetach = false
{{if .UseCredentialHelper}}
[credential]
  helper = luci
{{end -}}
{{- range .KnownGerritHosts}}
[url "https://{{.}}/a/"]
  insteadOf = https://{{.}}/a/
  insteadOf = https://{{.}}/
{{end -}}
`))

// gitConfig is used to setup .gitconfig used by subprocesses.
//
// This is used only if git authentication is enabled. Otherwise we don't mess
// with git config at all.
//
// Assumes 'git' binary is actually gitwrapper and that 'git-credential-luci'
// binary is in PATH.
type gitConfig struct {
	IsWindows           bool     // true if running on Windows
	UserEmail           string   // value of user.email
	UserName            string   // value of user.name
	UseCredentialHelper bool     // if true, use git-credential-luci helper for auth
	KnownGerritHosts    []string // hosts to use '/a/' paths on to force auth
}

// Write actually writes the config to 'path'.
func (gc *gitConfig) Write(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = gitConfigTempl.Execute(f, gc); err != nil {
		return err
	}
	return f.Close() // failure to close the file is an overall failure
}
