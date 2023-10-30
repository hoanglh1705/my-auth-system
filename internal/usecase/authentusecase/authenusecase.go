package authentusecase

import (
	"context"
	"my-auth-service/internal/usecase/model"
)

type AuthentUsecase interface {
	Login(ctx context.Context, Login model.LoginReq) (model.LoginResp, error)
}
