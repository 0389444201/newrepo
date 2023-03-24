package logic

import (
	"context"

	"user/user/internal/svc"
	"user/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountGetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountGetInfoLogic {
	return &AccountGetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountGetInfoLogic) AccountGetInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
