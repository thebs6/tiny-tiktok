package logic

import (
	"context"
	"fmt"

	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/pb/publish"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListResp, error) {
	fmt.Println(in.UserId)
	videoList, err := l.svcCtx.VideoModel.List(l.ctx, in.UserId)
	fmt.Println(len(videoList))
	if err != nil {
		logc.Alert(l.ctx, "DB List failed")
		return nil, err
	}
	userIdList := make([]int64, len(videoList))
	for i, video := range videoList {
		userIdList[i] = video.Author
	}
	userList := make([]*publish.User, len(videoList))
	queryUsersByIds(userIdList, userList)

	videos := make([]*publish.Video, len(videoList))
	for i, v := range videoList {
		videos[i] = &publish.Video{
			Id:            v.Id,
			Author:        userList[i],
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
			Title:         v.Title,
		}
	}

	return &publish.PublishListResp{
		StatusMsg: "Get publish list succesfully",
		VideoList: videos,
	}, nil
}

// stub code
func queryUsersByIds(userIds []int64, users []*publish.User) {
	for i, userId := range userIds {
		users[i] = &publish.User{
			Id:            userId,
			Name:          "gao",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
	}
}
