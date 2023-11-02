package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/comment/code"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanMuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDanMuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanMuListLogic {
	return &DanMuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DanMuListLogic) DanMuList(in *comment.DanmuListRequest) (*comment.DanmuListResponse, error) {

	videoId := in.VideoId

	// 检查视频id 是否存在
	_, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("弹幕视频不存在")
			return nil, code.DanMuVideoIdEmptyError
		}
		return nil, err
	}

	res, err := l.svcCtx.DanmuModel.FindsByVideoId(l.ctx, videoId)
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, err
	}
	danmuList := make([]*comment.DanMu, 0)
	for i := 0; i < len(res); i++ {
		singleinfo := &comment.DanMu{
			UserId:    int64(res[i].Uid),
			VideoId:   int64(res[i].Vid),
			DanmuText: res[i].Content,
			SendTime:  float32(res[i].SendTime),
		}
		danmuList = append(danmuList, singleinfo)
	}

	return &comment.DanmuListResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		DanmuList:  danmuList,
	}, nil
}
