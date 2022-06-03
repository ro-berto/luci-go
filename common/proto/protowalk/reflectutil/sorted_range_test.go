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

package reflectutil

import (
	"sort"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestMapRangeSorted(t *testing.T) {
	t.Parallel()

	Convey(`MapRangeSorted`, t, func() {
		Convey(`empty`, func() {
			msg := &TestMapMessage{}
			refl := msg.ProtoReflect()
			getField := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("string_map")).Map() }
			MapRangeSorted(getField(), protoreflect.StringKind, func(protoreflect.MapKey, protoreflect.Value) bool {
				panic("callback called for empty map?")
			})
		})

		Convey(`bool`, func() {
			type item struct {
				k bool
				v string
			}
			msg := &TestMapMessage{}
			refl := msg.ProtoReflect()
			getField := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("bool_map")).Map() }

			Convey(`true`, func() {
				msg.BoolMap = map[bool]string{true: "hey"}
				results := []item{}
				MapRangeSorted(getField(), protoreflect.BoolKind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Bool(), v.String()})
					return true
				})
				So(results, ShouldResemble, []item{{true, "hey"}})
			})
			Convey(`false`, func() {
				msg.BoolMap = map[bool]string{false: "hey"}
				results := []item{}
				MapRangeSorted(getField(), protoreflect.BoolKind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Bool(), v.String()})
					return true
				})
				So(results, ShouldResemble, []item{{false, "hey"}})
			})
			Convey(`both`, func() {
				msg.BoolMap = map[bool]string{false: "hey", true: "norp"}
				results := []item{}
				MapRangeSorted(getField(), protoreflect.BoolKind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Bool(), v.String()})
					return true
				})
				So(results, ShouldResemble, []item{{false, "hey"}, {true, "norp"}})
			})
		})

		Convey(`int`, func() {
			type item struct {
				k int64
				v string
			}
			msg := &TestMapMessage{}
			refl := msg.ProtoReflect()
			get32Field := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("int32_map")).Map() }
			get64Field := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("int64_map")).Map() }

			Convey(`int32`, func() {
				msg.Int32Map = map[int32]string{}
				expect := []item{}
				for i := int64(0); i < 100; i++ {
					entry := item{i, strconv.FormatInt(i, 10)}
					msg.Int32Map[int32(i)] = entry.v
					expect = append(expect, entry)
				}
				results := []item{}
				MapRangeSorted(get32Field(), protoreflect.Int32Kind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Int(), v.String()})
					return true
				})
				So(results, ShouldResemble, expect)
			})

			Convey(`int64`, func() {
				msg.Int64Map = map[int64]string{}
				expect := []item{}
				for i := int64(0); i < 100; i++ {
					entry := item{i, strconv.FormatInt(i, 10)}
					msg.Int64Map[i] = entry.v
					expect = append(expect, entry)
				}
				results := []item{}
				MapRangeSorted(get64Field(), protoreflect.Int64Kind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Int(), v.String()})
					return true
				})
				So(results, ShouldResemble, expect)
			})

		})

		Convey(`uint`, func() {
			type item struct {
				k uint64
				v string
			}
			msg := &TestMapMessage{}
			refl := msg.ProtoReflect()
			get32Field := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("uint32_map")).Map() }
			get64Field := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("uint64_map")).Map() }

			Convey(`uint32`, func() {
				msg.Uint32Map = map[uint32]string{}
				expect := []item{}
				for i := uint64(0); i < 100; i++ {
					entry := item{i, strconv.FormatUint(i, 10)}
					msg.Uint32Map[uint32(i)] = entry.v
					expect = append(expect, entry)
				}
				results := []item{}
				MapRangeSorted(get32Field(), protoreflect.Uint32Kind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Uint(), v.String()})
					return true
				})
				So(results, ShouldResemble, expect)
			})

			Convey(`uint64`, func() {
				msg.Uint64Map = map[uint64]string{}
				expect := []item{}
				for i := uint64(0); i < 100; i++ {
					entry := item{i, strconv.FormatUint(i, 10)}
					msg.Uint64Map[i] = entry.v
					expect = append(expect, entry)
				}
				results := []item{}
				MapRangeSorted(get64Field(), protoreflect.Uint64Kind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
					results = append(results, item{k.Uint(), v.String()})
					return true
				})
				So(results, ShouldResemble, expect)
			})
		})

		Convey(`string`, func() {
			msg := &TestMapMessage{}
			refl := msg.ProtoReflect()
			getField := func() protoreflect.Map { return refl.Get(refl.Descriptor().Fields().ByName("string_map")).Map() }

			msg.StringMap = map[string]string{}
			expect := []string{}
			for i := uint64(0); i < 100; i++ {
				entry := strconv.FormatUint(i, 10)
				msg.StringMap[entry] = entry
				expect = append(expect, entry)
			}
			// expect is sorted alphanumerically, not numerically
			sort.Strings(expect)

			results := []string{}
			MapRangeSorted(getField(), protoreflect.StringKind, func(k protoreflect.MapKey, v protoreflect.Value) bool {
				if k.String() != v.String() {
					panic("mismatched key/value")
				}
				results = append(results, v.String())
				return true
			})
			So(results, ShouldResemble, expect)
		})
	})
}
