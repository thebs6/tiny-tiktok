package core

import (
	"context"
	"net/http"
	"time"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/user/pb/user"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	logc.Debug(l.ctx, "LoginLogic.Login req")
	respRpc, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp = &types.LoginResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Login fail",
			UserID:     respRpc.UserId, // is -1
		}
		logc.Alert(l.ctx, "call UserRpc failed"+err.Error())
		err = nil
		return
	} else if respRpc.UserId == -1 {
		// the username does not exsit or the password is incorrect
		resp = &types.LoginResp{
			StatusCode: http.StatusOK,
			StatusMsg:  respRpc.StatusMsg,
			UserID:     respRpc.UserId, // is -1
		}
		err = nil
		return
	}

	secretKey := l.svcCtx.Config.Auth.AccessSecret
	iat := time.Now().Unix() // maybe not word on Windows OS
	seconds := l.svcCtx.Config.Auth.AccessExpire
	payload := respRpc.UserId

	token, err := getJwtToken(secretKey, iat, seconds, payload)
	if err != nil {
		resp = &types.LoginResp{
			StatusCode: http.StatusOK,
			StatusMsg:  "Login fail",
			UserID:     respRpc.UserId, // is -1
		}
		logc.Alert(l.ctx, "getJwtToken() "+err.Error())
		err = nil
		return
	}

	resp = &types.LoginResp{
		StatusCode: http.StatusOK,
		StatusMsg:  respRpc.StatusMsg,
		UserID:     respRpc.UserId,
		Token:      token,
	}

	return
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func getJwtToken(secretKey string, iat, seconds, payload int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
