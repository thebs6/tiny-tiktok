package repository

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var db sqlx.SqlConn

func Init() {
	dsn := "thebs:gogotiktok@tcp(124.71.9.116:3306)/gotiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db = sqlx.NewMysql(dsn)
}
