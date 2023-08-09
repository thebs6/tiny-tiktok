package logic

import (
	"context"
	"regexp"

	"tiny-tiktok/service/user/internal/model"
	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	data := &model.User{
		Username: in.Username,
		Password: in.Password,
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, data)

	switch err {
	case nil:
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		return &user.RegisterResp{
			StatusMsg: "Register successfully",
			UserId:    id,
		}, nil
	default:
		if match, _ := regexp.MatchString(".*(23000).*", err.Error()); match {
			return &user.RegisterResp{
				StatusMsg: "The username has been used",
				UserId:    -1,
			}, nil
		}

		return nil, err
	}
}
