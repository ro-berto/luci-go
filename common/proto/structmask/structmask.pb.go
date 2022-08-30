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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/common/proto/structmask/structmask.proto

package structmask

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StructMask selects a subset of a google.protobuf.Struct.
//
// Usually used as a repeated field, to allow specifying a union of different
// subsets.
type StructMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A field path inside the struct to select.
	//
	// Each item can be:
	//   * `some_value` - a concrete dict key to follow (unless it is a number or
	//     includes `*`, use quotes in this case).
	//   * `"some_value"` - same, but quoted. Useful for selecting `*` or numbers
	//     literally. See https://pkg.go.dev/strconv#Unquote for syntax.
	//   * `<number>` (e.g. `0`) - a zero-based list index to follow.
	//     **Not implemented**.
	//   *  `*` - follow all dict keys and all list elements. Applies **only** to
	//     dicts and lists. Trying to recurse into a number or a string results
	//     in an empty match.
	//
	// When examining a value the following exceptional conditions result in
	// an empty match, which is represented by `null` for list elements or
	// omissions of the field for dicts:
	//   * Trying to follow a dict key while examining a list.
	//   * Trying to follow a key which is not present in the dict.
	//   * Trying to use `*` mask with values that aren't dicts or lists.
	//
	// When using `*`, the result is always a subset of the input. In particular
	// this is important when filtering lists: if a list of size N is selected by
	// the mask, then the filtered result will also always be a list of size N,
	// with elements filtered further according to the rest of the mask (perhaps
	// resulting in `null` elements on type mismatches, as explained above).
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *StructMask) Reset() {
	*x = StructMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_structmask_structmask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructMask) ProtoMessage() {}

func (x *StructMask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_structmask_structmask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructMask.ProtoReflect.Descriptor instead.
func (*StructMask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescGZIP(), []int{0}
}

func (x *StructMask) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_go_chromium_org_luci_common_proto_structmask_structmask_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x20, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescData = file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_structmask_structmask_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_common_proto_structmask_structmask_proto_goTypes = []interface{}{
	(*StructMask)(nil), // 0: structmask.StructMask
}
var file_go_chromium_org_luci_common_proto_structmask_structmask_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_structmask_structmask_proto_init() }
func file_go_chromium_org_luci_common_proto_structmask_structmask_proto_init() {
	if File_go_chromium_org_luci_common_proto_structmask_structmask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_structmask_structmask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructMask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_structmask_structmask_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_structmask_structmask_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_common_proto_structmask_structmask_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_structmask_structmask_proto = out.File
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_structmask_structmask_proto_depIdxs = nil
}
