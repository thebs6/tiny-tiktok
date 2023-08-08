// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	commentModel interface {
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Comment, error)
		Update(ctx context.Context, data *Comment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comment struct {
		Id        int64     `db:"id"`
		Video     int64     `db:"video"`
		User      int64     `db:"user"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
		Content   string    `db:"content"`
		Date      time.Time `db:"date"`
	}
)

func newCommentModel(conn sqlx.SqlConn) *defaultCommentModel {
	return &defaultCommentModel{
		conn:  conn,
		table: "`comment`",
	}
}

func (m *defaultCommentModel) withSession(session sqlx.Session) *defaultCommentModel {
	return &defaultCommentModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	var resp Comment
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Video, data.User, data.DeletedAt, data.Content, data.Date)
	return ret, err
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Video, data.User, data.DeletedAt, data.Content, data.Date, data.Id)
	return err
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}