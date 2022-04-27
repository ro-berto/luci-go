// Copyright 2022 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/server/quota/proto/config.proto

package proto

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

// A Policy represents an abstract quota policy.
//
// Policies should be defined in ways that are relevant to each service. For
// example, a policy may represent "builds for builder B", in which case
// resources may be interpreted as "number of builds", or a policy may represent
// "storage in database D", in which case resources may be interpreted as
// "stored bytes".
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A globally unique name.
	//
	// Must start with a letter. Allowed characters (no spaces): A-Z a-z 0-9 - _ /
	// The substring "${user}" can be used in this name to define identical
	// per-user policies. Must not exceed 64 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The amount of resources available. Must not be negative.
	Resources int64 `protobuf:"varint,2,opt,name=resources,proto3" json:"resources,omitempty"`
	// The amount of resources to replenish every second. Must not be negative.
	Replenishment int64 `protobuf:"varint,3,opt,name=replenishment,proto3" json:"replenishment,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetResources() int64 {
	if x != nil {
		return x.Resources
	}
	return 0
}

func (x *Policy) GetReplenishment() int64 {
	if x != nil {
		return x.Replenishment
	}
	return 0
}

// A Config encapsulates a set of quota policies.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy []*Policy `protobuf:"bytes,1,rep,name=policy,proto3" json:"policy,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetPolicy() []*Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_go_chromium_org_luci_server_quota_proto_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_server_quota_proto_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x2f, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescData = file_go_chromium_org_luci_server_quota_proto_config_proto_rawDesc
)

func file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_server_quota_proto_config_proto_rawDescData
}

var file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_server_quota_proto_config_proto_goTypes = []interface{}{
	(*Policy)(nil), // 0: proto.Policy
	(*Config)(nil), // 1: proto.Config
}
var file_go_chromium_org_luci_server_quota_proto_config_proto_depIdxs = []int32{
	0, // 0: proto.Config.policy:type_name -> proto.Policy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_server_quota_proto_config_proto_init() }
func file_go_chromium_org_luci_server_quota_proto_config_proto_init() {
	if File_go_chromium_org_luci_server_quota_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_go_chromium_org_luci_server_quota_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_server_quota_proto_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_server_quota_proto_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_server_quota_proto_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_server_quota_proto_config_proto = out.File
	file_go_chromium_org_luci_server_quota_proto_config_proto_rawDesc = nil
	file_go_chromium_org_luci_server_quota_proto_config_proto_goTypes = nil
	file_go_chromium_org_luci_server_quota_proto_config_proto_depIdxs = nil
}
