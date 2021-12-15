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

// +build !windows

package invoke

import (
	"os/exec"
	"syscall"

	"go.chromium.org/luci/common/errors"
)

func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func (s *Subprocess) terminate() error {
	if err := syscall.Kill(-s.cmd.Process.Pid, syscall.SIGTERM); err != nil {
		return errors.Annotate(err, "send SIGTERM").Err()
	}
	return nil
}
