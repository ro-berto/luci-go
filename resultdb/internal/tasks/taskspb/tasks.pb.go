// Copyright 2020 The LUCI Authors.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/resultdb/internal/tasks/taskspb/tasks.proto

package taskspb

import (
	v1 "go.chromium.org/luci/resultdb/proto/v1"
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

type TryFinalizeInvocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
}

func (x *TryFinalizeInvocation) Reset() {
	*x = TryFinalizeInvocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TryFinalizeInvocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TryFinalizeInvocation) ProtoMessage() {}

func (x *TryFinalizeInvocation) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TryFinalizeInvocation.ProtoReflect.Descriptor instead.
func (*TryFinalizeInvocation) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *TryFinalizeInvocation) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

type ExportInvocationTestResultsToBQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvocationId string             `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	BqExport     *v1.BigQueryExport `protobuf:"bytes,2,opt,name=bq_export,json=bqExport,proto3" json:"bq_export,omitempty"`
}

func (x *ExportInvocationTestResultsToBQ) Reset() {
	*x = ExportInvocationTestResultsToBQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportInvocationTestResultsToBQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvocationTestResultsToBQ) ProtoMessage() {}

func (x *ExportInvocationTestResultsToBQ) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportInvocationTestResultsToBQ.ProtoReflect.Descriptor instead.
func (*ExportInvocationTestResultsToBQ) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *ExportInvocationTestResultsToBQ) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *ExportInvocationTestResultsToBQ) GetBqExport() *v1.BigQueryExport {
	if x != nil {
		return x.BqExport
	}
	return nil
}

type ExportInvocationArtifactsToBQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvocationId string             `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	BqExport     *v1.BigQueryExport `protobuf:"bytes,2,opt,name=bq_export,json=bqExport,proto3" json:"bq_export,omitempty"`
}

func (x *ExportInvocationArtifactsToBQ) Reset() {
	*x = ExportInvocationArtifactsToBQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportInvocationArtifactsToBQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportInvocationArtifactsToBQ) ProtoMessage() {}

func (x *ExportInvocationArtifactsToBQ) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportInvocationArtifactsToBQ.ProtoReflect.Descriptor instead.
func (*ExportInvocationArtifactsToBQ) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *ExportInvocationArtifactsToBQ) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *ExportInvocationArtifactsToBQ) GetBqExport() *v1.BigQueryExport {
	if x != nil {
		return x.BqExport
	}
	return nil
}

var File_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64,
	0x62, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x1a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x15, 0x54, 0x72, 0x79,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x1f, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x54, 0x6f, 0x42, 0x51, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x09, 0x62, 0x71, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x62, 0x71, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x83, 0x01, 0x0a, 0x1d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x42,
	0x51, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x62, 0x71, 0x5f, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x62, 0x71, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescData = file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_goTypes = []interface{}{
	(*TryFinalizeInvocation)(nil),           // 0: luci.resultdb.internal.tasks.TryFinalizeInvocation
	(*ExportInvocationTestResultsToBQ)(nil), // 1: luci.resultdb.internal.tasks.ExportInvocationTestResultsToBQ
	(*ExportInvocationArtifactsToBQ)(nil),   // 2: luci.resultdb.internal.tasks.ExportInvocationArtifactsToBQ
	(*v1.BigQueryExport)(nil),               // 3: luci.resultdb.v1.BigQueryExport
}
var file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_depIdxs = []int32{
	3, // 0: luci.resultdb.internal.tasks.ExportInvocationTestResultsToBQ.bq_export:type_name -> luci.resultdb.v1.BigQueryExport
	3, // 1: luci.resultdb.internal.tasks.ExportInvocationArtifactsToBQ.bq_export:type_name -> luci.resultdb.v1.BigQueryExport
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_init() }
func file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_init() {
	if File_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TryFinalizeInvocation); i {
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
		file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportInvocationTestResultsToBQ); i {
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
		file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportInvocationArtifactsToBQ); i {
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
			RawDescriptor: file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto = out.File
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_internal_tasks_taskspb_tasks_proto_depIdxs = nil
}
