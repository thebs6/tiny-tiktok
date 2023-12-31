// Code generated by goctl. DO NOT EDIT.
// Source: publish.proto

package publishservice

import (
	"context"

	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PublishActionReq  = publish.PublishActionReq
	PublishActionResp = publish.PublishActionResp
	PublishListReq    = publish.PublishListReq
	PublishListResp   = publish.PublishListResp
	User              = publish.User
	Video             = publish.Video
	VideoListReq      = publish.VideoListReq
	VideoListResp     = publish.VideoListResp

	PublishService interface {
		PublishAction(ctx context.Context, in *PublishActionReq, opts ...grpc.CallOption) (*PublishActionResp, error)
		PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
		VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error)
	}

	defaultPublishService struct {
		cli zrpc.Client
	}
)

func NewPublishService(cli zrpc.Client) PublishService {
	return &defaultPublishService{
		cli: cli,
	}
}

func (m *defaultPublishService) PublishAction(ctx context.Context, in *PublishActionReq, opts ...grpc.CallOption) (*PublishActionResp, error) {
	client := publish.NewPublishServiceClient(m.cli.Conn())
	return client.PublishAction(ctx, in, opts...)
}

func (m *defaultPublishService) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	client := publish.NewPublishServiceClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}

func (m *defaultPublishService) VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error) {
	client := publish.NewPublishServiceClient(m.cli.Conn())
	return client.VideoList(ctx, in, opts...)
}
