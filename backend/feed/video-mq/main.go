package main

import (
	"context"
	"flag"

	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/config"
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/logic"
	"github.com/huangsihao7/scooter-WSVA/feed/video-mq/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/video-mq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
