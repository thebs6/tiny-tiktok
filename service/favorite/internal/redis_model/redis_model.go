package redis_model

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

// key video_favored_cnt_{video_id} : cnt
// set user_favor_set_{user_id} : video_id, video_id, ...
// list favor_op_list : video_id::user_id::{action_type}, video_id::user_id::{action_type}, ...
var (
	ErrRecordRepeated = errors.New("record repeated")
	ErrRecordNonExist = errors.New("record not exist")
	maxRetries        = 3
	videoFavCntPrefix = "video_favored_cnt_"
	userFavSetPrefix  = "user_favor_set_"
	favoriteOpList    = "favor_op_list_"
	favorStream       = "favor_stream"
)

type RedisModel struct {
	redcli *redis.Client
	mutex  sync.Mutex
}

func NewRedisModel(redcli *redis.Client) RedisModel {
	return RedisModel{
		redcli: redcli,
		mutex:  sync.Mutex{},
	}
}

func (r *RedisModel) XADD(ctx context.Context, userId, videoId int64, actionType int32) error {
	return r.redcli.XAdd(ctx, &redis.XAddArgs{
		Stream: favorStream,
		Values: map[string]interface{}{
			"userId":     userId,
			"videoId":    videoId,
			"actionType": actionType,
		},
	}).Err()
}

// func (r *RedisModel) S(ctx context.Context, userId int64) error {

// }

func (r *RedisModel) Favor(ctx context.Context, user_id, video_id int64) error {
	user_id_str := strconv.FormatInt(user_id, 10)
	video_id_str := strconv.FormatInt(video_id, 10)

	// use optimistic lock here, because the same user
	// usually can not favor multi videos at the same time.
	txf := func(tx *redis.Tx) error {
		val, err := r.redcli.SIsMember(ctx, userFavSetPrefix+user_id_str, video_id_str).Result()
		if err != nil {
			return err
		}
		if val {
			return ErrRecordRepeated
		}
		_, err = r.redcli.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.SAdd(ctx, userFavSetPrefix+user_id_str, video_id_str)
			pipe.RPush(ctx, favoriteOpList+video_id_str, video_id_str+":"+user_id_str+":"+"1")
			pipe.Incr(ctx, videoFavCntPrefix+video_id_str)
			return nil
		})
		return err
	}

	var err error
	for i := 0; i < maxRetries; i++ {
		err = r.redcli.Watch(ctx, txf, userFavSetPrefix+user_id_str)
		if err == redis.TxFailedErr {
			continue
		}
		break
	}
	return err
}

func (r *RedisModel) CancelFavor(ctx context.Context, user_id, video_id int64) error {
	user_id_str := strconv.FormatInt(user_id, 10)
	video_id_str := strconv.FormatInt(video_id, 10)

	// use optimistic lock here, because a user
	// usually can not cancel more than one favor at the same time.
	txf := func(tx *redis.Tx) error {
		val, err := r.redcli.SIsMember(ctx, userFavSetPrefix+user_id_str, video_id_str).Result()
		if err != nil {
			return err
		}
		if !val {
			return ErrRecordNonExist
		}
		_, err = r.redcli.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.SRem(ctx, userFavSetPrefix+user_id_str, video_id_str)
			pipe.RPush(ctx, favoriteOpList+video_id_str, video_id_str+"::"+user_id_str+"::"+"0")
			pipe.Decr(ctx, videoFavCntPrefix+video_id_str)
			return nil
		})
		return err
	}

	var err error
	for i := 0; i < maxRetries; i++ {
		err = r.redcli.Watch(ctx, txf, userFavSetPrefix+user_id_str)
		if err == redis.TxFailedErr {
			continue
		}
		break
	}
	return err
}
