// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cvpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RunsClient is the client API for Runs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunsClient interface {
	// GetRun returns Run details.
	GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*Run, error)
}

type runsClient struct {
	cc grpc.ClientConnInterface
}

func NewRunsClient(cc grpc.ClientConnInterface) RunsClient {
	return &runsClient{cc}
}

func (c *runsClient) GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*Run, error) {
	out := new(Run)
	err := c.cc.Invoke(ctx, "/cv.v0.Runs/GetRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunsServer is the server API for Runs service.
// All implementations must embed UnimplementedRunsServer
// for forward compatibility
type RunsServer interface {
	// GetRun returns Run details.
	GetRun(context.Context, *GetRunRequest) (*Run, error)
	mustEmbedUnimplementedRunsServer()
}

// UnimplementedRunsServer must be embedded to have forward compatible implementations.
type UnimplementedRunsServer struct {
}

func (UnimplementedRunsServer) GetRun(context.Context, *GetRunRequest) (*Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRun not implemented")
}
func (UnimplementedRunsServer) mustEmbedUnimplementedRunsServer() {}

// UnsafeRunsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunsServer will
// result in compilation errors.
type UnsafeRunsServer interface {
	mustEmbedUnimplementedRunsServer()
}

func RegisterRunsServer(s grpc.ServiceRegistrar, srv RunsServer) {
	s.RegisterService(&Runs_ServiceDesc, srv)
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
		FullMethod: "/cv.v0.Runs/GetRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunsServer).GetRun(ctx, req.(*GetRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Runs_ServiceDesc is the grpc.ServiceDesc for Runs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cv.v0.Runs",
	HandlerType: (*RunsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRun",
			Handler:    _Runs_GetRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cv/api/v0/service_runs.proto",
}