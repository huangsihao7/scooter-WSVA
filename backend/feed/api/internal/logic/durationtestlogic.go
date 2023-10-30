package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DurationTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDurationTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DurationTestLogic {
	return &DurationTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DurationTestLogic) DurationTest(req *types.DurationTestReq) (resp *types.DurationTestResp, err error) {
	durationResp, err := l.svcCtx.FeedRpc.VideoDuration(l.ctx, &feed.VideoDurationReq{
		Duration: req.Duration,
		VideoId:  uint32(req.Vid),
	})
	if err != nil {
		return &types.DurationTestResp{StatusMsg: durationResp.StatusMsg, StatusCode: int(durationResp.StatusCode)}, err
	}

	return &types.DurationTestResp{StatusMsg: durationResp.StatusMsg, StatusCode: int(durationResp.StatusCode)}, nil
}
