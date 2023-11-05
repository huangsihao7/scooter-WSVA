package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/live/rpc/internal/common"
	"strconv"

	"github.com/huangsihao7/scooter-WSVA/live/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLiveLogic {
	return &StartLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StartLive 开启直播
func (l *StartLiveLogic) StartLive(in *live.StartLiveRequest) (*live.StartLiveResponse, error) {
	idStr := strconv.Itoa(int(in.Uid))
	err := common.CreatStream(idStr, l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey, l.svcCtx.Config.LiveBucket)
	if err != nil {
		return nil, err
	}
	return &live.StartLiveResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		StreamUrl:  l.svcCtx.Config.PublishUrl + "/" + l.svcCtx.Config.LiveBucket + "/" + idStr,
	}, nil
}
