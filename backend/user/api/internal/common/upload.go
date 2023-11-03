package common

import (
	"context"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"time"
)

func Upload(file multipart.File, handler *multipart.FileHeader, AccessKey string, SecretKey string, bucket string, baseURL string) (string, error) {
	filename := crypt.PasswordEncrypt(time.Now().String(), handler.Filename)
	key := filename + ".jpg"
	defer file.Close()
	// 创建七牛云的上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	uploadToken := putPolicy.UploadToken(mac)

	// 设置上传配置
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.Put(context.Background(), &ret, uploadToken, key, file, handler.Size, nil)
	if err != nil {
		return "", err
	}
	fileURL := baseURL + "/" + key
	//deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	//privateAccessURL := storage.MakePrivateURL(mac, baseURL, ret.Key, deadline)
	// 返回上传成功的信息
	return fileURL, nil
}
