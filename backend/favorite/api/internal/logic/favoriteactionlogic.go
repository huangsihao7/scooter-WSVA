package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"

	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.ActionReq) (resp *types.ActionResp, err error) {
	// todo: add your logic here and delete this line

	//拿到用户token

	//token解析得到用户id
	//token 解析
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()

	//请求服务
	r, err := l.svcCtx.Favor.FavoriteAction(l.ctx, &favorite.FavoriteActionRequest{
		UserId:     userId,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}
	return &types.ActionResp{
		StatusCode: int(r.StatusCode),
		StatusMsg:  r.StatusMsg,
	}, nil
}
