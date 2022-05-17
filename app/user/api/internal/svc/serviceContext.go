package svc

import (
	"archdesc-apis/app/user/api/internal/config"
	"archdesc-apis/app/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
