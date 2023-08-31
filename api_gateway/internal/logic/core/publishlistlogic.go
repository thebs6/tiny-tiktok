package core

import (
	"context"
	"net/http"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (resp *types.PublishListResp, err error) {
	respRpc, err := l.svcCtx.Publish.PublishList(l.ctx, &publish.PublishListReq{
		UserId: req.UserID,
	})
	if err != nil {
		logx.Error("svc.Publish.PublishList failed", err)
		return &types.PublishListResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Get publish list failed",
			VideoList:  nil,
		}, nil
	}

	videos := make([]types.Video, len(respRpc.VideoList))
	for i, v := range respRpc.VideoList {
		videos[i] = types.Video{
			Author: types.User{
				Id:   v.Author.Id,
				Name: v.Author.Name,
			},
			CommentCount:  v.CommentCount,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			ID:            v.Id,
			IsFavorite:    v.IsFavorite,
			PlayURL:       v.PlayUrl,
			Title:         v.Title,
		}
	}
	resp = &types.PublishListResp{
		StatusCode: http.StatusOK,
		StatusMsg:  respRpc.StatusMsg,
		VideoList:  videos,
	}
	return
}
