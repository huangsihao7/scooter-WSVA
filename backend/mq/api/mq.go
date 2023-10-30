package main

import (
	"context"
	"flag"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/mqs"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/mq/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mq-api-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	//循环
	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
