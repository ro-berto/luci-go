// Copyright 2020 The Swarming Authors. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: go.chromium.org/luci/buildbucket/proto/builder_service.proto

package buildbucketpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A request message for GetBuilder rpc.
type GetBuilderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the builder to return.
	Id *BuilderID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBuilderRequest) Reset() {
	*x = GetBuilderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuilderRequest) ProtoMessage() {}

func (x *GetBuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuilderRequest.ProtoReflect.Descriptor instead.
func (*GetBuilderRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBuilderRequest) GetId() *BuilderID {
	if x != nil {
		return x.Id
	}
	return nil
}

// A request message for ListBuilders.
type ListBuildersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LUCI project, e.g. "chromium".
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// A bucket in the project, e.g. "try".
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The maximum number of builders to return.
	//
	// The service may return fewer than this value.
	// If unspecified, at most 100 builders will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListBuilders` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListBuilders` MUST
	// match the call that provided the page token.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBuildersRequest) Reset() {
	*x = ListBuildersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildersRequest) ProtoMessage() {}

func (x *ListBuildersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildersRequest.ProtoReflect.Descriptor instead.
func (*ListBuildersRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBuildersRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ListBuildersRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ListBuildersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBuildersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// A response message for ListBuilders.
type ListBuildersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matched builders.
	Builders []*BuilderItem `protobuf:"bytes,1,rep,name=builders,proto3" json:"builders,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there were no subsequent pages at the time of
	// request.
	// If the invocation is not finalized, more results may appear later.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBuildersResponse) Reset() {
	*x = ListBuildersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildersResponse) ProtoMessage() {}

func (x *ListBuildersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildersResponse.ProtoReflect.Descriptor instead.
func (*ListBuildersResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBuildersResponse) GetBuilders() []*BuilderItem {
	if x != nil {
		return x.Builders
	}
	return nil
}

func (x *ListBuildersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_go_chromium_org_luci_buildbucket_proto_builder_service_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x77, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb7, 0x01, 0x0a, 0x08, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_goTypes = []interface{}{
	(*GetBuilderRequest)(nil),    // 0: buildbucket.v2.GetBuilderRequest
	(*ListBuildersRequest)(nil),  // 1: buildbucket.v2.ListBuildersRequest
	(*ListBuildersResponse)(nil), // 2: buildbucket.v2.ListBuildersResponse
	(*BuilderID)(nil),            // 3: buildbucket.v2.BuilderID
	(*BuilderItem)(nil),          // 4: buildbucket.v2.BuilderItem
}
var file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_depIdxs = []int32{
	3, // 0: buildbucket.v2.GetBuilderRequest.id:type_name -> buildbucket.v2.BuilderID
	4, // 1: buildbucket.v2.ListBuildersResponse.builders:type_name -> buildbucket.v2.BuilderItem
	0, // 2: buildbucket.v2.Builders.GetBuilder:input_type -> buildbucket.v2.GetBuilderRequest
	1, // 3: buildbucket.v2.Builders.ListBuilders:input_type -> buildbucket.v2.ListBuildersRequest
	4, // 4: buildbucket.v2.Builders.GetBuilder:output_type -> buildbucket.v2.BuilderItem
	2, // 5: buildbucket.v2.Builders.ListBuilders:output_type -> buildbucket.v2.ListBuildersResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_builder_service_proto != nil {
		return
	}
	file_go_chromium_org_luci_buildbucket_proto_builder_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuilderRequest); i {
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
		file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildersRequest); i {
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
		file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildersResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_builder_service_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_builder_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildersClient is the client API for Builders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildersClient interface {
	// Looks up one builder.
	GetBuilder(ctx context.Context, in *GetBuilderRequest, opts ...grpc.CallOption) (*BuilderItem, error)
	// Lists all builders of a project or a bucket.
	ListBuilders(ctx context.Context, in *ListBuildersRequest, opts ...grpc.CallOption) (*ListBuildersResponse, error)
}
type buildersPRPCClient struct {
	client *prpc.Client
}

func NewBuildersPRPCClient(client *prpc.Client) BuildersClient {
	return &buildersPRPCClient{client}
}

func (c *buildersPRPCClient) GetBuilder(ctx context.Context, in *GetBuilderRequest, opts ...grpc.CallOption) (*BuilderItem, error) {
	out := new(BuilderItem)
	err := c.client.Call(ctx, "buildbucket.v2.Builders", "GetBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildersPRPCClient) ListBuilders(ctx context.Context, in *ListBuildersRequest, opts ...grpc.CallOption) (*ListBuildersResponse, error) {
	out := new(ListBuildersResponse)
	err := c.client.Call(ctx, "buildbucket.v2.Builders", "ListBuilders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildersClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildersClient(cc grpc.ClientConnInterface) BuildersClient {
	return &buildersClient{cc}
}

func (c *buildersClient) GetBuilder(ctx context.Context, in *GetBuilderRequest, opts ...grpc.CallOption) (*BuilderItem, error) {
	out := new(BuilderItem)
	err := c.cc.Invoke(ctx, "/buildbucket.v2.Builders/GetBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildersClient) ListBuilders(ctx context.Context, in *ListBuildersRequest, opts ...grpc.CallOption) (*ListBuildersResponse, error) {
	out := new(ListBuildersResponse)
	err := c.cc.Invoke(ctx, "/buildbucket.v2.Builders/ListBuilders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildersServer is the server API for Builders service.
type BuildersServer interface {
	// Looks up one builder.
	GetBuilder(context.Context, *GetBuilderRequest) (*BuilderItem, error)
	// Lists all builders of a project or a bucket.
	ListBuilders(context.Context, *ListBuildersRequest) (*ListBuildersResponse, error)
}

// UnimplementedBuildersServer can be embedded to have forward compatible implementations.
type UnimplementedBuildersServer struct {
}

func (*UnimplementedBuildersServer) GetBuilder(context.Context, *GetBuilderRequest) (*BuilderItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuilder not implemented")
}
func (*UnimplementedBuildersServer) ListBuilders(context.Context, *ListBuildersRequest) (*ListBuildersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuilders not implemented")
}

func RegisterBuildersServer(s prpc.Registrar, srv BuildersServer) {
	s.RegisterService(&_Builders_serviceDesc, srv)
}

func _Builders_GetBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildersServer).GetBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builders/GetBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildersServer).GetBuilder(ctx, req.(*GetBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builders_ListBuilders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildersServer).ListBuilders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builders/ListBuilders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildersServer).ListBuilders(ctx, req.(*ListBuildersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbucket.v2.Builders",
	HandlerType: (*BuildersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuilder",
			Handler:    _Builders_GetBuilder_Handler,
		},
		{
			MethodName: "ListBuilders",
			Handler:    _Builders_ListBuilders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/buildbucket/proto/builder_service.proto",
}
