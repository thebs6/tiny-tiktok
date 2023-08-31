package svc

import (
	"tiny-tiktok/service/relation/internal/config"
	"tiny-tiktok/service/relation/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	RelationModel model.RelationModel
	UserModel model.UserModel

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RelationModel: model.NewRelationModel(sqlx.NewMysql(c.DB.DataSource)),
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
