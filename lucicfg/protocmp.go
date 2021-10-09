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

package lucicfg

import (
	"bytes"
	"fmt"
	"math"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	luciproto "go.chromium.org/luci/common/proto"
)

// semanticallyEqual is like proto.Equal, but it ignores fields annotated with
// lucicfg_ignore=true option and unknown fields.
//
// The implementation is heavily based on proto.Equal implementation.
//
// TODO(vadimsh): This can be optimized slightly. Given a MessageDescriptor, we
// can traverse its fields *once* to see if there are any E_LucicfgIgnore
// extensions at all. If there are none, we can just use proto.Equal and thus
// skip all 'shouldVisit' calls at once.
func semanticallyEqual(x, y proto.Message) bool {
	return equalMessage(x.ProtoReflect(), y.ProtoReflect())
}

func shouldVisit(fd protoreflect.FieldDescriptor) bool {
	ignore := false
	if opts, ok := fd.Options().(*descriptorpb.FieldOptions); ok {
		ignore = proto.GetExtension(opts, luciproto.E_LucicfgIgnore).(bool)
	}
	return !ignore
}

func equalMessage(x, y protoreflect.Message) bool {
	if x.Descriptor() != y.Descriptor() {
		return false
	}

	nx := 0
	equal := true
	x.Range(func(fd protoreflect.FieldDescriptor, vx protoreflect.Value) bool {
		if shouldVisit(fd) {
			nx++
			equal = y.Has(fd) && equalField(fd, vx, y.Get(fd))
		}
		return equal
	})
	if !equal {
		return false
	}

	ny := 0
	y.Range(func(fd protoreflect.FieldDescriptor, vy protoreflect.Value) bool {
		if shouldVisit(fd) {
			ny++
		}
		return true
	})

	return nx == ny
}

func equalField(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	switch {
	case fd.IsList():
		return equalList(fd, x.List(), y.List())
	case fd.IsMap():
		return equalMap(fd, x.Map(), y.Map())
	default:
		return equalValue(fd, x, y)
	}
}

func equalList(fd protoreflect.FieldDescriptor, x, y protoreflect.List) bool {
	if x.Len() != y.Len() {
		return false
	}
	for i := x.Len() - 1; i >= 0; i-- {
		if !equalValue(fd, x.Get(i), y.Get(i)) {
			return false
		}
	}
	return true
}

func equalMap(fd protoreflect.FieldDescriptor, x, y protoreflect.Map) bool {
	if x.Len() != y.Len() {
		return false
	}
	equal := true
	x.Range(func(k protoreflect.MapKey, vx protoreflect.Value) bool {
		equal = y.Has(k) && equalValue(fd.MapValue(), vx, y.Get(k))
		return equal
	})
	return equal
}

func equalValue(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
	// Panic if we see something unexpected as a precaution against protobuf lib
	// adding a new type. If this ever happens, the code below will likely need
	// to be updated.
	k := fd.Kind()
	if k > protoreflect.Sint64Kind {
		panic(fmt.Sprintf("unrecognized proto value kind %s in field %s", k, fd))
	}
	switch k {
	case protoreflect.BoolKind:
		return x.Bool() == y.Bool()
	case protoreflect.EnumKind:
		return x.Enum() == y.Enum()
	case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return x.Int() == y.Int()
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return x.Uint() == y.Uint()
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		fx := x.Float()
		fy := y.Float()
		if math.IsNaN(fx) || math.IsNaN(fy) {
			return math.IsNaN(fx) && math.IsNaN(fy)
		}
		return fx == fy
	case protoreflect.StringKind:
		return x.String() == y.String()
	case protoreflect.BytesKind:
		return bytes.Equal(x.Bytes(), y.Bytes())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return equalMessage(x.Message(), y.Message())
	default:
		return x.Interface() == y.Interface()
	}
}
