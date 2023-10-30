package svc

import (
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
