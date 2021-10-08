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

package longops

import (
	"errors"
)

// errCancelHonored is used inside this package to differentiate graceful honoring
// of request to cancel the long operation from other error conditions.
//
// It must never be returned outside this package.
var errCancelHonored = errors.New("cancel request was honored")

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
