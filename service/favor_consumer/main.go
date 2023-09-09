package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"tiny-tiktok/service/favor_consumer/internal/model"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	rdsCli        *redis.Client
	rdsOnce       sync.Once
	userModel     model.UserModel
	favoriteModel model.FavoriteModel
	videoModel    model.VideoModel
	dbOnce        sync.Once
	favorStream   = "favor_stream"
)

func NewRdsInstance() {
	rdsOnce.Do(
		func() {
			rdsCli = redis.NewClient(&redis.Options{
				Addr:     "127.0.0.1:6379",
				Password: "",
			})
		})
}
func NewDBInstance() {
	dbOnce.Do(
		func() {
			dsn := "thebs:gogotiktok@tcp(124.71.9.116:3306)/gotiktok?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
			userModel = model.NewUserModel(sqlx.NewMysql(dsn))
			favoriteModel = model.NewFavoriteModel(sqlx.NewMysql(dsn))
			videoModel = model.NewVideoModel(sqlx.NewMysql(dsn))
		})
}

func main() {
	NewRdsInstance()
	model.NewDBDao()
	ctx := context.Background()

	logc.Info(ctx, "favor_consumer start!")
	for {
		streams := rdsCli.XRead(ctx, &redis.XReadArgs{
			Streams: []string{favorStream, "0"},
			Count:   1,
			Block:   0,
		}).Val()
		if len(streams) == 0 {
			continue
		}
		for _, msg := range streams[0].Messages {
			userIdStr, ok := msg.Values["userId"].(string)
			if !ok {
				logc.Alert(ctx, "userId interface conversion error")
				break
			}
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				logc.Alert(ctx, "userId string conversion error")
				break
			}

			videoIdStr, ok := msg.Values["videoId"].(string)
			if !ok {
				logc.Alert(ctx, "videoId interface conversion error")
				break
			}

			videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
			if err != nil {
				logc.Alert(ctx, "userId string conversion error")
				break
			}

			actionTypeStr, ok := msg.Values["actionType"].(string)
			if !ok {
				logc.Alert(ctx, "actionType interface conversion error")
				break
			}
			actionType, err := strconv.ParseInt(actionTypeStr, 10, 32)
			if err != nil {
				logc.Alert(ctx, "userId string conversion error")
				break
			}
			fmt.Println("end conver")
			err = model.Trans(ctx, userId, videoId, actionType)
			if err != nil {
				logc.Alert(ctx, "persistance error: "+err.Error())
				break
			}
			rdsCli.XDel(ctx, favorStream, msg.ID)
		}
	}
}
