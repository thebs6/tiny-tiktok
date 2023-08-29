package svc

import (
	"tiny-tiktok/api_gateway/internal/config"
	"tiny-tiktok/service/comment/pb/comment"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Comment comment.CommentServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Comment: comment.NewCommentServiceClient(zrpc.MustNewClient(c.Comment).Conn()),
	}
}
