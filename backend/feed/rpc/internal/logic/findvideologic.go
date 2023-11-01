package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVideoLogic {
	return &FindVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindVideoLogic) FindVideo(in *feed.FindVideoReq) (*feed.FindVideoResp, error) {
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, int64(in.Vid))
	if err != nil {
		return &feed.FindVideoResp{
			StatusCode: constants.UnableToQueryVideoErrorCode,
			StatusMsg:  constants.UnableToQueryVideoError,
			Video:      nil,
		}, err
	}
	userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId:  int64(in.Uid),
		ActorId: int64(video.AuthorId),
	})
	if err != nil {
		return &feed.FindVideoResp{
			StatusCode: constants.UnableToQueryUserErrorCode,
			StatusMsg:  constants.UnableToQueryUserError,
			Video:      nil,
		}, err
	}
	IsFavorite, err := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.Uid), int64(in.Vid))
	IsStar, _ := l.svcCtx.StarModel.IsStarExist(l.ctx, int64(in.Uid), int64(in.Vid))
	userInfo := &feed.User{
		Id:              userRpcRes.User.Id,
		Name:            userRpcRes.User.Name,
		FollowCount:     userRpcRes.User.FollowCount,
		FollowerCount:   userRpcRes.User.FollowCount,
		IsFollow:        userRpcRes.User.IsFollow,
		Avatar:          userRpcRes.User.Avatar,
		BackgroundImage: userRpcRes.User.BackgroundImage,
		Signature:       userRpcRes.User.Signature,
		TotalFavorited:  userRpcRes.User.TotalFavorited,
		WorkCount:       userRpcRes.User.WorkCount,
		FavoriteCount:   userRpcRes.User.FavoriteCount,
		Gender:          userRpcRes.User.Gender,
	}
	videoInfo := &feed.VideoInfo{
		Id:            uint32(in.Vid),
		Author:        userInfo,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: uint32(video.FavoriteCount),
		CommentCount:  uint32(video.CommentCount),
		IsFavorite:    IsFavorite,
		Title:         video.Title,
		CreateTime:    video.CreatedAt.Time.Format(constants.TimeFormat),
		StarCount:     uint32(video.StarCount),
		IsStar:        IsStar,
		Duration:      video.Duration.String,
	}
	return &feed.FindVideoResp{
		StatusCode: 0,
		StatusMsg:  "",
		Video:      videoInfo,
	}, nil
}
