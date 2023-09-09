package extra_first

import (
	"context"
	"net/http"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/favorite/pb/favorite"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	respRpc, err := l.svcCtx.FavoriteRpc.FavoriteList(l.ctx, &favorite.FavoriteListReq{
		UserId: req.UserID,
	})
	if err != nil {
		logc.Errorf(l.ctx, "rpc favoritelist failed, err: "+err.Error())
		return &types.FavoriteListResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "点赞列表获取失败!",
		}, nil
	}

	var videoList = make([]types.Video, len(resp.VideoList))
	for i, video := range respRpc.VideoList {
		videoList[i] = types.Video{
			ID: video.Id,
			Author: types.User{
				Avatar:          video.Author.Avatar,
				BackgroundImage: video.Author.BackgroundImage,
				FavoriteCount:   video.Author.FavoriteCount,
				FollowCount:     video.Author.FollowCount,
				FollowerCount:   video.Author.FollowCount,
				Id:              video.Author.Id,
				IsFollow:        video.Author.IsFollow,
				Name:            video.Author.Name,
				Signature:       video.Author.Signature,
				TotalFavorited:  video.Author.TotalFavorited,
				WorkCount:       video.Author.WorkCount,
			},
			CommentCount:  video.CommentCount,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    video.IsFavorite,
			PlayURL:       video.PlayUrl,
			Title:         video.Title,
		}
	}

	return &types.FavoriteListResp{
		StatusCode: http.StatusOK,
		StatusMsg:  "点赞列表获取成功!",
		VideoList:  videoList,
	}, nil
}
