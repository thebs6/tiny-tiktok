package logic

import (
	"context"
	relationModel "tiny-tiktok/service/relation/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
	"tiny-tiktok/service/relation/internal/svc"
	"tiny-tiktok/service/relation/relation"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *relation.FriendListRequest) (*relation.FriendUserResponse, error) {
	// todo: add your logic here and delete this line
	userId := in.UserId

	followResp, err := l.svcCtx.UserModel.FindFollowList(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	followerResp, err := l.svcCtx.UserModel.FindFollowerList(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	resp1Map := make(map[int64]bool)
	for _, followUser := range followResp {
		resp1Map[followUser.Id] = true
	}

	// 创建一个切片来存储交集结果
	var friends []*relationModel.User

	// 遍历 resp2，检查元素是否也在 resp1 中
	for _, followerUser := range followerResp {
		if resp1Map[followerUser.Id] {
			friends = append(friends, followerUser)
		}
	}

	var respUserList []*relation.FriendUser
	//respUserList := make([]*relation.FriendUser)
	for _, friend := range friends {
		firstMsg, err := l.svcCtx.MessageModel.FindToUserFirstMsg(l.ctx, userId, friend.Id)
		if err != nil {
			return nil, err
		}

		var msgType int64
		msgType = 0
		if firstMsg.FromUserId == userId {
			msgType = 1
		}

		respUserList = append(respUserList, &relation.FriendUser{
			User: &relation.User{
				Id:              friend.Id,
				Name:            friend.Username,
				FollowCount:     friend.FollowCount,
				FollowerCount:   friend.FollowerCount,
				IsFollow:        friend.IsFollow.Valid,
				Avatar:          friend.Avatar.String,
				BackgroundImage: friend.BackgroundImage.String,
				Signature:       friend.Signature.String,
				TotalFavorited:  friend.TotalFavorited.Int64,
				WorkCount:       friend.WorkCount.Int64,
				FavoriteCount:   friend.FavoriteCount.Int64,
			},

			MsgType: msgType,
			Message: firstMsg.Content,
		})
	}

	return &relation.FriendUserResponse{
		StatusCode: 200,
		StatusMsg:  "查询成功",
		UserList:   respUserList,
	}, nil
}
