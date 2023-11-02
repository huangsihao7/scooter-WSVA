package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:            req.Name,
		Gender:          req.Gender,
		Mobile:          req.Mobile,
		Password:        req.Password,
		Avatar:          req.Avatar,
		Dec:             req.Dec,
		BackgroundImage: req.BackgroundImage,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil
}
