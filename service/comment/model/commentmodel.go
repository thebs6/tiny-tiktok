package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		List(ctx context.Context, vedioId int64) ([]*Comment, error)
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
	query := fmt.Sprintf("select %s from %s where video = ?", commentRows, c.table)
	err := c.conn.QueryRowsCtx(ctx, &comments, query, vedioId)
	if err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}
