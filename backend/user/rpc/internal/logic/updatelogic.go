package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/gmodel"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.UpdateReq) (*user.UpdateResp, error) {
	err := l.svcCtx.UserModel.UpdateUser(l.ctx, &gmodel.User{
		Id:            uint64(in.Uid),
		Name:          in.Name,
		Gender:        uint(in.Gender),
		Dec:           in.Dec,
		Avatar:        in.Avatar,
		BackgroundUrl: in.BackgroundImage,
	})
	if err != nil {
		return nil, err
	}
	return &user.UpdateResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
