package logic

import (
	"context"

	"user/user/internal/svc"
	"user/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountDeleteLogic {
	return &AccountDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountDeleteLogic) AccountDelete(req *types.DeleteRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
