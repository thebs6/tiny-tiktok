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
	// userId := in.UserId

	// resp, err := l.svcCtx.UserModel.FindFriendList(l.ctx, userId)

	// if err != nil {
	// 	return nil, err
	// }
	
	// var respUserList []*relation.User

	// for _, user := range resp {
	// 	respUserList = append(respUserList, &relation.User{
	// 		Id: user.Id,
	// 		Name: user.Username,
	// 		FollowCount: user.FollowCount,
	// 		FollowerCount: user.FollowerCount,
	// 		IsFollow: user.IsFollow.Valid,
	// 		Avatar: user.Avatar.String,
	// 		BackgroundImage: user.BackgroundImage.String,
	// 		Signature: user.Signature.String,
	// 		TotalFavorited: user.TotalFavorited.Int64,
	// 		WorkCount: user.WorkCount.Int64,
	// 		FavoriteCount: user.FavoriteCount.Int64,
	// 	})
	// }
	// TODO: fix 最新聊天信息
	// return &relation.FriendUserResponse{
	// 	StatusCode: 200,
	// 	StatusMsg: "查询成功",
	// 	UserList: respUserList,
	// }, nil
	return &relation.FriendUserResponse{}, nil
}
