// Code generated by svcdec; DO NOT EDIT.

package gofindit

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedGoFinditService struct {
	// Service is the service to decorate.
	Service GoFinditServiceServer
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

func (s *DecoratedGoFinditService) GetAnalysis(ctx context.Context, req *GetAnalysisRequest) (rsp *Analysis, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetAnalysis", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetAnalysis(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetAnalysis", rsp, err)
	}
	return
}

func (s *DecoratedGoFinditService) QueryAnalysis(ctx context.Context, req *QueryAnalysisRequest) (rsp *QueryAnalysisResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryAnalysis", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryAnalysis(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryAnalysis", rsp, err)
	}
	return
}

func (s *DecoratedGoFinditService) ListAnalyses(ctx context.Context, req *ListAnalysesRequest) (rsp *ListAnalysesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "ListAnalyses", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListAnalyses(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "ListAnalyses", rsp, err)
	}
	return
}

func (s *DecoratedGoFinditService) TriggerAnalysis(ctx context.Context, req *TriggerAnalysisRequest) (rsp *TriggerAnalysisResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "TriggerAnalysis", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.TriggerAnalysis(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "TriggerAnalysis", rsp, err)
	}
	return
}

func (s *DecoratedGoFinditService) UpdateAnalysis(ctx context.Context, req *UpdateAnalysisRequest) (rsp *Analysis, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateAnalysis", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateAnalysis(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateAnalysis", rsp, err)
	}
	return
}
