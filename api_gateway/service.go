package main

import (
	"flag"
	"fmt"

	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/api_gateway/internal/handler"
	"tiny-tiktok/api_gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/service.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logconf := logc.LogConf{
		Mode:        c.Log.Mode,
		ServiceName: c.Log.ServiceName,
		Path:        c.Log.Path,
	}
	logc.MustSetup(logconf)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
