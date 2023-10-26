package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosByRecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosByRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosByRecommendLogic {
	return &ListVideosByRecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideosByRecommendLogic) ListVideosByRecommend(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.ListFeedResponse{}, nil
}
