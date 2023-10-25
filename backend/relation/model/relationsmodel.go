package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RelationsModel = (*customRelationsModel)(nil)

type (
	// RelationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationsModel.
	RelationsModel interface {
		relationsModel
	}

	customRelationsModel struct {
		*defaultRelationsModel
	}
)

// NewRelationsModel returns a model for the database table.
func NewRelationsModel(conn sqlx.SqlConn) RelationsModel {
	return &customRelationsModel{
		defaultRelationsModel: newRelationsModel(conn),
	}
}
