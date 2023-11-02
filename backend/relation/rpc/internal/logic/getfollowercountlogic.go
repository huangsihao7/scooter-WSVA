package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerCountLogic {
	return &GetFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerCountLogic) GetFollowerCount(in *relation.FollowerCountReq) (*relation.FollowerCountResp, error) {
	count, err := l.svcCtx.RelationModel.GetFollowerCount(l.ctx, in.Uid)
	if err != nil {
		l.Logger.Error("获取粉丝数失败")
		return nil, err
	}
	return &relation.FollowerCountResp{
		Count:      count,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
