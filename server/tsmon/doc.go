// Copyright 2017 The LUCI Authors.
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

// Package tsmon adapts common/tsmon library to a server-side environment.
//
// It provides a bunch of useful things:
//   - Hooks up tsmon library configuration to the server settings so it can
//     be changed on the fly without restarts.
//   - Provides a middleware that captures request metrics.
//   - Periodically reports Go runtime memory stats and some other metrics.
package tsmon
