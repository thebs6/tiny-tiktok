package extra_second

import (
	"context"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListReq) (resp *types.RelationFriendListResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RelationRpc.FriendList(l.ctx, &relation.FriendListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		resp.StatusCode = "4040"
		resp.StatusMsg = "rpc调用错误"
		return
	}

	var respUser []types.User
	for _, user := range rpcResp.UserList {
		respUser = append(respUser, types.User{
			Avatar:          user.User.Avatar,
			BackgroundImage: user.User.BackgroundImage,
			FavoriteCount:   user.User.FavoriteCount,
			FollowCount:     user.User.FollowCount,
			FollowerCount:   user.User.FollowerCount,
			Id:              user.User.Id,
			IsFollow:        user.User.IsFollow,
			Name:            user.User.Name,
			Signature:       user.User.Signature,
			TotalFavorited:  user.User.TotalFavorited,
			WorkCount:       user.User.WorkCount,
		})
	}
	resp.UserList = respUser
	resp.StatusCode = "4200"
	resp.StatusMsg = "success"
	return
}
