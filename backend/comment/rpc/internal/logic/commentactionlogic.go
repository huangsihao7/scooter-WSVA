package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/model"
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
