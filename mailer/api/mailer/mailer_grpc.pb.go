// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mailer

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

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerClient interface {
	// SendMail enqueues an email for sending.
	//
	// On OK RPC status code, the email was enqueued and will eventually be sent.
	// The response contains an opaque identifier that can be used to trace this
	// operation in logs.
	//
	// Transient error status codes (INTERNAL, UNKNOWN, etc.) indicate that the
	// email may or may not been enqueued. The caller should retry, passing the
	// exact same `request_id`.
	//
	// Non-retryable errors (per https://google.aip.dev/194) indicate that the
	// email was rejected and retries won't help.
	SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error) {
	out := new(SendMailResponse)
	err := c.cc.Invoke(ctx, "/luci.mailer.v1.Mailer/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
// All implementations must embed UnimplementedMailerServer
// for forward compatibility
type MailerServer interface {
	// SendMail enqueues an email for sending.
	//
	// On OK RPC status code, the email was enqueued and will eventually be sent.
	// The response contains an opaque identifier that can be used to trace this
	// operation in logs.
	//
	// Transient error status codes (INTERNAL, UNKNOWN, etc.) indicate that the
	// email may or may not been enqueued. The caller should retry, passing the
	// exact same `request_id`.
	//
	// Non-retryable errors (per https://google.aip.dev/194) indicate that the
	// email was rejected and retries won't help.
	SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error)
	mustEmbedUnimplementedMailerServer()
}

// UnimplementedMailerServer must be embedded to have forward compatible implementations.
type UnimplementedMailerServer struct {
}

func (UnimplementedMailerServer) SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMailerServer) mustEmbedUnimplementedMailerServer() {}

// UnsafeMailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServer will
// result in compilation errors.
type UnsafeMailerServer interface {
	mustEmbedUnimplementedMailerServer()
}

func RegisterMailerServer(s grpc.ServiceRegistrar, srv MailerServer) {
	s.RegisterService(&Mailer_ServiceDesc, srv)
}

func _Mailer_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.mailer.v1.Mailer/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendMail(ctx, req.(*SendMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailer_ServiceDesc is the grpc.ServiceDesc for Mailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "luci.mailer.v1.Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _Mailer_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/mailer/api/mailer/mailer.proto",
}