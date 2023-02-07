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

//go:build luagentest

package main

import (
	"go.chromium.org/luci/common/proto/msgpackpb/luagen"
	"go.chromium.org/luci/server/quota/quotapb"
)

func main() {
	luagen.Main(
		&quotapb.Account{},
		&quotapb.ApplyOpsResponse{},
		&quotapb.OpError{},
		&quotapb.Policy{},
		&quotapb.UpdateAccountsInput{},
		quotapb.Op_NO_OPTIONS,
	)
}
