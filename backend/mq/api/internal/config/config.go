package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	KqConsumerConf kq.KqConf
}
