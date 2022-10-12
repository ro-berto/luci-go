// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: go.chromium.org/luci/server/redisconn/adminpb/admin.proto

package adminpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Deletes all the keys of all the existing databases.
	//
	// Issues "FLUSHALL" command, optionally with "ASYNC" argument.
	FlushAll(ctx context.Context, in *FlushAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) FlushAll(ctx context.Context, in *FlushAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/server.redisconn.Admin/FlushAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Deletes all the keys of all the existing databases.
	//
	// Issues "FLUSHALL" command, optionally with "ASYNC" argument.
	FlushAll(context.Context, *FlushAllRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) FlushAll(context.Context, *FlushAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushAll not implemented")
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

func _Admin_FlushAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).FlushAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.redisconn.Admin/FlushAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).FlushAll(ctx, req.(*FlushAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.redisconn.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FlushAll",
			Handler:    _Admin_FlushAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/server/redisconn/adminpb/admin.proto",
}
