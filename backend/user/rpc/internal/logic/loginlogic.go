package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"gorm.io/gorm"

	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &user.LoginResponse{
				StatusCode: constants.UserNotExistedCode,
				StatusMsg:  constants.UserNotExisted,
				UserId:     0,
				Avatar:     "",
			}, nil
		}
		return &user.LoginResponse{
			StatusCode: constants.FindDbErrorCode,
			StatusMsg:  constants.FindDbError,
			UserId:     0,
			Avatar:     "",
		}, nil
	}

	// 判断密码是否正确
	password := crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return &user.LoginResponse{
			StatusCode: constants.AuthUserLoginFailedCode,
			StatusMsg:  constants.AuthUserLoginFailed,
			UserId:     0,
			Avatar:     "",
		}, nil
	}

	return &user.LoginResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		UserId:     int64(res.Id),
		Avatar:     res.Avatar,
	}, nil
}
