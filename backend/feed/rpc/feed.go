package main

import (
	"flag"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huangsihao7/scooter-WSVA/feed/rpc/feed"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/server"
	"github.com/huangsihao7/scooter-WSVA/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/feed-dev.yaml", "the config file")

// feed
func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Timeout = 30000
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		feed.RegisterFeedServer(grpcServer, server.NewFeedServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
