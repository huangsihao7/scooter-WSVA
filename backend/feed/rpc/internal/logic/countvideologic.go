package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountVideoLogic {
	return &CountVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountVideoLogic) CountVideo(in *feed.CountVideoRequest) (*feed.CountVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &feed.CountVideoResponse{}, nil
}
