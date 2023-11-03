package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/comment/code"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
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

	if req.ActionType == 1 && req.CommentText == "" {
		return nil, code.CommentIsEmptyError
	}
	usrId, _ := l.ctx.Value("uid").(json.Number).Int64()

	if req.ActionType == 1 {
		format.Feedback(l.svcCtx.Config.RecommendUrl, "comment", int(req.VideoId), int(usrId))
	}
	res, err := l.svcCtx.Commenter.CommentAction(l.ctx, &comment.CommentActionRequest{
		UserId:      usrId,
		ActionType:  req.ActionType,
		VideoId:     req.VideoId,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})
	if err != nil {
		return nil, err
	}

	return &types.ActionResp{
		CommentId:  res.CommentId,
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil
}
