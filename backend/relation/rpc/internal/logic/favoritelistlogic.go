package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/relation/model"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *relation.FavoriteListReq) (*relation.FavoriteListResp, error) {
	favorite, err := l.svcCtx.RelationModel.FindFavorite(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return &relation.FavoriteListResp{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, err
		} else {
			return &relation.FavoriteListResp{
				StatusCode: constants.UnableToGetFollowListErrorCode,
				StatusMsg:  constants.UnableToGetFollowListError,
			}, err
		}
	}

	List := make([]*relation.UserInfo, 0)
	for _, item := range favorite {
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, item.FollowingId)
		if err != nil {
			return &relation.FavoriteListResp{
				StatusCode: constants.UnableToGetFollowListErrorCode,
				StatusMsg:  constants.UnableToGetFollowListError,
			}, err
		}
		List = append(List, &relation.UserInfo{
			Id:              one.Id,
			Name:            one.Name,
			Gender:          one.Gender,
			Mobile:          one.Mobile,
			Avatar:          one.Avatar,
			Dec:             one.Dec,
			BackgroundImage: one.BackgroundUrl,
		})
	}
	return &relation.FavoriteListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       List,
	}, nil
}
