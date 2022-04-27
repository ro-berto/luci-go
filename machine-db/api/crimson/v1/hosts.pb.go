// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/machine-db/api/crimson/v1/hosts.proto

package crimson

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

// A request to delete a physical or virtual host from the database.
type DeleteHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the host to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The VLAN the host belongs to.
	Vlan int64 `protobuf:"varint,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
}

func (x *DeleteHostRequest) Reset() {
	*x = DeleteHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHostRequest) ProtoMessage() {}

func (x *DeleteHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHostRequest.ProtoReflect.Descriptor instead.
func (*DeleteHostRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteHostRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteHostRequest) GetVlan() int64 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72,
	0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x6c,
	0x61, 0x6e, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_goTypes = []interface{}{
	(*DeleteHostRequest)(nil), // 0: crimson.DeleteHostRequest
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHostRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_hosts_proto_depIdxs = nil
}
