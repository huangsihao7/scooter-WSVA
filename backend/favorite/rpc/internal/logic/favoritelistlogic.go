package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/pkg/errors"
	"log"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	// todo: add your logic here and delete this line
	//得到喜欢列表
	userId := in.UserId
	actorId := in.ActorId

	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("用户不存在")
			return nil, errors.New("评论用户不存在")
		}
		return nil, err
	}

	//检查用户id 是否能存在
	_, err = l.svcCtx.UserModel.FindOne(l.ctx, actorId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("用户不存在")
			return nil, errors.New("评论用户不存在")
		}
		return nil, err
	}

	favorVideos, err := l.svcCtx.Model.FindOwnFavorites(l.ctx, actorId)
	if err != nil {
		return nil, err
	}
	//得到喜欢视频的id
	videoInfoList := make([]*favorite.Video, 0)
	for i := 0; i < len(favorVideos); i++ {

		videoId := favorVideos[i].Vid
		videoDetail, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
		if err != nil {
			return nil, err
		}
		//这个视频的作者
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userId, ActorId: videoDetail.AuthorId})
		if err != nil {
			return nil, err
		}

		//userID 是否喜欢该视频
		isFavorited, err := l.svcCtx.Model.IsFavorite(l.ctx, userId, videoId)
		if err != nil {
			return nil, err
		}
		userDetail := &favorite.User{
			Id:             userInfo.User.Id,
			Name:           userInfo.User.Name,
			FollowCount:    userInfo.User.FollowCount,
			FollowerCount:  userInfo.User.FollowerCount,
			IsFollow:       userInfo.User.IsFollow,
			Avatar:         userInfo.User.Avatar,
			Signature:      userInfo.User.Signature,
			TotalFavorited: userInfo.User.TotalFavorited,
			WorkCount:      userInfo.User.WorkCount,
			FavoriteCount:  userInfo.User.FavoriteCount,
		}

		videoInfo := &favorite.Video{
			Id:            videoDetail.Id,
			Author:        userDetail,
			PlayUrl:       videoDetail.PlayUrl,
			CoverUrl:      videoDetail.CoverUrl,
			FavoriteCount: videoDetail.FavoriteCount,
			CommentCount:  videoDetail.CommentCount,
			IsFavorite:    isFavorited,
			Title:         videoDetail.Title,
		}

		videoInfoList = append(videoInfoList, videoInfo)
	}

	return &favorite.FavoriteListResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  videoInfoList,
	}, nil
}
