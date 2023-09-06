package svc

import (
	"tiny-tiktok/service/comment/internal/config"
	"tiny-tiktok/service/comment/internal/model"
	"tiny-tiktok/service/comment/internal/redismodel"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	CommentRedis redismodel.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentRedis: redismodel.NewCommentModel(redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		})),
	}
}
