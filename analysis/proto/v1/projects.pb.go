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
// source: go.chromium.org/luci/analysis/proto/v1/projects.proto

package weetbixpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// A request object with data to fetch the list of projects configured
// by Weetbix.
type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP(), []int{0}
}

// A response containing the list of projects which are are using Weetbix.
type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of projects using Weetbix.
	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type GetProjectConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project configuration to retrieve.
	// Format: projects/{project}/config.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetProjectConfigRequest) Reset() {
	*x = GetProjectConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectConfigRequest) ProtoMessage() {}

func (x *GetProjectConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectConfigRequest.ProtoReflect.Descriptor instead.
func (*GetProjectConfigRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP(), []int{2}
}

func (x *GetProjectConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProjectConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the project configuration.
	// Format: projects/{project}/config.
	// See also https://google.aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Details about the monorail project used for this LUCI project.
	Monorail *ProjectConfig_Monorail `protobuf:"bytes,2,opt,name=monorail,proto3" json:"monorail,omitempty"`
}

func (x *ProjectConfig) Reset() {
	*x = ProjectConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectConfig) ProtoMessage() {}

func (x *ProjectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[3]
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
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectConfig) GetMonorail() *ProjectConfig_Monorail {
	if x != nil {
		return x.Monorail
	}
	return nil
}

type ProjectConfig_Monorail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The monorail project used for this LUCI project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The shortlink format used for this bug tracker.
	// For example, "crbug.com".
	DisplayPrefix string `protobuf:"bytes,2,opt,name=display_prefix,json=displayPrefix,proto3" json:"display_prefix,omitempty"`
}

func (x *ProjectConfig_Monorail) Reset() {
	*x = ProjectConfig_Monorail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectConfig_Monorail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectConfig_Monorail) ProtoMessage() {}

func (x *ProjectConfig_Monorail) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectConfig_Monorail.ProtoReflect.Descriptor instead.
func (*ProjectConfig_Monorail) Descriptor() ([]byte, []int) {
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ProjectConfig_Monorail) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectConfig_Monorail) GetDisplayPrefix() string {
	if x != nil {
		return x.DisplayPrefix
	}
	return ""
}

var File_infra_appengine_weetbix_proto_v1_projects_proto protoreflect.FileDescriptor

var file_infra_appengine_weetbix_proto_v1_projects_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77,
	0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x2d, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb0, 0x01, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x08, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61,
	0x69, 0x6c, 0x1a, 0x4b, 0x0a, 0x08, 0x4d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x32,
	0xa6, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x2e, 0x77, 0x65, 0x65, 0x74,
	0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x65, 0x65, 0x74, 0x62, 0x69, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x77, 0x65, 0x65, 0x74,
	0x62, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x65,
	0x74, 0x62, 0x69, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescOnce sync.Once
	file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescData = file_infra_appengine_weetbix_proto_v1_projects_proto_rawDesc
)

func file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescGZIP() []byte {
	file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescOnce.Do(func() {
		file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescData)
	})
	return file_infra_appengine_weetbix_proto_v1_projects_proto_rawDescData
}

