package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		ListByUserId(ctx context.Context, videoId int64) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn),
	}
}

func (c *customVideoModel) ListByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videos []*Video
	query := fmt.Sprintf("select %s from %s where author = ? and deleted_at IS NULL", videoRows, c.table)
	err := c.conn.QueryRowsCtx(ctx, &videos, query, userId)
	if err != nil {
		return nil, err
	} else {
		return videos, nil
	}
}
