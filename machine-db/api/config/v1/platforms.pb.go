// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/machine-db/api/config/v1/platforms.proto

package config

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

// A platform to store in the database.
type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this platform. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this platform.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The name of the hardware manufacturer of this platform.
	Manufacturer string `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescGZIP(), []int{0}
}

func (x *Platform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Platform) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Platform) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

// A list of platforms.
type Platforms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of platforms.
	Platform []*Platform `protobuf:"bytes,1,rep,name=platform,proto3" json:"platform,omitempty"`
}

func (x *Platforms) Reset() {
	*x = Platforms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platforms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platforms) ProtoMessage() {}

func (x *Platforms) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platforms.ProtoReflect.Descriptor instead.
func (*Platforms) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescGZIP(), []int{1}
}

func (x *Platforms) GetPlatform() []*Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

var File_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x64, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x22, 0x39, 0x0a,
	0x09, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescData = file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDesc
)

func file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescData)
	})
	return file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDescData
}

var file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_goTypes = []interface{}{
	(*Platform)(nil),  // 0: config.Platform
	(*Platforms)(nil), // 1: config.Platforms
}
var file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_depIdxs = []int32{
	0, // 0: config.Platforms.platform:type_name -> config.Platform
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_init() }
func file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_init() {
	if File_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platforms); i {
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
			RawDescriptor: file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto = out.File
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_rawDesc = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_goTypes = nil
	file_go_chromium_org_luci_machine_db_api_config_v1_platforms_proto_depIdxs = nil
}
