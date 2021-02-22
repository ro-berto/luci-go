// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/machine-db/api/crimson/v1/switches.proto

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

// A switch in the database.
type Switch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this switch. Uniquely identifies this switch.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this switch.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The number of ports this switch has.
	Ports int32 `protobuf:"varint,3,opt,name=ports,proto3" json:"ports,omitempty"`
	// The datacenter this switch belongs to.
	Datacenter string `protobuf:"bytes,4,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	// The rack this switch belongs to.
	Rack string `protobuf:"bytes,5,opt,name=rack,proto3" json:"rack,omitempty"`
	// The state of this switch.
	State v1.State `protobuf:"varint,6,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
}

func (x *Switch) Reset() {
	*x = Switch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Switch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Switch) ProtoMessage() {}

func (x *Switch) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Switch.ProtoReflect.Descriptor instead.
func (*Switch) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescGZIP(), []int{0}
}

func (x *Switch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Switch) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Switch) GetPorts() int32 {
	if x != nil {
		return x.Ports
	}
	return 0
}

func (x *Switch) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *Switch) GetRack() string {
	if x != nil {
		return x.Rack
	}
	return ""
}

func (x *Switch) GetState() v1.State {
	if x != nil {
		return x.State
	}
	return v1.State_STATE_UNSPECIFIED
}

// A request to list switches in the database.
type ListSwitchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The names of switches to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The datacenters to filter retrieved switches on.
	Datacenters []string `protobuf:"bytes,2,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
	// The racks to filter retrieved switches on.
	Racks []string `protobuf:"bytes,3,rep,name=racks,proto3" json:"racks,omitempty"`
}

func (x *ListSwitchesRequest) Reset() {
	*x = ListSwitchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSwitchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSwitchesRequest) ProtoMessage() {}

func (x *ListSwitchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSwitchesRequest.ProtoReflect.Descriptor instead.
func (*ListSwitchesRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescGZIP(), []int{1}
}

func (x *ListSwitchesRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *ListSwitchesRequest) GetDatacenters() []string {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

func (x *ListSwitchesRequest) GetRacks() []string {
	if x != nil {
		return x.Racks
	}
	return nil
}

// A response containing a list of switches in the database.
type ListSwitchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The switches matching the request.
	Switches []*Switch `protobuf:"bytes,1,rep,name=switches,proto3" json:"switches,omitempty"`
}

func (x *ListSwitchesResponse) Reset() {
	*x = ListSwitchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSwitchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSwitchesResponse) ProtoMessage() {}

func (x *ListSwitchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSwitchesResponse.ProtoReflect.Descriptor instead.
func (*ListSwitchesResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescGZIP(), []int{2}
}

func (x *ListSwitchesResponse) GetSwitches() []*Switch {
	if x != nil {
		return x.Switches
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x1a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12,
	0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x43, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x72, 0x69, 0x6d, 0x73, 0x6f, 0x6e, 0x2e, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x52, 0x08, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_goTypes = []interface{}{
	(*Switch)(nil),               // 0: crimson.Switch
	(*ListSwitchesRequest)(nil),  // 1: crimson.ListSwitchesRequest
	(*ListSwitchesResponse)(nil), // 2: crimson.ListSwitchesResponse
	(v1.State)(0),                // 3: common.State
}
var file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_depIdxs = []int32{
	3, // 0: crimson.Switch.state:type_name -> common.State
	0, // 1: crimson.ListSwitchesResponse.switches:type_name -> crimson.Switch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_init() }
func file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Switch); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSwitchesRequest); i {
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
		file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSwitchesResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto = out.File
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_crimson_v1_switches_proto_depIdxs = nil
}
