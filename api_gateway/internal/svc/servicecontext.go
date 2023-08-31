package svc

import (
	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/service/relation/relation"
	"tiny-tiktok/service/relation/relationservice"
	"tiny-tiktok/service/user/pb/user"
	"tiny-tiktok/service/user/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userservice.UserService
	RelationRpc relationservice.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
		RelationRpc: relation.NewServiceClient(zrpc.MustNewClient(c.RelationRpcConf).Conn()),
	}
}
