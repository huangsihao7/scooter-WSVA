package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoLogic) CreateVideo(req *types.CreateVideoReq) (resp *types.CreateVideoResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.FeedRpc.CreateVideo(l.ctx, &feed.CreateVideoRequest{
		ActorId:  uint32(uid),
		CoverUrl: req.CoverUrl,
		Url:      req.Url,
		Title:    req.Title,
		Category: uint32(req.Category),
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateVideoResp{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil
}
