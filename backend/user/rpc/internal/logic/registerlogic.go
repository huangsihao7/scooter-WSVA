package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"google.golang.org/grpc/status"

	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:          in.Name,
			Gender:        in.Gender,
			Mobile:        in.Mobile,
			Password:      crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Avatar:        in.Avatar,
			Dec:           in.Dec,
			BackgroundUrl: in.BackgroundImage,
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:              newUser.Id,
			Name:            newUser.Name,
			Gender:          newUser.Gender,
			Mobile:          newUser.Mobile,
			Avatar:          newUser.Avatar,
			Dec:             newUser.Dec,
			BackgroundImage: newUser.BackgroundUrl,
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
