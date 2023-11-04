package svc

import (
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/huangsihao7/scooter-WSVA/relation/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/relation/rpc/relationclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc relationclient.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	RelationRPC := zrpc.MustNewClient(c.RelationRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:      c,
		RelationRpc: relationclient.NewRelation(RelationRPC),
	}
}
