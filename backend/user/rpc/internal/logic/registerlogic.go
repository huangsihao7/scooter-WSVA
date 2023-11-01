package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

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
		return &user.RegisterResponse{
			StatusCode: constants.AuthUserExistedCode,
			StatusMsg:  constants.AuthUserExisted,
		}, nil
	}

	if err == gorm.ErrRecordNotFound {
		newUser := gmodel.User{
			Name:          in.Name,
			Gender:        uint(in.Gender),
			Mobile:        in.Mobile,
			Password:      crypt.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Avatar:        in.Avatar,
			Dec:           in.Dec,
			BackgroundUrl: in.BackgroundImage,
		}

		err = l.svcCtx.UserModel.InsertUser(l.ctx, &newUser)
		if err != nil {
			return &user.RegisterResponse{
				StatusCode: constants.InsertUserErrorCode,
				StatusMsg:  constants.InsertUserError,
			}, nil
		}

		postbaseurl := "http://172.22.121.54:8088/api/user"
		userReq := format.UserGoresBody{UserId: fmt.Sprintf("%d", newUser.Id)}
		jsonData, err := json.Marshal(userReq)
		if err != nil {
			fmt.Println("JSON编码失败:", err)
			return nil, status.Error(500, "JSON编码失败")
		}
		_, err = format.QiNiuPost(postbaseurl, jsonData)
		return &user.RegisterResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
	}
	return &user.RegisterResponse{
		StatusCode: constants.FindDbErrorCode,
		StatusMsg:  constants.FindDbError,
	}, nil
}
