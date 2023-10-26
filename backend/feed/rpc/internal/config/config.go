package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataSource string // 手动代码
	UserRpc    zrpc.RpcClientConf
}
