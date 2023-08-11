package model

import (
	"context"
	"fmt"

	"github.com/vmihailenco/msgpack"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		List(ctx context.Context, vedioId int64) ([]*Comment, error)
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn),
	}
}

func (c *customCommentModel) List(ctx context.Context, vedioId int64) ([]*Comment, error) {
	var comments []*Comment
	query := fmt.Sprintf("select %s from %s where video_id = ?", commentRows, c.table)
	err := c.conn.QueryRowsCtx(ctx, &comments, query, vedioId)
	if err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}

func (c *customCommentModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return c.conn.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}

// for redis ZAdd
func (c *Comment) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(c)
}

func (c *Comment) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, c)
}
