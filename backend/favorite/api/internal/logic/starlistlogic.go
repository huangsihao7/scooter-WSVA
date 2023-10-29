package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarListLogic {
	return &StarListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarListLogic) StarList(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
