// Copyright 2016 The LUCI Authors.
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

package tsmon

import (
	"context"
	"os"
	"testing"

	"go.chromium.org/luci/common/tsmon/target"

	. "github.com/smartystreets/goconvey/convey"
)

// newTestingFlags returns new Flags that are appropriate for testing. These
// stub out relevant system-local information so that system specifics don't
// affect things.
func newTestingFlags() Flags {
	fl := NewFlags()
	fl.ConfigFile = "" // Stub out so we don't load local system config.
	fl.Endpoint = "file://fake/path"
	fl.Target.SysInfo = &target.SysInfo{
		Hostname: "host-sys",
		Region:   "region-sys",
	}
	return fl
}

func TestInitializeFromFlags(t *testing.T) {
	t.Parallel()

	Convey("Initialize tsmon", t, func() {
		c := WithState(context.Background(), NewState())

		tsmonFlags := newTestingFlags()
		tsmonFlags.Target.TargetType = target.DeviceType
		tsmonFlags.Flush = FlushManual

		Convey("with autogenerated hostnames", func() {
			tsmonFlags.Target.AutoGenHostname = true
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.NetworkDevice)(nil))
			target := s.DefaultTarget().(*target.NetworkDevice)
			So(target.Hostname, ShouldEqual, "autogen:host-sys")
			So(target.Metro, ShouldEqual, "region-sys")
		})

		Convey("with predefined autogenerated hostnames (Device)", func() {
			tsmonFlags.Target.AutoGenHostname = true
			tsmonFlags.Target.DeviceHostname = "test-m5"
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.NetworkDevice)(nil))
			target := s.DefaultTarget().(*target.NetworkDevice)
			So(target.Hostname, ShouldEqual, "autogen:test-m5")
			So(target.Metro, ShouldEqual, "region-sys")
		})

		Convey("with predefined autogenerated hostnames (Task)", func() {
			tsmonFlags.Target.AutoGenHostname = true
			tsmonFlags.Target.TaskHostname = "test-m5"
			tsmonFlags.Target.TargetType = target.TaskType
			tsmonFlags.Target.TaskServiceName = "test-service"
			tsmonFlags.Target.TaskJobName = "test-job"
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.Task)(nil))
			target := s.DefaultTarget().(*target.Task)
			So(target.HostName, ShouldEqual, "autogen:test-m5")
			So(target.DataCenter, ShouldEqual, "region-sys")
		})

		Convey("with static hostnames", func() {
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.NetworkDevice)(nil))
			target := s.DefaultTarget().(*target.NetworkDevice)
			So(target.Hostname, ShouldEqual, "host-sys")
			So(target.Metro, ShouldEqual, "region-sys")
		})

		Convey("with predefined static hostnames (Device)", func() {
			tsmonFlags.Target.DeviceHostname = "host-flag"
			tsmonFlags.Target.DeviceRegion = "region-flag"
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.NetworkDevice)(nil))
			target := s.DefaultTarget().(*target.NetworkDevice)
			So(target.Hostname, ShouldEqual, "host-flag")
			So(target.Metro, ShouldEqual, "region-flag")
		})

		Convey("with predefined static hostnames (Task)", func() {
			tsmonFlags.Target.TaskHostname = "host-flag"
			tsmonFlags.Target.TaskRegion = "region-flag"
			tsmonFlags.Target.TargetType = target.TaskType
			tsmonFlags.Target.TaskServiceName = "test-service"
			tsmonFlags.Target.TaskJobName = "test-job"
			err := InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.Task)(nil))
			target := s.DefaultTarget().(*target.Task)
			So(target.HostName, ShouldEqual, "host-flag")
			So(target.DataCenter, ShouldEqual, "region-flag")
		})

		Convey("with region in config and hostname in a flag", func() {
			tf, err := os.CreateTemp("", "config_test")
			if err != nil {
				t.Fail()
			}
			defer tf.Close()
			defer os.Remove(tf.Name())

			tf.WriteString(`
			{"endpoint":         "foo",
			 "credentials":      "bar",
			 "autogen_hostname": true,
			 "hostname":         "host-config",
			 "region":           "region-config"
			}`)
			tf.Sync()
			tsmonFlags.ConfigFile = tf.Name()
			tsmonFlags.Target.DeviceHostname = "host-flag"
			err = InitializeFromFlags(c, &tsmonFlags)
			So(err, ShouldBeNil)
			defer Shutdown(c)

			s := GetState(c).Store()
			So(s.DefaultTarget(), ShouldHaveSameTypeAs, (*target.NetworkDevice)(nil))
			target := s.DefaultTarget().(*target.NetworkDevice)
			So(target.Hostname, ShouldEqual, "autogen:host-flag")
			So(target.Metro, ShouldEqual, "region-config")
		})

	})
}
