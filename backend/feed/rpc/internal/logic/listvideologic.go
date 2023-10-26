package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideoLogic) ListVideo(in *feed.ListVideoRequest) (*feed.ListVideoResponse, error) {
	Feeds, err := l.svcCtx.FeedModel.FindOwnFeed(l.ctx, int64(in.ToUid))
	if err != nil {
		if err == model.ErrNotFound {
			return &feed.ListVideoResponse{
				StatusCode: constants.UserDoNotExistedCode,
				StatusMsg:  constants.UserDoNotExisted,
			}, nil
		} else {
			return &feed.ListVideoResponse{
				StatusCode: constants.FindUserVideosErrorCode,
				StatusMsg:  constants.FindUserVideosError,
			}, nil
		}
	}

	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range Feeds {
		IsFavorite, _ := l.svcCtx.FavoriteModel.IsFavorite(l.ctx, int64(in.Uid), item.Id)
		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        nil,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			IsFavorite:    IsFavorite,
			Title:         item.Title,
		})
	}
	return &feed.ListVideoResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
