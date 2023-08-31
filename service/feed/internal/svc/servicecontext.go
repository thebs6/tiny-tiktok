package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tiny-tiktok/service/feed/internal/config"
	model "tiny-tiktok/service/feed/internal/model/genModel"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	UserModel  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
		UserModel:  model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
