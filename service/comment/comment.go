package main

import (
	"flag"
	"fmt"

	"tiny-tiktok/service/comment/internal/config"
	"tiny-tiktok/service/comment/internal/server"
	"tiny-tiktok/service/comment/internal/svc"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comment.RegisterCommentServiceServer(grpcServer, server.NewCommentServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// logconf := logc.LogConf{
	// 	ServiceName: c.LOG.ServiceName,
	// 	Mode:        c.LOG.Mode,
	// }
	// logc.MustSetup(logconf)

	fmt.Printf("Starting comment service rpc server at %s...\n", c.ListenOn)
	s.Start()
}
