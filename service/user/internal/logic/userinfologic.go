package logic

import (
	"context"
	"errors"

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
		return nil, errors.New("查询数据失败")
	}

	if resp == nil {
		return nil, errors.New("用户不存在")
	}

	return &user.UserInfoResp{}, nil
}
