// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto

package e2etest

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x70, 0x63,
	0x2f, 0x65, 0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x65, 0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x3c, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x33, 0x0a, 0x05,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65,
	0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x70, 0x63, 0x2f, 0x65, 0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescData = file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDesc
)

func file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescData)
	})
	return file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDescData
}

var file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: e2etest.HelloRequest
	(*HelloReply)(nil),   // 1: e2etest.HelloReply
}
var file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_depIdxs = []int32{
	0, // 0: e2etest.Hello.Greet:input_type -> e2etest.HelloRequest
	1, // 1: e2etest.Hello.Greet:output_type -> e2etest.HelloReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_init() }
func file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_init() {
	if File_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto = out.File
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_rawDesc = nil
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_goTypes = nil
	file_go_chromium_org_luci_grpc_prpc_e2etest_helloworld_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
type helloPRPCClient struct {
	client *prpc.Client
}

func NewHelloPRPCClient(client *prpc.Client) HelloClient {
	return &helloPRPCClient{client}
}

func (c *helloPRPCClient) Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.client.Call(ctx, "e2etest.Hello", "Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Greet(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/e2etest.Hello/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Greet(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Greet(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterHelloServer(s prpc.Registrar, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/e2etest.Hello/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Greet(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "e2etest.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Hello_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/prpc/e2etest/helloworld_test.proto",
}
