package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoTestLogic {
	return &CreateVideoTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoTestLogic) CreateVideoTest(in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	msg := &feed.CreateVideoRequest{
		ActorId:  1,
		CoverUrl: "www.123",
		Url:      "www.456",
		Title:    "test queue",
		Category: 1,
	}

	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Video] marshal msg: %v error: %v", msg, err)
			return
		}
		err = l.svcCtx.KqPusherTestClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Video] kq push data: %s error: %v", data, err)
		}

	})
	return &feed.CreateVideoResponse{
		StatusCode: 200,
		StatusMsg:  "成功",
	}, nil
}
