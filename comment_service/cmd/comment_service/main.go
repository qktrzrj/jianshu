package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/shyptr/graphql/federation"
	"github.com/shyptr/hello-world-web/comment_service/handler"
	"github.com/shyptr/hello-world-web/comment_service/model"
	"github.com/shyptr/hello-world-web/middleware"
	"github.com/shyptr/hello-world-web/setting"
	"github.com/shyptr/hello-world-web/util"
	"time"
)

func main() {
	setting.InitConfig()
	util.InitLogger(nil)
	model.Init(nil)
	etcdReg := etcd.NewRegistry(func(options *registry.Options) {
		etcdCfg := setting.GetETCDConfig()
		options.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
		options.Timeout = time.Second * 15
	})
	service := micro.NewService(
		micro.Name("comment_service"),
		micro.Registry(etcdReg),
		micro.Version("v1"),
		micro.WrapHandler(middleware.Logger, middleware.Recovery, middleware.Tx(model.DB)),
	)
	service.Init()

	federation.RegisterFederationServiceHandler(service.Server(), handler.NewService())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
