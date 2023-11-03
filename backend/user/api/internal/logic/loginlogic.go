package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/jwtx"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"time"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, _ := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.UserId)

	return &types.LoginResponse{
		StatusCode:      int(res.StatusCode),
		StatusMsg:       res.StatusMsg,
		Avatar:          res.Avatar,
		AccessToken:     accessToken,
		UserID:          res.UserId,
		Name:            res.Name,
		Gender:          uint32(res.Gender),
		Signature:       res.Signature,
		BackgroundImage: res.BackgroundImage,
	}, nil
}
