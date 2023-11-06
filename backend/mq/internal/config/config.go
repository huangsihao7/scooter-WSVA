package config

import (
	"github.com/huangsihao7/scooter-WSVA/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	KqConsumerConf    kq.KqConf
	Comment           zrpc.RpcClientConf
	Label             zrpc.RpcClientConf
	Feed              zrpc.RpcClientConf
	KqConsumerJobConf kq.KqConf
	KqPusherJobConf   struct {
		Brokers []string
		Topic   string
	}
	Es struct {
		Addresses []string
		Username  string
		Password  string
	}
	KqVideoConsumerConf kq.KqConf
	KqPusherConf        struct {
		Brokers []string
		Topic   string
	}
	RecommendUrl string
	AIUrl        string
	SecretKey    string
	AccessKey    string
}
