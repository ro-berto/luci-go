// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Launches a mapping job that examines and/or fixes datastore entities.
	LaunchJob(ctx context.Context, in *JobConfig, opts ...grpc.CallOption) (*JobID, error)
	// Initiates an abort of a mapping job.
	AbortJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns state of a mapping job.
	GetJobState(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*JobState, error)
	// Fixes (right inside the handler) tags marked by the given mapper job.
	FixMarkedTags(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*TagFixReport, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) LaunchJob(ctx context.Context, in *JobConfig, opts ...grpc.CallOption) (*JobID, error) {
	out := new(JobID)
	err := c.cc.Invoke(ctx, "/cipd.Admin/LaunchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AbortJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cipd.Admin/AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetJobState(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*JobState, error) {
	out := new(JobState)
	err := c.cc.Invoke(ctx, "/cipd.Admin/GetJobState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) FixMarkedTags(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*TagFixReport, error) {
	out := new(TagFixReport)
	err := c.cc.Invoke(ctx, "/cipd.Admin/FixMarkedTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Launches a mapping job that examines and/or fixes datastore entities.
	LaunchJob(context.Context, *JobConfig) (*JobID, error)
	// Initiates an abort of a mapping job.
	AbortJob(context.Context, *JobID) (*empty.Empty, error)
	// Returns state of a mapping job.
	GetJobState(context.Context, *JobID) (*JobState, error)
	// Fixes (right inside the handler) tags marked by the given mapper job.
	FixMarkedTags(context.Context, *JobID) (*TagFixReport, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) LaunchJob(context.Context, *JobConfig) (*JobID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchJob not implemented")
}
func (UnimplementedAdminServer) AbortJob(context.Context, *JobID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortJob not implemented")
}
func (UnimplementedAdminServer) GetJobState(context.Context, *JobID) (*JobState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobState not implemented")
}
func (UnimplementedAdminServer) FixMarkedTags(context.Context, *JobID) (*TagFixReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FixMarkedTags not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_LaunchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).LaunchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/LaunchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).LaunchJob(ctx, req.(*JobConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AbortJob(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetJobState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/GetJobState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetJobState(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_FixMarkedTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).FixMarkedTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/FixMarkedTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).FixMarkedTags(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cipd.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchJob",
			Handler:    _Admin_LaunchJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _Admin_AbortJob_Handler,
		},
		{
			MethodName: "GetJobState",
			Handler:    _Admin_GetJobState_Handler,
		},
		{
			MethodName: "FixMarkedTags",
			Handler:    _Admin_FixMarkedTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cipd/api/admin/v1/admin.proto",
}
