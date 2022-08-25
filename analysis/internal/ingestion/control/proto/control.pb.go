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
// source: go.chromium.org/luci/analysis/internal/ingestion/control/proto/control.proto

package controlpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	v1 "go.chromium.org/luci/analysis/proto/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BuildResult represents the result from the buildbucket pub/sub
// that should be passed to the result ingestion task.
type BuildResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Buildbucket build ID, unique per Buildbucket instance.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Buildbucket host, e.g. "cr-buildbucket.appspot.com".
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// The time the build was created.
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	// The LUCI Project to which the build belongs.
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *BuildResult) Reset() {
	*x = BuildResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResult) ProtoMessage() {}

func (x *BuildResult) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResult.ProtoReflect.Descriptor instead.
func (*BuildResult) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescGZIP(), []int{0}
}

func (x *BuildResult) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BuildResult) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *BuildResult) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *BuildResult) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

// PresubmitResult represents the result from the presubmit pub/sub
// that should be passed to the result ingestion task.
type PresubmitResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identity of the presubmit run. If the ingestion does not relate to test
	// results obtained in a presubmit run, this field should not be set.
	PresubmitRunId *v1.PresubmitRunId `protobuf:"bytes,1,opt,name=presubmit_run_id,json=presubmitRunId,proto3" json:"presubmit_run_id,omitempty"`
	// The ending status of the presubmit run. E.g. Canceled, Success, Failure.
	Status v1.PresubmitRunStatus `protobuf:"varint,9,opt,name=status,proto3,enum=weetbix.v1.PresubmitRunStatus" json:"status,omitempty"`
	// The presubmit run mode.
	// E.g. FULL_RUN, DRY_RUN, QUICK_DRY_RUN.
	Mode v1.PresubmitRunMode `protobuf:"varint,8,opt,name=mode,proto3,enum=weetbix.v1.PresubmitRunMode" json:"mode,omitempty"`
	// The owner of the presubmit run (if any).
	// This is the owner of the CL on which CQ+1/CQ+2 was clicked
	// (even in case of presubmit run with multiple CLs).
	// There is scope for this field to become an email address if privacy
	// approval is obtained, until then it is "automation" (for automation
	// service accounts) and "user" otherwise.
	Owner string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	// The time the presubmit was created.
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	// Whether the build was critical to the completion of the presubmit run.
	// True if the failure of the build would cause the presubmit run to fail.
	Critical bool `protobuf:"varint,7,opt,name=critical,proto3" json:"critical,omitempty"`
}

func (x *PresubmitResult) Reset() {
	*x = PresubmitResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresubmitResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresubmitResult) ProtoMessage() {}

func (x *PresubmitResult) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresubmitResult.ProtoReflect.Descriptor instead.
func (*PresubmitResult) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescGZIP(), []int{1}
}

func (x *PresubmitResult) GetPresubmitRunId() *v1.PresubmitRunId {
	if x != nil {
		return x.PresubmitRunId
	}
	return nil
}

func (x *PresubmitResult) GetStatus() v1.PresubmitRunStatus {
	if x != nil {
		return x.Status
	}
	return v1.PresubmitRunStatus(0)
}

func (x *PresubmitResult) GetMode() v1.PresubmitRunMode {
	if x != nil {
		return x.Mode
	}
	return v1.PresubmitRunMode(0)
}

func (x *PresubmitResult) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PresubmitResult) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *PresubmitResult) GetCritical() bool {
	if x != nil {
		return x.Critical
	}
	return false
}

var File_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDesc = []byte{
	0x0a, 0x46, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69,
	0x78, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77,
	0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a,
	0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xc6, 0x02, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x65, 0x65, 0x74,
	0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x77, 0x65,
	0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04,
	0x08, 0x06, 0x10, 0x07, 0x42, 0x44, 0x5a, 0x42, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescData = file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDesc
)

func file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDescData
}

var file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_goTypes = []interface{}{
	(*BuildResult)(nil),           // 0: weetbix.internal.ingestion.control.BuildResult
	(*PresubmitResult)(nil),       // 1: weetbix.internal.ingestion.control.PresubmitResult
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*v1.PresubmitRunId)(nil),     // 3: weetbix.v1.PresubmitRunId
	(v1.PresubmitRunStatus)(0),    // 4: weetbix.v1.PresubmitRunStatus
	(v1.PresubmitRunMode)(0),      // 5: weetbix.v1.PresubmitRunMode
}
var file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_depIdxs = []int32{
	2, // 0: weetbix.internal.ingestion.control.BuildResult.creation_time:type_name -> google.protobuf.Timestamp
	3, // 1: weetbix.internal.ingestion.control.PresubmitResult.presubmit_run_id:type_name -> weetbix.v1.PresubmitRunId
	4, // 2: weetbix.internal.ingestion.control.PresubmitResult.status:type_name -> weetbix.v1.PresubmitRunStatus
	5, // 3: weetbix.internal.ingestion.control.PresubmitResult.mode:type_name -> weetbix.v1.PresubmitRunMode
	2, // 4: weetbix.internal.ingestion.control.PresubmitResult.creation_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_init() }
func file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_init() {
	if File_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResult); i {
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
		file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresubmitResult); i {
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
			RawDescriptor: file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_depIdxs,
		MessageInfos:      file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto = out.File
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_rawDesc = nil
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_goTypes = nil
	file_infra_appengine_weetbix_internal_ingestion_control_proto_control_proto_depIdxs = nil
}
