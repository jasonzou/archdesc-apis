package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TermModel = (*customTermModel)(nil)

type (
	// TermModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTermModel.
	TermModel interface {
		termModel
	}

	customTermModel struct {
		*defaultTermModel
	}
)

// NewTermModel returns a model for the database table.
func NewTermModel(conn sqlx.SqlConn) TermModel {
	return &customTermModel{
		defaultTermModel: newTermModel(conn),
	}
}
