package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendVideosLogic {
	return &RecommendVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendVideosLogic) RecommendVideos(req *types.RecommendVideosListReq) (resp *types.RecommendVideosListResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var recommend *feed.ListFeedResponse
	if req.Offset == 0 {
		recommend, err = l.svcCtx.FeedRpc.ListVideosByRecommend(l.ctx, &feed.ListFeedRequest{
			ActorId: uint32(uid),
			Num:     5,
			Offset:  req.Offset,
			ReadVid: req.ReadedVideoId,
		})
	} else {
		recommend, err = l.svcCtx.FeedRpc.ListVideosByRecommend(l.ctx, &feed.ListFeedRequest{
			ActorId: uint32(uid),
			Num:     1,
			Offset:  req.Offset,
			ReadVid: req.ReadedVideoId,
		})
	}
	if err != nil {
		return &types.RecommendVideosListResp{
			StatusCode: int(recommend.StatusCode),
			StatusMsg:  recommend.StatusMsg,
			Videos:     nil,
		}, nil
	}
	resList := make([]types.VideoInfo, 0)
	for _, item := range recommend.VideoList {
		resList = append(resList, types.VideoInfo{
			VideoId: int64(item.Id),
			Author: types.UserInfo{
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				Avatar:          *item.Author.Avatar,
				Gender:          item.Author.Gender,
				Signature:       *item.Author.Signature,
				BackgroundImage: *item.Author.BackgroundImage,
				FollowCount:     *item.Author.FollowerCount,
				FollowerCount:   *item.Author.FollowCount,
				TotalFavorited:  *item.Author.TotalFavorited,
				WorkCount:       *item.Author.WorkCount,
				FavoriteCount:   *item.Author.FavoriteCount,
				IsFollow:        item.Author.IsFollow,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: int64(item.FavoriteCount),
			CommentCount:  int64(item.CommentCount),
			StarCount:     int64(item.StarCount),
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
			CreateTime:    item.CreateTime,
			Duration:      item.Duration,
		})
	}
	return &types.RecommendVideosListResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		Videos:     resList,
	}, nil

}
