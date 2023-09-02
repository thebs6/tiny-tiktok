package svc

import (
	"tiny-tiktok/service/favorite/internal/config"
	"tiny-tiktok/service/favorite/internal/model"
	"tiny-tiktok/service/favorite/internal/redis_model"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	Redis         redis_model.RedisModel
	FavoriteModel model.FavoriteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis_model.NewRedisModel(redis.NewClient(&redis.Options{
			Addr:     c.RedisConf.Host,
			Password: c.RedisConf.Pass,
		})),
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
