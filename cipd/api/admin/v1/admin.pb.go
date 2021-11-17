// Copyright 2018 The LUCI Authors.
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/cipd/api/admin/v1/admin.proto

package api

import (
	dsmapperpb "go.chromium.org/luci/server/dsmapper/dsmapperpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Supported mapping jobs.
type MapperKind int32

const (
	MapperKind_MAPPER_KIND_UNSPECIFIED MapperKind = 0
	// Dump names of all packages to GAE logs, to test mapping jobs framework.
	MapperKind_ENUMERATE_PACKAGES MapperKind = 1
	// Find tags that don't pass ValidateInstanceTag and marks them.
	MapperKind_FIND_MALFORMED_TAGS MapperKind = 2
	// Exports all tags into a BigQuery table 'exported_tags'.
	MapperKind_EXPORT_TAGS_TO_BQ MapperKind = 3
)

// Enum value maps for MapperKind.
var (
	MapperKind_name = map[int32]string{
		0: "MAPPER_KIND_UNSPECIFIED",
		1: "ENUMERATE_PACKAGES",
		2: "FIND_MALFORMED_TAGS",
		3: "EXPORT_TAGS_TO_BQ",
	}
	MapperKind_value = map[string]int32{
		"MAPPER_KIND_UNSPECIFIED": 0,
		"ENUMERATE_PACKAGES":      1,
		"FIND_MALFORMED_TAGS":     2,
		"EXPORT_TAGS_TO_BQ":       3,
	}
)

func (x MapperKind) Enum() *MapperKind {
	p := new(MapperKind)
	*p = x
	return p
}

func (x MapperKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapperKind) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_enumTypes[0].Descriptor()
}

func (MapperKind) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_enumTypes[0]
}

func (x MapperKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapperKind.Descriptor instead.
func (MapperKind) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

// Defines what a mapping job should do.
type JobConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    MapperKind `protobuf:"varint,1,opt,name=kind,proto3,enum=cipd.MapperKind" json:"kind,omitempty"`
	Comment string     `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"` // arbitrary human readable string
	DryRun  bool       `protobuf:"varint,3,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *JobConfig) Reset() {
	*x = JobConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobConfig) ProtoMessage() {}

func (x *JobConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobConfig.ProtoReflect.Descriptor instead.
func (*JobConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *JobConfig) GetKind() MapperKind {
	if x != nil {
		return x.Kind
	}
	return MapperKind_MAPPER_KIND_UNSPECIFIED
}

func (x *JobConfig) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *JobConfig) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

// Identifies an instance of a mapping job.
type JobID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId int64 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *JobID) Reset() {
	*x = JobID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobID) ProtoMessage() {}

func (x *JobID) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobID.ProtoReflect.Descriptor instead.
func (*JobID) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *JobID) GetJobId() int64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

// Details about a mapping job.
type JobState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Original job config, exactly as it was submitted to LaunchJob.
	Config *JobConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// Current state of the job and all its shards.
	Info *dsmapperpb.JobInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *JobState) Reset() {
	*x = JobState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobState) ProtoMessage() {}

func (x *JobState) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobState.ProtoReflect.Descriptor instead.
func (*JobState) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *JobState) GetConfig() *JobConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *JobState) GetInfo() *dsmapperpb.JobInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Result of running FixMarkedTags.
type TagFixReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fixed []*TagFixReport_Tag `protobuf:"bytes,1,rep,name=fixed,proto3" json:"fixed,omitempty"`
}

func (x *TagFixReport) Reset() {
	*x = TagFixReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagFixReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFixReport) ProtoMessage() {}

func (x *TagFixReport) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFixReport.ProtoReflect.Descriptor instead.
func (*TagFixReport) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *TagFixReport) GetFixed() []*TagFixReport_Tag {
	if x != nil {
		return x.Fixed
	}
	return nil
}

type TagFixReport_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pkg       string `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Instance  string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	BrokenTag string `protobuf:"bytes,3,opt,name=broken_tag,json=brokenTag,proto3" json:"broken_tag,omitempty"`
	FixedTag  string `protobuf:"bytes,4,opt,name=fixed_tag,json=fixedTag,proto3" json:"fixed_tag,omitempty"` // or "" if it was deleted
}

