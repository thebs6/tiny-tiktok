package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
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

func (c *customVideoModel) GetVideoIdByUserid(ctx context.Context, userId int64) (int64, error) {
	query := fmt.Sprintf("select id from %s where `author` = ? limit 1", c.table)
	var resp Video
	err := c.conn.QueryRowCtx(ctx, resp, query, userId)
	switch err {
	case nil:
		return resp.Id, nil
	case sqlc.ErrNotFound:
		return -1, ErrNotFound
	default:
		return -1, err
	}
}
