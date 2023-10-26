package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/common/constants"

	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.ActionReq) (resp *types.ActionResp, err error) {
	// todo: add your logic here and delete this line

	//验证token是否有效
	//token 解析

	usrId := 10
	_, err = l.svcCtx.Commenter.CommentAction(l.ctx, &comment.CommentActionRequest{
		UserId:      int64(usrId),
		ActionType:  req.ActionType,
		VideoId:     req.VideoId,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})
	if err != nil {
		return &types.ActionResp{
			StatusCode: constants.UnableToQueryCommentErrorCode,
			StatusMsg:  constants.UnableToCreateCommentError,
		}, err
	}
	output := &types.ActionResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
	}
	return output, nil
}
