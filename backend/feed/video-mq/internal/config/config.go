package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	KqConsumerConf kq.KqConf
	Es             struct {
		Addresses []string
		Username  string
		Password  string
	}
	Datasource string
	UserRpc    zrpc.RpcClientConf
}
