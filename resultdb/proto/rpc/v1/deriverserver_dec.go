// Code generated by svcdec; DO NOT EDIT.

package resultpb

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedDeriver struct {
	// Service is the service to decorate.
	Service DeriverServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedDeriver) DeriveChromiumInvocation(ctx context.Context, req *DeriveChromiumInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeriveChromiumInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeriveChromiumInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeriveChromiumInvocation", rsp, err)
	}
	return
}
