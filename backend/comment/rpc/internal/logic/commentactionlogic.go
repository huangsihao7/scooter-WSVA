package logic

import (
	"context"
	"errors"
	"github.com/huangsihao7/scooter-WSVA/comment/gmodel"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	constants "github.com/huangsihao7/scooter-WSVA/common/constants"
	gmodel3 "github.com/huangsihao7/scooter-WSVA/feed/gmodel"
	"gorm.io/gorm"
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
	logx.DisableStat()
	// todo: add your logic here and delete this line
	userId := in.UserId
	videoId := in.VideoId
	contents := in.CommentText
	actionType := in.ActionType
	commentId := in.CommentId

	//检查用户id 是否能存在
	_, err := l.svcCtx.UserModel.GetUserByID(l.ctx, uint(in.UserId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("用户不存在")
			return nil, errors.New("评论用户不存在")
		}
		return nil, err
	}
	// 检查视频id 是否存在
	videoDetail, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("视频不存在")
			return nil, errors.New("评论视频不存在")
		}
		return nil, err
	}

	switch actionType {
	//添加评论
	case 1:
		newComment := gmodel.Comments{
			Uid:     uint(userId),
			Vid:     uint(videoId),
			Content: contents,
		}
		err := l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		//添加video的评论数
		err = l.svcCtx.VideoModel.Update(l.ctx, &gmodel3.Videos{
			Id:            uint(videoId),
			AuthorId:      videoDetail.AuthorId,
			Title:         videoDetail.Title,
			CoverUrl:      videoDetail.CoverUrl,
			PlayUrl:       videoDetail.PlayUrl,
			FavoriteCount: videoDetail.FavoriteCount,
			CommentCount:  videoDetail.CommentCount + 1,
			Category:      videoDetail.Category,
			CreatedAt:     videoDetail.CreatedAt,
			Duration:      videoDetail.Duration,
		})
		if err != nil {
			log.Println(err.Error())
			return &comment.CommentActionResponse{
				StatusCode: constants.UnableToCreateCommentErrorCode,
				StatusMsg:  constants.UnableToCreateCommentError,
			}, nil
		}
		return &comment.CommentActionResponse{
			StatusCode: constants.ServiceOKCode,
			StatusMsg:  constants.ServiceOK,
		}, nil
		//删除评论
	case 2:
		err := l.svcCtx.CommentModel.Delete(l.ctx, commentId)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		//减少video的评论数
		err = l.svcCtx.VideoModel.Update(l.ctx, &gmodel3.Videos{
			Id:            uint(videoId),
			AuthorId:      videoDetail.AuthorId,
			Title:         videoDetail.Title,
			CoverUrl:      videoDetail.CoverUrl,
			PlayUrl:       videoDetail.PlayUrl,
			FavoriteCount: videoDetail.FavoriteCount,
			CommentCount:  videoDetail.CommentCount - 1,
			Category:      videoDetail.Category,
			CreatedAt:     videoDetail.CreatedAt,
			Duration:      videoDetail.Duration,
		})
		if err != nil {
			log.Println(err.Error())
			return &comment.CommentActionResponse{
				StatusCode: constants.UnableToCreateCommentErrorCode,
				StatusMsg:  constants.UnableToCreateCommentError,
			}, nil
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
