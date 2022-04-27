// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/machine-db/api/crimson/v1/ips.proto

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

// An IP address in the database.
type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IPv4 address. Uniquely identifies this IP address.
	Ipv4 string `protobuf:"bytes,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The VLAN this IP address belongs to.
	Vlan int64 `protobuf:"varint,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// The hostname this IP address is assigned to.
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescGZIP(), []int{0}
}

func (x *IP) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *IP) GetVlan() int64 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

func (x *IP) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// A request to list free IP addresses in the database.
type ListFreeIPsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The VLAN to list free IP addresses on.
	Vlan int64 `protobuf:"varint,1,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// The maximum number of free IP addresses to return, or 0 to let the server decide.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListFreeIPsRequest) Reset() {
	*x = ListFreeIPsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFreeIPsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFreeIPsRequest) ProtoMessage() {}

func (x *ListFreeIPsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFreeIPsRequest.ProtoReflect.Descriptor instead.
func (*ListFreeIPsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescGZIP(), []int{1}
}

func (x *ListFreeIPsRequest) GetVlan() int64 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

func (x *ListFreeIPsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// A response containing a list of IP addresses in the database.
type ListIPsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP addresses matching this request.
	Ips []*IP `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"` // TODO(smut): Support page tokens.
}

func (x *ListIPsResponse) Reset() {
	*x = ListIPsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIPsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIPsResponse) ProtoMessage() {}

func (x *ListIPsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIPsResponse.ProtoReflect.Descriptor instead.
func (*ListIPsResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescGZIP(), []int{2}
}

func (x *ListIPsResponse) GetIps() []*IP {
	if x != nil {
		return x.Ips
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72, 0x69, 0x6d,
	0x73, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76,
	0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x6c, 0x61,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x65, 0x65, 0x49, 0x50, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x30, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x50, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x49,
	0x50, 0x52, 0x03, 0x69, 0x70, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72,
	0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_goTypes = []interface{}{
	(*IP)(nil),                 // 0: crimson.IP
	(*ListFreeIPsRequest)(nil), // 1: crimson.ListFreeIPsRequest
	(*ListIPsResponse)(nil),    // 2: crimson.ListIPsResponse
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_depIdxs = []int32{
	0, // 0: crimson.ListIPsResponse.ips:type_name -> crimson.IP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFreeIPsRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIPsResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_ips_proto_depIdxs = nil
}
