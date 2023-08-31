package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	model "tiny-tiktok/service/feed/internal/model/genModel"
	"tiny-tiktok/service/feed/internal/svc"
	"tiny-tiktok/service/feed/pb/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *feed.FeedRequest) (*feed.FeedResponse, error) {
	// todo: add your logic here and delete this line
	//查询所有视频
	video, err := l.svcCtx.VideoModel.FindVideoList(l.ctx, in.LatestTime)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if video == nil {
		return nil, errors.New("id不存在")
	}
	var VideoList []*feed.Video
	var video2 feed.Video
	//查询所有作者
	for _, ele := range video {
		//封装作者
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, ele.Author)
		if err != nil || err == model.ErrNotFound {
			return nil, errors.New("作者查询失败")
		}
		if user == nil {
			return nil, errors.New("id不存在")
		}
		var user2 feed.User
		copier.Copy(&user2, &user)
		user2.Name = user.Username

		//将作者封装到视频，并拼接到视频列表里
		copier.Copy(&video2, &ele)
		video2.Author = &user2

		VideoList = append(VideoList, &video2)
	}
	logx.Info("封装成功")

	nextTime := video[0].CreatedAt.Unix()

	return &feed.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "查找成功",
		VideoList:  VideoList,
		NextTime:   nextTime,
	}, nil
}
