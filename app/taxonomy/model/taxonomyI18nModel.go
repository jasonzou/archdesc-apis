package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaxonomyI18nModel = (*customTaxonomyI18nModel)(nil)

type (
	// TaxonomyI18nModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaxonomyI18nModel.
	TaxonomyI18nModel interface {
		taxonomyI18nModel
		FindOneByCulture(ctx context.Context, id int64) (*TaxonomyI18n, error)
	}

	customTaxonomyI18nModel struct {
		*defaultTaxonomyI18nModel
	}
)

// NewTaxonomyI18nModel returns a model for the database table.
func NewTaxonomyI18nModel(conn sqlx.SqlConn, c cache.CacheConf) TaxonomyI18nModel {
	return &customTaxonomyI18nModel{
		defaultTaxonomyI18nModel: newTaxonomyI18nModel(conn, c),
	}
}

func (m *defaultTaxonomyI18nModel) FindOneByCulture(ctx context.Context, id int64) (*TaxonomyI18n, error) {
	taxonomyI18nIdKey := fmt.Sprintf("%s%v", cacheTaxonomyI18nIdPrefix, id)
	var resp TaxonomyI18n
	err := m.QueryRowCtx(ctx, &resp, taxonomyI18nIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `culture` = \"en\" limit 1", taxonomyI18nRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
