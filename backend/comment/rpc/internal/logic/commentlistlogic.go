package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/code"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListRequest) (*comment.CommentListResponse, error) {

	userId := in.UserId
	videoId := in.VideoId

	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("评论用户不存在")
			return nil, code.CommentUserIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	// 检查视频id 是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("评论视频不存在")
			return nil, code.CommentVideoIdEmptyError
		}
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	comments, err := l.svcCtx.CommentModel.FindComments(l.ctx, videoId)
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Logger.Errorf(err.Error())
		return nil, err
	}

	reslist := make([]*comment.Comment, 0)
	for i := 0; i < len(comments); i++ {
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userId, ActorId: int64(comments[i].Uid)})
		if err != nil {
			l.Logger.Errorf(err.Error())
			return nil, err
		}
		userDetail := &comment.User{
			Id:             userInfo.User.Id,
			Name:           userInfo.User.Name,
			FollowCount:    userInfo.User.FollowCount,
			Gender:         userInfo.User.Gender,
			FollowerCount:  userInfo.User.FollowerCount,
			IsFollow:       userInfo.User.IsFollow,
			Avatar:         userInfo.User.Avatar,
			Signature:      userInfo.User.Signature,
			TotalFavorited: userInfo.User.TotalFavorited,
			WorkCount:      userInfo.User.WorkCount,
			FavoriteCount:  userInfo.User.FavoriteCount,
			FriendCount:    userInfo.User.FriendCount,
		}
		res := &comment.Comment{
			Id:         int64(comments[i].Id),
			User:       userDetail,
			Content:    comments[i].Content,
			CreateDate: comments[i].CreatedAt.Format(constants.TimeFormat),
		}

		reslist = append(reslist, res)
	}

	return &comment.CommentListResponse{
		StatusCode:  constants.ServiceOKCode,
		StatusMsg:   constants.ServiceOK,
		CommentList: reslist,
	}, nil
}
