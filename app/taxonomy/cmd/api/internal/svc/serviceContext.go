package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/api/internal/middleware"
	"archdesc-apis/app/taxonomy/cmd/rpc/taxonomyservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	TaxonomyMiddleware rest.Middleware
	TaxonomyRpcClient  taxonomyservice.TaxonomyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		TaxonomyMiddleware: middleware.NewTaxonomyMiddleware().Handle,
		TaxonomyRpcClient:  taxonomyservice.NewTaxonomyService(zrpc.MustNewClient(c.TaxonomyRpcConf)),
	}
}
