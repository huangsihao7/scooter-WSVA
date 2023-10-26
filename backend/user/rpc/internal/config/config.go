package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource  string          // 手动代码
	Cache       cache.CacheConf // 手动代码
	Salt        string
	FavoriteRpc zrpc.RpcClientConf //手动代码
	RelationRpc zrpc.RpcClientConf //手动代码
}
