package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/huangsihao7/scooter-WSVA/user/code"
	"gorm.io/gorm"
	"log"

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
			l.Logger.Infof("用户不存在")
			return nil, code.UserNotExistError
		}
		return nil, err
	}

	// 判断密码是否
	password := crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		log.Println("密码错误")
		return nil, code.UserNotExistError
	}

	return &user.LoginResponse{
		StatusCode:      constants.ServiceOKCode,
		StatusMsg:       constants.ServiceOK,
		UserId:          int64(res.Id),
		Avatar:          res.Avatar,
		Name:            res.Name,
		Gender:          int64(res.Gender),
		BackgroundImage: res.BackgroundUrl,
		Signature:       res.Dec,
	}, nil
}
