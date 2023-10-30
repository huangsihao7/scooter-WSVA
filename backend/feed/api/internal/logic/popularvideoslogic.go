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

type PopularVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPopularVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PopularVideosLogic {
	return &PopularVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PopularVideosLogic) PopularVideos(req *types.PopularVideosListReq) (resp *types.PopularVideosListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var popular *feed.ListFeedResponse
	if req.Offset == 0 {
		popular, err = l.svcCtx.FeedRpc.ListPopularVideos(l.ctx, &feed.ListFeedRequest{
			ActorId: uint32(uid),
			Num:     5,
			Offset:  req.Offset,
		})
	} else {
		popular, err = l.svcCtx.FeedRpc.ListPopularVideos(l.ctx, &feed.ListFeedRequest{
			ActorId: uint32(uid),
			Num:     1,
			Offset:  req.Offset,
		})
	}
	if err != nil {
		return &types.PopularVideosListResp{
			StatusCode: int(popular.StatusCode),
			StatusMsg:  popular.StatusMsg,
			Videos:     nil,
		}, nil
	}
	resList := make([]types.VideoInfo, 0)
	for _, item := range popular.VideoList {
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
			StarCount:     int64(item.StarCount),
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
			CreateTime:    item.CreateTime,
		})
	}
	return &types.PopularVideosListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		Videos:     resList,
	}, nil
}