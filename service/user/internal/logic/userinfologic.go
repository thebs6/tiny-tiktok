package logic

import (
	"context"

	"tiny-tiktok/service/user/internal/model"
	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	resp, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	if resp == nil {
		return &user.UserInfoResp{
			StatusCode: 1004,
			StatusMsg: "User Not Found",
			User: nil,
		}, nil
	}

	var respUser user.User
	
	respUser.Avatar = resp.Avatar.String
	respUser.BackgroundImage = resp.BackgroundImage.String
	respUser.FavoriteCount = resp.FavoriteCount.Int64
	respUser.FollowCount = resp.FollowCount
	respUser.FollowerCount = resp.FollowerCount
	respUser.Id = resp.Id
	respUser.IsFollow = resp.IsFollow.Int64 != 0
	respUser.Name = resp.Username
	respUser.Signature = resp.Signature.String
	respUser.TotalFavorited = resp.TotalFavorited.Int64
	respUser.WorkCount = resp.WorkCount.Int64

	return &user.UserInfoResp{
		StatusCode: 200,
		StatusMsg: "Success",
		User: &respUser,
	}, nil
}
