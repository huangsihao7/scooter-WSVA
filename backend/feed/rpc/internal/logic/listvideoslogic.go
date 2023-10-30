package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosLogic {
	return &ListVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListVideosLogic) ListVideos(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	Feeds, err := l.svcCtx.FeedModel.FindFeeds(l.ctx)
	if err != nil {
		if err == model.ErrNotFound {
			return &feed.ListFeedResponse{
				StatusCode: constants.UserVideosDoNotExistedCode,
				StatusMsg:  constants.UserVideosDoNotExisted,
			}, nil
		} else {
			return &feed.ListFeedResponse{
				StatusCode: constants.FindUserVideosErrorCode,
				StatusMsg:  constants.FindUserVideosError,
			}, nil
		}
	}
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range Feeds {
		userRpcRes, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: item.AuthorId})

		userInfo := &feed.User{
			Id:             userRpcRes.User.Id,
			Name:           userRpcRes.User.Name,
			FollowCount:    userRpcRes.User.FollowCount,
			FollowerCount:  userRpcRes.User.FollowCount,
			IsFollow:       userRpcRes.User.IsFollow,
			Avatar:         userRpcRes.User.Avatar,
			Signature:      userRpcRes.User.Signature,
			TotalFavorited: userRpcRes.User.TotalFavorited,
			WorkCount:      userRpcRes.User.WorkCount,
			FavoriteCount:  userRpcRes.User.FavoriteCount,
		}
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.ActorId), item.Id)
		IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.ActorId), item.Id)

		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        userInfo,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			StarCount:     uint32(item.StarCount),
			IsFavorite:    IsFavorite,
			IsStar:        IsStar,
			Title:         item.Title,
			CreateTime:    item.CreatedAt.Format(constants.TimeFormat),
		})
	}
	return &feed.ListFeedResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
