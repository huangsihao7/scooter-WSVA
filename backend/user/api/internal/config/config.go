package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
	//User      zrpc.RpcClientConf
	AccessKey string
	SecretKey string
	Bucket    string
	OssUrl    string  `json:",env=OSS_URL"`
}
