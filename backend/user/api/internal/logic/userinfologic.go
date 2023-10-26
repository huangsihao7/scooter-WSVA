package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId:  req.Uid,
		ActorId: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:   int64(res.User.Id),
		Name: res.User.Name,
	}, nil
}
