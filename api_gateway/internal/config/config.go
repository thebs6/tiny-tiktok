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

	LogConf struct {
		ServiceName string
		Mode        string
		Path        string
	}
	RedisConf struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
		DB   int
	}
	Cos struct {
		URL       string
		SecretId  string
		SecretKey string
	}
	FeedRpcConf    zrpc.RpcClientConf
	UserRpcConf    zrpc.RpcClientConf
	PublishRpcConf zrpc.RpcClientConf
	CommentRpcConf zrpc.RpcClientConf
}
