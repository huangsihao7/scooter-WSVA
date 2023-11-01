package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	userId := in.UserId
	actionId := in.ActorId
	//查询用户是否存在的
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &user.UserInfoResponse{
				StatusCode: constants.UserNotExistedCode,
				StatusMsg:  constants.UserNotExisted,
				User:       nil,
			}, nil
		}
		return &user.UserInfoResponse{
			StatusCode: constants.FindDbErrorCode,
			StatusMsg:  constants.FindDbError,
			User:       nil,
		}, nil
	}

	//查询用户是否存在的
	res, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(actionId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &user.UserInfoResponse{
				StatusCode: constants.UserNotExistedCode,
				StatusMsg:  constants.UserNotExisted,
				User:       nil,
			}, nil
		}
		return &user.UserInfoResponse{
			StatusCode: constants.FindDbErrorCode,
			StatusMsg:  constants.FindDbError,
			User:       nil,
		}, nil
	}

	//查询关注数
	followCount, err := l.svcCtx.RelationModel.GetFollowCount(l.ctx, actionId)
	if err != nil {
		return &user.UserInfoResponse{
			StatusCode: constants.UnableToGetFollowCountErrorCode,
			StatusMsg:  constants.UnableToGetFollowCountError,
			User:       nil,
		}, nil
	}
	uint32followCount := uint32(followCount)

	//查询粉丝数
	followerCount, err := l.svcCtx.RelationModel.GetFollowerCount(l.ctx, actionId)
	if err != nil {
		return &user.UserInfoResponse{
			StatusCode: constants.UnableToGetFollowerCountErrorCode,
			StatusMsg:  constants.UnableToGetFollowerCountError,
			User:       nil,
		}, nil
	}
	uint32followerCount := uint32(followerCount)

	//查询点赞视频数
	favorCount, err := l.svcCtx.FavorModel.GetFavoriteCount(l.ctx, actionId)
	if err != nil {
		return &user.UserInfoResponse{
			StatusCode: constants.UnableToGetFavoriteVideosCountErrorCode,
			StatusMsg:  constants.UnableToGetFavoriteVideosCountError,
			User:       nil,
		}, nil
	}
	uint32favorCount := uint32(len(favorCount))

	//查询视频获赞数
	favorVideoCount, err := l.svcCtx.FavorModel.GetVideoCount(l.ctx, actionId)
	if err != nil {
		return &user.UserInfoResponse{
			StatusCode: constants.UnableToGetVideosFavoriteCountErrorCode,
			StatusMsg:  constants.UnableToGetVideosFavoriteCountError,
			User:       nil,
		}, nil
	}
	uint32FavorVideoCount := uint32(len(favorVideoCount))

	//查询A是否关注B
	isFavorite := false
	//如果查询的是自己 那就不做这个查询了
	if userId != actionId {
		isFavorite = l.svcCtx.RelationModel.FindIsFavorite(l.ctx, userId, actionId)
		if err != nil {
			return &user.UserInfoResponse{
				StatusCode: constants.UnableToGetFavoriteErrorCode,
				StatusMsg:  constants.UnableToGetFavoriteError,
				User:       nil,
			}, nil
		}
	}

	voideoList, err := l.svcCtx.VideoModel.FindOwnFeed(l.ctx, actionId)
	if err != nil {
		return &user.UserInfoResponse{
			StatusCode: constants.UserServiceInnerErrorCode,
			StatusMsg:  constants.UserServiceInnerError,
		}, nil
	}
	workCount := uint32(len(voideoList))
	users := user.UserInfo{
		Name:            res.Name,
		Id:              uint32(actionId),
		Avatar:          &res.Avatar,
		Gender:          uint32(res.Gender),
		Signature:       &res.Dec,
		FollowCount:     &uint32followCount,
		FollowerCount:   &uint32followerCount,
		FavoriteCount:   &uint32favorCount,
		TotalFavorited:  &uint32FavorVideoCount,
		IsFollow:        isFavorite,
		WorkCount:       &workCount,
		BackgroundImage: &res.BackgroundUrl,
	}

	return &user.UserInfoResponse{
		User:       &users,
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}, nil
}
