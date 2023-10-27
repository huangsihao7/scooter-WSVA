package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"
	"log"

	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	constants "github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId
	contents := in.CommentText
	actionType := in.ActionType
	commentId := in.CommentId

	//检查用户id 是否能存在
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		ActorId: userId,
	})
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		log.Println("用户不存在")
	}
	// 检查视频id 是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		return nil, err
	}

	switch actionType {
	//添加评论
	case 1:
		newComment := model.Comments{
			Uid:     userId,
			Vid:     videoId,
			Content: contents,
		}
		_, err := l.svcCtx.Model.Insert(l.ctx, &newComment)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return &comment.CommentActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
		//删除评论
	case 2:
		err := l.svcCtx.Model.Delete(l.ctx, commentId)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return &comment.CommentActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
	}
	return &comment.CommentActionResponse{
		StatusCode: constants.UnableToCreateCommentErrorCode,
		StatusMsg:  constants.UnableToCreateCommentError,
	}, nil
}
