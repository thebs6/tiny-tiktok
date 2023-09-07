// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	relationFieldNames          = builder.RawFieldNames(&Relation{})
	relationRows                = strings.Join(relationFieldNames, ",")
	relationWithoutId           = strings.Join(stringx.Remove(relationFieldNames, "`id`"), ",")
	relationRowsExpectAutoSet   = strings.Join(stringx.Remove(relationFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	relationRowsWithPlaceHolder = strings.Join(stringx.Remove(relationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	relationModel interface {
		Insert(ctx context.Context, data *Relation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Relation, error)
		Update(ctx context.Context, data *Relation) error
		Delete(ctx context.Context, id int64) error
		FindRelationByTwoId(ctx context.Context, follow_id int64, follower_id int64) (*Relation, error)
		FindFriendRelation(ctx context.Context, user_id int64) ([]*Relation, error)
	}

	defaultRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Relation struct {
		Id         int64     `db:"id"`
		FollowId   int64     `db:"follow_id"`
		FollowerId int64     `db:"follower_id"`
		CreatedAt  time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
	}
)

func newRelationModel(conn sqlx.SqlConn) *defaultRelationModel {
	return &defaultRelationModel{
		conn:  conn,
		table: "`relation`",
	}
}

func (m *defaultRelationModel) withSession(session sqlx.Session) *defaultRelationModel {
	return &defaultRelationModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`relation`",
	}
}

func (m *defaultRelationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRelationModel) FindOne(ctx context.Context, id int64) (*Relation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relationRows, m.table)
	var resp Relation
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

func (m *defaultRelationModel) Insert(ctx context.Context, data *Relation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, relationWithoutId)
	ret, err := m.conn.ExecCtx(ctx, query, data.FollowId, data.FollowerId, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultRelationModel) Update(ctx context.Context, data *Relation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, relationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FollowId, data.FollowerId, data.Id)
	return err
}

func (m *defaultRelationModel) tableName() string {
	return m.table
}

func (m *defaultRelationModel) FindRelationByTwoId(ctx context.Context, follow_id int64, follower_id int64) (*Relation, error) {
	query := fmt.Sprintf("select %s from %s where `follow_id` = ? and `follower_id` = ? limit 1", relationRows, m.table)
	var resp Relation
	err := m.conn.QueryRowCtx(ctx, &resp, query, follow_id, follower_id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRelationModel)FindFriendRelation(ctx context.Context, user_id int64) ([]*Relation, error) {
	builder := squirrel.SelectBuilder{}
	builder = builder.Columns(relationRows)
	builder = builder.From(m.table).Where("follow_id = ?", user_id)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp1 []*Relation
	err = m.conn.QueryRowCtx(ctx, &resp1, query, args...)

	builder = squirrel.SelectBuilder{}
	builder = builder.Columns(relationRows)
	builder = builder.From(m.table).Where("follow_id = ?", user_id)

	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp2 []*Relation
	err = m.conn.QueryRowCtx(ctx, &resp1, query, args...)

	resp1Map := make(map[int64]bool)
	for _, relation := range resp1 {
		resp1Map[relation.FollowId] = true
	}

	// 创建一个切片来存储交集结果
	var intersection []*Relation

	// 遍历 resp2，检查元素是否也在 resp1 中
	for _, relation := range resp2 {
		if resp1Map[relation.FollowerId] {
			intersection = append(intersection, relation)
		}
	}

	switch err {
	case nil:
		return intersection, nil
	default:
		return nil, err
	}
}