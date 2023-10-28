package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/recommend/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/recommend/rpc/recommend"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoLogic) CreateVideo(in *recommend.CreateVideoRequest) (*recommend.CreateVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &recommend.CreateVideoResponse{}, nil
}
