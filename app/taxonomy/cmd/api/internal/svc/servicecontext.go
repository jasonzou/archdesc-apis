package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	TaxonomyModel model.TaxonomyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		TaxonomyModel: model.NewTaxonomyModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
