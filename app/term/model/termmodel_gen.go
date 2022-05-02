// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	termFieldNames          = builder.RawFieldNames(&Term{})
	termRows                = strings.Join(termFieldNames, ",")
	termRowsExpectAutoSet   = strings.Join(stringx.Remove(termFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	termRowsWithPlaceHolder = strings.Join(stringx.Remove(termFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	termModel interface {
		Insert(ctx context.Context, data *Term) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Term, error)
		Update(ctx context.Context, data *Term) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTermModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Term struct {
		Id            int64          `db:"id"`
		TaxonomyId    int64          `db:"taxonomy_id"`
		Code          sql.NullString `db:"code"`
		ParentId      sql.NullInt64  `db:"parent_id"`
		Lft           sql.NullInt64  `db:"lft"`
		Rgt           sql.NullInt64  `db:"rgt"`
		SourceCulture string         `db:"source_culture"`
	}
)

func newTermModel(conn sqlx.SqlConn) *defaultTermModel {
	return &defaultTermModel{
		conn:  conn,
		table: "`term`",
	}
}

func (m *defaultTermModel) Insert(ctx context.Context, data *Term) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, termRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TaxonomyId, data.Code, data.ParentId, data.Lft, data.Rgt, data.SourceCulture)
	return ret, err
}

func (m *defaultTermModel) FindOne(ctx context.Context, id int64) (*Term, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", termRows, m.table)
	var resp Term
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

func (m *defaultTermModel) Update(ctx context.Context, data *Term) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, termRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.TaxonomyId, data.Code, data.ParentId, data.Lft, data.Rgt, data.SourceCulture, data.Id)
	return err
}

func (m *defaultTermModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTermModel) tableName() string {
	return m.table
}