package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
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
	list, err := l.svcCtx.RelationRpc.FavoriteList(l.ctx, &relation.FavoriteListReq{
		Uid: req.Uid,
	})

	if err != nil {
		return &types.FavoriteListResp{
			StatusCode: int(list.StatusCode),
			StatusMsg:  list.StatusMsg,
		}, err
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
	return &types.FavoriteListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       resList}, nil
}
