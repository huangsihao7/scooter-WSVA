package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	list, err := l.svcCtx.RelationRpc.FriendList(l.ctx, &relation.FriendListReq{
		Uid: req.Uid,
	})

	if err != nil {
		return &types.FriendListResp{
			StatusCode: int(list.StatusCode),
			StatusMsg:  list.StatusMsg,
		}, nil
	}
	resList := make([]types.UserInfo, 0)
	for _, item := range list.List {
		resList = append(resList, types.UserInfo{
			Id:      item.Id,
			Name:    item.Name,
			Gender:  item.Gender,
			Mobile:  item.Mobile,
			Avatar:  item.Avatar,
			Dec:     item.Dec,
			BackImg: item.BackImg,
		})
	}
	return &types.FriendListResp{
		StatusCode: int(list.StatusCode),
		StatusMsg:  list.GetStatusMsg(),
		List:       resList,
	}, nil
}
