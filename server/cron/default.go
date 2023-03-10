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

package cron

// Default is a dispatcher installed into the server when using NewModule or
// NewModuleFromFlags.
//
// The module takes care of configuring this dispatcher based on the server
// environment and module's options.
//
// You still need to register your handlers in it using RegisterHandler and
// configure Cloud Scheduler jobs (or cron.yaml when running on Appengine) to
// actually invoke them.
var Default Dispatcher

// RegisterHandler is a shortcut for Default.RegisterHandler.
func RegisterHandler(id string, h Handler) {
	Default.RegisterHandler(id, h)
}
