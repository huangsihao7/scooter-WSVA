package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/zeromicro/go-zero/core/logx"
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
	newVideo := model.Videos{
		AuthorId:      int64(in.ActorId),
		Title:         in.Title,
		CoverUrl:      in.CoverUrl,
		PlayUrl:       in.Url,
		FavoriteCount: 0,
		CommentCount:  0,
		Category:      int64(in.Category),
	}
	res, err := l.svcCtx.FeedModel.Insert(l.ctx, &newVideo)
	//newVideo.Id,_ = res.LastInsertId()
	if err != nil {
		return &feed.CreateVideoResponse{
			StatusCode: constants.VideoServiceInnerErrorCode,
			StatusMsg:  constants.VideoServiceInnerError,
		}, nil
	}
	newVideo.Id, err = res.LastInsertId()

	//将文件信息传入mq
	messagekq := format.UploadFile{
		Id:  newVideo.Id,
		Url: in.Url,
	}
	jsonString, err := json.Marshal(messagekq)
	if err != nil {
		return &feed.CreateVideoResponse{
			StatusCode: constants.VideoServiceInnerErrorCode,
			StatusMsg:  constants.VideoServiceInnerError,
		}, nil
	}
	if err = l.svcCtx.KqPusherClient.Push(string(jsonString)); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}

	return &feed.CreateVideoResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
