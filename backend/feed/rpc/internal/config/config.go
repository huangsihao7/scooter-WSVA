package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataSource       string // 手动代码
	UserRpc          zrpc.RpcClientConf
	KqPusherTesTConf struct {
		Brokers []string
		Topic   string
	}
	DB struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	KqPusherJobConf struct {
		Brokers []string
		Topic   string
	}

	Es struct {
		Addresses []string
		Username  string
		Password  string
	}
	SecretKey    string
	AccessKey    string
	AIUrl        string
	RecommendUrl string
}
