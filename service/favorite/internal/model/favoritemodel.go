package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoriteModel = (*customFavoriteModel)(nil)

type (
	// FavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteModel.
	FavoriteModel interface {
		favoriteModel
		ListByUserId(ctx context.Context, userId int64) ([]*Favorite, error)
	}

	customFavoriteModel struct {
		*defaultFavoriteModel
	}
)

// NewFavoriteModel returns a model for the database table.
func NewFavoriteModel(conn sqlx.SqlConn) FavoriteModel {
	return &customFavoriteModel{
		defaultFavoriteModel: newFavoriteModel(conn),
	}
}

func (c *customFavoriteModel) ListByUserId(ctx context.Context, userId int64) ([]*Favorite, error) {
	var favorite []*Favorite
	query := fmt.Sprintf("select %s from %s where `id` = ?", favoriteRows, c.table)
	err := c.conn.QueryRowsCtx(ctx, &favorite, query, userId)
	return favorite, err
}
