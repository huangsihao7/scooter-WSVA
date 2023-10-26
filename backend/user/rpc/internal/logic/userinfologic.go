package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"time"

	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	resp := new(user.UserInfoResponse)
	userId := in.UserId
	actionId := in.ActorId
	// 查询用户是否存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	followCountCh := make(chan int32)
	followerCountCh := make(chan int32)
	workCountCh := make(chan int32)
	favoriteCountCh := make(chan int32)
	totalFavoriteCountCh := make(chan int32)
	isFollowCh := make(chan bool)

	defer func() {
		close(followCountCh)
		close(followerCountCh)
		close(workCountCh)
		close(favoriteCountCh)
		close(totalFavoriteCountCh)
		close(isFollowCh)
	}()

	// 关注数
	go func() {
		followCount, _ := relationClient.GetFollowListCount(l.ctx, userId)
		followCountCh <- followCount
	}()

	// 粉丝数
	go func() {
		followerCount, _ := relationClient.GetFollowerListCount(l.ctx, userId)
		followerCountCh <- followerCount
	}()

	// 作品数
	go func() {
		workCount, _ := videoClient.GetWorkCount(l.ctx, userId)
		workCountCh <- workCount
	}()

	// 喜欢数
	go func() {
		favoriteCount, _ := favoriteClient.GetUserFavoriteCount(l.ctx, userId)
		favoriteCountCh <- favoriteCount
	}()

	// 总的被点赞数
	go func() {
		totalFavoriteCount, _ := favoriteClient.GetUserTotalFavoritedCount(l.ctx, userId)
		totalFavoriteCountCh <- totalFavoriteCount
	}()

	//是否关注该用户
	go func(isLog bool) {
		if isLog {
			isFollow, _ := relationClient.IsFollowing(ctx, &relation.IsFollowingRequest{
				ActorId: actionId,
				UserId:  userId,
			})
			isFollowCh <- isFollow
			return
		}
		isFollowCh <- false
	}(isLogged)

	var followCount, followerCount, workCount, favoriteCount, totalFavoriteCount uint32
	var isFollow bool

	// 从通道接收结果
	for receivedCount := 0; receivedCount < 6; receivedCount++ {
		select {
		case followCount = <-followCountCh:
			resp.User.FollowCount = &followCount
		case followerCount = <-followerCountCh:
			resp.User.FollowerCount = &followerCount
		case workCount = <-workCountCh:
			resp.User.WorkCount = &workCount
		case favoriteCount = <-favoriteCountCh:
			resp.User.FavoriteCount = &favoriteCount
		case totalFavoriteCount = <-totalFavoriteCountCh:
			resp.User.TotalFavorited = &totalFavoriteCount
		case isFollow = <-isFollowCh:
			resp.User.IsFollow = isFollow
		case <-time.After(3 * time.Second):
			zap.L().Error("3s overtime.")
			break
		}
	}

	resp.StatusCode = constants.ServiceOKCode
	resp.StatusMsg = constants.ServiceOK
	return resp, nil
}
