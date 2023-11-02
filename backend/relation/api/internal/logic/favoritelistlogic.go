package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	list, err := l.svcCtx.RelationRpc.FavoriteList(l.ctx, &relation.FavoriteListReq{
		Uid:     uid,
		ActUser: req.Uid,
	})

	if err != nil {
		return nil, err
	}
	resList := make([]types.UserInfo, 0)
	for _, item := range list.List {
		resList = append(resList, types.UserInfo{
			Id:              item.Id,
			Name:            item.Name,
			Gender:          item.Gender,
			Avatar:          item.Avatar,
			Dec:             item.Dec,
			BackgroundImage: item.BackgroundImage,
			FollowCount:     item.FollowCount,
			FollowerCount:   item.FollowerCount,
			TotalFavorited:  item.TotalFavorited,
			WorkCount:       item.WorkCount,
			FavoriteCount:   item.FavoriteCount,
			IsFollow:        item.IsFollow,
			CoverUrl:        item.CoverUrl,
			VideoId:         item.VideoId,
			FriendCount:     item.FriendCount,
		})
	}
	return &types.FavoriteListResp{
		StatusCode: int(list.StatusCode),
		StatusMsg:  list.StatusMsg,
		List:       resList}, nil
}
