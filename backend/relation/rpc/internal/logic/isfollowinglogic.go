package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/huangsihao7/scooter-WSVA/relation/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowingLogic {
	return &IsFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowingLogic) IsFollowing(in *relation.IsFollowingReq) (*relation.IsFollowingResp, error) {
	// todo: add your logic here and delete this line
	flag, err := l.svcCtx.RelationModel.IsFollowing(l.ctx, in.ActorId, in.UserId)
	if err != nil {
		l.Logger.Error("获取是否关注失败")
		return nil, err
	}
	return &relation.IsFollowingResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		Flag:       flag,
	}, nil
}
