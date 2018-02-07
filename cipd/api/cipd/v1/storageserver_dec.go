// Code generated by svcdec; DO NOT EDIT

package api

import (
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
)

type DecoratedStorage struct {
	// Service is the service to decorate.
	Service StorageServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedStorage) GetObjectURL(c context.Context, req *GetObjectURLRequest) (rsp *ObjectURL, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "GetObjectURL", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.GetObjectURL(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "GetObjectURL", rsp, err)
	}
	return
}

func (s *DecoratedStorage) BeginUpload(c context.Context, req *BeginUploadRequest) (rsp *UploadOperation, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "BeginUpload", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.BeginUpload(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "BeginUpload", rsp, err)
	}
	return
}

func (s *DecoratedStorage) FinishUpload(c context.Context, req *FinishUploadRequest) (rsp *UploadOperation, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "FinishUpload", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.FinishUpload(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "FinishUpload", rsp, err)
	}
	return
}
