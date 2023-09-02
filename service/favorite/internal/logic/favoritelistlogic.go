package logic

import (
	"context"

	"tiny-tiktok/service/favorite/internal/svc"
	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favorite.FavoriteListReq) (*favorite.FavoriteListResp, error) {
	// todo: add your logic here and delete this line

	return &favorite.FavoriteListResp{}, nil
}
