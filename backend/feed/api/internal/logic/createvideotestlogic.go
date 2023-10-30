package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoTestLogic {
	return &CreateVideoTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoTestLogic) CreateVideoTest(req *types.CreateVideoReq) (resp *types.CreateVideoResp, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.FeedRpc.CreateVideoTest(l.ctx, &feed.CreateVideoRequest{
		ActorId:  1,
		CoverUrl: "2",
		Url:      "3",
		Title:    "4",
		Category: 1,
	})
	return &types.CreateVideoResp{
		StatusCode: 200,
		StatusMsg:  "OK",
	}, nil
}
