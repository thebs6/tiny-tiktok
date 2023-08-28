package logic

import (
	"context"

	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoListLogic {
	return &UserInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoListLogic) UserInfoList(in *user.UserInfoListReq) (*user.UserInfoListResp, error) {
	userList := make([]*user.User, 0, len(in.UserIdList))
	for _, userId := range in.UserIdList {
		userInfo, err := l.getUserInfo(userId)
		if err != nil {
			userList = append(userList, &user.User{Id: userId})
			continue
		}
		userList = append(userList, userInfo)
	}

	return &user.UserInfoListResp{
		StatusCode: 0,
		StatusMsg:  "success",
		UserList:   userList,
	}, nil
}

func (l *UserInfoListLogic) getUserInfo(userId int64) (*user.User, error) {

	resp, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)

	if err != nil {
		return nil, err
	}

	userInfo := &user.User{
		Avatar:          resp.Avatar.String,
		BackgroundImage: resp.BackgroundImage.String,
		FavoriteCount:   resp.FavoriteCount.Int64,
		FollowCount:     resp.FollowCount,
		FollowerCount:   resp.FollowerCount,
		Id:              resp.Id,
		IsFollow:        resp.IsFollow.Int64 != 0,
		Name:            resp.Username,
		Signature:       resp.Signature.String,
		TotalFavorited:  resp.TotalFavorited.Int64,
		WorkCount:       resp.WorkCount.Int64,
	}

	return userInfo, nil
}
