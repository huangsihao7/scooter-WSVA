package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"gorm.io/gorm"
	"log"

	"github.com/huangsihao7/scooter-WSVA/comment/rpc/comment"
	"github.com/huangsihao7/scooter-WSVA/comment/rpc/internal/svc"

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
	logx.DisableStat()
	// todo: add your logic here and delete this line
	videoId := in.VideoId

	// 检查视频id 是否存在
	_, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("视频不存在")
			return &comment.DanmuListResponse{
				StatusCode: constants.UserVideosDoNotExistedCode,
				StatusMsg:  constants.FindUserVideosError,
			}, nil
		}
		return nil, err
	}

	res, err := l.svcCtx.DanmuModel.FindsByVideoId(l.ctx, videoId)
	if err != nil {
		logx.Infof(err.Error())
		return &comment.DanmuListResponse{
			StatusCode: constants.DanmuCanNotEmptyCode,
			StatusMsg:  constants.DanmuCanNotEmptyError,
		}, nil
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
