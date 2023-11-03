package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/mq/format"

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

	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	if req.ActionType == 1 {
		format.Feedback(l.svcCtx.Config.RecommendUrl, "like", int(req.VideoId), int(userId))
	}
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
