package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowCountLogic {
	return &GetFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowCountLogic) GetFollowCount(in *relation.FollowCountReq) (*relation.FollowCountResp, error) {
	count, err := l.svcCtx.RelationModel.GetFollowCount(l.ctx, in.Uid)
	if err != nil {
		return &relation.FollowCountResp{
			StatusCode: constants.UnableToGetFollowCountErrorCode,
			StatusMsg:  constants.UnableToGetFollowCountError,
		}, err
	}
	return &relation.FollowCountResp{
		Count:      count,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
