package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTotalFavoritedCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTotalFavoritedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTotalFavoritedCountLogic {
	return &GetUserTotalFavoritedCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserTotalFavoritedCountLogic) GetUserTotalFavoritedCount(in *favorite.QueryId) (*favorite.QueryCount, error) {
	// todo: add your logic here and delete this line

	return &favorite.QueryCount{}, nil
}
