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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: go.chromium.org/luci/milo/api/service/v1/rpc.proto

package milopb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	proto "go.chromium.org/luci/buildbucket/proto"
	git "go.chromium.org/luci/common/proto/git"
	config "go.chromium.org/luci/milo/api/config"
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

// A request message for `QueryBlamelist` RPC.
type QueryBlamelistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Gitiles commit of the build.
	//
	// This defines the end_commit of the blamelist.
	// It should be set to the output Gitiles commit of the build.
	// Input Gitiles commit should be used when output gitiles commit is not
	// available.
	GitilesCommit *proto.GitilesCommit `protobuf:"bytes,1,opt,name=gitiles_commit,json=gitilesCommit,proto3" json:"gitiles_commit,omitempty"`
	// The context builder of the blamelist.
	//
	// The start commit of the blamelist is the closest ancestor commit with an
	// associated build that is from the same builder and is not expired,
	// cancelled, or infra-failed.
	Builder *proto.BuilderID `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	// Optional. The maximum number of commits to return.
	//
	// The service may return fewer than this value.
	// If unspecified, at most 100 commits will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A page token, received from a previous `QueryBlamelist` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all parameters provided to `QueryBlamelist`, with the
	// exception of page_size and page_token, must match the call that provided
	// the page token.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Enable multi-project support.
	//
	// When set to false (default), BuildSummary.BuildSets will be used to find
	// the closest ancestor commit with an associated build.
	// When set to true, BuildSummary.BlamelistPins will be used instead. Older
	// builds may not have BlamelistPins populated.
	// TODO(crbugs/1047893): once all recent builds have BlamelistPins populated,
	// remove this flag and use BlamelistPins unconditionally.
	MultiProjectSupport bool `protobuf:"varint,5,opt,name=multi_project_support,json=multiProjectSupport,proto3" json:"multi_project_support,omitempty"`
}

func (x *QueryBlamelistRequest) Reset() {
	*x = QueryBlamelistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlamelistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlamelistRequest) ProtoMessage() {}

func (x *QueryBlamelistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlamelistRequest.ProtoReflect.Descriptor instead.
func (*QueryBlamelistRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryBlamelistRequest) GetGitilesCommit() *proto.GitilesCommit {
	if x != nil {
		return x.GitilesCommit
	}
	return nil
}

func (x *QueryBlamelistRequest) GetBuilder() *proto.BuilderID {
	if x != nil {
		return x.Builder
	}
	return nil
}

func (x *QueryBlamelistRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryBlamelistRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *QueryBlamelistRequest) GetMultiProjectSupport() bool {
	if x != nil {
		return x.MultiProjectSupport
	}
	return false
}

// A response message for QueryBlamelist RPC.
type QueryBlamelistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The commits from the blamelist of the build, in reverse chronological
	// order.
	Commits []*git.Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The repo commit immediately preceding |commits|. Useful for creating
	// git log queries, which are exclusive of the first commit.
	// Unset when |commits| includes the first commit in the repository.
	PrecedingCommit *git.Commit `protobuf:"bytes,3,opt,name=preceding_commit,json=precedingCommit,proto3" json:"preceding_commit,omitempty"`
}

func (x *QueryBlamelistResponse) Reset() {
	*x = QueryBlamelistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlamelistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlamelistResponse) ProtoMessage() {}

func (x *QueryBlamelistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlamelistResponse.ProtoReflect.Descriptor instead.
func (*QueryBlamelistResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryBlamelistResponse) GetCommits() []*git.Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

func (x *QueryBlamelistResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *QueryBlamelistResponse) GetPrecedingCommit() *git.Commit {
	if x != nil {
		return x.PrecedingCommit
	}
	return nil
}

// A stateless page token for QueryBlamelist RPC.
type QueryBlamelistPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The first commit in the next page.
	NextCommitId string `protobuf:"bytes,2,opt,name=next_commit_id,json=nextCommitId,proto3" json:"next_commit_id,omitempty"`
}

func (x *QueryBlamelistPageToken) Reset() {
	*x = QueryBlamelistPageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlamelistPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlamelistPageToken) ProtoMessage() {}

func (x *QueryBlamelistPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlamelistPageToken.ProtoReflect.Descriptor instead.
func (*QueryBlamelistPageToken) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryBlamelistPageToken) GetNextCommitId() string {
	if x != nil {
		return x.NextCommitId
	}
	return ""
}

type GetProjectCfgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project name.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetProjectCfgRequest) Reset() {
	*x = GetProjectCfgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectCfgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectCfgRequest) ProtoMessage() {}

func (x *GetProjectCfgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectCfgRequest.ProtoReflect.Descriptor instead.
func (*GetProjectCfgRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetProjectCfgRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

var File_go_chromium_org_luci_milo_api_service_v1_rpc_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x69, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e,
	0x76, 0x31, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75,
	0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x69, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x6c, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0d, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x16, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x67, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0f, 0x70, 0x72, 0x65,
	0x63, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x3f, 0x0a, 0x17,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x66, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32,
	0xb3, 0x01, 0x0a, 0x0c, 0x4d, 0x69, 0x6c, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x12, 0x5d, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x23, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d,
	0x69, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x61, 0x6d,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x66, 0x67,
	0x12, 0x22, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x66, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x69, 0x6c, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6d, 0x69,
	0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x69, 0x6c, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescData = file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDesc
)

func file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescData)
	})
	return file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDescData
}

var file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_goTypes = []interface{}{
	(*QueryBlamelistRequest)(nil),   // 0: luci.milo.v1.QueryBlamelistRequest
	(*QueryBlamelistResponse)(nil),  // 1: luci.milo.v1.QueryBlamelistResponse
	(*QueryBlamelistPageToken)(nil), // 2: luci.milo.v1.QueryBlamelistPageToken
	(*GetProjectCfgRequest)(nil),    // 3: luci.milo.v1.GetProjectCfgRequest
	(*proto.GitilesCommit)(nil),     // 4: buildbucket.v2.GitilesCommit
	(*proto.BuilderID)(nil),         // 5: buildbucket.v2.BuilderID
	(*git.Commit)(nil),              // 6: git.Commit
	(*config.Project)(nil),          // 7: milo.Project
}
var file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_depIdxs = []int32{
	4, // 0: luci.milo.v1.QueryBlamelistRequest.gitiles_commit:type_name -> buildbucket.v2.GitilesCommit
	5, // 1: luci.milo.v1.QueryBlamelistRequest.builder:type_name -> buildbucket.v2.BuilderID
	6, // 2: luci.milo.v1.QueryBlamelistResponse.commits:type_name -> git.Commit
	6, // 3: luci.milo.v1.QueryBlamelistResponse.preceding_commit:type_name -> git.Commit
	0, // 4: luci.milo.v1.MiloInternal.QueryBlamelist:input_type -> luci.milo.v1.QueryBlamelistRequest
	3, // 5: luci.milo.v1.MiloInternal.GetProjectCfg:input_type -> luci.milo.v1.GetProjectCfgRequest
	1, // 6: luci.milo.v1.MiloInternal.QueryBlamelist:output_type -> luci.milo.v1.QueryBlamelistResponse
	7, // 7: luci.milo.v1.MiloInternal.GetProjectCfg:output_type -> milo.Project
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_init() }
func file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_init() {
	if File_go_chromium_org_luci_milo_api_service_v1_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlamelistRequest); i {
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
		file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlamelistResponse); i {
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
		file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlamelistPageToken); i {
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
		file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectCfgRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_milo_api_service_v1_rpc_proto = out.File
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_rawDesc = nil
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_goTypes = nil
	file_go_chromium_org_luci_milo_api_service_v1_rpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MiloInternalClient is the client API for MiloInternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MiloInternalClient interface {
	// Retrieves blamelist of a build.
	//
	// The blamelist of a build is defined as [end_commit, start_commit)
	// end_commit is the Gitiles commit of the build (specified in gitiles
	// buildset tag).
	// start_commit is the closest ancestor commit with an associated build that
	// is from the same builder and is not expired, cancelled, or infra-failed.
	QueryBlamelist(ctx context.Context, in *QueryBlamelistRequest, opts ...grpc.CallOption) (*QueryBlamelistResponse, error)
	// Gets the project config.
	//
	// Return the config of the project.
	GetProjectCfg(ctx context.Context, in *GetProjectCfgRequest, opts ...grpc.CallOption) (*config.Project, error)
}
type miloInternalPRPCClient struct {
	client *prpc.Client
}

func NewMiloInternalPRPCClient(client *prpc.Client) MiloInternalClient {
	return &miloInternalPRPCClient{client}
}

func (c *miloInternalPRPCClient) QueryBlamelist(ctx context.Context, in *QueryBlamelistRequest, opts ...grpc.CallOption) (*QueryBlamelistResponse, error) {
	out := new(QueryBlamelistResponse)
	err := c.client.Call(ctx, "luci.milo.v1.MiloInternal", "QueryBlamelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miloInternalPRPCClient) GetProjectCfg(ctx context.Context, in *GetProjectCfgRequest, opts ...grpc.CallOption) (*config.Project, error) {
	out := new(config.Project)
	err := c.client.Call(ctx, "luci.milo.v1.MiloInternal", "GetProjectCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type miloInternalClient struct {
	cc grpc.ClientConnInterface
}

func NewMiloInternalClient(cc grpc.ClientConnInterface) MiloInternalClient {
	return &miloInternalClient{cc}
}

func (c *miloInternalClient) QueryBlamelist(ctx context.Context, in *QueryBlamelistRequest, opts ...grpc.CallOption) (*QueryBlamelistResponse, error) {
	out := new(QueryBlamelistResponse)
	err := c.cc.Invoke(ctx, "/luci.milo.v1.MiloInternal/QueryBlamelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *miloInternalClient) GetProjectCfg(ctx context.Context, in *GetProjectCfgRequest, opts ...grpc.CallOption) (*config.Project, error) {
	out := new(config.Project)
	err := c.cc.Invoke(ctx, "/luci.milo.v1.MiloInternal/GetProjectCfg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiloInternalServer is the server API for MiloInternal service.
type MiloInternalServer interface {
	// Retrieves blamelist of a build.
	//
	// The blamelist of a build is defined as [end_commit, start_commit)
	// end_commit is the Gitiles commit of the build (specified in gitiles
	// buildset tag).
	// start_commit is the closest ancestor commit with an associated build that
	// is from the same builder and is not expired, cancelled, or infra-failed.
	QueryBlamelist(context.Context, *QueryBlamelistRequest) (*QueryBlamelistResponse, error)
	// Gets the project config.
	//
	// Return the config of the project.
	GetProjectCfg(context.Context, *GetProjectCfgRequest) (*config.Project, error)
}

// UnimplementedMiloInternalServer can be embedded to have forward compatible implementations.
type UnimplementedMiloInternalServer struct {
}

func (*UnimplementedMiloInternalServer) QueryBlamelist(context.Context, *QueryBlamelistRequest) (*QueryBlamelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBlamelist not implemented")
}
func (*UnimplementedMiloInternalServer) GetProjectCfg(context.Context, *GetProjectCfgRequest) (*config.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectCfg not implemented")
}

func RegisterMiloInternalServer(s prpc.Registrar, srv MiloInternalServer) {
	s.RegisterService(&_MiloInternal_serviceDesc, srv)
}

func _MiloInternal_QueryBlamelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlamelistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiloInternalServer).QueryBlamelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.milo.v1.MiloInternal/QueryBlamelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiloInternalServer).QueryBlamelist(ctx, req.(*QueryBlamelistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MiloInternal_GetProjectCfg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectCfgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiloInternalServer).GetProjectCfg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.milo.v1.MiloInternal/GetProjectCfg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiloInternalServer).GetProjectCfg(ctx, req.(*GetProjectCfgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MiloInternal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luci.milo.v1.MiloInternal",
	HandlerType: (*MiloInternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryBlamelist",
			Handler:    _MiloInternal_QueryBlamelist_Handler,
		},
		{
			MethodName: "GetProjectCfg",
			Handler:    _MiloInternal_GetProjectCfg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/milo/api/service/v1/rpc.proto",
}
