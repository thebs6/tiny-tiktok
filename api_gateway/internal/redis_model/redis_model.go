package redis_model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisModel struct {
	redcli   *redis.Client
	videoKey string
}

func NewRedisModel(redcli *redis.Client) RedisModel {
	videoKeyDuration := time.Duration(60 * time.Second)
	redcli.Set(context.Background(), "video_key", 0, videoKeyDuration)
	return RedisModel{
		redcli:   redcli,
		videoKey: "video_key",
	}
}

func (r RedisModel) INCR(ctx context.Context) (int64, error) {
	key, err := r.redcli.Incr(ctx, r.videoKey).Result()
	if err != nil {
		return -1, err
	}
	return key, nil
}
