// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package favoriteservice

import (
	"context"

	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FavoriteActionReq  = favorite.FavoriteActionReq
	FavoriteActionResp = favorite.FavoriteActionResp
	FavoriteListReq    = favorite.FavoriteListReq
	FavoriteListResp   = favorite.FavoriteListResp
	User               = favorite.User
	Video              = favorite.Video

	FavoriteService interface {
		FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error)
		FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
	}

	defaultFavoriteService struct {
		cli zrpc.Client
	}
)

func NewFavoriteService(cli zrpc.Client) FavoriteService {
	return &defaultFavoriteService{
		cli: cli,
	}
}

func (m *defaultFavoriteService) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	client := favorite.NewFavoriteServiceClient(m.cli.Conn())
	return client.FavoriteAction(ctx, in, opts...)
}

func (m *defaultFavoriteService) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	client := favorite.NewFavoriteServiceClient(m.cli.Conn())
	return client.FavoriteList(ctx, in, opts...)
}
