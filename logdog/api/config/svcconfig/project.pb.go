// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/logdog/api/config/svcconfig/project.proto

package svcconfig

import (
	_ "go.chromium.org/luci/common/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProjectConfig is a set of per-project configuration parameters. Each
// luci-config project must include one of these configs in order to register
// or view log streams in that project's log stream space.
//
// A project's configuration should reside in the "projects/<project>" config
// set and be named "<app-id>.cfg".
//
// Many of the parameters here can be bounded by GlobalConfig parameters.
type ProjectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated and unused. Use 'role/logdog.reader' realms role instead.
	//
	// Deprecated: Do not use.
	ReaderAuthGroups []string `protobuf:"bytes,2,rep,name=reader_auth_groups,json=readerAuthGroups,proto3" json:"reader_auth_groups,omitempty"`
	// Deprecated and unused. Use 'role/logdog.writer' realms role instead.
	//
	// Deprecated: Do not use.
	WriterAuthGroups []string `protobuf:"bytes,3,rep,name=writer_auth_groups,json=writerAuthGroups,proto3" json:"writer_auth_groups,omitempty"`
	// The maximum amount of time after a prefix has been registered when log
	// streams may also be registered under that prefix.
	//
	// See Config's "prefix_expiration" for more information.
	PrefixExpiration *durationpb.Duration `protobuf:"bytes,5,opt,name=prefix_expiration,json=prefixExpiration,proto3" json:"prefix_expiration,omitempty"`
	// The archival Google Storage bucket name.
	//
	// Log streams artifacts will be stored in a subdirectory of this bucket:
	// gs://<archive_gs_bucket>/<app-id>/<project-name>/<log-path>/artifact...
	//
	// Note that the Archivist microservice must have WRITE access to this
	// bucket, and the Coordinator must have READ access.
	//
	// If this is not set, the logs will be archived in a project-named
	// subdirectory in the global "archive_gs_base" location.
	ArchiveGsBucket string `protobuf:"bytes,10,opt,name=archive_gs_bucket,json=archiveGsBucket,proto3" json:"archive_gs_bucket,omitempty"`
	// Project-specific archive index configuration.
	//
	// Any unspecified index configuration will default to the service archival
	// config.
	ArchiveIndexConfig *ArchiveIndexConfig `protobuf:"bytes,12,opt,name=archive_index_config,json=archiveIndexConfig,proto3" json:"archive_index_config,omitempty"`
	// Project-specific CloudLogging configuration.
	//
	// If not specified, no Logdog logs will be exported to Cloud Logging.
	CloudLoggingConfig *CloudLoggingConfig `protobuf:"bytes,13,opt,name=cloud_logging_config,json=cloudLoggingConfig,proto3" json:"cloud_logging_config,omitempty"`
}

func (x *ProjectConfig) Reset() {
	*x = ProjectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectConfig) ProtoMessage() {}

func (x *ProjectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectConfig.ProtoReflect.Descriptor instead.
func (*ProjectConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *ProjectConfig) GetReaderAuthGroups() []string {
	if x != nil {
		return x.ReaderAuthGroups
	}
	return nil
}

// Deprecated: Do not use.
func (x *ProjectConfig) GetWriterAuthGroups() []string {
	if x != nil {
		return x.WriterAuthGroups
	}
	return nil
}

func (x *ProjectConfig) GetPrefixExpiration() *durationpb.Duration {
	if x != nil {
		return x.PrefixExpiration
	}
	return nil
}

func (x *ProjectConfig) GetArchiveGsBucket() string {
	if x != nil {
		return x.ArchiveGsBucket
	}
	return ""
}

func (x *ProjectConfig) GetArchiveIndexConfig() *ArchiveIndexConfig {
	if x != nil {
		return x.ArchiveIndexConfig
	}
	return nil
}

func (x *ProjectConfig) GetCloudLoggingConfig() *CloudLoggingConfig {
	if x != nil {
		return x.CloudLoggingConfig
	}
	return nil
}

var File_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x12, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x30, 0x0a, 0x12, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x10, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x73, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x47, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4f, 0x0a,
	0x14, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x76,
	0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4f,
	0x0a, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x76, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x52, 0x0e, 0x6d, 0x61, 0x78,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x61, 0x67, 0x65, 0x52, 0x12, 0x72, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x42,
	0x78, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x76, 0x63, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0xa2, 0xfe, 0x23, 0x42, 0x0a, 0x40, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x3a, 0x6c, 0x75, 0x63, 0x69, 0x2d, 0x6c,
	0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x66, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescData = file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_goTypes = []interface{}{
	(*ProjectConfig)(nil),       // 0: svcconfig.ProjectConfig
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
	(*ArchiveIndexConfig)(nil),  // 2: svcconfig.ArchiveIndexConfig
	(*CloudLoggingConfig)(nil),  // 3: svcconfig.CloudLoggingConfig
}
var file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_depIdxs = []int32{
	1, // 0: svcconfig.ProjectConfig.prefix_expiration:type_name -> google.protobuf.Duration
	2, // 1: svcconfig.ProjectConfig.archive_index_config:type_name -> svcconfig.ArchiveIndexConfig
	3, // 2: svcconfig.ProjectConfig.cloud_logging_config:type_name -> svcconfig.CloudLoggingConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_init() }
func file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_init() {
	if File_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto != nil {
		return
	}
	file_go_chromium_org_luci_logdog_api_config_svcconfig_archival_proto_init()
	file_go_chromium_org_luci_logdog_api_config_svcconfig_cloud_logging_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectConfig); i {
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
			RawDescriptor: file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto = out.File
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_config_svcconfig_project_proto_depIdxs = nil
}
