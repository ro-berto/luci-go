// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package snooperpb

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

// SelfReportClient is the client API for SelfReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SelfReportClient interface {
	// Interface to report cipd packages admitted on the local machine. This
	// should be used by Admission plugin only.
	ReportCipd(ctx context.Context, in *ReportCipdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Interface to report git repos checked out  on the local machine. This
	// should be used by git plugin only.
	ReportGit(ctx context.Context, in *ReportGitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Reports running task's stage. A task is typically a collection of
	// workflows/statements. Some of these statements can be grouped together
	// to define a stage, e.g. when a task is downloading sources/deps, it is
	// known as `FETCH` stage. For details read go/snoopy-design.
	ReportTaskStage(ctx context.Context, in *ReportTaskStageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Reports digest of produced artifact from a task.
	ReportArtifactDigest(ctx context.Context, in *ReportArtifactDigestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type selfReportClient struct {
	cc grpc.ClientConnInterface
}

func NewSelfReportClient(cc grpc.ClientConnInterface) SelfReportClient {
	return &selfReportClient{cc}
}

func (c *selfReportClient) ReportCipd(ctx context.Context, in *ReportCipdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/provenance.snooperpb.SelfReport/ReportCipd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfReportClient) ReportGit(ctx context.Context, in *ReportGitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/provenance.snooperpb.SelfReport/ReportGit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfReportClient) ReportTaskStage(ctx context.Context, in *ReportTaskStageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/provenance.snooperpb.SelfReport/ReportTaskStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfReportClient) ReportArtifactDigest(ctx context.Context, in *ReportArtifactDigestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/provenance.snooperpb.SelfReport/ReportArtifactDigest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SelfReportServer is the server API for SelfReport service.
// All implementations must embed UnimplementedSelfReportServer
// for forward compatibility
type SelfReportServer interface {
	// Interface to report cipd packages admitted on the local machine. This
	// should be used by Admission plugin only.
	ReportCipd(context.Context, *ReportCipdRequest) (*emptypb.Empty, error)
	// Interface to report git repos checked out  on the local machine. This
	// should be used by git plugin only.
	ReportGit(context.Context, *ReportGitRequest) (*emptypb.Empty, error)
	// Reports running task's stage. A task is typically a collection of
	// workflows/statements. Some of these statements can be grouped together
	// to define a stage, e.g. when a task is downloading sources/deps, it is
	// known as `FETCH` stage. For details read go/snoopy-design.
	ReportTaskStage(context.Context, *ReportTaskStageRequest) (*emptypb.Empty, error)
	// Reports digest of produced artifact from a task.
	ReportArtifactDigest(context.Context, *ReportArtifactDigestRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSelfReportServer()
}

// UnimplementedSelfReportServer must be embedded to have forward compatible implementations.
type UnimplementedSelfReportServer struct {
}

func (UnimplementedSelfReportServer) ReportCipd(context.Context, *ReportCipdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportCipd not implemented")
}
func (UnimplementedSelfReportServer) ReportGit(context.Context, *ReportGitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportGit not implemented")
}
func (UnimplementedSelfReportServer) ReportTaskStage(context.Context, *ReportTaskStageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTaskStage not implemented")
}
func (UnimplementedSelfReportServer) ReportArtifactDigest(context.Context, *ReportArtifactDigestRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportArtifactDigest not implemented")
}
func (UnimplementedSelfReportServer) mustEmbedUnimplementedSelfReportServer() {}

// UnsafeSelfReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SelfReportServer will
// result in compilation errors.
type UnsafeSelfReportServer interface {
	mustEmbedUnimplementedSelfReportServer()
}

func RegisterSelfReportServer(s grpc.ServiceRegistrar, srv SelfReportServer) {
	s.RegisterService(&SelfReport_ServiceDesc, srv)
}

func _SelfReport_ReportCipd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportCipdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfReportServer).ReportCipd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.snooperpb.SelfReport/ReportCipd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfReportServer).ReportCipd(ctx, req.(*ReportCipdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfReport_ReportGit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportGitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfReportServer).ReportGit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.snooperpb.SelfReport/ReportGit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfReportServer).ReportGit(ctx, req.(*ReportGitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfReport_ReportTaskStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportTaskStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfReportServer).ReportTaskStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.snooperpb.SelfReport/ReportTaskStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfReportServer).ReportTaskStage(ctx, req.(*ReportTaskStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfReport_ReportArtifactDigest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportArtifactDigestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfReportServer).ReportArtifactDigest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.snooperpb.SelfReport/ReportArtifactDigest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfReportServer).ReportArtifactDigest(ctx, req.(*ReportArtifactDigestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SelfReport_ServiceDesc is the grpc.ServiceDesc for SelfReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SelfReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.snooperpb.SelfReport",
	HandlerType: (*SelfReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportCipd",
			Handler:    _SelfReport_ReportCipd_Handler,
		},
		{
			MethodName: "ReportGit",
			Handler:    _SelfReport_ReportGit_Handler,
		},
		{
			MethodName: "ReportTaskStage",
			Handler:    _SelfReport_ReportTaskStage_Handler,
		},
		{
			MethodName: "ReportArtifactDigest",
			Handler:    _SelfReport_ReportArtifactDigest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/provenance/api/snooperpb/v1/report.proto",
}
