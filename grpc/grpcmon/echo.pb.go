// Copyright 2021 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/grpc/grpcmon/echo.proto

package grpcmon

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

type SayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SayRequest) Reset() {
	*x = SayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayRequest) ProtoMessage() {}

func (x *SayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayRequest.ProtoReflect.Descriptor instead.
func (*SayRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescGZIP(), []int{0}
}

func (x *SayRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SayResponse) Reset() {
	*x = SayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayResponse) ProtoMessage() {}

func (x *SayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayResponse.ProtoReflect.Descriptor instead.
func (*SayResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescGZIP(), []int{1}
}

func (x *SayResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_go_chromium_org_luci_grpc_grpcmon_echo_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x3a, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f,
	0x12, 0x32, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescData = file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDesc
)

func file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescData)
	})
	return file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDescData
}

var file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_grpc_grpcmon_echo_proto_goTypes = []interface{}{
	(*SayRequest)(nil),  // 0: grpcmon.SayRequest
	(*SayResponse)(nil), // 1: grpcmon.SayResponse
}
var file_go_chromium_org_luci_grpc_grpcmon_echo_proto_depIdxs = []int32{
	0, // 0: grpcmon.Echo.Say:input_type -> grpcmon.SayRequest
	1, // 1: grpcmon.Echo.Say:output_type -> grpcmon.SayResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_grpc_grpcmon_echo_proto_init() }
func file_go_chromium_org_luci_grpc_grpcmon_echo_proto_init() {
	if File_go_chromium_org_luci_grpc_grpcmon_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayRequest); i {
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
		file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_grpc_grpcmon_echo_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_grpc_grpcmon_echo_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_grpc_grpcmon_echo_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_grpc_grpcmon_echo_proto = out.File
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_rawDesc = nil
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_goTypes = nil
	file_go_chromium_org_luci_grpc_grpcmon_echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}
type echoPRPCClient struct {
	client *prpc.Client
}

func NewEchoPRPCClient(client *prpc.Client) EchoClient {
	return &echoPRPCClient{client}
}

func (c *echoPRPCClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.client.Call(ctx, "grpcmon.Echo", "Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/grpcmon.Echo/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterEchoServer(s prpc.Registrar, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcmon.Echo/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcmon.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Echo_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/grpcmon/echo.proto",
}
