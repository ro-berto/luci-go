// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.0
// source: go.chromium.org/luci/tokenserver/api/admin/v1/admin.proto

package admin

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
	// ImportCAConfigs makes the server read 'tokenserver.cfg'.
	ImportCAConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// ImportDelegationConfigs makes the server read 'delegation.cfg'.
	ImportDelegationConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// ImportServiceAccountsConfigs makes the server read 'service_accounts.cfg'.
	ImportServiceAccountsConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// ImportProjectIdentityConfigs makes the server read 'projects.cfg'.
	ImportProjectIdentityConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// ImportProjectOwnedAccountsConfigs makes the server read 'project_owned_accounts.cfg'.
	ImportProjectOwnedAccountsConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// InspectMachineToken decodes a machine token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time and revocation status.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. it checks for revocation status even if token is already
	// expired).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectMachineTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectMachineToken(ctx context.Context, in *InspectMachineTokenRequest, opts ...grpc.CallOption) (*InspectMachineTokenResponse, error)
	// InspectDelegationToken decodes a delegation token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. attempts to decode the body even if the signing key has been
	// rotated already).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectDelegationTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectDelegationToken(ctx context.Context, in *InspectDelegationTokenRequest, opts ...grpc.CallOption) (*InspectDelegationTokenResponse, error)
	// InspectOAuthTokenGrant decodes OAuth token grant and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. attempts to decode the body even if the signing key has been
	// rotated already).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectOAuthTokenGrantResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectOAuthTokenGrant(ctx context.Context, in *InspectOAuthTokenGrantRequest, opts ...grpc.CallOption) (*InspectOAuthTokenGrantResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) ImportCAConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/ImportCAConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ImportDelegationConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/ImportDelegationConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ImportServiceAccountsConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/ImportServiceAccountsConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ImportProjectIdentityConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/ImportProjectIdentityConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ImportProjectOwnedAccountsConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/ImportProjectOwnedAccountsConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InspectMachineToken(ctx context.Context, in *InspectMachineTokenRequest, opts ...grpc.CallOption) (*InspectMachineTokenResponse, error) {
	out := new(InspectMachineTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/InspectMachineToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InspectDelegationToken(ctx context.Context, in *InspectDelegationTokenRequest, opts ...grpc.CallOption) (*InspectDelegationTokenResponse, error) {
	out := new(InspectDelegationTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/InspectDelegationToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InspectOAuthTokenGrant(ctx context.Context, in *InspectOAuthTokenGrantRequest, opts ...grpc.CallOption) (*InspectOAuthTokenGrantResponse, error) {
	out := new(InspectOAuthTokenGrantResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.admin.Admin/InspectOAuthTokenGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// ImportCAConfigs makes the server read 'tokenserver.cfg'.
	ImportCAConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error)
	// ImportDelegationConfigs makes the server read 'delegation.cfg'.
	ImportDelegationConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error)
	// ImportServiceAccountsConfigs makes the server read 'service_accounts.cfg'.
	ImportServiceAccountsConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error)
	// ImportProjectIdentityConfigs makes the server read 'projects.cfg'.
	ImportProjectIdentityConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error)
	// ImportProjectOwnedAccountsConfigs makes the server read 'project_owned_accounts.cfg'.
	ImportProjectOwnedAccountsConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error)
	// InspectMachineToken decodes a machine token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time and revocation status.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. it checks for revocation status even if token is already
	// expired).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectMachineTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectMachineToken(context.Context, *InspectMachineTokenRequest) (*InspectMachineTokenResponse, error)
	// InspectDelegationToken decodes a delegation token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. attempts to decode the body even if the signing key has been
	// rotated already).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectDelegationTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectDelegationToken(context.Context, *InspectDelegationTokenRequest) (*InspectDelegationTokenResponse, error)
	// InspectOAuthTokenGrant decodes OAuth token grant and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. attempts to decode the body even if the signing key has been
	// rotated already).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectOAuthTokenGrantResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectOAuthTokenGrant(context.Context, *InspectOAuthTokenGrantRequest) (*InspectOAuthTokenGrantResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) ImportCAConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCAConfigs not implemented")
}
func (UnimplementedAdminServer) ImportDelegationConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportDelegationConfigs not implemented")
}
func (UnimplementedAdminServer) ImportServiceAccountsConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportServiceAccountsConfigs not implemented")
}
func (UnimplementedAdminServer) ImportProjectIdentityConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProjectIdentityConfigs not implemented")
}
func (UnimplementedAdminServer) ImportProjectOwnedAccountsConfigs(context.Context, *emptypb.Empty) (*ImportedConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProjectOwnedAccountsConfigs not implemented")
}
func (UnimplementedAdminServer) InspectMachineToken(context.Context, *InspectMachineTokenRequest) (*InspectMachineTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectMachineToken not implemented")
}
func (UnimplementedAdminServer) InspectDelegationToken(context.Context, *InspectDelegationTokenRequest) (*InspectDelegationTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectDelegationToken not implemented")
}
func (UnimplementedAdminServer) InspectOAuthTokenGrant(context.Context, *InspectOAuthTokenGrantRequest) (*InspectOAuthTokenGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InspectOAuthTokenGrant not implemented")
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

func _Admin_ImportCAConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportCAConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportCAConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportCAConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ImportDelegationConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportDelegationConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportDelegationConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportDelegationConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ImportServiceAccountsConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportServiceAccountsConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportServiceAccountsConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportServiceAccountsConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ImportProjectIdentityConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportProjectIdentityConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportProjectIdentityConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportProjectIdentityConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ImportProjectOwnedAccountsConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportProjectOwnedAccountsConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportProjectOwnedAccountsConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportProjectOwnedAccountsConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InspectMachineToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectMachineTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InspectMachineToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/InspectMachineToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InspectMachineToken(ctx, req.(*InspectMachineTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InspectDelegationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectDelegationTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InspectDelegationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/InspectDelegationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InspectDelegationToken(ctx, req.(*InspectDelegationTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InspectOAuthTokenGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectOAuthTokenGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InspectOAuthTokenGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/InspectOAuthTokenGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InspectOAuthTokenGrant(ctx, req.(*InspectOAuthTokenGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportCAConfigs",
			Handler:    _Admin_ImportCAConfigs_Handler,
		},
		{
			MethodName: "ImportDelegationConfigs",
			Handler:    _Admin_ImportDelegationConfigs_Handler,
		},
		{
			MethodName: "ImportServiceAccountsConfigs",
			Handler:    _Admin_ImportServiceAccountsConfigs_Handler,
		},
		{
			MethodName: "ImportProjectIdentityConfigs",
			Handler:    _Admin_ImportProjectIdentityConfigs_Handler,
		},
		{
			MethodName: "ImportProjectOwnedAccountsConfigs",
			Handler:    _Admin_ImportProjectOwnedAccountsConfigs_Handler,
		},
		{
			MethodName: "InspectMachineToken",
			Handler:    _Admin_InspectMachineToken_Handler,
		},
		{
			MethodName: "InspectDelegationToken",
			Handler:    _Admin_InspectDelegationToken_Handler,
		},
		{
			MethodName: "InspectOAuthTokenGrant",
			Handler:    _Admin_InspectOAuthTokenGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/tokenserver/api/admin/v1/admin.proto",
}
