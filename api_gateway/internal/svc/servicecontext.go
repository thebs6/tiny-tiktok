package svc

import (
	"net/http"
	"net/url"
	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/service/favorite/pb/favorite"
	"tiny-tiktok/service/feed/pb/feed"
	"tiny-tiktok/service/relation/relation"
	"tiny-tiktok/service/relation/relationservice"
	"tiny-tiktok/service/user/pb/user"
	"tiny-tiktok/service/user/userservice"

	"tiny-tiktok/service/publish/pb/publish"

	"tiny-tiktok/api_gateway/internal/redis_model"

	"tiny-tiktok/service/comment/pb/comment"

	"github.com/redis/go-redis/v9"
	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc relationservice.Service
	Redis       redis_model.RedisModel
	CosClient   *cos.Client
	FeedRpc     feed.FeedServiceClient
	UserRpc     userservice.UserService
	PublishRpc  publish.PublishServiceClient
	CommentRpc  comment.CommentServiceClient
	FavoriteRpc favorite.FavoriteServiceClient
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
		FeedRpc:     feed.NewFeedServiceClient(zrpc.MustNewClient(c.FeedRpcConf).Conn()),
		UserRpc:     user.NewUserServiceClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
		PublishRpc:  publish.NewPublishServiceClient(zrpc.MustNewClient(c.PublishRpcConf).Conn()),
		CommentRpc:  comment.NewCommentServiceClient(zrpc.MustNewClient(c.CommentRpcConf).Conn()),
		FavoriteRpc: favorite.NewFavoriteServiceClient(zrpc.MustNewClient(c.FavoriteRpcConf).Conn()),
		RelationRpc: relation.NewServiceClient(zrpc.MustNewClient(c.RelationRpcConf).Conn()),
	}
}
