// Copyright 2015 The LUCI Authors.
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

//go:build !appengine
// +build !appengine

package prod

import (
	"context"

	"google.golang.org/appengine"
)

// UseBackground is the same as Use except that it activates production
// implementations which aren't associated with any particular request.
//
// This is only available on Managed VMs.
//
// It is important to note that this DOES NOT install the AppEngine SDK into the
// supplied Context. See the warning in Use for more information.
func UseBackground(c context.Context) context.Context {
	return setupAECtx(c, appengine.BackgroundContext())
}
