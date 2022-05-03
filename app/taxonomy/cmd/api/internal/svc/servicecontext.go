package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/rpc/taxonomies"
	"archdesc-apis/app/taxonomy/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	TaxonomyModel model.TaxonomyModel
	TaxonomyRpc   taxonomies.Taxonomies
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		TaxonomyRpc:   taxonomies.NewTaxonomies(zrpc.MustNewClient(c.TaxonomyRpcConf)),
		TaxonomyModel: model.NewTaxonomyModel(sqlx.NewMysql(c.DB.DataSource), cache.ClusterConf{}),
	}
}
