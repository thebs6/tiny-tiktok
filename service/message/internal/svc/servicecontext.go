package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiny-tiktok/service/message/internal/config"
	model "tiny-tiktok/service/message/internal/genModel"
)

type ServiceContext struct {
	Config config.Config
	Model  model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewMessageModel(sqlx.NewMysql(c.DB.Datasource)),
	}
}
