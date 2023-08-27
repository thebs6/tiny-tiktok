package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	LogConf struct {
		ServiceName string
		Mode        string
		Path        string
	}
	DB struct {
		DataSource string
	}
}
