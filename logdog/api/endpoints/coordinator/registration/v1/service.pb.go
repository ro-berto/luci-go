// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto

package logdog

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RegisterPrefixRequest registers a new Prefix with the Coordinator.
type RegisterPrefixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The log stream's project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The log stream prefix to register.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// The realm name (within the project) to associate the stream prefix with.
	//
	// This realm contains ACLs defining who will be able to read logs under this
	// prefix.
	//
	// The caller should have "logdog.logs.create" permission in this realm.
	Realm string `protobuf:"bytes,5,opt,name=realm,proto3" json:"realm,omitempty"`
	// Optional information about the registering agent.
	SourceInfo []string `protobuf:"bytes,3,rep,name=source_info,json=sourceInfo,proto3" json:"source_info,omitempty"`
	// Optional nonce to allow retries of this RPC. ALL CLIENTS SHOULD PROVIDE
	// THIS. The client should generate the nonce once while preparing the request
	// message, and then re-use the same nonce for retries of the request.
	//
	// The nonce should be 32 bytes of random data.
	// The nonce must not be reused between different requests (only for retries
	//   of the same request).
	//
	// NOTE: This is currently optional, but once all clients have upgraded to
	// this scheme, it will become mandatory. During the transition if this is
	// omitted, then NO RETRIES will be allowed for this request, if the server
	// processes it correctly but the client fails to get the response from the
	// server.
	OpNonce []byte `protobuf:"bytes,4,opt,name=op_nonce,json=opNonce,proto3" json:"op_nonce,omitempty"`
	// The prefix expiration time. If <= 0, the project's default prefix
	// expiration period will be applied.
	//
	// The prefix will be closed by the Coordinator after its expiration period.
	// Once closed, new stream registration requests will no longer be accepted.
	//
	// If supplied, this value should exceed the timeout of the local task, else
	// some of the task's streams may be dropped due to failing registration.
	Expiration *durationpb.Duration `protobuf:"bytes,10,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *RegisterPrefixRequest) Reset() {
	*x = RegisterPrefixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPrefixRequest) ProtoMessage() {}

func (x *RegisterPrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPrefixRequest.ProtoReflect.Descriptor instead.
func (*RegisterPrefixRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterPrefixRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *RegisterPrefixRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *RegisterPrefixRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *RegisterPrefixRequest) GetSourceInfo() []string {
	if x != nil {
		return x.SourceInfo
	}
	return nil
}

func (x *RegisterPrefixRequest) GetOpNonce() []byte {
	if x != nil {
		return x.OpNonce
	}
	return nil
}

func (x *RegisterPrefixRequest) GetExpiration() *durationpb.Duration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// The response message for the RegisterPrefix RPC.
type RegisterPrefixResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Secret is the prefix's secret. This must be included verbatim in Butler
	// bundles to assert ownership of this prefix.
	Secret []byte `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	// The name of the Pub/Sub topic to publish butlerproto-formatted Butler log
	// bundles to.
	LogBundleTopic string `protobuf:"bytes,2,opt,name=log_bundle_topic,json=logBundleTopic,proto3" json:"log_bundle_topic,omitempty"`
}

func (x *RegisterPrefixResponse) Reset() {
	*x = RegisterPrefixResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPrefixResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPrefixResponse) ProtoMessage() {}

func (x *RegisterPrefixResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPrefixResponse.ProtoReflect.Descriptor instead.
func (*RegisterPrefixResponse) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterPrefixResponse) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *RegisterPrefixResponse) GetLogBundleTopic() string {
	if x != nil {
		return x.LogBundleTopic
	}
	return ""
}

var File_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x6f, 0x70, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x5f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x32, 0x5f, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x64, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x64,
	0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67,
	0x64, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescData = file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDesc
)

func file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescData)
	})
	return file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDescData
}

var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_goTypes = []interface{}{
	(*RegisterPrefixRequest)(nil),  // 0: logdog.RegisterPrefixRequest
	(*RegisterPrefixResponse)(nil), // 1: logdog.RegisterPrefixResponse
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_depIdxs = []int32{
	2, // 0: logdog.RegisterPrefixRequest.expiration:type_name -> google.protobuf.Duration
	0, // 1: logdog.Registration.RegisterPrefix:input_type -> logdog.RegisterPrefixRequest
	1, // 2: logdog.Registration.RegisterPrefix:output_type -> logdog.RegisterPrefixResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_init()
}
func file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_init() {
	if File_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPrefixRequest); i {
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
		file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPrefixResponse); i {
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
			RawDescriptor: file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto = out.File
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_rawDesc = nil
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_goTypes = nil
	file_go_chromium_org_luci_logdog_api_endpoints_coordinator_registration_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationClient interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error)
}
type registrationPRPCClient struct {
	client *prpc.Client
}

func NewRegistrationPRPCClient(client *prpc.Client) RegistrationClient {
	return &registrationPRPCClient{client}
}

func (c *registrationPRPCClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := c.client.Call(ctx, "logdog.Registration", "RegisterPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type registrationClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationClient(cc grpc.ClientConnInterface) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := c.cc.Invoke(ctx, "/logdog.Registration/RegisterPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
type RegistrationServer interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(context.Context, *RegisterPrefixRequest) (*RegisterPrefixResponse, error)
}

// UnimplementedRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedRegistrationServer struct {
}

func (*UnimplementedRegistrationServer) RegisterPrefix(context.Context, *RegisterPrefixRequest) (*RegisterPrefixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPrefix not implemented")
}

func RegisterRegistrationServer(s prpc.Registrar, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_RegisterPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).RegisterPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Registration/RegisterPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).RegisterPrefix(ctx, req.(*RegisterPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logdog.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPrefix",
			Handler:    _Registration_RegisterPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto",
}
