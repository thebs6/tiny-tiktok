package redismodel

import (
	"context"
	"strconv"
	"tiny-tiktok/service/comment/internal/model"

	"github.com/redis/go-redis/v9"
)

// format: comment-video-{video_id}

type CommentModel struct {
	redcli    *redis.Client
	table     string
	keyPrefix string
}

func NewCommentModel(redcli *redis.Client) CommentModel {
	return CommentModel{
		redcli:    redcli,
		table:     "`comment`",
		keyPrefix: "comment-video-",
	}
}

func (m *CommentModel) Exists(ctx context.Context, video_id int64) int64 {
	key := m.keyPrefix + strconv.FormatInt(video_id, 10)
	return m.redcli.Exists(ctx, key).Val()
}

func (m *CommentModel) ZAdd(ctx context.Context, video_id int64, score int64, comment *model.Comment) error {
	key := m.keyPrefix + strconv.FormatInt(video_id, 10)

	// Question: why use comment's pointer as the Member?
	_, err := m.redcli.ZAdd(ctx, key, redis.Z{
		Score:  float64(score),
		Member: comment,
	}).Result()
	return err
}

// func (m *CommentModel) ZRevRangeWithScores(ctx context.Context, video_id int64) ([]redis.Z, error) {
func (m *CommentModel) ZRevRangeWithScores(ctx context.Context, video_id int64) ([]model.Comment, error) {
	key := m.keyPrefix + strconv.FormatInt(video_id, 10)
	// resList, err := m.redcli.ZRevRangeWithScores(ctx, key, 0, -1).Result()

	var commentList []model.Comment
	err := m.redcli.ZRevRange(ctx, key, 0, -1).ScanSlice(&commentList)
	return commentList, err
}

func (m *CommentModel) ZRem(ctx context.Context, video_id int64, comment *model.Comment) error {
	key := m.keyPrefix + strconv.FormatInt(video_id, 10)
	_, err := m.redcli.ZRem(ctx, key, comment).Result()
	return err
}
