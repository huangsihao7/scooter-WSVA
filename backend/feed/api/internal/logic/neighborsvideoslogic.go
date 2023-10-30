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

type NeighborsVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNeighborsVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NeighborsVideosLogic {
	return &NeighborsVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NeighborsVideosLogic) NeighborsVideos(req *types.NeighborsVideoReq) (resp *types.NeighborsVideoResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	neighbors, err := l.svcCtx.FeedRpc.ListNeighborVideos(l.ctx, &feed.NeighborsReq{
		Vid: int32(req.Vid),
		Uid: int32(uid),
	})
	if err != nil {
		println("-------------222----------------------")
		return &types.NeighborsVideoResp{
			StatusCode: int(neighbors.StatusCode),
			StatusMsg:  neighbors.StatusMsg,
			VideoList:  nil,
		}, nil
	}
	resList := make([]types.VideoInfo, 0)
	for _, item := range neighbors.VideoList {
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
	return &types.NeighborsVideoResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  resList,
	}, nil
}
