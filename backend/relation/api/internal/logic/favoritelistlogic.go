package logic

import (
	"context"
	"fmt"
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
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.RelationRpc.FavoriteList(l.ctx, &relation.FavoriteListReq{
		Uid: req.Uid,
	})

	if err != nil {
		fmt.Println(222222)
		return nil, err
	}
	resList := make([]types.UserInfo, 0)
	for _, item := range list.List {
		resList = append(resList, types.UserInfo{
			Id:     item.Id,
			Name:   item.Name,
			Gender: item.Gender,
			Mobile: item.Mobile,
			Avatar: item.Avatar,
			Dec:    item.Dec,
		})
	}
	return &types.FavoriteListResp{List: resList}, nil
}
