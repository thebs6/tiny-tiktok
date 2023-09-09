// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	//"tiny-tiktok/service/relation/internal/model"

	// "tiny-tiktok/service/relation/internal/model"

	// "github.com/zeromicro/go-zero/core/logx"

	//"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	//"github.com/go-playground/locales/rm"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
	userRowsWithOutPWD      = strings.Join(stringx.Remove(userFieldNames, "`password`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userIdSet               = ""
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
		FindFollowerList(ctx context.Context, id int64) ([]*User, error)
		FindFollowList(ctx context.Context, id int64) ([]*User, error)
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id              int64          `db:"id"`
		Username        string         `db:"username"`
		Password        string         `db:"password"`
		FollowCount     int64          `db:"follow_count"`
		FollowerCount   int64          `db:"follower_count"`
		CreatedAt       time.Time      `db:"created_at"`
		UpdatedAt       time.Time      `db:"updated_at"`
		DeletedAt       sql.NullTime   `db:"deleted_at"`
		IsFollow        sql.NullInt64  `db:"is_follow"`
		Avatar          sql.NullString `db:"avatar"`
		BackgroundImage sql.NullString `db:"background_image"`
		Signature       sql.NullString `db:"signature"`
		TotalFavorited  sql.NullInt64  `db:"total_favorited"`
		WorkCount       sql.NullInt64  `db:"work_count"`
		FavoriteCount   sql.NullInt64  `db:"favorite_count"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) withSession(session sqlx.Session) *defaultUserModel {
	return &defaultUserModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
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

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.FollowCount, data.FollowerCount, data.DeletedAt, data.IsFollow, data.Avatar, data.BackgroundImage, data.Signature, data.TotalFavorited, data.WorkCount, data.FavoriteCount)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.FollowCount, newData.FollowerCount, newData.DeletedAt, newData.IsFollow, newData.Avatar, newData.BackgroundImage, newData.Signature, newData.TotalFavorited, newData.WorkCount, newData.FavoriteCount, newData.Id)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}

func (m *defaultUserModel) FindFollowerList(ctx context.Context, id int64) ([]*User, error) {
	q := squirrel.Select("u.*").From(m.table+" u").Join("relation r ON u.id = r.follower_id").Where("r.follow_id = ?", id)
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindFollowList(ctx context.Context, id int64) ([]*User, error) {
	q := squirrel.Select("u.*").From(m.table+" u").Join("relation r ON u.id = r.follow_id").Where("r.follow_id = ?", id)

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//func (m *defaultUserModel) FindFriendList(ctx context.Context, id int64) ([]*User, error) {
//	rr := newRelationModel(m.conn)
//	relationResp, err := rr.FindFriendRelation(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//
//	var resp []*User
//	for _, r := range(relationResp) {
//		fid := r.FollowId
//		if fid == id {
//			fid = r.FollowerId
//		}
//
//		friend, err := m.FindOne(ctx, fid)
//		if err != nil {
//			return nil, err
//		}
//
//		resp = append(resp, friend)
//	}
//
//	return resp, nil
//}
