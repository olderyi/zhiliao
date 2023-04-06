package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
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

	// router
	ginRouter := gin.Default()
	// router组

	// Create service
	srv := web.NewService(
		web.Name(service),
		web.Version(version),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	srv.Run()

}
