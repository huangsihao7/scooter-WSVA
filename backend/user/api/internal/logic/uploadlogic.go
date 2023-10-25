package logic

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/api/internal/types"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

// tip
const maxUploadSize = 100 * 1024 * 1024 // 10MB
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

func (l *UploadLogic) Upload(w http.ResponseWriter, req *http.Request) (resp *types.UserUploadResponse, err error) {
	// todo: add your logic here and delete this line
	req.Body = http.MaxBytesReader(w, req.Body, maxUploadSize)
	file, handler, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建七牛云的上传凭证
	putPolicy := storage.PutPolicy{
		Scope: l.svcCtx.Config.Bucket,
	}
	mac := qbox.NewMac(l.svcCtx.Config.AccessKey, l.svcCtx.Config.SecretKey)
	uploadToken := putPolicy.UploadToken(mac)

	// 设置上传配置
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件

	err = formUploader.PutWithoutKey(context.Background(), &ret, uploadToken, file, handler.Size, nil)
	if err != nil {
		return nil, err
	}
	baseURL := "s327crbzf.hn-bkt.clouddn.com" // 替换为您的文件访问域名
	fileURL := baseURL + "/" + ret.Key
	//deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	//privateAccessURL := storage.MakePrivateURL(mac, baseURL, ret.Key, deadline)
	// 返回上传成功的信息
	return &types.UserUploadResponse{
		Url: fileURL,
	}, nil
}
