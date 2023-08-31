package core

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/user/pb/user"

	// "github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	// "github.com/zeromicro/go-zero/zrpc"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	rpcResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		UserId: req.UserID,
		Token: req.Token,
	})

	if err != nil {
		return nil, err
	}

	var respUser types.User
	// _ = copier.Copy(respUser, rpcResp.User)
	respUser.Avatar = rpcResp.User.Avatar
	respUser.BackgroundImage = rpcResp.User.BackgroundImage
	respUser.FavoriteCount = rpcResp.User.FavoriteCount
	respUser.FollowCount = rpcResp.User.FollowCount
	respUser.FollowerCount = rpcResp.User.FollowerCount
	respUser.Id = rpcResp.User.Id
	respUser.IsFollow = rpcResp.User.IsFollow
	respUser.Name = rpcResp.User.Name
	respUser.Signature = rpcResp.User.Signature
	respUser.TotalFavorited = rpcResp.User.TotalFavorited
	respUser.WorkCount = rpcResp.User.WorkCount
	
	return &types.UserInfoResp{
		StatusCode: 200,
		StatusMsg: "success",
		User: respUser,
	}, nil
}
