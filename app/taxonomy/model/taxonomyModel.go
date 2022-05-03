package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TaxonomyModel = (*customTaxonomyModel)(nil)

type (
	// TaxonomyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaxonomyModel.
	TaxonomyModel interface {
		taxonomyModel
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
