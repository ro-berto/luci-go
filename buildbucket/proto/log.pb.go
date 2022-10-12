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
// source: go.chromium.org/luci/buildbucket/proto/log.proto

package buildbucketpb

import (
	_ "go.chromium.org/luci/common/bq/pb"
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

// A pRPC request log.
type PRPCRequestLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An ID associated with this request log.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Any parent ID associated with this request log.
	// Typically a parent will be the batch request containing this request.
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// When the request being logged was received (microseconds since epoch).
	CreationTime int64 `protobuf:"varint,3,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	// Duration of the request in microseconds.
	Duration int64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// Method called. e.g. "Builds.GetBuild".
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	// Caller auth identity (e.g. "<kind>:<id>"). Only set when the caller is a
	// robot or anonymous.
	User string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PRPCRequestLog) Reset() {
	*x = PRPCRequestLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PRPCRequestLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PRPCRequestLog) ProtoMessage() {}

func (x *PRPCRequestLog) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PRPCRequestLog.ProtoReflect.Descriptor instead.
func (*PRPCRequestLog) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescGZIP(), []int{0}
}

func (x *PRPCRequestLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PRPCRequestLog) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PRPCRequestLog) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *PRPCRequestLog) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *PRPCRequestLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PRPCRequestLog) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_go_chromium_org_luci_buildbucket_proto_log_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x71, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x50, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x34,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xe2, 0xbc, 0x24, 0x0b, 0x0a, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_buildbucket_proto_log_proto_goTypes = []interface{}{
	(*PRPCRequestLog)(nil), // 0: buildbucket.v2.PRPCRequestLog
}
var file_go_chromium_org_luci_buildbucket_proto_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_log_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_log_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PRPCRequestLog); i {
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
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_log_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_log_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_log_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_log_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_log_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_log_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_log_proto_depIdxs = nil
}
