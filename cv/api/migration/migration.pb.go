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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cv/api/migration/migration.proto

package migrationpb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "go.chromium.org/luci/cv/api/bigquery/v1"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReportRunsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runs []*ReportedRun `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
}

func (x *ReportRunsRequest) Reset() {
	*x = ReportRunsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRunsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRunsRequest) ProtoMessage() {}

func (x *ReportRunsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRunsRequest.ProtoReflect.Descriptor instead.
func (*ReportRunsRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRunsRequest) GetRuns() []*ReportedRun {
	if x != nil {
		return x.Runs
	}
	return nil
}

type ReportedRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Re-use BQ attempt for now.
	Attempt *v1.Attempt `protobuf:"bytes,1,opt,name=attempt,proto3" json:"attempt,omitempty"`
}

func (x *ReportedRun) Reset() {
	*x = ReportedRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportedRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportedRun) ProtoMessage() {}

func (x *ReportedRun) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportedRun.ProtoReflect.Descriptor instead.
func (*ReportedRun) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{1}
}

func (x *ReportedRun) GetAttempt() *v1.Attempt {
	if x != nil {
		return x.Attempt
	}
	return nil
}

type ReportFinishedRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Run *ReportedRun `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *ReportFinishedRunRequest) Reset() {
	*x = ReportFinishedRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFinishedRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFinishedRunRequest) ProtoMessage() {}

func (x *ReportFinishedRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFinishedRunRequest.ProtoReflect.Descriptor instead.
func (*ReportFinishedRunRequest) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP(), []int{2}
}

func (x *ReportFinishedRunRequest) GetRun() *ReportedRun {
	if x != nil {
		return x.Run
	}
	return nil
}

var File_go_chromium_org_luci_cv_api_migration_migration_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x72,
	0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x75,
	0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x22, 0x3a, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x22, 0x44, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x03, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x52, 0x75, 0x6e, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x32, 0xa1, 0x01, 0x0a, 0x09, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x11, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x75, 0x6e,
	0x12, 0x23, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData = file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes = []interface{}{
	(*ReportRunsRequest)(nil),        // 0: migration.ReportRunsRequest
	(*ReportedRun)(nil),              // 1: migration.ReportedRun
	(*ReportFinishedRunRequest)(nil), // 2: migration.ReportFinishedRunRequest
	(*v1.Attempt)(nil),               // 3: bigquery.Attempt
	(*empty.Empty)(nil),              // 4: google.protobuf.Empty
}
var file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs = []int32{
	1, // 0: migration.ReportRunsRequest.runs:type_name -> migration.ReportedRun
	3, // 1: migration.ReportedRun.attempt:type_name -> bigquery.Attempt
	1, // 2: migration.ReportFinishedRunRequest.run:type_name -> migration.ReportedRun
	0, // 3: migration.Migration.ReportRuns:input_type -> migration.ReportRunsRequest
	2, // 4: migration.Migration.ReportFinishedRun:input_type -> migration.ReportFinishedRunRequest
	4, // 5: migration.Migration.ReportRuns:output_type -> google.protobuf.Empty
	4, // 6: migration.Migration.ReportFinishedRun:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_migration_migration_proto_init() }
func file_go_chromium_org_luci_cv_api_migration_migration_proto_init() {
	if File_go_chromium_org_luci_cv_api_migration_migration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRunsRequest); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportedRun); i {
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
		file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFinishedRunRequest); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_api_migration_migration_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_migration_migration_proto = out.File
	file_go_chromium_org_luci_cv_api_migration_migration_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_migration_migration_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_migration_migration_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MigrationClient is the client API for Migration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MigrationClient interface {
	// ReportRuns notifies CV of the Runs CQDaemon is currently working with.
	//
	// Used to determine whether CV's view of the world matches that of CQDaemon.
	// Initially, this is just FYI for CV.
	ReportRuns(ctx context.Context, in *ReportRunsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// ReportFinishedRun notifies CV of the Run CQDaemon has just finalized.
	ReportFinishedRun(ctx context.Context, in *ReportFinishedRunRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}
type migrationPRPCClient struct {
	client *prpc.Client
}

func NewMigrationPRPCClient(client *prpc.Client) MigrationClient {
	return &migrationPRPCClient{client}
}

func (c *migrationPRPCClient) ReportRuns(ctx context.Context, in *ReportRunsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "migration.Migration", "ReportRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationPRPCClient) ReportFinishedRun(ctx context.Context, in *ReportFinishedRunRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "migration.Migration", "ReportFinishedRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type migrationClient struct {
	cc grpc.ClientConnInterface
}

func NewMigrationClient(cc grpc.ClientConnInterface) MigrationClient {
	return &migrationClient{cc}
}

func (c *migrationClient) ReportRuns(ctx context.Context, in *ReportRunsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/migration.Migration/ReportRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) ReportFinishedRun(ctx context.Context, in *ReportFinishedRunRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/migration.Migration/ReportFinishedRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServer is the server API for Migration service.
type MigrationServer interface {
	// ReportRuns notifies CV of the Runs CQDaemon is currently working with.
	//
	// Used to determine whether CV's view of the world matches that of CQDaemon.
	// Initially, this is just FYI for CV.
	ReportRuns(context.Context, *ReportRunsRequest) (*empty.Empty, error)
	// ReportFinishedRun notifies CV of the Run CQDaemon has just finalized.
	ReportFinishedRun(context.Context, *ReportFinishedRunRequest) (*empty.Empty, error)
}

// UnimplementedMigrationServer can be embedded to have forward compatible implementations.
type UnimplementedMigrationServer struct {
}

func (*UnimplementedMigrationServer) ReportRuns(context.Context, *ReportRunsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportRuns not implemented")
}
func (*UnimplementedMigrationServer) ReportFinishedRun(context.Context, *ReportFinishedRunRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportFinishedRun not implemented")
}

func RegisterMigrationServer(s prpc.Registrar, srv MigrationServer) {
	s.RegisterService(&_Migration_serviceDesc, srv)
}

func _Migration_ReportRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).ReportRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/ReportRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).ReportRuns(ctx, req.(*ReportRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_ReportFinishedRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportFinishedRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).ReportFinishedRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/ReportFinishedRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).ReportFinishedRun(ctx, req.(*ReportFinishedRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Migration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "migration.Migration",
	HandlerType: (*MigrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportRuns",
			Handler:    _Migration_ReportRuns_Handler,
		},
		{
			MethodName: "ReportFinishedRun",
			Handler:    _Migration_ReportFinishedRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cv/api/migration/migration.proto",
}
