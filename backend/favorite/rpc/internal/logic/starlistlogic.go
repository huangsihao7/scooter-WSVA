package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/favorite/code"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"gorm.io/gorm"
	"log"

	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/favorite"
	"github.com/huangsihao7/scooter-WSVA/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStarListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarListLogic {
	return &StarListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StarListLogic) StarList(in *favorite.StarListRequest) (*favorite.StarListResponse, error) {

	userId := in.UserId
	actorId := in.ActorId

	//检查用户id 是否存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("收藏用户不存在")
			return nil, code.StarUserIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	//检查用户id 是否存在
	_, err = l.svcCtx.UserModel.GetUserByID(l.ctx, uint(actorId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("收藏用户不存在")
			return nil, code.StarUserIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	StarVideos, err := l.svcCtx.StarModel.FindsByUserId(l.ctx, actorId)
	if err != nil {
		l.Logger.Errorf(err.Error())
		return nil, err
	}
	//得到喜欢视频的id
	videoInfoList := make([]*favorite.Video, 0)
	for i := 0; i < len(StarVideos); i++ {

		videoId := StarVideos[i].Vid
		videoDetail, err := l.svcCtx.VideoModel.FindById(l.ctx, int64(videoId))
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		//这个视频的作者
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userId, ActorId: int64(videoDetail.AuthorId)})
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}

		//userID 是否喜欢该视频
		isFavorited, err := l.svcCtx.FavorModel.IsFavorite(l.ctx, userId, int64(videoId))
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		//userId 是否收藏该视频
		isStar, err := l.svcCtx.StarModel.IsStarExist(l.ctx, userId, int64(videoId))
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		userDetail := &favorite.User{
			Id:              userInfo.User.Id,
			Name:            userInfo.User.Name,
			Gender:          userInfo.User.Gender,
			FollowCount:     userInfo.User.FollowCount,
			FollowerCount:   userInfo.User.FollowerCount,
			BackgroundImage: userInfo.User.BackgroundImage,
			IsFollow:        userInfo.User.IsFollow,
			Avatar:          userInfo.User.Avatar,
			Signature:       userInfo.User.Signature,
			TotalFavorited:  userInfo.User.TotalFavorited,
			WorkCount:       userInfo.User.WorkCount,
			FavoriteCount:   userInfo.User.FavoriteCount,
			FriendCount:     userInfo.User.FriendCount,
		}

		videoInfo := &favorite.Video{
			Id:            int64(videoDetail.Id),
			Author:        userDetail,
			PlayUrl:       videoDetail.PlayUrl,
			CoverUrl:      videoDetail.CoverUrl,
			FavoriteCount: int64(videoDetail.FavoriteCount),
			CommentCount:  int64(videoDetail.CommentCount),
			StarCount:     int64(videoDetail.StarCount),
			IsFavorite:    isFavorited,
			IsStar:        isStar,
			Title:         videoDetail.Title,
			CreateTime:    videoDetail.CreatedAt.Time.Format(constants.TimeFormat),
			Duration:      videoDetail.Duration.String,
		}
		//
		videoInfoList = append(videoInfoList, videoInfo)
	}

	return &favorite.StarListResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  videoInfoList,
	}, nil
}
