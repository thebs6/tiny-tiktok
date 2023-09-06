package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
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
}
