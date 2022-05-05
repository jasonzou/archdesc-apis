package main

import (
	"flag"
	"fmt"

	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/api/internal/handler"
	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/taxonomy-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewGlobalMiddleware("Jason").Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
