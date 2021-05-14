// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.0
// source: go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto

package minter

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

// TokenMinterClient is the client API for TokenMinter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenMinterClient interface {
	// MintMachineToken generates a new token for an authenticated machine.
	//
	// It checks that provided certificate was signed by some trusted CA, and it
	// is still valid (non-expired and hasn't been revoked). It then checks that
	// the request was signed by the corresponding private key. Finally it checks
	// that the caller is authorized to generate requested kind of token.
	//
	// If everything checks out, it generates and returns a new machine token.
	//
	// On fatal error it returns detailed error response via same
	// MintMachineTokenResponse. On transient errors it returns generic
	// grpc.Internal error.
	MintMachineToken(ctx context.Context, in *MintMachineTokenRequest, opts ...grpc.CallOption) (*MintMachineTokenResponse, error)
	// MintDelegationToken generates a new bearer delegation token.
	//
	// Such token can be sent in 'X-Delegation-Token-V1' header (alongside regular
	// credentials like OAuth2 access token) to convey that the caller should be
	// authentication as 'delegated_identity' specified in the token.
	//
	// The delegation tokens are subject to multiple restrictions (embedded in
	// the token):
	//   * They have expiration time.
	//   * They are usable only if presented with a credential of someone from
	//     the 'audience' list.
	//   * They are usable only on services specified in the 'services' list.
	//
	// The token server must be configured in advance with all expected
	// combinations of (caller identity, delegated identity, audience, service)
	// tuples. See DelegationRule in config.proto.
	MintDelegationToken(ctx context.Context, in *MintDelegationTokenRequest, opts ...grpc.CallOption) (*MintDelegationTokenResponse, error)
	// MintOAuthTokenGrant generates a new grant for getting an OAuth2 token.
	//
	// This is a special (opaque for clients) token that asserts that the caller
	// at the time of the call was allowed to act as a particular service account
	// to perform a task authorized by an end-user.
	//
	// The returned grant can be used later (when the end-user is no longer
	// present) to get a real OAuth2 access token via MintOAuthTokenViaGrant call.
	//
	// This pair of RPCs is used to "delay" generation of service account OAuth
	// token until some later time, when it is actually needed. This is used by
	// Swarming:
	//   1. When the task is posted, Swarming calls MintOAuthTokenGrant to verify
	//      that the end-user is allowed to act as the requested service account
	//      on Swarming. On success, Swarming stores the grant in the task
	//      metadata.
	//   2. At a later time, when the task is executing and it needs an access
	//      token, Swarming calls MintOAuthTokenViaGrant to convert the grant into
	//      a real OAuth2 token.
	//
	// The returned grant can be used multiple times (as long as its validity
	// duration and the token server policy allows).
	//
	// The token server must be configured in advance with all expected
	// combinations of (caller identity, service account name, end users) tuples.
	// See ServiceAccountRule in config.proto.
	//
	// MintOAuthTokenGrant will check that the requested usage is allowed by the
	// rules. Later, MintOAuthTokenViaGrant will recheck this too.
	//
	// Note: in the process of being replaced with MintServiceAccountToken.
	MintOAuthTokenGrant(ctx context.Context, in *MintOAuthTokenGrantRequest, opts ...grpc.CallOption) (*MintOAuthTokenGrantResponse, error)
	// MintOAuthTokenViaGrant converts an OAuth2 token grant into an access token.
	//
	// The grant must be previously generated by MintOAuthTokenGrant function, see
	// its docs for more details.
	//
	// Note: in the process of being replaced with MintServiceAccountToken.
	MintOAuthTokenViaGrant(ctx context.Context, in *MintOAuthTokenViaGrantRequest, opts ...grpc.CallOption) (*MintOAuthTokenViaGrantResponse, error)
	// MintProjectToken mints an OAuth2 access token that represents an identity
	// associated with a LUCI project.
	//
	// Project-scoped tokens prevent accidental cross-project identity confusion
	// when LUCI services access project specific resources such as a source code
	// repository.
	MintProjectToken(ctx context.Context, in *MintProjectTokenRequest, opts ...grpc.CallOption) (*MintProjectTokenResponse, error)
	// MintServiceAccountToken mints an OAuth2 access token or OpenID ID token
	// that belongs to some service account using LUCI Realms for authorization.
	//
	// As an input it takes a service account email and a name of a LUCI Realm the
	// caller is operating in. To authorize the call the token server checks the
	// following conditions:
	//   1. The caller has luci.serviceAccounts.mintToken permission in the
	//      realm, allowing them to "impersonate" all service accounts belonging
	//      to this realm.
	//   2. The service account has luci.serviceAccounts.existInRealm permission
	//      in the realm. This makes the account "belong" to the realm.
	//   3. Realm's LUCI project has the service account associated with it in
	//      the project_owned_accounts.cfg global config file. This makes sure
	//      different LUCI projects can't just arbitrary use each others accounts
	//      by adding them to their respective realms.cfg. See also comments for
	//      ServiceAccountsProjectMapping in api/admin/v1/config.proto.
	MintServiceAccountToken(ctx context.Context, in *MintServiceAccountTokenRequest, opts ...grpc.CallOption) (*MintServiceAccountTokenResponse, error)
}

type tokenMinterClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenMinterClient(cc grpc.ClientConnInterface) TokenMinterClient {
	return &tokenMinterClient{cc}
}

func (c *tokenMinterClient) MintMachineToken(ctx context.Context, in *MintMachineTokenRequest, opts ...grpc.CallOption) (*MintMachineTokenResponse, error) {
	out := new(MintMachineTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintMachineToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMinterClient) MintDelegationToken(ctx context.Context, in *MintDelegationTokenRequest, opts ...grpc.CallOption) (*MintDelegationTokenResponse, error) {
	out := new(MintDelegationTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintDelegationToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMinterClient) MintOAuthTokenGrant(ctx context.Context, in *MintOAuthTokenGrantRequest, opts ...grpc.CallOption) (*MintOAuthTokenGrantResponse, error) {
	out := new(MintOAuthTokenGrantResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintOAuthTokenGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMinterClient) MintOAuthTokenViaGrant(ctx context.Context, in *MintOAuthTokenViaGrantRequest, opts ...grpc.CallOption) (*MintOAuthTokenViaGrantResponse, error) {
	out := new(MintOAuthTokenViaGrantResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintOAuthTokenViaGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMinterClient) MintProjectToken(ctx context.Context, in *MintProjectTokenRequest, opts ...grpc.CallOption) (*MintProjectTokenResponse, error) {
	out := new(MintProjectTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintProjectToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMinterClient) MintServiceAccountToken(ctx context.Context, in *MintServiceAccountTokenRequest, opts ...grpc.CallOption) (*MintServiceAccountTokenResponse, error) {
	out := new(MintServiceAccountTokenResponse)
	err := c.cc.Invoke(ctx, "/tokenserver.minter.TokenMinter/MintServiceAccountToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenMinterServer is the server API for TokenMinter service.
// All implementations must embed UnimplementedTokenMinterServer
// for forward compatibility
type TokenMinterServer interface {
	// MintMachineToken generates a new token for an authenticated machine.
	//
	// It checks that provided certificate was signed by some trusted CA, and it
	// is still valid (non-expired and hasn't been revoked). It then checks that
	// the request was signed by the corresponding private key. Finally it checks
	// that the caller is authorized to generate requested kind of token.
	//
	// If everything checks out, it generates and returns a new machine token.
	//
	// On fatal error it returns detailed error response via same
	// MintMachineTokenResponse. On transient errors it returns generic
	// grpc.Internal error.
	MintMachineToken(context.Context, *MintMachineTokenRequest) (*MintMachineTokenResponse, error)
	// MintDelegationToken generates a new bearer delegation token.
	//
	// Such token can be sent in 'X-Delegation-Token-V1' header (alongside regular
	// credentials like OAuth2 access token) to convey that the caller should be
	// authentication as 'delegated_identity' specified in the token.
	//
	// The delegation tokens are subject to multiple restrictions (embedded in
	// the token):
	//   * They have expiration time.
	//   * They are usable only if presented with a credential of someone from
	//     the 'audience' list.
	//   * They are usable only on services specified in the 'services' list.
	//
	// The token server must be configured in advance with all expected
	// combinations of (caller identity, delegated identity, audience, service)
	// tuples. See DelegationRule in config.proto.
	MintDelegationToken(context.Context, *MintDelegationTokenRequest) (*MintDelegationTokenResponse, error)
	// MintOAuthTokenGrant generates a new grant for getting an OAuth2 token.
	//
	// This is a special (opaque for clients) token that asserts that the caller
	// at the time of the call was allowed to act as a particular service account
	// to perform a task authorized by an end-user.
	//
	// The returned grant can be used later (when the end-user is no longer
	// present) to get a real OAuth2 access token via MintOAuthTokenViaGrant call.
	//
	// This pair of RPCs is used to "delay" generation of service account OAuth
	// token until some later time, when it is actually needed. This is used by
	// Swarming:
	//   1. When the task is posted, Swarming calls MintOAuthTokenGrant to verify
	//      that the end-user is allowed to act as the requested service account
	//      on Swarming. On success, Swarming stores the grant in the task
	//      metadata.
	//   2. At a later time, when the task is executing and it needs an access
	//      token, Swarming calls MintOAuthTokenViaGrant to convert the grant into
	//      a real OAuth2 token.
	//
	// The returned grant can be used multiple times (as long as its validity
	// duration and the token server policy allows).
	//
	// The token server must be configured in advance with all expected
	// combinations of (caller identity, service account name, end users) tuples.
	// See ServiceAccountRule in config.proto.
	//
	// MintOAuthTokenGrant will check that the requested usage is allowed by the
	// rules. Later, MintOAuthTokenViaGrant will recheck this too.
	//
	// Note: in the process of being replaced with MintServiceAccountToken.
	MintOAuthTokenGrant(context.Context, *MintOAuthTokenGrantRequest) (*MintOAuthTokenGrantResponse, error)
	// MintOAuthTokenViaGrant converts an OAuth2 token grant into an access token.
	//
	// The grant must be previously generated by MintOAuthTokenGrant function, see
	// its docs for more details.
	//
	// Note: in the process of being replaced with MintServiceAccountToken.
	MintOAuthTokenViaGrant(context.Context, *MintOAuthTokenViaGrantRequest) (*MintOAuthTokenViaGrantResponse, error)
	// MintProjectToken mints an OAuth2 access token that represents an identity
	// associated with a LUCI project.
	//
	// Project-scoped tokens prevent accidental cross-project identity confusion
	// when LUCI services access project specific resources such as a source code
	// repository.
	MintProjectToken(context.Context, *MintProjectTokenRequest) (*MintProjectTokenResponse, error)
	// MintServiceAccountToken mints an OAuth2 access token or OpenID ID token
	// that belongs to some service account using LUCI Realms for authorization.
	//
	// As an input it takes a service account email and a name of a LUCI Realm the
	// caller is operating in. To authorize the call the token server checks the
	// following conditions:
	//   1. The caller has luci.serviceAccounts.mintToken permission in the
	//      realm, allowing them to "impersonate" all service accounts belonging
	//      to this realm.
	//   2. The service account has luci.serviceAccounts.existInRealm permission
	//      in the realm. This makes the account "belong" to the realm.
	//   3. Realm's LUCI project has the service account associated with it in
	//      the project_owned_accounts.cfg global config file. This makes sure
	//      different LUCI projects can't just arbitrary use each others accounts
	//      by adding them to their respective realms.cfg. See also comments for
	//      ServiceAccountsProjectMapping in api/admin/v1/config.proto.
	MintServiceAccountToken(context.Context, *MintServiceAccountTokenRequest) (*MintServiceAccountTokenResponse, error)
	mustEmbedUnimplementedTokenMinterServer()
}

// UnimplementedTokenMinterServer must be embedded to have forward compatible implementations.
type UnimplementedTokenMinterServer struct {
}

func (UnimplementedTokenMinterServer) MintMachineToken(context.Context, *MintMachineTokenRequest) (*MintMachineTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintMachineToken not implemented")
}
func (UnimplementedTokenMinterServer) MintDelegationToken(context.Context, *MintDelegationTokenRequest) (*MintDelegationTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintDelegationToken not implemented")
}
func (UnimplementedTokenMinterServer) MintOAuthTokenGrant(context.Context, *MintOAuthTokenGrantRequest) (*MintOAuthTokenGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintOAuthTokenGrant not implemented")
}
func (UnimplementedTokenMinterServer) MintOAuthTokenViaGrant(context.Context, *MintOAuthTokenViaGrantRequest) (*MintOAuthTokenViaGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintOAuthTokenViaGrant not implemented")
}
func (UnimplementedTokenMinterServer) MintProjectToken(context.Context, *MintProjectTokenRequest) (*MintProjectTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintProjectToken not implemented")
}
func (UnimplementedTokenMinterServer) MintServiceAccountToken(context.Context, *MintServiceAccountTokenRequest) (*MintServiceAccountTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintServiceAccountToken not implemented")
}
func (UnimplementedTokenMinterServer) mustEmbedUnimplementedTokenMinterServer() {}

// UnsafeTokenMinterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenMinterServer will
// result in compilation errors.
type UnsafeTokenMinterServer interface {
	mustEmbedUnimplementedTokenMinterServer()
}

func RegisterTokenMinterServer(s grpc.ServiceRegistrar, srv TokenMinterServer) {
	s.RegisterService(&TokenMinter_ServiceDesc, srv)
}

func _TokenMinter_MintMachineToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintMachineTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintMachineToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintMachineToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintMachineToken(ctx, req.(*MintMachineTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMinter_MintDelegationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintDelegationTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintDelegationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintDelegationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintDelegationToken(ctx, req.(*MintDelegationTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMinter_MintOAuthTokenGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintOAuthTokenGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintOAuthTokenGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintOAuthTokenGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintOAuthTokenGrant(ctx, req.(*MintOAuthTokenGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMinter_MintOAuthTokenViaGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintOAuthTokenViaGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintOAuthTokenViaGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintOAuthTokenViaGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintOAuthTokenViaGrant(ctx, req.(*MintOAuthTokenViaGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMinter_MintProjectToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintProjectTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintProjectToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintProjectToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintProjectToken(ctx, req.(*MintProjectTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMinter_MintServiceAccountToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MintServiceAccountTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMinterServer).MintServiceAccountToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.minter.TokenMinter/MintServiceAccountToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMinterServer).MintServiceAccountToken(ctx, req.(*MintServiceAccountTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenMinter_ServiceDesc is the grpc.ServiceDesc for TokenMinter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenMinter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.minter.TokenMinter",
	HandlerType: (*TokenMinterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintMachineToken",
			Handler:    _TokenMinter_MintMachineToken_Handler,
		},
		{
			MethodName: "MintDelegationToken",
			Handler:    _TokenMinter_MintDelegationToken_Handler,
		},
		{
			MethodName: "MintOAuthTokenGrant",
			Handler:    _TokenMinter_MintOAuthTokenGrant_Handler,
		},
		{
			MethodName: "MintOAuthTokenViaGrant",
			Handler:    _TokenMinter_MintOAuthTokenViaGrant_Handler,
		},
		{
			MethodName: "MintProjectToken",
			Handler:    _TokenMinter_MintProjectToken_Handler,
		},
		{
			MethodName: "MintServiceAccountToken",
			Handler:    _TokenMinter_MintServiceAccountToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/tokenserver/api/minter/v1/token_minter.proto",
}
