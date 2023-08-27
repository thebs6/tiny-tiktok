package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Cos struct {
		URL       string
		SecretId  string
		SecretKey string
	}
	Publish zrpc.RpcClientConf
}
