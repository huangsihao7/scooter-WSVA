package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NullLogic {
	return &NullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NullLogic) Null(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
