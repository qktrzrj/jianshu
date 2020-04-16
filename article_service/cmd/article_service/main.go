package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/shyptr/graphql/federation"
	"github.com/shyptr/hello-world-web/article_service/handler"
	"github.com/shyptr/hello-world-web/article_service/model"
	"github.com/shyptr/hello-world-web/middleware"
	"github.com/shyptr/hello-world-web/setting"
	"github.com/shyptr/hello-world-web/util"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	setting.InitConfig()
	util.InitLogger()
	model.Init()

	etcdReg := etcd.NewRegistry(func(options *registry.Options) {
		etcdCfg := setting.GetETCDConfig()
		options.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
		options.Timeout = time.Second * 15
	})
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.Logger(), middleware.Recovery(), middleware.Tx(model.DB),
		)),
	)
	service := micro.NewService(
		micro.Name("article_service"),
		micro.Registry(etcdReg),
		micro.Version("v1"),
	)
	service.Init()

	federation.RegisterFederationServiceServer(server, handler.NewService())

	li, err := net.Listen("tcp", ":30001")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(li); err != nil {
		panic(err)
	}
}
