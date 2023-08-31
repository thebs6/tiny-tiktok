package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RelationModel = (*customRelationModel)(nil)

type (
	// RelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationModel.
	RelationModel interface {
		relationModel
	}

	customRelationModel struct {
		*defaultRelationModel
	}
)

// NewRelationModel returns a model for the database table.
func NewRelationModel(conn sqlx.SqlConn) RelationModel {
	return &customRelationModel{
		defaultRelationModel: newRelationModel(conn),
	}
}
