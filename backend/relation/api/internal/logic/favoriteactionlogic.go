package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relation"

	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/types"

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

	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.RelationRpc.Favorite(l.ctx, &relation.FavoriteRequest{
		Uid:    uid,
		ToUid:  req.ToUserId,
		Action: req.Action,
	})
	if err != nil {
		return nil, err
	}
	return &types.ActionResp{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil

}
