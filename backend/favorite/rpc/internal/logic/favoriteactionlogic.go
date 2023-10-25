package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/favorite/model"
	"google.golang.org/grpc/status"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId
	actionType := in.ActionType
	fmt.Println("actionType:", actionType)

	// Check if feed exists

	//判断是否重复点赞

	//将点赞信息添加到数据库
	switch actionType {
	case 1, 2:
		newFavorite := model.Favorites{
			Uid: userId,
			Vid: videoId,
		}
		_, err := l.svcCtx.Model.Insert(l.ctx, &newFavorite)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &favorite.FavoriteActionResponse{
			StatusCode: 200,
			StatusMsg:  "favor success",
		}, nil
	}

	return &favorite.FavoriteActionResponse{
		StatusCode: 500,
		StatusMsg:  "invalid actionType",
	}, nil
}
