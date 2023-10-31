package logic

import (
	"context"
	"encoding/json"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuActionLogic {
	return &DanmuActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuActionLogic) DanmuAction(req *types.DanmuActionReq) (resp *types.DanmuActionResp, err error) {
	// todo: add your logic here and delete this line
	//验证token是否有效
	//token 解析
	userId, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.Commenter.DanMuAction(l.ctx, &comment.DanmuActionRequest{
		UserId:    userId,
		VideoId:   req.VideoId,
		DanmuText: req.DanmuText,
		SendTime:  strconv.FormatInt(req.SendTime, 10),
	})
	if err != nil {
		return &types.DanmuActionResp{
			StatusCode: constants.UnableToQueryCommentErrorCode,
			StatusMsg:  constants.UnableToCreateCommentError,
		}, err
	}
	return &types.DanmuActionResp{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil
}
