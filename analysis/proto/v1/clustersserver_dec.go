// Code generated by svcdec; DO NOT EDIT.

package analysispb

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedClusters struct {
	// Service is the service to decorate.
	Service ClustersServer
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

func (s *DecoratedClusters) Cluster(ctx context.Context, req *ClusterRequest) (rsp *ClusterResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "Cluster", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.Cluster(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "Cluster", rsp, err)
	}
	return
}

func (s *DecoratedClusters) BatchGet(ctx context.Context, req *BatchGetClustersRequest) (rsp *BatchGetClustersResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "BatchGet", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.BatchGet(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "BatchGet", rsp, err)
	}
	return
}

func (s *DecoratedClusters) GetReclusteringProgress(ctx context.Context, req *GetReclusteringProgressRequest) (rsp *ReclusteringProgress, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetReclusteringProgress", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetReclusteringProgress(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetReclusteringProgress", rsp, err)
	}
	return
}

func (s *DecoratedClusters) QueryClusterSummaries(ctx context.Context, req *QueryClusterSummariesRequest) (rsp *QueryClusterSummariesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryClusterSummaries", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryClusterSummaries(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryClusterSummaries", rsp, err)
	}
	return
}

func (s *DecoratedClusters) QueryClusterFailures(ctx context.Context, req *QueryClusterFailuresRequest) (rsp *QueryClusterFailuresResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "QueryClusterFailures", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.QueryClusterFailures(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "QueryClusterFailures", rsp, err)
	}
	return
}