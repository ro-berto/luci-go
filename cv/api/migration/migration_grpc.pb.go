// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: go.chromium.org/luci/cv/api/migration/migration.proto

package migrationpb

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

// MigrationClient is the client API for Migration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MigrationClient interface {
	// ReportVerifiedRun notifies CV of the Run CQDaemon has just finished
	// verifying.
	//
	// The Run may not contain CV's id, but CV can figure out the ID using
	// Run.Attempt.Key.
	//
	// Called by CQDaemon when CV is in charge of run management.
	ReportVerifiedRun(ctx context.Context, in *ReportVerifiedRunRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// FetchRunStatus provides to CQDaemon info about a completed Run in order to
	// send to CQ Status app.
	//
	// The Run may not contain CV's id, but CV can figure out the ID using
	// Run.Attempt.Key.
	//
	// Called by CQDaemon when CV is in charge of run management.
	FetchRunStatus(ctx context.Context, in *FetchRunStatusRequest, opts ...grpc.CallOption) (*FetchRunStatusResponse, error)
	// PostGerritMessage posts a unique per run message to Gerrit.
	//
	// Best effort, since Gerrit doesn't provide for idempotent or conditional
	// (etag like) updates.
	//
	// Use-cases:
	//  * CQDaemon linter posting a warning/error.
	//  * GerritCQAbility verifier posting error on each of Run's CL before
	//    failing entire attempt.
	//
	// Error handling:
	//  * If presumably transient Gerrit error, fails with Internal error (for
	//    simplicity). CQDaemon will have to retry.
	//  * If Gerrit error is 403, 404 or 412 (Precondition error), responds with
	//    corresponding gRPC code.
	PostGerritMessage(ctx context.Context, in *PostGerritMessageRequest, opts ...grpc.CallOption) (*PostGerritMessageResponse, error)
	// FetchActiveRuns returns all currently RUNNING runs in CV for the given
	// project.
	FetchActiveRuns(ctx context.Context, in *FetchActiveRunsRequest, opts ...grpc.CallOption) (*FetchActiveRunsResponse, error)
	// ReportTryjobs notifies CV of the tryjobs applicable to a Run.
	ReportTryjobs(ctx context.Context, in *ReportTryjobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type migrationClient struct {
	cc grpc.ClientConnInterface
}

func NewMigrationClient(cc grpc.ClientConnInterface) MigrationClient {
	return &migrationClient{cc}
}

func (c *migrationClient) ReportVerifiedRun(ctx context.Context, in *ReportVerifiedRunRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/migration.Migration/ReportVerifiedRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) FetchRunStatus(ctx context.Context, in *FetchRunStatusRequest, opts ...grpc.CallOption) (*FetchRunStatusResponse, error) {
	out := new(FetchRunStatusResponse)
	err := c.cc.Invoke(ctx, "/migration.Migration/FetchRunStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) PostGerritMessage(ctx context.Context, in *PostGerritMessageRequest, opts ...grpc.CallOption) (*PostGerritMessageResponse, error) {
	out := new(PostGerritMessageResponse)
	err := c.cc.Invoke(ctx, "/migration.Migration/PostGerritMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) FetchActiveRuns(ctx context.Context, in *FetchActiveRunsRequest, opts ...grpc.CallOption) (*FetchActiveRunsResponse, error) {
	out := new(FetchActiveRunsResponse)
	err := c.cc.Invoke(ctx, "/migration.Migration/FetchActiveRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) ReportTryjobs(ctx context.Context, in *ReportTryjobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/migration.Migration/ReportTryjobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServer is the server API for Migration service.
// All implementations must embed UnimplementedMigrationServer
// for forward compatibility
type MigrationServer interface {
	// ReportVerifiedRun notifies CV of the Run CQDaemon has just finished
	// verifying.
	//
	// The Run may not contain CV's id, but CV can figure out the ID using
	// Run.Attempt.Key.
	//
	// Called by CQDaemon when CV is in charge of run management.
	ReportVerifiedRun(context.Context, *ReportVerifiedRunRequest) (*emptypb.Empty, error)
	// FetchRunStatus provides to CQDaemon info about a completed Run in order to
	// send to CQ Status app.
	//
	// The Run may not contain CV's id, but CV can figure out the ID using
	// Run.Attempt.Key.
	//
	// Called by CQDaemon when CV is in charge of run management.
	FetchRunStatus(context.Context, *FetchRunStatusRequest) (*FetchRunStatusResponse, error)
	// PostGerritMessage posts a unique per run message to Gerrit.
	//
	// Best effort, since Gerrit doesn't provide for idempotent or conditional
	// (etag like) updates.
	//
	// Use-cases:
	//  * CQDaemon linter posting a warning/error.
	//  * GerritCQAbility verifier posting error on each of Run's CL before
	//    failing entire attempt.
	//
	// Error handling:
	//  * If presumably transient Gerrit error, fails with Internal error (for
	//    simplicity). CQDaemon will have to retry.
	//  * If Gerrit error is 403, 404 or 412 (Precondition error), responds with
	//    corresponding gRPC code.
	PostGerritMessage(context.Context, *PostGerritMessageRequest) (*PostGerritMessageResponse, error)
	// FetchActiveRuns returns all currently RUNNING runs in CV for the given
	// project.
	FetchActiveRuns(context.Context, *FetchActiveRunsRequest) (*FetchActiveRunsResponse, error)
	// ReportTryjobs notifies CV of the tryjobs applicable to a Run.
	ReportTryjobs(context.Context, *ReportTryjobsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMigrationServer()
}

// UnimplementedMigrationServer must be embedded to have forward compatible implementations.
type UnimplementedMigrationServer struct {
}

func (UnimplementedMigrationServer) ReportVerifiedRun(context.Context, *ReportVerifiedRunRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportVerifiedRun not implemented")
}
func (UnimplementedMigrationServer) FetchRunStatus(context.Context, *FetchRunStatusRequest) (*FetchRunStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRunStatus not implemented")
}
func (UnimplementedMigrationServer) PostGerritMessage(context.Context, *PostGerritMessageRequest) (*PostGerritMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGerritMessage not implemented")
}
func (UnimplementedMigrationServer) FetchActiveRuns(context.Context, *FetchActiveRunsRequest) (*FetchActiveRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchActiveRuns not implemented")
}
func (UnimplementedMigrationServer) ReportTryjobs(context.Context, *ReportTryjobsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTryjobs not implemented")
}
func (UnimplementedMigrationServer) mustEmbedUnimplementedMigrationServer() {}

// UnsafeMigrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MigrationServer will
// result in compilation errors.
type UnsafeMigrationServer interface {
	mustEmbedUnimplementedMigrationServer()
}

func RegisterMigrationServer(s grpc.ServiceRegistrar, srv MigrationServer) {
	s.RegisterService(&Migration_ServiceDesc, srv)
}

func _Migration_ReportVerifiedRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportVerifiedRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).ReportVerifiedRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/ReportVerifiedRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).ReportVerifiedRun(ctx, req.(*ReportVerifiedRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_FetchRunStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRunStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).FetchRunStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/FetchRunStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).FetchRunStatus(ctx, req.(*FetchRunStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_PostGerritMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostGerritMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).PostGerritMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/PostGerritMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).PostGerritMessage(ctx, req.(*PostGerritMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_FetchActiveRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchActiveRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).FetchActiveRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/FetchActiveRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).FetchActiveRuns(ctx, req.(*FetchActiveRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_ReportTryjobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportTryjobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).ReportTryjobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.Migration/ReportTryjobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).ReportTryjobs(ctx, req.(*ReportTryjobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Migration_ServiceDesc is the grpc.ServiceDesc for Migration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Migration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "migration.Migration",
	HandlerType: (*MigrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportVerifiedRun",
			Handler:    _Migration_ReportVerifiedRun_Handler,
		},
		{
			MethodName: "FetchRunStatus",
			Handler:    _Migration_FetchRunStatus_Handler,
		},
		{
			MethodName: "PostGerritMessage",
			Handler:    _Migration_PostGerritMessage_Handler,
		},
		{
			MethodName: "FetchActiveRuns",
			Handler:    _Migration_FetchActiveRuns_Handler,
		},
		{
			MethodName: "ReportTryjobs",
			Handler:    _Migration_ReportTryjobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cv/api/migration/migration.proto",
}
