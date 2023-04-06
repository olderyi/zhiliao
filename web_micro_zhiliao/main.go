package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/web"
	"web_micro_zhiliao/router"
)

var (
	service = "web_micro_zhiliao"
	version = "latest"
)

func main() {

	//consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// Create selector
	mySelector := selector.NewSelector(
		selector.SetStrategy(selector.RoundRobin),
		selector.Registry(consulReg),
	)
	// Create service
	rpcSerivce := micro.NewService(
		micro.Client(client.NewClient()),
		micro.Selector(mySelector),
		//micro.Registry(consulReg),
	)
	rpcSerivce.Init()

	rpcSerivce.Run()

	// router
	ginRouter := gin.Default()
	// router组
	router.SetRouter(ginRouter)

	srv := web.NewService(
		web.Name(service),
		web.Version(version),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	srv.Run()

}
