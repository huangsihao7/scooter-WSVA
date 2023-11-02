package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Salt string
	//FavoriteRpc zrpc.RpcClientConf //手动代码 fweqfqw
	//RelationRpc zrpc.RpcClientConf //手动代码
	DB struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
}
