package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/api/internal/middleware"
	"archdesc-apis/app/taxonomy/cmd/rpc/taxonomyservice"
	"archdesc-apis/app/taxonomy/model"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	TaxonomyMiddleware rest.Middleware
	TaxonomyModel      model.TaxonomyModel
	TaxonomyI18nModel  model.TaxonomyI18nModel
	TaxonomyRpcClient  taxonomyservice.TaxonomyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		TaxonomyMiddleware: middleware.NewTaxonomyMiddleware().Handle,
		TaxonomyModel:      model.NewTaxonomyModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
		TaxonomyI18nModel:  model.NewTaxonomyI18nModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
		TaxonomyRpcClient:  taxonomyservice.NewTaxonomyService(zrpc.MustNewClient(c.TaxonomyRpcConf)),
	}
}