func (x *TagFixReport_Tag) Reset() {
	*x = TagFixReport_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagFixReport_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFixReport_Tag) ProtoMessage() {}

func (x *TagFixReport_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFixReport_Tag.ProtoReflect.Descriptor instead.
func (*TagFixReport_Tag) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TagFixReport_Tag) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *TagFixReport_Tag) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *TagFixReport_Tag) GetBrokenTag() string {
	if x != nil {
		return x.BrokenTag
	}
	return ""
}

func (x *TagFixReport_Tag) GetFixedTag() string {
	if x != nil {
		return x.FixedTag
	}
	return ""
}

var File_go_chromium_org_luci_cipd_api_admin_v1_admin_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69, 0x70, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x73, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x64, 0x73,
	0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x22, 0x1e, 0x0a,
	0x05, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x66, 0x0a,
	0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64,
	0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x64,
	0x73, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x46, 0x69, 0x78,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x69, 0x78, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x54, 0x61, 0x67,
	0x46, 0x69, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x05, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x1a, 0x6f, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x6b, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x54, 0x61, 0x67, 0x2a, 0x71, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x4b, 0x49,
	0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41,
	0x43, 0x4b, 0x41, 0x47, 0x45, 0x53, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4e, 0x44,
	0x5f, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x54, 0x41, 0x47, 0x53, 0x10,
	0x02, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x53,
	0x5f, 0x54, 0x4f, 0x5f, 0x42, 0x51, 0x10, 0x03, 0x32, 0xc1, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12,
	0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x0b, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x2f, 0x0a,
	0x08, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0b, 0x2e, 0x63, 0x69, 0x70, 0x64,
	0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e,
	0x63, 0x69, 0x70, 0x64, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x63, 0x69, 0x70,
	0x64, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x0d, 0x46, 0x69,
	0x78, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x0b, 0x2e, 0x63, 0x69,
	0x70, 0x64, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e,
	0x54, 0x61, 0x67, 0x46, 0x69, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescData = file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_goTypes = []interface{}{
	(MapperKind)(0),            // 0: cipd.MapperKind
	(*JobConfig)(nil),          // 1: cipd.JobConfig
	(*JobID)(nil),              // 2: cipd.JobID
	(*JobState)(nil),           // 3: cipd.JobState
	(*TagFixReport)(nil),       // 4: cipd.TagFixReport
	(*TagFixReport_Tag)(nil),   // 5: cipd.TagFixReport.Tag
	(*dsmapperpb.JobInfo)(nil), // 6: luci.server.dsmapper.JobInfo
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_depIdxs = []int32{
	0, // 0: cipd.JobConfig.kind:type_name -> cipd.MapperKind
	1, // 1: cipd.JobState.config:type_name -> cipd.JobConfig
	6, // 2: cipd.JobState.info:type_name -> luci.server.dsmapper.JobInfo
	5, // 3: cipd.TagFixReport.fixed:type_name -> cipd.TagFixReport.Tag
	1, // 4: cipd.Admin.LaunchJob:input_type -> cipd.JobConfig
	2, // 5: cipd.Admin.AbortJob:input_type -> cipd.JobID
	2, // 6: cipd.Admin.GetJobState:input_type -> cipd.JobID
	2, // 7: cipd.Admin.FixMarkedTags:input_type -> cipd.JobID
	2, // 8: cipd.Admin.LaunchJob:output_type -> cipd.JobID
	7, // 9: cipd.Admin.AbortJob:output_type -> google.protobuf.Empty
	3, // 10: cipd.Admin.GetJobState:output_type -> cipd.JobState
	4, // 11: cipd.Admin.FixMarkedTags:output_type -> cipd.TagFixReport
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_init() }
func file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_init() {
	if File_go_chromium_org_luci_cipd_api_admin_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobConfig); i {
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
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobID); i {
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
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobState); i {
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
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagFixReport); i {
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
		file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagFixReport_Tag); i {
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
			RawDescriptor: file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_api_admin_v1_admin_proto = out.File
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_api_admin_v1_admin_proto_depIdxs = nil
}
