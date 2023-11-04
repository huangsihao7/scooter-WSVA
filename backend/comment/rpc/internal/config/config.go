package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string             // 手动代码
	Cache      cache.CacheConf    // 手动代码
	UserRpc    zrpc.RpcClientConf //手动代码
	DB         struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	AIUrl  string
	Consul consul.Conf
}
