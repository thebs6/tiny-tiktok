package svc

import (
	"tiny-tiktok/service/publish/internal/config"
	"tiny-tiktok/service/user/pb/user"
	"tiny-tiktok/service/user/userservice"

	"tiny-tiktok/service/publish/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	UserRpc    userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRpc:    user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
	}
}
