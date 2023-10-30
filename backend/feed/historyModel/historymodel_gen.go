// Code generated by goctl. DO NOT EDIT.

package historyModel

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
	historyFieldNames          = builder.RawFieldNames(&History{})
	historyRows                = strings.Join(historyFieldNames, ",")
	historyRowsExpectAutoSet   = strings.Join(stringx.Remove(historyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	historyRowsWithPlaceHolder = strings.Join(stringx.Remove(historyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	historyModel interface {
		Insert(ctx context.Context, data *History) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*History, error)
		Update(ctx context.Context, data *History) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	History struct {
		Id        int64     `db:"id"`
		Uid       int64     `db:"uid"`
		Vid       int64     `db:"vid"`
		CreatedAt time.Time `db:"created_at"`
	}
)

func newHistoryModel(conn sqlx.SqlConn) *defaultHistoryModel {
	return &defaultHistoryModel{
		conn:  conn,
		table: "`history`",
	}
}

func (m *defaultHistoryModel) withSession(session sqlx.Session) *defaultHistoryModel {
	return &defaultHistoryModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`history`",
	}
}

func (m *defaultHistoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultHistoryModel) FindOne(ctx context.Context, id int64) (*History, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", historyRows, m.table)
	var resp History
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

func (m *defaultHistoryModel) Insert(ctx context.Context, data *History) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, historyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Vid)
	return ret, err
}

func (m *defaultHistoryModel) Update(ctx context.Context, data *History) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, historyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.Vid, data.Id)
	return err
}

func (m *defaultHistoryModel) tableName() string {
	return m.table
}