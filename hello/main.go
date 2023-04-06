package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"hello/router"
)

var (
	service = "hello"
	version = "latest"
)

func main() {

	// ----------------consul------------------
	consulReg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{"127.0.0.1:8500"}
		})
	// ----------------grpc--------------------
	//grpcServer := grpc.NewServer()

	// router
	ginRouter := gin.Default()
	// 设置路由组
	router.SetRouterGroup(ginRouter)
	// Create service
	srv := web.NewService(
		web.Name(service),
		web.Version(version),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	srv.Run()

	// Register handler
	//if err := pb.RegisterHelloHandler(srv.Server(), new(handler.Hello)); err != nil {
	//	logger.Fatal(err)
	//}
	//// Run service
	//if err := srv.Run(); err != nil {
	//	logger.Fatal(err)
	//}
}
