package svc

import (
	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/api/internal/middleware"
	"archdesc-apis/app/taxonomy/cmd/rpc/taxonomyservice"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
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
		TaxonomyRpcClient:  taxonomyservice.NewTaxonomyService(zrpc.MustNewClient(c.TaxonomyRpcConf, zrpc.WithUnaryClientInterceptor(TestInterceptor))),
	}
}

func TestInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("Test Client Interceptor ================> Begin\n")
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Printf("Test Client Interceptor ================> END\n")

	return nil
}
