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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/bisection/task/proto/task.proto

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

// Payload of the FailedBuildIngestionTask task.
type FailedBuildIngestionTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bbid of the failed build.
	Bbid int64 `protobuf:"varint,1,opt,name=bbid,proto3" json:"bbid,omitempty"`
}

func (x *FailedBuildIngestionTask) Reset() {
	*x = FailedBuildIngestionTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedBuildIngestionTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedBuildIngestionTask) ProtoMessage() {}

func (x *FailedBuildIngestionTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedBuildIngestionTask.ProtoReflect.Descriptor instead.
func (*FailedBuildIngestionTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *FailedBuildIngestionTask) GetBbid() int64 {
	if x != nil {
		return x.Bbid
	}
	return 0
}

// Payload of the RevertCulpritTask
type RevertCulpritTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the CompileFailureAnalysis associated with the culprit
	AnalysisId int64 `protobuf:"varint,1,opt,name=analysis_id,json=analysisId,proto3" json:"analysis_id,omitempty"`
	// The ID of the Suspect which is the culprit
	CulpritId int64 `protobuf:"varint,2,opt,name=culprit_id,json=culpritId,proto3" json:"culprit_id,omitempty"`
}

func (x *RevertCulpritTask) Reset() {
	*x = RevertCulpritTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertCulpritTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertCulpritTask) ProtoMessage() {}

func (x *RevertCulpritTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertCulpritTask.ProtoReflect.Descriptor instead.
func (*RevertCulpritTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *RevertCulpritTask) GetAnalysisId() int64 {
	if x != nil {
		return x.AnalysisId
	}
	return 0
}

func (x *RevertCulpritTask) GetCulpritId() int64 {
	if x != nil {
		return x.CulpritId
	}
	return 0
}

// Payload of the CancelAnalysis
type CancelAnalysisTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The analysis ID that we need to cancel.
	AnalysisId int64 `protobuf:"varint,1,opt,name=analysis_id,json=analysisId,proto3" json:"analysis_id,omitempty"`
}

func (x *CancelAnalysisTask) Reset() {
	*x = CancelAnalysisTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelAnalysisTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelAnalysisTask) ProtoMessage() {}

func (x *CancelAnalysisTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelAnalysisTask.ProtoReflect.Descriptor instead.
func (*CancelAnalysisTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescGZIP(), []int{2}
}

func (x *CancelAnalysisTask) GetAnalysisId() int64 {
	if x != nil {
		return x.AnalysisId
	}
	return 0
}

// Payload for Culprit Verification
type CulpritVerificationTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The analysis ID
	AnalysisId int64 `protobuf:"varint,1,opt,name=analysis_id,json=analysisId,proto3" json:"analysis_id,omitempty"`
	// The ID of the suspect for culprit verification
	SuspectId int64 `protobuf:"varint,2,opt,name=suspect_id,json=suspectId,proto3" json:"suspect_id,omitempty"`
	// The encoded datastore key of suspect parent
	ParentKey string `protobuf:"bytes,3,opt,name=parent_key,json=parentKey,proto3" json:"parent_key,omitempty"`
}

func (x *CulpritVerificationTask) Reset() {
	*x = CulpritVerificationTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CulpritVerificationTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CulpritVerificationTask) ProtoMessage() {}

func (x *CulpritVerificationTask) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CulpritVerificationTask.ProtoReflect.Descriptor instead.
func (*CulpritVerificationTask) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescGZIP(), []int{3}
}

func (x *CulpritVerificationTask) GetAnalysisId() int64 {
	if x != nil {
		return x.AnalysisId
	}
	return 0
}

func (x *CulpritVerificationTask) GetSuspectId() int64 {
	if x != nil {
		return x.SuspectId
	}
	return 0
}

func (x *CulpritVerificationTask) GetParentKey() string {
	if x != nil {
		return x.ParentKey
	}
	return ""
}

var File_go_chromium_org_luci_bisection_task_proto_task_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x69, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x18, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x62, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x62, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x11, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74,
	0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x17, 0x43, 0x75, 0x6c,
	0x70, 0x72, 0x69, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x69, 0x73, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescData = file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDesc
)

func file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescData)
	})
	return file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDescData
}

var file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_bisection_task_proto_task_proto_goTypes = []interface{}{
	(*FailedBuildIngestionTask)(nil), // 0: proto.FailedBuildIngestionTask
	(*RevertCulpritTask)(nil),        // 1: proto.RevertCulpritTask
	(*CancelAnalysisTask)(nil),       // 2: proto.CancelAnalysisTask
	(*CulpritVerificationTask)(nil),  // 3: proto.CulpritVerificationTask
}
var file_go_chromium_org_luci_bisection_task_proto_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_bisection_task_proto_task_proto_init() }
func file_go_chromium_org_luci_bisection_task_proto_task_proto_init() {
	if File_go_chromium_org_luci_bisection_task_proto_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedBuildIngestionTask); i {
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
		file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertCulpritTask); i {
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
		file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelAnalysisTask); i {
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
		file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CulpritVerificationTask); i {
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
			RawDescriptor: file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_bisection_task_proto_task_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_bisection_task_proto_task_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_bisection_task_proto_task_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_bisection_task_proto_task_proto = out.File
	file_go_chromium_org_luci_bisection_task_proto_task_proto_rawDesc = nil
	file_go_chromium_org_luci_bisection_task_proto_task_proto_goTypes = nil
	file_go_chromium_org_luci_bisection_task_proto_task_proto_depIdxs = nil
}
