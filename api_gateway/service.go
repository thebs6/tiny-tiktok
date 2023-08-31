package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"

	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/api_gateway/internal/handler"
	"tiny-tiktok/api_gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "api_gateway/etc/service.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var c2 logc.LogConf
	c2.Encoding = "plain"
	logc.MustSetup(c2)

	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
