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
// 	protoc        v3.17.0
// source: go.chromium.org/luci/resultdb/internal/proto/ui/ui.proto

package uipb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	v1 "go.chromium.org/luci/resultdb/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x69, 0x2f, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6c, 0x75, 0x63, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x75, 0x69, 0x1a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x74, 0x0a, 0x02,
	0x55, 0x49, 0x12, 0x6e, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x69, 0x3b, 0x75, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_goTypes = []interface{}{
	(*v1.QueryTestVariantsRequest)(nil),  // 0: luci.resultdb.v1.QueryTestVariantsRequest
	(*v1.QueryTestVariantsResponse)(nil), // 1: luci.resultdb.v1.QueryTestVariantsResponse
}
var file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_depIdxs = []int32{
	0, // 0: luci.resultdb.internal.ui.UI.QueryTestVariants:input_type -> luci.resultdb.v1.QueryTestVariantsRequest
	1, // 1: luci.resultdb.internal.ui.UI.QueryTestVariants:output_type -> luci.resultdb.v1.QueryTestVariantsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_init() }
func file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_init() {
	if File_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_depIdxs,
	}.Build()
	File_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto = out.File
	file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_rawDesc = nil
	file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_goTypes = nil
	file_go_chromium_org_luci_resultdb_internal_proto_ui_ui_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UIClient is the client API for UI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UIClient interface {
	// Retrieves test variants from an invocation, recursively.
	// Supports invocation inclusions.
	// For displaying test variants in the UI.
	QueryTestVariants(ctx context.Context, in *v1.QueryTestVariantsRequest, opts ...grpc.CallOption) (*v1.QueryTestVariantsResponse, error)
}
type uIPRPCClient struct {
	client *prpc.Client
}

func NewUIPRPCClient(client *prpc.Client) UIClient {
	return &uIPRPCClient{client}
}

func (c *uIPRPCClient) QueryTestVariants(ctx context.Context, in *v1.QueryTestVariantsRequest, opts ...grpc.CallOption) (*v1.QueryTestVariantsResponse, error) {
	out := new(v1.QueryTestVariantsResponse)
	err := c.client.Call(ctx, "luci.resultdb.internal.ui.UI", "QueryTestVariants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type uIClient struct {
	cc grpc.ClientConnInterface
}

func NewUIClient(cc grpc.ClientConnInterface) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) QueryTestVariants(ctx context.Context, in *v1.QueryTestVariantsRequest, opts ...grpc.CallOption) (*v1.QueryTestVariantsResponse, error) {
	out := new(v1.QueryTestVariantsResponse)
	err := c.cc.Invoke(ctx, "/luci.resultdb.internal.ui.UI/QueryTestVariants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UIServer is the server API for UI service.
type UIServer interface {
	// Retrieves test variants from an invocation, recursively.
	// Supports invocation inclusions.
	// For displaying test variants in the UI.
	QueryTestVariants(context.Context, *v1.QueryTestVariantsRequest) (*v1.QueryTestVariantsResponse, error)
}

// UnimplementedUIServer can be embedded to have forward compatible implementations.
type UnimplementedUIServer struct {
}

func (*UnimplementedUIServer) QueryTestVariants(context.Context, *v1.QueryTestVariantsRequest) (*v1.QueryTestVariantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTestVariants not implemented")
}

func RegisterUIServer(s prpc.Registrar, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_QueryTestVariants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.QueryTestVariantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).QueryTestVariants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultdb.internal.ui.UI/QueryTestVariants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).QueryTestVariants(ctx, req.(*v1.QueryTestVariantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luci.resultdb.internal.ui.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTestVariants",
			Handler:    _UI_QueryTestVariants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/resultdb/internal/proto/ui/ui.proto",
}
