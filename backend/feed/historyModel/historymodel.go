package historyModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

// NewHistoryModel returns a model for the database table.
func NewHistoryModel(conn sqlx.SqlConn) HistoryModel {
	return &customHistoryModel{
		defaultHistoryModel: newHistoryModel(conn),
	}
}
