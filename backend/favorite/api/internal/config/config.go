package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Favorite zrpc.RpcClientConf // 手动代码
	Auth     struct {
		AccessSecret string
	}
	RecommendUrl string
}
