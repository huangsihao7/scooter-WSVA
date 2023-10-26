package main

import (
	"flag"
	"fmt"

	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/config"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/handler"
	"github.com/huangsihao7/scooter-WSVA/feed/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/feed-api.yaml", "the config file")

func main() {
	flag.Parse()
	//dfwsf
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
