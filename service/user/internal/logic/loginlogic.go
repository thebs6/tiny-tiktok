package logic

import (
	"context"
	"encoding/base64"

	"tiny-tiktok/service/user/internal/model"
	"tiny-tiktok/service/user/internal/svc"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/scrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var PW_HASH_BYTES = 32
var SALT = []byte{126, 145, 58, 233, 153, 107, 4, 231}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	resp, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Username)

	switch err {
	case nil:
		hash, err := scrypt.Key([]byte(in.Password), SALT, 1<<15, 8, 1, PW_HASH_BYTES)
		encodedHash := base64.StdEncoding.EncodeToString(hash)
		if err != nil {
			return nil, err
		}
		if resp.Password != encodedHash {
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
