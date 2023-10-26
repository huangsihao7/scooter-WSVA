package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	userId := in.UserId
	//actionId := in.ActorId
	//查询用户是否存在的
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	users := &user.UserInfo{
		Name: res.Name,
		Id:   uint32(userId),
	}
	return &user.UserInfoResponse{
		User:       users,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
