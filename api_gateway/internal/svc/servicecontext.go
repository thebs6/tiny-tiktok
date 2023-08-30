package svc

import (
	"net/http"
	"net/url"
	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/service/publish/pb/publish"

	"tiny-tiktok/api_gateway/internal/redis_model"

	"github.com/redis/go-redis/v9"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Redis     redis_model.RedisModel
	CosClient *cos.Client
	Publish   publish.PublishServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	// tencent COS
	BucketURL, _ := url.Parse(c.Cos.URL)
	baseurl := &cos.BaseURL{BucketURL: BucketURL}
	cosclient := cos.NewClient(baseurl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.Cos.SecretId,
			SecretKey: c.Cos.SecretKey,
		},
	})

	return &ServiceContext{
		Config:    c,
		CosClient: cosclient,
		Redis: redis_model.NewRedisModel(redis.NewClient(&redis.Options{
			Addr:     c.RedisConf.Host,
			Password: c.RedisConf.Pass,
		})),
		Publish: publish.NewPublishServiceClient(zrpc.MustNewClient(c.Publish).Conn()),
	}
}
