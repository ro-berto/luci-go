// Copyright 2021 The LUCI Authors.
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
// source: go.chromium.org/luci/cv/api/rpc/v0/runs.proto

package rpcpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	v1 "go.chromium.org/luci/cv/api/common/v1"
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

type GetRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is CV Run ID, e.g. "projects/chromium/runs/9991234120-1-badcafe"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRunRequest) Reset() {
	*x = GetRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRunRequest) ProtoMessage() {}

func (x *GetRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRunRequest.ProtoReflect.Descriptor instead.
func (*GetRunRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescGZIP(), []int{0}
}

func (x *GetRunRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_go_chromium_org_luci_cv_api_rpc_v0_runs_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x30, 0x1a, 0x2f, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x3d, 0x0a, 0x04,
	0x52, 0x75, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x18,
	0x2e, 0x63, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x76, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x30, 0x3b, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescData = file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_goTypes = []interface{}{
	(*GetRunRequest)(nil), // 0: cv.rpc.v0.GetRunRequest
	(*v1.Run)(nil),        // 1: cv.common.v1.Run
}
var file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_depIdxs = []int32{
	0, // 0: cv.rpc.v0.Runs.GetRun:input_type -> cv.rpc.v0.GetRunRequest
	1, // 1: cv.rpc.v0.Runs.GetRun:output_type -> cv.common.v1.Run
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_init() }
func file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_init() {
	if File_go_chromium_org_luci_cv_api_rpc_v0_runs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRunRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_rpc_v0_runs_proto = out.File
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_rpc_v0_runs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RunsClient is the client API for Runs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunsClient interface {
	// GetRun returns Run details.
	GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*v1.Run, error)
}
type runsPRPCClient struct {
	client *prpc.Client
}

func NewRunsPRPCClient(client *prpc.Client) RunsClient {
	return &runsPRPCClient{client}
}

func (c *runsPRPCClient) GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*v1.Run, error) {
	out := new(v1.Run)
	err := c.client.Call(ctx, "cv.rpc.v0.Runs", "GetRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type runsClient struct {
	cc grpc.ClientConnInterface
}

func NewRunsClient(cc grpc.ClientConnInterface) RunsClient {
	return &runsClient{cc}
}

func (c *runsClient) GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*v1.Run, error) {
	out := new(v1.Run)
	err := c.cc.Invoke(ctx, "/cv.rpc.v0.Runs/GetRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunsServer is the server API for Runs service.
type RunsServer interface {
	// GetRun returns Run details.
	GetRun(context.Context, *GetRunRequest) (*v1.Run, error)
}

// UnimplementedRunsServer can be embedded to have forward compatible implementations.
type UnimplementedRunsServer struct {
}

func (*UnimplementedRunsServer) GetRun(context.Context, *GetRunRequest) (*v1.Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRun not implemented")
}

func RegisterRunsServer(s prpc.Registrar, srv RunsServer) {
	s.RegisterService(&_Runs_serviceDesc, srv)
}

func _Runs_GetRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunsServer).GetRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cv.rpc.v0.Runs/GetRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunsServer).GetRun(ctx, req.(*GetRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cv.rpc.v0.Runs",
	HandlerType: (*RunsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRun",
			Handler:    _Runs_GetRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cv/api/rpc/v0/runs.proto",
}
