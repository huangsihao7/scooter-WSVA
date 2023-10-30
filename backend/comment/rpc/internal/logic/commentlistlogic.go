package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"github.com/pkg/errors"
	"log"

	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"

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
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId

	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("用户不存在")
			return nil, errors.New("评论用户不存在")
		}
		return nil, err
	}

	// 检查视频id 是否存在 ddd
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("视频不存在")
			return &comment.CommentListResponse{
				StatusCode: constants.UnableToQueryCommentErrorCode,
				StatusMsg:  constants.UnableToQueryVideoError,
			}, nil
		}
		return nil, err
	}
	fmt.Println("+++++++++++++++++++++++")
	comments, err := l.svcCtx.Model.FindComments(l.ctx, videoId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	fmt.Println("-----------------------------", comments)

	reslist := make([]*comment.Comment, 0)
	for i := 0; i < len(comments); i++ {
		userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: userId, ActorId: comments[i].Uid})
		if err != nil {
			return nil, err
		}
		userDetail := &comment.User{
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
		res := &comment.Comment{
			Id:         comments[i].Id,
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
