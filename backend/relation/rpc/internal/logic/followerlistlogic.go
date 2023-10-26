package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/relation/model"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *relation.FollowerListReq) (*relation.FollowerListResp, error) {
	follower, err := l.svcCtx.RelationModel.FindFollower(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return &relation.FollowerListResp{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, err
		} else {
			return &relation.FollowerListResp{
				StatusCode: constants.UnableToGetFollowerListErrorCode,
				StatusMsg:  constants.UnableToGetFollowerListError,
			}, err
		}
	}

	List := make([]*relation.UserInfo, 0)
	for _, item := range follower {
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, item.FollowerId)
		if err != nil {
			return &relation.FollowerListResp{
				StatusCode: constants.UnableToGetFollowerListErrorCode,
				StatusMsg:  constants.UnableToGetFollowerListError,
			}, err
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
	return &relation.FollowerListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		List:       List,
	}, nil
}
