// Copyright 2019 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/server/auth/service/protocol/components/auth/proto/security_config.proto

package protocol

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

// SecurityConfig is read from 'security.cfg' by Auth Service and distributed to
// all linked services (in its serialized form) as part of AuthDB proto.
//
// See AuthDB.security_config in replication.proto.
type SecurityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of regular expressions matching hostnames that should be recognized
	// as being a part of single LUCI deployment.
	//
	// Different microservices within a single LUCI deployment may trust each
	// other. This setting (coupled with the TLS certificate check) allows
	// a service to recognize that a target of an RPC is another internal service
	// belonging to the same LUCI deployment.
	//
	// '^' and '$' are implied. The regexp language is intersection of Python and
	// Golang regexp languages and thus should use only very standard features
	// common to both.
	//
	// Example: "(.*-dot-)?chromium-swarm\.appspot\.com".
	InternalServiceRegexp []string `protobuf:"bytes,1,rep,name=internal_service_regexp,json=internalServiceRegexp,proto3" json:"internal_service_regexp,omitempty"`
}

func (x *SecurityConfig) Reset() {
	*x = SecurityConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityConfig) ProtoMessage() {}

func (x *SecurityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityConfig.ProtoReflect.Descriptor instead.
func (*SecurityConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityConfig) GetInternalServiceRegexp() []string {
	if x != nil {
		return x.InternalServiceRegexp
	}
	return nil
}

var File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescData = file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDesc
)

func file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDescData
}

var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_goTypes = []interface{}{
	(*SecurityConfig)(nil), // 0: components.auth.SecurityConfig
}
var file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_init()
}
func file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_init() {
	if File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityConfig); i {
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
			RawDescriptor: file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto = out.File
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_rawDesc = nil
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_goTypes = nil
	file_go_chromium_org_luci_server_auth_service_protocol_components_auth_proto_security_config_proto_depIdxs = nil
}
