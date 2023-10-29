package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/mq/format"

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

	//验证token是否有效
	//token 解析
	usrId, _ := l.ctx.Value("uid").(json.Number).Int64()
	if req.ActionType == 1 {
		format.Feedback("comment", int(req.VideoId), int(usrId))
	}
	_, err = l.svcCtx.Commenter.CommentAction(l.ctx, &comment.CommentActionRequest{
		UserId:      usrId,
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
