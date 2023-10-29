package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	list, err := l.svcCtx.RelationRpc.FollowerList(l.ctx, &relation.FollowerListReq{
		Uid: req.Uid,
	})

	if err != nil {
		return &types.FollowerListResp{
			StatusCode: int(list.StatusCode),
			StatusMsg:  list.StatusMsg,
		}, err
	}
	resList := make([]types.UserInfo, 0)
	for _, item := range list.List {
		resList = append(resList, types.UserInfo{
			Id:              item.Id,
			Name:            item.Name,
			Gender:          item.Gender,
			Mobile:          item.Mobile,
			Avatar:          item.Avatar,
			Dec:             item.Dec,
			BackgroundImage: item.BackgroundImage,
		})
	}
	return &types.FollowerListResp{
		StatusCode: int(list.StatusCode),
		StatusMsg:  list.GetStatusMsg(),
		List:       resList,
	}, nil
}
