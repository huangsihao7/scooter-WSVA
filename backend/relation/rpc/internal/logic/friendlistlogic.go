package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/relation/model"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *relation.FriendListReq) (*relation.FriendListResp, error) {
	friend, err := l.svcCtx.RelationModel.FindFriend(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return &relation.FriendListResp{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		} else {
			return &relation.FriendListResp{
				StatusCode: constants.UnableToGetFriendListErrorCode,
				StatusMsg:  constants.UnableToGetFriendListError,
			}, nil
		}
	}

	List := make([]*relation.UserInfo, 0)
	for _, item := range friend {
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, item.FollowingId)
		if err != nil {
			return &relation.FriendListResp{
				StatusCode: constants.UnableToGetFriendListErrorCode,
				StatusMsg:  constants.UnableToGetFriendListError,
			}, nil
		}
		List = append(List, &relation.UserInfo{
			Id:     one.Id,
			Name:   one.Name,
			Gender: one.Gender,
			Mobile: one.Mobile,
			Avatar: one.Avatar,
			Dec:    one.Dec,
		})
	}
	return &relation.FriendListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       List,
	}, nil
}
