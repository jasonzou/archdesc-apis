package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaxonomyModel = (*customTaxonomyModel)(nil)

type (
	// TaxonomyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaxonomyModel.
	TaxonomyModel interface {
		taxonomyModel
		FindAll(ctx context.Context) ([]*Taxonomy, error)
	}

	customTaxonomyModel struct {
		*defaultTaxonomyModel
	}
)

// NewTaxonomyModel returns a model for the database table.
func NewTaxonomyModel(conn sqlx.SqlConn) TaxonomyModel {
	return &customTaxonomyModel{
		defaultTaxonomyModel: newTaxonomyModel(conn),
	}
}

func (m *defaultTaxonomyModel) FindAll(ctx context.Context) ([]*Taxonomy, error) {
	query := fmt.Sprintf("select %s from %s", taxonomyRows, m.table)
	var resp []*Taxonomy
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
