package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/common"
	"net/http"

	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(req *http.Request) (resp *types.UploadImageResponse, err error) {
	imgFile, imgHandler, err := req.FormFile("img")
	imgUrl, err := common.Upload(imgFile, imgHandler, l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey, l.svcCtx.Config.Bucket)
	if err != nil {
		return &types.UploadImageResponse{
			StatusCode: constants.UploadErrCode,
			StatusMsg:  constants.UploadErr,
		}, nil
	}

	return &types.UploadImageResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		Url:        imgUrl,
	}, nil
}
