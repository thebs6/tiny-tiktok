package model

import (
	"context"
	"errors"
	"sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	dbDao  sqlx.SqlConn
	dbOnce sync.Once
)

func NewDBDao() *sqlx.SqlConn {
	dbOnce.Do(
		func() {
			dsn := "thebs:gogotiktok@tcp(124.71.9.116:3306)/gotiktok?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
			dbDao = sqlx.NewMysql(dsn)
		})
	return &dbDao
}

func Trans(ctx context.Context, userId, videoId int64, actionType int64) error {
	var increment = 1
	var isFavorite int64 = 1
	if actionType == 2 {
		increment = -1
		isFavorite = 0
	}

	err := dbDao.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		query := "select * from favorite where `user_id` = ? and `video_id` = ?"
		var favorite Favorite
		err := session.QueryRowCtx(ctx, &favorite, query, userId, videoId)
		if err == ErrNotFound {
			query = "insert into favorite(`user_id`, `video_id`, `is_favorite`) values(?, ?, ?)"
			_, err = session.ExecCtx(ctx, query, userId, videoId, isFavorite)
			if err != nil {
				return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
			}
		} else if err != nil {
			return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
		} else {
			if favorite.IsFavorite == isFavorite {
				return nil
			}
			query = "update favorite set is_favorite=? where `id`=?"
			_, err = session.ExecCtx(ctx, query, isFavorite, favorite.Id)
			if err != nil {
				return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
			}
		}

		query = "select * from `video` where `id` = ?"
		var video Video
		err = session.QueryRowCtx(ctx, &video, query, videoId)
		if err != nil {
			return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
		}

		query = "update video set `favorite_count`=`favorite_count`+? where `id` = ?"
		_, err = session.ExecCtx(ctx, query, increment, videoId)
		if err != nil {
			return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
		}

		query = "update user set `total_favorited`=`total_favorited`+? where `id` = ?"
		_, err = session.ExecCtx(ctx, query, increment, video.Author)
		if err != nil {
			return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
		}

		query = "update user set `favorite_count`=`favorite_count`+? where `id` = ?"
		_, err = session.ExecCtx(ctx, query, increment, userId)
		if err != nil {
			return errors.New("\"" + query + "\"" + " in Trans error: " + err.Error())
		}

		return nil
	})
	return err
}
