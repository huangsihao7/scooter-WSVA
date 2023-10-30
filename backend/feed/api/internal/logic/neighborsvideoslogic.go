package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NeighborsVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNeighborsVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NeighborsVideosLogic {
	return &NeighborsVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NeighborsVideosLogic) NeighborsVideos(req *types.NeighborsVideoReq) (resp *types.NeighborsVideoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
