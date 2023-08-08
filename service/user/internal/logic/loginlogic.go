package logic

import (
	"context"

	"tiny-tiktok/service/user/internal/model"
	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	resp, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)
	// resp, err := l.svcCtx.UserModel.FindOne(l.ctx, 2)
	// resp := model.User{
	// 	Id:       1,
	// 	Username: "gao",
	// }
	// var err error = nil

	switch err {
	case nil:
		if resp.Password != in.Password {
			return &user.LoginResp{
				StatusMsg: "Password is incorrect",
				UserId:    -1,
			}, nil
		}
		return &user.LoginResp{
			StatusMsg: "Login successfully",
			UserId:    resp.Id,
		}, nil
	case model.ErrNotFound:
		return &user.LoginResp{
			StatusMsg: "The username does not exsit",
			UserId:    -1,
		}, nil
	default:
		return nil, err
	}
}
