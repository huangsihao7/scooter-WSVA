package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}
)

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn, c, opts...),
	}
}
