package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	update, _ := l.svcCtx.UserRpc.Update(l.ctx, &user.UpdateReq{
		Name:            req.Name,
		Gender:          req.Gender,
		Avatar:          req.Avatar,
		Dec:             req.Dec,
		BackgroundImage: req.BackgroundImage,
		Uid:             uid,
	})
	return &types.UpdateResponse{
		StatusCode: int(update.StatusCode),
		StatusMsg:  update.StatusMsg,
	}, nil
}
