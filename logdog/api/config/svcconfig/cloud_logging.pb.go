// Copyright 2021 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/logdog/api/config/svcconfig/cloud_logging.proto

package svcconfig

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

// CloudLoggingConfig specifies how Logdog should export log entries to
// Cloud Logging.
type CloudLoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project, where Logdog logs are exported.
	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// If true, Logdog uses its default credential for log exports
	// to the destination project. If false, it uses the project-scoped account.
	//
	// NOTE: Using the global logdog account is insecure from a client-isolation
	// point of view, and means that any LUCI client could send their logs to
	// your project. This likely wouldn't happen maliciously, but if someone
	// copy-pasted your configuration into their own project, you may end up
	// paying to index and store their logs.
	UseGlobalLogdogAccount bool `protobuf:"varint,2,opt,name=use_global_logdog_account,json=useGlobalLogdogAccount,proto3" json:"use_global_logdog_account,omitempty"`
}

func (x *CloudLoggingConfig) Reset() {
	*x = CloudLoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudLoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudLoggingConfig) ProtoMessage() {}

func (x *CloudLoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudLoggingConfig.ProtoReflect.Descriptor instead.
func (*CloudLoggingConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescGZIP(), []int{0}
}

func (x *CloudLoggingConfig) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *CloudLoggingConfig) GetUseGlobalLogdogAccount() bool {
	if x != nil {
		return x.UseGlobalLogdogAccount
	}
	return false
}

var File_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x71, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x19, 0x75, 0x73, 0x65,
	0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x73,
	0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67,
	0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73,
	0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescData = file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_goTypes = []interface{}{
	(*CloudLoggingConfig)(nil), // 0: svcconfig.CloudLoggingConfig
}
var file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_init() }
func file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_init() {
	if File_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudLoggingConfig); i {
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
			RawDescriptor: file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto = out.File
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_depIdxs = nil
}
