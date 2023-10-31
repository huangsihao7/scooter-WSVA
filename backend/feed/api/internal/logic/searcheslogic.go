package logic

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchEsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchEsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchEsLogic {
	return &SearchEsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchEsLogic) SearchEs(req *types.SearchEsReq) (resp *types.SearchEsResp, err error) {
	// todo: add your logic here and delete this line
	//uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.FeedRpc.SearchES(l.ctx, &feed.EsSearchReq{
		Content: req.Content,
	})
	fmt.Println(res)
	if err != nil {
		logx.Infof(err.Error())
	}
	return &types.SearchEsResp{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		VideoList:  nil,
	}, nil
}
