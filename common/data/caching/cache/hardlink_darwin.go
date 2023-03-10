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

package cache

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/sys/unix"

	"go.chromium.org/luci/common/errors"
)

var cachedMajorVersion int
var cachedMajorVersionOnce sync.Once

func mustGetDarwinMajorVersion() int {
	cachedMajorVersionOnce.Do(func() {
		var utsname unix.Utsname
		if err := unix.Uname(&utsname); err != nil {
			panic(fmt.Sprintf("failed to call uname(): %v", err))
		}

		release := string(utsname.Release[:])
		vers := strings.Split(release, ".")
		if len(vers) < 2 {
			panic(fmt.Sprintf("unexpected release from uname: %s", string(release)))
		}

		var err error
		cachedMajorVersion, err = strconv.Atoi(vers[0])
		if err != nil {
			panic(fmt.Sprintf("failed to parse version %s: %v", vers[0], err))
		}
	})
	return cachedMajorVersion
}

func clonefile(src, dst string) error {
	srcByte, err := unix.BytePtrFromString(src)
	if err != nil {
		log.Fatalf("unix.BytePtrFromString(%s): %v", src, err)
	}

	dstByte, err := unix.BytePtrFromString(dst)
	if err != nil {
		log.Fatalf("unix.BytePtrFromString(%s): %v", dst, err)
	}

	atFDCwd := unix.AT_FDCWD

	if _, _, err := unix.Syscall6(unix.SYS_CLONEFILEAT, uintptr(atFDCwd), uintptr(unsafe.Pointer(srcByte)), uintptr(atFDCwd), uintptr(unsafe.Pointer(dstByte)), unix.CLONE_NOFOLLOW|unix.CLONE_NOOWNERCOPY, 0); err != 0 {
		return errors.Annotate(err, "failed to call clonefile").Err()
	}
	return nil
}

func makeHardLinkOrClone(src, dst string) error {
	// Hardlinked executables don't work well with dyld.
	// Use clonefile instead on macOS 12 (Darwin 21) or newer to workaround that.
	// ref: https://crbug.com/1296318#c54
	if mustGetDarwinMajorVersion() >= 21 {
		// TODO(crbug.com/1315077): we can use unix.Clonefile when we stop supporting macOS 10.11.
		if err := clonefile(src, dst); err != nil {
			return errors.Annotate(err, "failed to call clonefile(%s, %s, CLONE_NOFOLLOW|CLONE_NOOWNERCOPY)", src, dst).Err()
		}
		return nil
	}

	if err := os.Link(src, dst); err != nil {
		return errors.Annotate(err, "failed to call os.Link(%s, %s)", src, dst).Err()
	}
	return nil
}
