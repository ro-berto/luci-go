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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/resultdb/proto/bq/common.proto

package resultpb

import (
	_ "go.chromium.org/luci/common/bq/pb"
	v1 "go.chromium.org/luci/resultdb/proto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InvocationRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the invocation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Tags represents Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*v1.StringPair `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	// The LUCI Realm the invocation exists under.
	Realm string `protobuf:"bytes,3,opt,name=realm,proto3" json:"realm,omitempty"`
	// Arbitrary JSON object that contains structured, domain-specific properties
	// of the invocation.
	Properties *structpb.Struct `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *InvocationRecord) Reset() {
	*x = InvocationRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_proto_bq_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvocationRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvocationRecord) ProtoMessage() {}

func (x *InvocationRecord) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_proto_bq_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvocationRecord.ProtoReflect.Descriptor instead.
func (*InvocationRecord) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescGZIP(), []int{0}
}

func (x *InvocationRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvocationRecord) GetTags() []*v1.StringPair {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *InvocationRecord) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *InvocationRecord) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_proto_bq_common_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2e, 0x62, 0x71, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x71, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x0a, 0xe2, 0xbc, 0x24, 0x06, 0x0a, 0x04, 0x4a, 0x53, 0x4f,
	0x4e, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x71, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescData = file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_proto_bq_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_resultdb_proto_bq_common_proto_goTypes = []interface{}{
	(*InvocationRecord)(nil), // 0: luci.resultdb.bq.InvocationRecord
	(*v1.StringPair)(nil),    // 1: luci.resultdb.v1.StringPair
	(*structpb.Struct)(nil),  // 2: google.protobuf.Struct
}
var file_go_chromium_org_luci_resultdb_proto_bq_common_proto_depIdxs = []int32{
	1, // 0: luci.resultdb.bq.InvocationRecord.tags:type_name -> luci.resultdb.v1.StringPair
	2, // 1: luci.resultdb.bq.InvocationRecord.properties:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_proto_bq_common_proto_init() }
func file_go_chromium_org_luci_resultdb_proto_bq_common_proto_init() {
	if File_go_chromium_org_luci_resultdb_proto_bq_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_proto_bq_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvocationRecord); i {
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
			RawDescriptor: file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_proto_bq_common_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_proto_bq_common_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_resultdb_proto_bq_common_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_proto_bq_common_proto = out.File
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_proto_bq_common_proto_depIdxs = nil
}
