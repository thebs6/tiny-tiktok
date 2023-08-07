package core

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"tiny-tiktok/api_gateway/internal/svc"
	"tiny-tiktok/api_gateway/internal/types"
	"tiny-tiktok/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
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
	// todo: add your logic here and delete this line
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc",
		},
	})
	client := user.NewUserServiceClient(conn.Conn())
	respRpc, err := client.Login(context.Background(), &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	fmt.Println(req.Username, req.Password)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp = &types.LoginResp{
		StatusCode: http.StatusOK,
		StatusMsg:  "login success",
		UserID:     respRpc.UserId,
		Token:      "token",
	}

	return
}
