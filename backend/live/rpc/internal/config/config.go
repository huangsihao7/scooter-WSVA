package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	UserRpc      zrpc.RpcClientConf
	Consul       consul.Conf
	LiveBucket   string
	PublishUrl   string
	LiveUrl      string
	LiveCoverUrl string
	SecretKey    string
	AccessKey    string
}
