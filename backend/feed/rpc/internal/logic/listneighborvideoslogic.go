package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNeighborVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNeighborVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNeighborVideosLogic {
	return &ListNeighborVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNeighborVideosLogic) ListNeighborVideos(in *feed.NeighborsReq) (*feed.NeighborsResp, error) {
	// todo: add your logic here and delete this line

	return &feed.NeighborsResp{}, nil
}
