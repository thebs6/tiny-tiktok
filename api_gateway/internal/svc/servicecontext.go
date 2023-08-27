package svc

import (
	"net/http"
	"net/url"
	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type ServiceContext struct {
	Config    config.Config
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
		// Publish:   publish.NewPublishServiceClient(zrpc.MustNewClient(c.Publish).Conn()),
	}
}
