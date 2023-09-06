package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/api_gateway/internal/handler"
	"tiny-tiktok/api_gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
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

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		// 自定义处理返回
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	var logconf logc.LogConf

	_ = conf.FillDefault(&logconf)
	logconf.ServiceName = c.Log.ServiceName
	logconf.Mode = c.Log.Mode
	logconf.Path = c.Log.Path

	logc.MustSetup(logconf)
	defer logc.Close()
	logc.Info(context.Background(), "api_gateway start")

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
