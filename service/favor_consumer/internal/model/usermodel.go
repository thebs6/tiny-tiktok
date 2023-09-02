package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		IncrTotalFavoritedBy(ctx context.Context, id int64, num int) error
		IncrFavoriteCntBy(ctx context.Context, id int64, num int) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) IncrTotalFavoritedBy(ctx context.Context, id int64, num int) error {
	query := fmt.Sprintf("update %s set total_favorited=total_favorited+? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, num, id)
	return err
}

func (m *customUserModel) IncrFavoriteCntBy(ctx context.Context, id int64, num int) error {
	query := fmt.Sprintf("update %s set favorite_count=favorite_count+? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, num, id)
	return err
}