var file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infra_appengine_weetbix_proto_v1_projects_proto_goTypes = []interface{}{
	(*ListProjectsRequest)(nil),     // 0: weetbix.v1.ListProjectsRequest
	(*ListProjectsResponse)(nil),    // 1: weetbix.v1.ListProjectsResponse
	(*GetProjectConfigRequest)(nil), // 2: weetbix.v1.GetProjectConfigRequest
	(*ProjectConfig)(nil),           // 3: weetbix.v1.ProjectConfig
	(*ProjectConfig_Monorail)(nil),  // 4: weetbix.v1.ProjectConfig.Monorail
	(*Project)(nil),                 // 5: weetbix.v1.Project
}
var file_infra_appengine_weetbix_proto_v1_projects_proto_depIdxs = []int32{
	5, // 0: weetbix.v1.ListProjectsResponse.projects:type_name -> weetbix.v1.Project
	4, // 1: weetbix.v1.ProjectConfig.monorail:type_name -> weetbix.v1.ProjectConfig.Monorail
	2, // 2: weetbix.v1.Projects.GetConfig:input_type -> weetbix.v1.GetProjectConfigRequest
	0, // 3: weetbix.v1.Projects.List:input_type -> weetbix.v1.ListProjectsRequest
	3, // 4: weetbix.v1.Projects.GetConfig:output_type -> weetbix.v1.ProjectConfig
	1, // 5: weetbix.v1.Projects.List:output_type -> weetbix.v1.ListProjectsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_appengine_weetbix_proto_v1_projects_proto_init() }
func file_infra_appengine_weetbix_proto_v1_projects_proto_init() {
	if File_infra_appengine_weetbix_proto_v1_projects_proto != nil {
		return
	}
	file_infra_appengine_weetbix_proto_v1_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectConfigRequest); i {
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
		file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectConfig_Monorail); i {
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
			RawDescriptor: file_infra_appengine_weetbix_proto_v1_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_appengine_weetbix_proto_v1_projects_proto_goTypes,
		DependencyIndexes: file_infra_appengine_weetbix_proto_v1_projects_proto_depIdxs,
		MessageInfos:      file_infra_appengine_weetbix_proto_v1_projects_proto_msgTypes,
	}.Build()
	File_infra_appengine_weetbix_proto_v1_projects_proto = out.File
	file_infra_appengine_weetbix_proto_v1_projects_proto_rawDesc = nil
	file_infra_appengine_weetbix_proto_v1_projects_proto_goTypes = nil
	file_infra_appengine_weetbix_proto_v1_projects_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectsClient interface {
	// Gets Weetbix configuration for a LUCI Project.
	//
	// RPC desigend to comply with https://google.aip.dev/131.
	GetConfig(ctx context.Context, in *GetProjectConfigRequest, opts ...grpc.CallOption) (*ProjectConfig, error)
	// Lists LUCI Projects visible to the user.
	//
	// RPC compliant with https://google.aip.dev/132.
	// This RPC is incomplete. Future breaking changes are
	// expressly flagged.
	List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
}
type projectsPRPCClient struct {
	client *prpc.Client
}

func NewProjectsPRPCClient(client *prpc.Client) ProjectsClient {
	return &projectsPRPCClient{client}
}

func (c *projectsPRPCClient) GetConfig(ctx context.Context, in *GetProjectConfigRequest, opts ...grpc.CallOption) (*ProjectConfig, error) {
	out := new(ProjectConfig)
	err := c.client.Call(ctx, "weetbix.v1.Projects", "GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsPRPCClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.client.Call(ctx, "weetbix.v1.Projects", "List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) GetConfig(ctx context.Context, in *GetProjectConfigRequest, opts ...grpc.CallOption) (*ProjectConfig, error) {
	out := new(ProjectConfig)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Projects/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/weetbix.v1.Projects/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
type ProjectsServer interface {
	// Gets Weetbix configuration for a LUCI Project.
	//
	// RPC desigend to comply with https://google.aip.dev/131.
	GetConfig(context.Context, *GetProjectConfigRequest) (*ProjectConfig, error)
	// Lists LUCI Projects visible to the user.
	//
	// RPC compliant with https://google.aip.dev/132.
	// This RPC is incomplete. Future breaking changes are
	// expressly flagged.
	List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
}

// UnimplementedProjectsServer can be embedded to have forward compatible implementations.
type UnimplementedProjectsServer struct {
}

func (*UnimplementedProjectsServer) GetConfig(context.Context, *GetProjectConfigRequest) (*ProjectConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (*UnimplementedProjectsServer) List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterProjectsServer(s prpc.Registrar, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Projects/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetConfig(ctx, req.(*GetProjectConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weetbix.v1.Projects/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).List(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weetbix.v1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Projects_GetConfig_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Projects_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/analysis/proto/v1/projects.proto",
}
