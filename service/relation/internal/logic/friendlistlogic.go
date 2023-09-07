package logic

import (
	"context"

	"tiny-tiktok/service/relation/internal/svc"
	"tiny-tiktok/service/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
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

	friendIdList, err := l.svcCtx.UserModel.FindFriendList(l.ctx, userId)

	if err != nil {
		return nil, err
	}

	var respUserList []*relation.FriendUser

	for _, friend := range friendIdList {
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
