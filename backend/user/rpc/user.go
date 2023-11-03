package main

import (
	"flag"
	"fmt"
	"github.com/huangsihao7/scooter-WSVA/pkg/interceptors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/config"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/server"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/internal/svc"
	"github.com/huangsihao7/scooter-WSVA/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user-dev.yaml", "the config file")

// user rpc
func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	c.Timeout = 20000
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUseSrvServer(grpcServer, server.NewUseSrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	fmt.Printf("Starting user rpc server at %s...\n", c.ListenOn)
	s.Start()
}
