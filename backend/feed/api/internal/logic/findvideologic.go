package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVideoLogic {
	return &FindVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 查询视频
func (l *FindVideoLogic) FindVideo(req *types.FindVideoByIdReq) (resp *types.FindVideoByIdResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	video, err := l.svcCtx.FeedRpc.FindVideo(l.ctx, &feed.FindVideoReq{
		Uid: int32(uid),
		Vid: int32(req.Vid),
	})
	if err != nil {
		return nil, err
	}
	return &types.FindVideoByIdResp{
		StatusCode: int(video.StatusCode),
		StatusMsg:  video.StatusMsg,
		Video: types.VideoInfo{
			VideoId: int64(video.Video.Id),
			Author: types.UserInfo{
				Id:              video.Video.Author.Id,
				Name:            video.Video.Author.Name,
				Avatar:          *video.Video.Author.Avatar,
				Gender:          video.Video.Author.Gender,
				Signature:       *video.Video.Author.Signature,
				BackgroundImage: *video.Video.Author.BackgroundImage,
				FollowCount:     *video.Video.Author.FollowerCount,
				FollowerCount:   *video.Video.Author.FollowCount,
				TotalFavorited:  *video.Video.Author.TotalFavorited,
				WorkCount:       *video.Video.Author.WorkCount,
				FavoriteCount:   *video.Video.Author.FavoriteCount,
				IsFollow:        video.Video.Author.IsFollow,
				FriendCount:     video.Video.Author.FriendCount,
			},
			PlayUrl:       video.Video.PlayUrl,
			CoverUrl:      video.Video.CoverUrl,
			FavoriteCount: int64(video.Video.FavoriteCount),
			CommentCount:  int64(video.Video.CommentCount),
			StarCount:     int64(video.Video.StarCount),
			IsFavorite:    video.Video.IsFavorite,
			Title:         video.Video.Title,
			CreateTime:    video.Video.CreateTime,
			Duration:      video.Video.Duration,
			IsStar:        video.Video.IsStar,
		},
	}, nil

}
