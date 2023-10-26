package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideosModel = (*customVideosModel)(nil)

type (
	// VideosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideosModel.
	VideosModel interface {
		videosModel
	}

	customVideosModel struct {
		*defaultVideosModel
	}
)

// NewVideosModel returns a model for the database table.
func NewVideosModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideosModel {
	return &customVideosModel{
		defaultVideosModel: newVideosModel(conn, c, opts...),
	}
}
