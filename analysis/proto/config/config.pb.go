// Copyright 2022 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: go.chromium.org/luci/analysis/proto/config/config.proto

package configpb

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

// Config is the service-wide configuration data for LUCI Analysis.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The endpoint for Monorail APIs.
	MonorailHostname string `protobuf:"bytes,1,opt,name=monorail_hostname,json=monorailHostname,proto3" json:"monorail_hostname,omitempty"`
	// The GCS bucket that chunk contents should be archived to.
	ChunkGcsBucket string `protobuf:"bytes,2,opt,name=chunk_gcs_bucket,json=chunkGcsBucket,proto3" json:"chunk_gcs_bucket,omitempty"`
	// The number of workers to use when re-clustering. Maximum value is 1000,
	// which is the default max_concurrent_requests on the reclustering queue:
	// https://cloud.google.com/appengine/docs/standard/go111/config/queueref.
	//
	// If this is unset or zero, re-clustering is disabled.
	ReclusteringWorkers int64 `protobuf:"varint,3,opt,name=reclustering_workers,json=reclusteringWorkers,proto3" json:"reclustering_workers,omitempty"`
	// The frequency by which to re-cluster. This is specified as a
	// number of minutes. Maximum value is 9, which is one minute less than
	// the 10 minute hard request deadline for autoscaled GAE instances:
	// https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed.
	//
	// If this is unset or zero, re-clustering is disabled.
	ReclusteringIntervalMinutes int64 `protobuf:"varint,4,opt,name=reclustering_interval_minutes,json=reclusteringIntervalMinutes,proto3" json:"reclustering_interval_minutes,omitempty"`
	// Controls whether LUCI Analysis will interact with bug-filing systems.
	// Can be used to stop LUCI Analysis auto-bug filing and updates in
	// response to a problem.
	BugUpdatesEnabled bool `protobuf:"varint,5,opt,name=bug_updates_enabled,json=bugUpdatesEnabled,proto3" json:"bug_updates_enabled,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_analysis_proto_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_analysis_proto_config_config_proto_msgTypes[0]
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
	return file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetMonorailHostname() string {
	if x != nil {
		return x.MonorailHostname
	}
	return ""
}

func (x *Config) GetChunkGcsBucket() string {
	if x != nil {
		return x.ChunkGcsBucket
	}
	return ""
}

func (x *Config) GetReclusteringWorkers() int64 {
	if x != nil {
		return x.ReclusteringWorkers
	}
	return 0
}

func (x *Config) GetReclusteringIntervalMinutes() int64 {
	if x != nil {
		return x.ReclusteringIntervalMinutes
	}
	return 0
}

func (x *Config) GetBugUpdatesEnabled() bool {
	if x != nil {
		return x.BugUpdatesEnabled
	}
	return false
}

var File_go_chromium_org_luci_analysis_proto_config_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6c, 0x75, 0x63, 0x69, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x86, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f,
	0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x67, 0x63, 0x73, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x47, 0x63, 0x73, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x72, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x1d, 0x72, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x72, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x75, 0x67, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x62, 0x75, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescData = file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDesc
)

func file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDescData
}

var file_go_chromium_org_luci_analysis_proto_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_analysis_proto_config_config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: luci.analysis.config.Config
}
var file_go_chromium_org_luci_analysis_proto_config_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_analysis_proto_config_config_proto_init() }
func file_go_chromium_org_luci_analysis_proto_config_config_proto_init() {
	if File_go_chromium_org_luci_analysis_proto_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_analysis_proto_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_analysis_proto_config_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_analysis_proto_config_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_analysis_proto_config_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_analysis_proto_config_config_proto = out.File
	file_go_chromium_org_luci_analysis_proto_config_config_proto_rawDesc = nil
	file_go_chromium_org_luci_analysis_proto_config_config_proto_goTypes = nil
	file_go_chromium_org_luci_analysis_proto_config_config_proto_depIdxs = nil
}
