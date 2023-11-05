package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/qiniu/go-sdk/v7/pili"
	"strconv"

	"github.com/huangsihao7/scooter-WSVA/live/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/live/rpc/live"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看直播列表
func (l *ListVideoLogic) ListVideo(in *live.ListLiveRequest) (*live.ListLiveResponse, error) {
	man := pili.ManagerConfig{
		AccessKey: l.svcCtx.Config.AccessKey,
		SecretKey: l.svcCtx.Config.SecretKey,
		Transport: nil,
	}
	manger := pili.NewManager(man)
	list, err := manger.GetStreamsList(l.ctx, pili.GetStreamListRequest{
		Hub:      "scooter",
		LiveOnly: true,
	})
	if err != nil {
		return nil, err
	}
	if len(list.Items) == 0 {
		return &live.ListLiveResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
			UserList:   nil,
		}, nil
	}
	userlist := make([]*live.User, 0)
	for _, item := range list.Items {
		uid, err := strconv.Atoi(item.Key)
		if err != nil {
			return nil, err
		}
		info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId:  in.Uid,
			ActorId: int64(uid),
		})
		userlist = append(userlist, &live.User{
			Id:           int64(uid),
			Name:         info.User.Name,
			IsFollow:     info.User.IsFollow,
			Avatar:       *info.User.Avatar,
			Signature:    *info.User.Signature,
			LiveUrl:      fmt.Sprintf("http://%s/%s/%s.m3u8", l.svcCtx.Config.LiveUrl, l.svcCtx.Config.LiveBucket, item.Key),
			LiveCoverUrl: fmt.Sprintf("http://%s/%s/%s.jpg", l.svcCtx.Config.LiveCoverUrl, l.svcCtx.Config.LiveBucket, item.Key),
		})

	}

	return &live.ListLiveResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		UserList:   userlist,
	}, nil
}
