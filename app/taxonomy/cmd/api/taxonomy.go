package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"archdesc-apis/app/taxonomy/cmd/api/internal/config"
	"archdesc-apis/app/taxonomy/cmd/api/internal/handler"
	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/logx/zapx"
)

var configFile = flag.String("f", "etc/taxonomy-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewGlobalMiddleware("Jason").Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	logx.Infow("infow foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Errorw("errorw foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Sloww("sloww foo",
		logx.Field("url", "http://localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.Error("error")
	logx.Infov(map[string]interface{}{
		"url":     "localhost:8080/hello",
		"attempt": 3,
		"backoff": time.Second,
		"value":   "foo",
	})
	logx.WithDuration(1100*time.Microsecond).Infow("infow withduration",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.WithContext(context.Background()).WithDuration(1100*time.Microsecond).Errorw(
		"errorw withcontext withduration",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	logx.WithDuration(1100*time.Microsecond).WithContext(context.Background()).Errorw(
		"errorw withduration withcontext",
		logx.Field("url", "localhost:8080/hello"),
		logx.Field("attempt", 3),
		logx.Field("backoff", time.Second),
	)
	server.Start()
}
