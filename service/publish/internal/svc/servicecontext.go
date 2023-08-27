package svc

import (
	"tiny-tiktok/service/publish/internal/config"

	"tiny-tiktok/service/publish/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
