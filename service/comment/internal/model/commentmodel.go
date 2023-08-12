package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
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
		SoftDel(ctx context.Context, comment_id int64) error
		TransInsert(ctx context.Context, session sqlx.Session, data *Comment) (sql.Result, error)
		TransFindone(ctx context.Context, session sqlx.Session, id int64) (*Comment, error)
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
	query := fmt.Sprintf("select %s from %s where video_id = ? and deleted_at IS NULL", commentRows, c.table)
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

func (m *defaultCommentModel) SoftDel(ctx context.Context, comment_id int64) error {
	query := fmt.Sprintf("update %s set `deleted_at` = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, sql.NullTime{Time: time.Now(), Valid: true}, comment_id)
	return err
}

func (m *defaultCommentModel) TransInsert(ctx context.Context, session sqlx.Session, data *Comment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.VideoId, data.UserId, data.DeletedAt, data.Content, data.Date)
	return ret, err
}

func (m *defaultCommentModel) TransFindone(ctx context.Context, session sqlx.Session, id int64) (*Comment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	var resp Comment
	err := session.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// for redis ZAdd
func (c *Comment) MarshalBinary() ([]byte, error) {
	fmt.Println("MarshalBinary")
	// return msgpack.Marshal(c)
	return json.Marshal(c)
}

func (c *Comment) UnmarshalBinary(data []byte) error {
	fmt.Println("UnmarshalBinary")
	// return msgpack.Unmarshal(data, c)
	return json.Unmarshal(data, c)
}
