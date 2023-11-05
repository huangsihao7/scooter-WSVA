package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/live/rpc/live"

	"github.com/huangsihao7/scooter-WSVA/live/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/live/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLiveListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveListLogic {
	return &LiveListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LiveListLogic) LiveList() (resp *types.LiveListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	info, err := l.svcCtx.LiveRpc.ListVideo(l.ctx, &live.ListLiveRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	resList := make([]types.User, 0)
	for _, item := range info.UserList {
		resList = append(resList, types.User{
			UId:          item.Id,
			Name:         item.Name,
			Avatar:       item.Avatar,
			Signature:    item.Signature,
			IsFollow:     item.IsFollow,
			LiveCoverUrl: item.LiveCoverUrl,
			LiveUrl:      item.LiveUrl,
		})
	}
	return &types.LiveListResp{
		StatusCode: int(info.StatusCode),
		StatusMsg:  info.StatusMsg,
		UserList:   resList,
	}, nil
}
