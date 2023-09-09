package logic

import (
	"context"

	"tiny-tiktok/service/publish/internal/svc"
	"tiny-tiktok/service/publish/pb/publish"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoListLogic) VideoList(in *publish.VideoListReq) (*publish.VideoListResp, error) {
	videoList := make([]*publish.Video, len(in.VideoIdList))
	authorIdList := make([]int64, len(in.VideoIdList))
	for i, id := range in.VideoIdList {
		v, err := l.svcCtx.VideoModel.FindOne(l.ctx, id)
		if err != nil {
			return &publish.VideoListResp{
				StatusCode: 0,
				StatusMsg:  "fail to get full video list",
				VideoList:  videoList,
			}, err
		}
		authorIdList[i] = v.Author
		videoList[i] = &publish.Video{
			Id:            v.Id,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
		}
	}
	respRpc, err := l.svcCtx.UserRpc.UserInfoList(l.ctx, &user.UserInfoListReq{
		UserIdList: authorIdList,
	})
	if err == nil {
		return &publish.VideoListResp{
			StatusCode: 0,
			StatusMsg:  "fail to get full video list",
			VideoList:  videoList,
		}, err
	}

	for i, u := range respRpc.UserList {
		videoList[i].Author = &publish.User{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        u.IsFollow,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
	}

	return &publish.VideoListResp{
		StatusCode: 0,
		StatusMsg:  "success to get video list",
		VideoList:  videoList,
	}, nil
}
