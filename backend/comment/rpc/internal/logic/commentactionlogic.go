package logic

import (
	"context"
	"errors"
	"github.com/huangsihao7/scooter-WSVA/comment/model"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	constants "github.com/huangsihao7/scooter-WSVA/common/constants"
	"log"

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
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("用户不存在")
			return nil, errors.New("评论用户不存在")
		}
		return nil, err
	}
	// 检查视频id 是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == model.ErrNotFound {
			log.Println("视频不存在")
			return nil, errors.New("评论视频不存在")
		}
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
