package api

import (
	"context"
	"my-auth-service/config"
)

type api struct {
	cfg *config.Configuration
}

type MASAPIServer interface {
}

func NewAPI(cfg *config.Configuration) MASAPIServer {
	return api{cfg: cfg}
}

func (a *api) Readiness(ctx context.Context) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	return resp, nil
}

func (a *api) Liveness(ctx context.Context, req *emptypb.Empty) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	return resp, nil
}
