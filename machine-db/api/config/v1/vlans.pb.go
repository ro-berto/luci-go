// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/machine-db/api/config/v1/vlans.proto

package config

import (
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
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

// A VLAN to store in the database.
type VLAN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of this VLAN. Must be unique.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An alias for this VLAN.
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	// The block of IPv4 addresses belonging to this VLAN.
	CidrBlock string `protobuf:"bytes,3,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	// The state of this VLAN.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
}

func (x *VLAN) Reset() {
	*x = VLAN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VLAN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VLAN) ProtoMessage() {}

func (x *VLAN) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VLAN.ProtoReflect.Descriptor instead.
func (*VLAN) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescGZIP(), []int{0}
}

func (x *VLAN) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VLAN) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *VLAN) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *VLAN) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A list of VLANs.
type VLANs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of VLANs.
	Vlan []*VLAN `protobuf:"bytes,1,rep,name=vlan,proto3" json:"vlan,omitempty"`
}

func (x *VLANs) Reset() {
	*x = VLANs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VLANs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VLANs) ProtoMessage() {}

func (x *VLANs) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VLANs.ProtoReflect.Descriptor instead.
func (*VLANs) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescGZIP(), []int{1}
}

func (x *VLANs) GetVlan() []*VLAN {
	if x != nil {
		return x.Vlan
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x6c, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x70, 0x0a, 0x04, 0x56, 0x4c, 0x41, 0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x29, 0x0a, 0x05, 0x56, 0x4c, 0x41, 0x4e, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x76, 0x6c,
	0x61, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x56, 0x4c, 0x41, 0x4e, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_goTypes = []interface{}{
	(*VLAN)(nil),  // 0: config.VLAN
	(*VLANs)(nil), // 1: config.VLANs
	(v1.State)(0), // 2: common.State
}
var file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_depIdxs = []int32{
	2, // 0: config.VLAN.state:type_name -> common.State
	0, // 1: config.VLANs.vlan:type_name -> config.VLAN
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_init() }
func file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VLAN); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VLANs); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto = out.File
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_vlans_proto_depIdxs = nil
}
