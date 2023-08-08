package logic

import (
	"context"
	"fmt"

	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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
	// resp := model.User{
	// 	Id:       1,
	// 	Username: "gao",
	// }
	// var err error = nil

	fmt.Println(resp.Id, resp.Username, resp.Password)
	switch err {
	case nil:
		return &user.LoginResp{
			StatusMsg: "login successfully",
			UserId:    resp.Id,
		}, nil
	case sqlc.ErrNotFound:
		return &user.LoginResp{
			StatusMsg: "the username does not exsit",
			UserId:    -1,
		}, nil
	default:
		return nil, err
	}
}
