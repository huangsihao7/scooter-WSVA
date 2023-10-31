package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"

	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDanmuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDanmuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDanmuListLogic {
	return &GetDanmuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDanmuListLogic) GetDanmuList(req *types.DanmulistReq) (resp *types.DanmulistResp, err error) {
	//token 解析
	//usrId, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.Commenter.DanMuList(l.ctx, &comment.DanmuListRequest{VideoId: req.VideoId})
	resLists := make([]types.DanmuInfo, 0)

	for i := 0; i < len(res.DanmuList); i++ {

		danMuDetail := types.DanmuInfo{
			UserId:   res.DanmuList[i].UserId,
			VideoId:  res.DanmuList[i].VideoId,
			Content:  res.DanmuList[i].DanmuText,
			SendTime: res.DanmuList[i].SendTime,
		}
		resLists = append(resLists, danMuDetail)

	}
	return &types.DanmulistResp{
		StatusCode: int(res.StatusCode),
		StatusMsg:  res.StatusMsg,
		DanmuList:  resLists,
	}, nil
}
