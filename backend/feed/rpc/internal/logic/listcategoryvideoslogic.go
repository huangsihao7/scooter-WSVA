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

type ListCategoryVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoryVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryVideosLogic {
	return &ListCategoryVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCategoryVideosLogic) ListCategoryVideos(in *feed.CategoryFeedRequest) (*feed.CategoryFeedResponse, error) {
	Feeds, err := l.svcCtx.FeedModel.FindCategoryFeeds(l.ctx, int64(in.Category))
	if err != nil {
		if err == model.ErrNotFound {
			return &feed.CategoryFeedResponse{
				StatusCode: constants.CategoryVideosDoNotExistedCode,
				StatusMsg:  constants.CategoryVideosDoNotExisted,
			}, err
		} else {
			return &feed.CategoryFeedResponse{
				StatusCode: constants.FindCategoryVideosErrorCode,
				StatusMsg:  constants.FindCategoryVideosError,
			}, err
		}
	}

	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range Feeds {
		userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: item.AuthorId})
		if err != nil {
			return &feed.CategoryFeedResponse{
				StatusCode: constants.FindUserErrorCode,
				StatusMsg:  constants.FindUserError,
			}, err
		}

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
		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(item.Id),
			Author:        userInfo,
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: uint32(item.FavoriteCount),
			CommentCount:  uint32(item.CommentCount),
			IsFavorite:    IsFavorite,
			Title:         item.Title,
			CreateTime:    item.CreatedAt.Unix(),
		})
	}
	return &feed.CategoryFeedResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
