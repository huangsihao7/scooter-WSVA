package common

import (
	"context"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/common/crypt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UserUpload(accessKey, secretKey, bucket string, filepath string) (string, error) {
	fileURL := ""
	mac := qbox.NewMac(accessKey, secretKey)

	//path := filepath.Join("./", fileName)

	putPolicy := storage.PutPolicy{
		Scope: bucket,
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

	//key为上传的文件名
	key := crypt.PasswordEncrypt("wy", filepath) + ".jpg"
	go func() {
		err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, filepath, &putExtra)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			println("上传成功")
		}
	}()

	baseURL := "s327crbzf.hn-bkt.clouddn.com"
	fileURL = baseURL + "/" + key

	return fileURL, nil
}
