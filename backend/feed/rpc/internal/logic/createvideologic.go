package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/common"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
)

type CreateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoLogic) CreateVideo(in *feed.CreateVideoRequest) (*feed.CreateVideoResponse, error) {
	newVideo := gmodel.Videos{
		AuthorId:      uint(in.ActorId),
		Title:         in.Title,
		CoverUrl:      in.CoverUrl,
		PlayUrl:       in.Url,
		FavoriteCount: 0,
		CommentCount:  0,
		StarCount:     0,
		Category:      int(in.Category),
	}
	err := l.svcCtx.VideoModel.Insert(l.ctx, &newVideo)
	if err != nil {
		return &feed.CreateVideoResponse{
			StatusCode: constants.VideoServiceInnerErrorCode,
			StatusMsg:  constants.VideoServiceInnerError,
		}, nil
	}

	JobId := common.IsSafeJobId(in.Url, strconv.Itoa(int(newVideo.Id)))

	//将文件信息传入mq
	messagekq := format.UploadFile{
		Id:  int64(newVideo.Id),
		Url: in.Url,
		Uid: int64(newVideo.AuthorId),
	}
	jobKq := format.JobBody{
		Job: JobId,
	}

	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data, err := json.Marshal(messagekq)
		if err != nil {
			l.Logger.Errorf("[Video] marshal msg: %v error: %v", messagekq, err)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Video] kq push data: %s error: %v", data, err)
		}
	})
	threading.GoSafe(func() {
		data, err := json.Marshal(jobKq)
		if err != nil {
			l.Logger.Errorf("[Video] marshal msg: %v error: %v", jobKq, err)
			return
		}
		err = l.svcCtx.KqPusherJobClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Video] kq push data: %s error: %v", data, err)
		}
	})
	return &feed.CreateVideoResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
