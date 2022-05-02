package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/rpc/taxonomies"

  "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
  TaxonomyRpc taxonomies.Taxonomies
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
    TaxonomyRpc: taxonomies.NewTaxonomies(zrpc.MustNewClient(c.TaxonomyRpcConf)),
	}
}
