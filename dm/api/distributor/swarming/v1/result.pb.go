// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/dm/api/distributor/swarming/v1/result.proto

package swarmingV1

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

// This is the swarming-specific result for Executions run via swarming.
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode int64 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// The isolated hash of the output directory
	IsolatedOutdir *IsolatedRef `protobuf:"bytes,2,opt,name=isolated_outdir,json=isolatedOutdir,proto3" json:"isolated_outdir,omitempty"`
	// The pinned cipd packages that this task actually used.
	CipdPins *CipdSpec `protobuf:"bytes,3,opt,name=cipd_pins,json=cipdPins,proto3" json:"cipd_pins,omitempty"`
	// The captured snapshot dimensions that the bot actually had.
	SnapshotDimensions map[string]string `protobuf:"bytes,4,rep,name=snapshot_dimensions,json=snapshotDimensions,proto3" json:"snapshot_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Result) GetIsolatedOutdir() *IsolatedRef {
	if x != nil {
		return x.IsolatedOutdir
	}
	return nil
}

func (x *Result) GetCipdPins() *CipdSpec {
	if x != nil {
		return x.CipdPins
	}
	return nil
}

func (x *Result) GetSnapshotDimensions() map[string]string {
	if x != nil {
		return x.SnapshotDimensions
	}
	return nil
}

var File_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x1a, 0x3e,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x64, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a,
	0x0f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x64, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e,
	0x67, 0x56, 0x31, 0x2e, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x66, 0x52,
	0x0e, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x64, 0x69, 0x72, 0x12,
	0x31, 0x0a, 0x09, 0x63, 0x69, 0x70, 0x64, 0x5f, 0x70, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x2e,
	0x43, 0x69, 0x70, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x63, 0x69, 0x70, 0x64, 0x50, 0x69,
	0x6e, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x45, 0x0a, 0x17, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescData = file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDesc
)

func file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescData)
	})
	return file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDescData
}

var file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_goTypes = []interface{}{
	(*Result)(nil),      // 0: swarmingV1.Result
	nil,                 // 1: swarmingV1.Result.SnapshotDimensionsEntry
	(*IsolatedRef)(nil), // 2: swarmingV1.IsolatedRef
	(*CipdSpec)(nil),    // 3: swarmingV1.CipdSpec
}
var file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_depIdxs = []int32{
	2, // 0: swarmingV1.Result.isolated_outdir:type_name -> swarmingV1.IsolatedRef
	3, // 1: swarmingV1.Result.cipd_pins:type_name -> swarmingV1.CipdSpec
	1, // 2: swarmingV1.Result.snapshot_dimensions:type_name -> swarmingV1.Result.SnapshotDimensionsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_init() }
func file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_init() {
	if File_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto != nil {
		return
	}
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_cipd_proto_init()
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_isolate_ref_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto = out.File
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_rawDesc = nil
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_goTypes = nil
	file_go_chromium_org_luci_dm_api_distributor_swarming_v1_result_proto_depIdxs = nil
}
