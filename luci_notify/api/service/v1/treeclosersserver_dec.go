// Code generated by svcdec; DO NOT EDIT.

package lucinotifypb

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedTreeClosers struct {
	// Service is the service to decorate.
	Service TreeClosersServer
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

func (s *DecoratedTreeClosers) QueryTreeClosers(ctx context.Context, req *QueryTreeClosersRequest) (rsp *QueryTreeClosersResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryTreeClosers", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryTreeClosers(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryTreeClosers", rsp, err)
	}
	return
}

func (s *DecoratedTreeClosers) CheckTreeCloser(ctx context.Context, req *CheckTreeCloserRequest) (rsp *CheckTreeCloserResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CheckTreeCloser", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CheckTreeCloser(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CheckTreeCloser", rsp, err)
	}
	return
}
