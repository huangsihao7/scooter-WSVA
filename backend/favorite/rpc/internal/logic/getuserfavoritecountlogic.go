package logic

import (
	"context"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteCountLogic {
	return &GetUserFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavoriteCountLogic) GetUserFavoriteCount(in *favorite.QueryId) (*favorite.QueryCount, error) {
	// todo: add your logic here and delete this line
	// 用户给多少个视频点赞了
	//res, err := l.svcCtx.Model.GetFavoriteCount(l.ctx, in.Id)
	//if err != nil {
	//	return nil, err
	//}
	//return &favorite.QueryCount{
	//	Count: int64(len(res)),
	//}, nil
	return nil, nil
}
