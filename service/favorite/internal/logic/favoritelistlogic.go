package logic

import (
	"context"

	"tiny-tiktok/service/favorite/internal/svc"
	"tiny-tiktok/service/favorite/pb/favorite"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logc"
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

	favorites, err := l.svcCtx.FavoriteModel.ListByUserId(l.ctx, in.UserId)
	if err != nil {
		logc.Alert(l.ctx, "ListByUserId() err: "+err.Error())
	}
	videoIdList := make([]int64, len(favorites))
	for i, f := range favorites {
		videoIdList[i] = f.VideoId
	}

	videoList := make([]*favorite.Video, len(favorites))
	respRpc, err := l.svcCtx.PublishRpc.VideoList(l.ctx, &publish.VideoListReq{
		VideoIdList: videoIdList,
	})
	if err != nil {
		logc.Alert(l.ctx, "Rpc videoList() error: "+err.Error())
		return &favorite.FavoriteListResp{
			StatusCode: 0,
			StatusMsg:  "fail to get favorite list",
			VideoList:  nil,
		}, err
	}
	for i, v := range respRpc.VideoList {
		videoList[i] = &favorite.Video{
			Id: v.Id,
			Author: &favorite.User{
				Id:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				TotalFavorited:  v.Author.TotalFavorited,
				WorkCount:       v.Author.WorkCount,
				FavoriteCount:   v.Author.FavoriteCount,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
			Title:         v.Title,
		}

	}

	return &favorite.FavoriteListResp{
		StatusCode: 0,
		StatusMsg:  "success to get favorite list",
		VideoList:  videoList,
	}, nil
}
