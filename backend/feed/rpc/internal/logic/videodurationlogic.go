package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoDurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoDurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoDurationLogic {
	return &VideoDurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoDurationLogic) VideoDuration(in *feed.VideoDurationReq) (*feed.VideoDurationResp, error) {
	err := l.svcCtx.VideoModel.UpdateDuration(l.ctx, int64(in.VideoId), in.Duration)
	if err != nil {
		return &feed.VideoDurationResp{
			StatusCode: constants.DurationErrorCode,
			StatusMsg:  constants.DurationError,
		}, err
	}

	return &feed.VideoDurationResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
