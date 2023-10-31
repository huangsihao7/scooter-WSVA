package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"strconv"

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
	resLists := make([][]interface{}, 0)

	for i := 0; i < len(res.DanmuList); i++ {

		intTime, _ := strconv.Atoi(res.DanmuList[i].SendTime)
		details := make([]interface{}, 0)
		details = append(details, intTime)
		details = append(details, 0)
		details = append(details, 16777215)
		details = append(details, "yuding")
		details = append(details, res.DanmuList[i].DanmuText)

		resLists = append(resLists, details)
	}
	return &types.DanmulistResp{
		StatusCode: 0,
		DanmuList:  resLists,
	}, nil
}
