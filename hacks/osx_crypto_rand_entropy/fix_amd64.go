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

//go:build darwin && amd64 && go1.17
// +build darwin,amd64,go1.17

package osx_crypto_rand_entropy

import (
	"fmt"
	"os"

	// To use go:linkname.
	_ "unsafe"
)

//go:linkname altGetRandom crypto/rand.altGetRandom
var altGetRandom func(p []byte) error

func init() {
	// Unset "optimized" implementation, making crypto/rand fallback to the
	// general /dev/urandom implementation, as it did prior go1.17.
	//
	// Note: we assume nothing is calling crypto/rand during init() time. If it
	// does, it should import `osx_crypto_rand_entropy` package explicitly first
	// to make sure it installs the hack. This is all very fragile.
	altGetRandom = nil

	if os.Getenv("LUCI_GO_CHECK_HACKS") == "1" {
		fmt.Fprintf(os.Stderr, "LUCI_GO_CHECK_HACKS: osx_crypto_rand_entropy is enabled\n")
	}
}
