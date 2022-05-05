package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TaxonomyI18nModel = (*customTaxonomyI18nModel)(nil)

type (
	// TaxonomyI18nModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaxonomyI18nModel.
	TaxonomyI18nModel interface {
		taxonomyI18nModel
	}

	customTaxonomyI18nModel struct {
		*defaultTaxonomyI18nModel
	}
)

// NewTaxonomyI18nModel returns a model for the database table.
func NewTaxonomyI18nModel(conn sqlx.SqlConn) TaxonomyI18nModel {
	return &customTaxonomyI18nModel{
		defaultTaxonomyI18nModel: newTaxonomyI18nModel(conn),
	}
}
