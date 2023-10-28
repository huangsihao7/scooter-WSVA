package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideosListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideosListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideosListLogic {
	return &VideosListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideosListLogic) VideosList() (resp *types.VideosListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	videos, err := l.svcCtx.FeedRpc.ListVideos(l.ctx, &feed.ListFeedRequest{ActorId: uint32(uid)})
	if err != nil {
		return &types.VideosListResp{
			StatusCode: int(videos.StatusCode),
			StatusMsg:  videos.StatusMsg,
			Videos:     nil,
		}, nil
	}
	resList := make([]types.VideoInfo, 0)
	for _, item := range videos.VideoList {
		resList = append(resList, types.VideoInfo{
			Id: int64(item.Id),
			Author: types.UserInfo{
				Id:             item.Author.Id,
				Name:           item.Author.Name,
				Avatar:         *item.Author.Avatar,
				Signature:      *item.Author.Signature,
				FollowCount:    *item.Author.FollowerCount,
				FollowerCount:  *item.Author.FollowCount,
				TotalFavorited: *item.Author.TotalFavorited,
				WorkCount:      *item.Author.WorkCount,
				FavoriteCount:  *item.Author.FavoriteCount,
				IsFollow:       item.Author.IsFollow,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: int64(item.FavoriteCount),
			CommentCount:  int64(item.CommentCount),
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
			CreateTime:    item.CreateTime,
		})
	}
	return &types.VideosListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		Videos:     resList,
	}, nil
}
