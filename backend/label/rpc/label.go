package main

import (
	"flag"
	"fmt"

	"github.com/huangsihao7/scooter-WSVA/label/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/internal/server"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/label/rpc/label"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/label.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		label.RegisterLabelServer(grpcServer, server.NewLabelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}