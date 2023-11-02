package logic

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/constants"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"
	"github.com/huangsihao7/scooter-WSVA/user/code"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *http.Request) (resp *types.UserUploadResponse, err error) {
	accessKey := l.svcCtx.Config.AccessKey
	secretKey := l.svcCtx.Config.SecretKey
	bucket := l.svcCtx.Config.Bucket
	fileURL := ""
	file, handler, err := req.FormFile("file")
	//key为上传的文件名
	filename := crypt.PasswordEncrypt(time.Now().String(), handler.Filename)
	key := filename + ".mp4"
	seveas := filename + "-1.mp4"
	seveasstr := l.svcCtx.Config.Bucket + ":" + seveas
	saceas := base64.StdEncoding.EncodeToString([]byte(seveasstr))
	mac := qbox.NewMac(accessKey, secretKey)
	//path := filepath.Join("./", fileName)

	putPolicy := storage.PutPolicy{
		Scope:               bucket,
		PersistentOps:       "avthumb/mp4/wmImage/aHR0cDovL3MzZWplbjBnMi5oZC1ia3QuY2xvdWRkbi5jb20vbG9nLnBuZw==/wmGravity/NorthWest|saveas/" + saceas,
		PersistentNotifyURL: "http://fake.com/qiniu/notify",
	}
	upToken := putPolicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)

	ret := storage.PutRet{}           // 上传后返回的结果
	putExtra := storage.RputV2Extra{} // 额外参数

	err = resumeUploader.Put(context.Background(), &ret, upToken, key, file, handler.Size, &putExtra)
	if err != nil {
		return nil, code.UserUploadVideoError
	}
	operationManager := storage.NewOperationManager(mac, &cfg)
	fopVframe := fmt.Sprintf("vframe/jpg/offset/1|saveas/%s",
		storage.EncodedEntry(bucket, strings.TrimSuffix(key, filepath.Ext(key))+".jpg"))
	fops := fopVframe
	_, err = operationManager.Pfop(bucket, key, fops, "", "", true)
	if err != nil {
		l.Logger.Error("截帧失败")
		return nil, code.UserUploadVideoError
	}

	baseURL := l.svcCtx.Config.OssUrl
	fileURL = baseURL + "/" + key
	coverUrl := baseURL + "/" + strings.TrimSuffix(key, filepath.Ext(key)) + ".jpg"
	println(ret.PersistentID)

	return &types.UserUploadResponse{
		StatusCode: constants.ServiceOKCode,
		StatusMsg:  constants.ServiceOK,
		CoverUrl:   coverUrl,
		Url:        fileURL,
	}, nil
}
