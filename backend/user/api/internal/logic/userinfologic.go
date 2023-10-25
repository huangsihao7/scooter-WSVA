package logic

import (
	"context"
	"encoding/json"
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

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
		Dec:    res.Dec,
		Avatar: res.Avatar,
	}, nil
}
