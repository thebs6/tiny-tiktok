package svc

import (
	"tiny-tiktok/service/comment/internal/config"
	"tiny-tiktok/service/comment/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
