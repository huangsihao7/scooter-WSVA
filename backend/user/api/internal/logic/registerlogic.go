package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/common"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

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
	avatarUrl, err := common.UserUpload(l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey, l.svcCtx.Config.Bucket, req.Avatar)
	backUrl, err := common.UserUpload(l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey, l.svcCtx.Config.Bucket, req.BackgroundImage)
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:            req.Name,
		Gender:          req.Gender,
		Mobile:          req.Mobile,
		Password:        req.Password,
		Avatar:          avatarUrl,
		Dec:             req.Dec,
		BackgroundImage: backUrl,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Id:              res.Id,
		Name:            res.Name,
		Gender:          res.Gender,
		Mobile:          res.Mobile,
		Avatar:          res.Avatar,
		Dec:             res.Dec,
		BackgroundImage: res.BackgroundImage,
	}, nil
}
