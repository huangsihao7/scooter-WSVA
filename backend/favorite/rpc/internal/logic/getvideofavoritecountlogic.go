package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFavoriteCountLogic {
	return &GetVideoFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoFavoriteCountLogic) GetVideoFavoriteCount(in *favorite.QueryId) (*favorite.QueryCount, error) {
	// todo: add your logic here and delete this line
	//获取video的点赞数
	//res, err := l.svcCtx.Model.GetVideoCount(l.ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	//return &favorite.QueryCount{
	//	Count: int64(len(res)),
	//}, nil
	return nil, nil
}
