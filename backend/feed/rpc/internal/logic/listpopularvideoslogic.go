package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/mq/format"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"strconv"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPopularVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPopularVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPopularVideosLogic {
	return &ListPopularVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPopularVideosLogic) ListPopularVideos(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	baseurl := "http://172.22.121.54:8088/api/popular"
	url := fmt.Sprintf("%s?user-id=%d&n=%d&offset=%d", baseurl, in.ActorId, in.Num, in.Offset)

	getresponse, err := format.QiNiuGet(url)
	if err != nil {
		return &feed.ListFeedResponse{
			StatusCode: constants.RecommendServiceInnerErrorCode,
			StatusMsg:  constants.RecommendServiceInnerError,
			VideoList:  nil,
		}, err
	}
	var result []format.PopularResp
	println(string(getresponse))
	err = json.Unmarshal(getresponse, &result)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return &feed.ListFeedResponse{
			StatusCode: constants.JsonErrorCode,
			StatusMsg:  constants.JsonError,
			VideoList:  nil,
		}, err
	}
	VideoList := make([]*feed.VideoInfo, 0)
	for _, item := range result {
		id, err := strconv.Atoi(item.Id)
		if err != nil {
			return &feed.ListFeedResponse{
				StatusCode: constants.StringToIntErrorCode,
				StatusMsg:  constants.StringToIntError,
				VideoList:  nil,
			}, err
		}
		video, err := l.svcCtx.FeedModel.FindOne(l.ctx, int64(id))
		userRpcRes, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: int64(in.ActorId), ActorId: video.AuthorId})

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
		IsFavorite, _ := l.svcCtx.FavorModel.IsFavorite(l.ctx, int64(in.ActorId), video.Id)
		VideoList = append(VideoList, &feed.VideoInfo{
			Id:            uint32(video.Id),
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: uint32(video.FavoriteCount),
			CommentCount:  uint32(video.CommentCount),
			StarCount:     uint32(video.StarCount),
			IsFavorite:    IsFavorite,
			Title:         video.Title,
			CreateTime:    video.CreatedAt.Unix(),
		})
	}
	println("------------------------")
	println(VideoList)
	return &feed.ListFeedResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  VideoList,
	}, nil
}
