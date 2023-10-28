package svc

import (
	"github.com/huangsihao7/scooter-WSVA/recommend/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
