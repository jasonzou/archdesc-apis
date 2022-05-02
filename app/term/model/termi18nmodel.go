package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TermI18nModel = (*customTermI18nModel)(nil)

type (
	// TermI18nModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTermI18nModel.
	TermI18nModel interface {
		termI18nModel
	}

	customTermI18nModel struct {
		*defaultTermI18nModel
	}
)

// NewTermI18nModel returns a model for the database table.
func NewTermI18nModel(conn sqlx.SqlConn) TermI18nModel {
	return &customTermI18nModel{
		defaultTermI18nModel: newTermI18nModel(conn),
	}
}
