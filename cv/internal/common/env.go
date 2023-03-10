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

package common

import "go.chromium.org/luci/server"

// Env describes where CV runs at.
type Env struct {
	// LogicalHostname is CV hostname referred to in configs.
	//
	// On GAE, this is something like "luci-change-verifier-dev.appspot.com"
	// and it is part of the HTTPAddressBase.
	//
	// Under local development, this is usually set to GAE-looking hostname, while
	// keeping HTTPAddressBase a real localhost URL.
	LogicalHostname string

	// HTTPAddressBase can be used to generate URLs to this CV service.
	//
	// Doesn't have a trailing slash.
	//
	// For example,
	//   * "https://luci-change-verifier-dev.appspot.com"
	//   * "http://localhost:8080"
	HTTPAddressBase string

	// IsGAEDev is true if this is a -dev GAE environment.
	//
	// Deprecated. Do not use in new code. It should only be used during migration
	// from CQDaemon which doesn't have equivalent -dev environment.
	IsGAEDev bool

	// GAEInfo is populated if LUCI CV runs on GAE.
	GAEInfo struct {
		// CloudProject is the name of the Google Cloud Project LUCI CV runs in.
		CloudProject string
		// ServiceName is the name of the micro-service in the GAE app.
		ServiceName string
		// InstanceID is the ID of the instance that runs LUCI CV.
		InstanceID string
	}
}

// MakeEnv creates a new `Env` from server options.
func MakeEnv(opts server.Options) *Env {
	env := &Env{
		LogicalHostname: opts.CloudProject + ".appspot.com",
		IsGAEDev:        opts.CloudProject == "luci-change-verifier-dev",
		GAEInfo: struct {
			CloudProject string
			ServiceName  string
			InstanceID   string
		}{
			CloudProject: opts.CloudProject,
			// TODO(yiwzhang): have a more reliable way to get the GAE service name.
			ServiceName: opts.TsMonJobName,
			InstanceID:  opts.Hostname,
		},
	}
	env.HTTPAddressBase = "https://" + env.LogicalHostname
	if !opts.Prod {
		// Local development.
		env.HTTPAddressBase = "http://" + opts.HTTPAddr
	}
	return env
}
