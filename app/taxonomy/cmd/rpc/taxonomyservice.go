package main

import (
	"context"
	"flag"
	"fmt"

	"archdesc-apis/app/taxonomy/cmd/rpc/internal/config"
	"archdesc-apis/app/taxonomy/cmd/rpc/internal/server"
	"archdesc-apis/app/taxonomy/cmd/rpc/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/taxonomyservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewTaxonomyServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		taxonomyservice.RegisterTaxonomyServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(TestServerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("TestServerInterceptor ================> Begin\n")
	fmt.Printf("req ================> %+v \n", req)
	fmt.Printf("req ================> %+v \n", info)

	resp, err = handler(ctx, req)

	fmt.Printf("TestServerInterceptor ================> END\n")

	return resp, err
}
