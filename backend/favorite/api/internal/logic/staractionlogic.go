package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/mq/format"

	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarActionLogic {
	return &StarActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarActionLogic) StarAction(req *types.ActionReq) (resp *types.ActionResp, err error) {

	//token 解析
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()
	if req.ActionType == 1 {
		format.Feedback(l.svcCtx.Config.RecommendUrl, "star", int(req.VideoId), int(userId))
	}
	//请求服务
	r, err := l.svcCtx.Favor.StarAction(l.ctx, &favorite.StarActionRequest{
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
