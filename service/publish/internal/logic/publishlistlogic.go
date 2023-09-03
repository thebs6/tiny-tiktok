package logic

import (
	"context"
	"net/http"

	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/pb/publish"
	"tiny-tiktok/service/user/pb/user"

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
	videoList, err := l.svcCtx.VideoModel.ListByUserId(l.ctx, in.UserId)
	if err != nil {
		logc.Alert(l.ctx, "DB List failed "+err.Error())
		return nil, err
	}
	userIdList := make([]int64, len(videoList))
	for i, video := range videoList {
		userIdList[i] = video.Author
	}
	// userList := make([]*publish.User, len(videoList))
	// queryUsersByIds(userIdList, userList)
	userResp, err := l.svcCtx.UserRpc.UserInfoList(l.ctx, &user.UserInfoListReq{
		UserIdList: userIdList,
	})
	if err != nil {
		logc.Alert(l.ctx, "UserRpc UserInfoList failed "+err.Error())
		return &publish.PublishListResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Get publish list failed",
			VideoList:  nil,
		}, nil
	}

	videos := make([]*publish.Video, len(videoList))
	for i, v := range videoList {
		videos[i] = &publish.Video{
			Id: v.Id,
			// Author:        userList[i],
			Author: &publish.User{
				Id:              userResp.UserList[i].Id,
				Name:            userResp.UserList[i].Name,
				FollowCount:     userResp.UserList[i].FollowCount,
				FollowerCount:   userResp.UserList[i].FollowerCount,
				IsFollow:        userResp.UserList[i].IsFollow,
				Avatar:          userResp.UserList[i].Avatar,
				BackgroundImage: userResp.UserList[i].BackgroundImage,
				Signature:       userResp.UserList[i].Signature,
				TotalFavorited:  userResp.UserList[i].TotalFavorited,
				WorkCount:       userResp.UserList[i].WorkCount,
				FavoriteCount:   userResp.UserList[i].FavoriteCount,
			},
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
