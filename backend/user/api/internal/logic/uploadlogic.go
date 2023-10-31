package logic

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"
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

//func (l *UploadLogic) Upload(req *http.Request) (resp *types.UserUploadResponse, err error) {
//	// todo: add your logic here and delete this line
//	//req.ParseMultipartForm(200 * 1024 * 1024)
//	file, handler, err := req.FormFile("file")
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	// 创建七牛云的上传凭证
//	putPolicy := storage.PutPolicy{
//		Scope: l.svcCtx.Config.Bucket,
//	}
//	mac := qbox.NewMac(l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey)
//	uploadToken := putPolicy.UploadToken(mac)
//
//	// 设置上传配置
//	cfg := storage.Config{}
//	//formUploader := storage.NewFormUploader(&cfg)
//	formUploader := storage.NewResumeUploaderV2(&cfg)
//	ret := storage.PutRet{}
//	putExtra := storage.RputV2Extra{}
//	// 上传文件
//
//	err = formUploader.PutWithoutKey(context.Background(), &ret, uploadToken, file, handler.Size, &putExtra)
//	if err != nil {
//		return nil, err
//	}
//	baseURL := "s327crbzf.hn-bkt.clouddn.com" // 替换为您的文件访问域名
//	fileURL := baseURL + "/" + ret.Key
//	//deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
//	//privateAccessURL := storage.MakePrivateURL(mac, baseURL, ret.Key, deadline)
//	// 返回上传成功的信息
//	return &types.UserUploadResponse{
//		Url: fileURL,
//	}, nil
//}

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
		PersistentOps:       "avthumb/mp4/wmImage/aHR0cDovL3Rlc3QtMi5xaW5pdWRuLmNvbS9sb2dvLnBuZw==/wmGravity/NorthWest|saveas/" + saceas,
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
		return
	}
	operationManager := storage.NewOperationManager(mac, &cfg)
	fopVframe := fmt.Sprintf("vframe/jpg/offset/1|saveas/%s",
		storage.EncodedEntry(bucket, strings.TrimSuffix(key, filepath.Ext(key))+".jpg"))

	fops := fopVframe

	_, err = operationManager.Pfop(bucket, key, fops, "", "", true)
	if err != nil {
		fmt.Println(err)
		return
	}

	baseURL := "http://s327crbzf.hn-bkt.clouddn.com"
	fileURL = baseURL + "/" + key
	coverUrl := baseURL + "/" + strings.TrimSuffix(key, filepath.Ext(key)) + ".jpg"
	println(ret.PersistentID)

	return &types.UserUploadResponse{
		CoverUrl: coverUrl,
		Url:      fileURL,
	}, nil
}
