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
	videosFieldNames          = builder.RawFieldNames(&Videos{})
	videosRows                = strings.Join(videosFieldNames, ",")
	videosRowsExpectAutoSet   = strings.Join(stringx.Remove(videosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videosRowsWithPlaceHolder = strings.Join(stringx.Remove(videosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	videosModel interface {
		Insert(ctx context.Context, data *Videos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Videos, error)
		Update(ctx context.Context, data *Videos) error
		Delete(ctx context.Context, id int64) error
		FindOwnFeed(ctx context.Context, uid int64) ([]*Videos, error)
		FindFeeds(ctx context.Context) ([]*Videos, error)
		FindCategoryFeeds(ctx context.Context, category int64) ([]*Videos, error)
	}

	defaultVideosModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Videos struct {
		Id            int64        `db:"id"`             // 主键
		AuthorId      int64        `db:"author_id"`      // 上传用户Id
		Title         string       `db:"title"`          // 视频标题
		CoverUrl      string       `db:"cover_url"`      // 封面url
		PlayUrl       string       `db:"play_url"`       // 视频播放url
		FavoriteCount int64        `db:"favorite_count"` // 点赞数
		StarCount     int64        `db:"star_count"`     // 收藏数
		CommentCount  int64        `db:"comment_count"`  // 评论数目
		CreatedAt     time.Time    `db:"created_at"`
		UpdatedAt     sql.NullTime `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
		Category      int64        `db:"category"` // 视频分类
	}
)

func newVideosModel(conn sqlx.SqlConn) *defaultVideosModel {
	return &defaultVideosModel{
		conn:  conn,
		table: "`videos`",
	}
}

func (m *defaultVideosModel) withSession(session sqlx.Session) *defaultVideosModel {
	return &defaultVideosModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`videos`",
	}
}

func (m *defaultVideosModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVideosModel) FindOne(ctx context.Context, id int64) (*Videos, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videosRows, m.table)
	var resp Videos
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

func (m *defaultVideosModel) Insert(ctx context.Context, data *Videos) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, videosRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AuthorId, data.Title, data.CoverUrl, data.PlayUrl, data.FavoriteCount, data.StarCount, data.CommentCount, data.DeletedAt, data.Category)
	return ret, err
}

func (m *defaultVideosModel) Update(ctx context.Context, data *Videos) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videosRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AuthorId, data.Title, data.CoverUrl, data.PlayUrl, data.FavoriteCount, data.StarCount, data.CommentCount, data.DeletedAt, data.Category, data.Id)
	return err
}

func (m *defaultVideosModel) tableName() string {
	return m.table
}
func (m *defaultVideosModel) FindOwnFeed(ctx context.Context, uid int64) ([]*Videos, error) {
	query := fmt.Sprintf("select %s from %s where `author_id` = ?", videosRows, m.table)
	var resp []*Videos
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultVideosModel) FindFeeds(ctx context.Context) ([]*Videos, error) {
	query := fmt.Sprintf("select %s from %s ORDER BY `id` DESC", videosRows, m.table)
	var resp []*Videos
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultVideosModel) FindCategoryFeeds(ctx context.Context, category int64) ([]*Videos, error) {
	query := fmt.Sprintf("select %s from %s where `category` = ? ORDER BY `id` DESC", videosRows, m.table)
	var resp []*Videos
	err := m.conn.QueryRowsCtx(ctx, &resp, query, category)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
