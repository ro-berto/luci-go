// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/machine-db/api/crimson/v1/vlans.proto

package crimson

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

// A VLAN in the database.
type VLAN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of this VLAN. Uniquely identifies this VLAN.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An alias for this VLAN.
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	// A description of this VLAN.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The state of this VLAN.
	State v1.State `protobuf:"varint,4,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The block of IPv4 addresses belonging to this VLAN.
	CidrBlock string `protobuf:"bytes,5,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
}

func (x *VLAN) Reset() {
	*x = VLAN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VLAN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VLAN) ProtoMessage() {}

func (x *VLAN) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescGZIP(), []int{0}
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

func (x *VLAN) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VLAN) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (x *VLAN) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

// A request to list VLANs in the database.
type ListVLANsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IDs of VLANs to retrieve.
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// The aliases of VLANs to retrieve.
	Aliases []string `protobuf:"bytes,2,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *ListVLANsRequest) Reset() {
	*x = ListVLANsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVLANsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVLANsRequest) ProtoMessage() {}

func (x *ListVLANsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVLANsRequest.ProtoReflect.Descriptor instead.
func (*ListVLANsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescGZIP(), []int{1}
}

func (x *ListVLANsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ListVLANsRequest) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

// A response containing a list of VLANs in the database.
type ListVLANsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The VLANs matching the request.
	Vlans []*VLAN `protobuf:"bytes,1,rep,name=vlans,proto3" json:"vlans,omitempty"`
}

func (x *ListVLANsResponse) Reset() {
	*x = ListVLANsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVLANsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVLANsResponse) ProtoMessage() {}

func (x *ListVLANsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVLANsResponse.ProtoReflect.Descriptor instead.
func (*ListVLANsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescGZIP(), []int{2}
}

func (x *ListVLANsResponse) GetVlans() []*VLAN {
	if x != nil {
		return x.Vlans
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x6c, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72,
	0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x56, 0x4c, 0x41, 0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64,
	0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x3e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4c,
	0x41, 0x4e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x4c,
	0x41, 0x4e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x76,
	0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x69,
	0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x56, 0x4c, 0x41, 0x4e, 0x52, 0x05, 0x76, 0x6c, 0x61, 0x6e, 0x73,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_goTypes = []interface{}{
	(*VLAN)(nil),              // 0: crimson.VLAN
	(*ListVLANsRequest)(nil),  // 1: crimson.ListVLANsRequest
	(*ListVLANsResponse)(nil), // 2: crimson.ListVLANsResponse
	(v1.State)(0),             // 3: common.State
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_depIdxs = []int32{
	3, // 0: crimson.VLAN.state:type_name -> common.State
	0, // 1: crimson.ListVLANsResponse.vlans:type_name -> crimson.VLAN
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVLANsRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVLANsResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_vlans_proto_depIdxs = nil
}
