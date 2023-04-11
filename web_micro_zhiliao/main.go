package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/selector"
	"go-micro.dev/v4/web"
	"net/http"
	_ "user_micro_zhiliao/database"
	userPb "user_micro_zhiliao/proto"
	busRouter "web_micro_zhiliao/router"
)

var (
	service = "web_micro_zhiliao"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// router
	router := gin.Default()
	busRouter.InitRouter(router)
	//router.GET("/test", ServiceOne)
	// router组
	//router := router.InitrRouter(rpcService.Client())

	srv := web.NewService(
		web.Address(":8088"),
		web.Name(service),
		web.Version(version),
		web.Handler(router),
		web.Registry(consulReg),
	)

	srv.Init()
	srv.Run()
	//ginRouter.Run(":8080")
}

func ServiceOne(c *gin.Context) {
	////consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	//Create selector
	mySelector := selector.NewSelector(
		selector.SetStrategy(selector.RoundRobin),
		selector.Registry(consulReg),
	)
	// Create service
	rpcService := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Selector(mySelector),
	)
	rpcService.Init()

	//userClient := pb.NewWebmicrozhiliaoService("user_micro_zhiliao", rpcService.Client())
	userClient := userPb.NewUsermicrozhiliaoService("user_micro_zhiliao", rpcService.Client())

	resp, err := userClient.Call(context.TODO(), &userPb.CallRequest{Name: "userJson.Name"})
	if err != nil {
		fmt.Println("调用微服务出错", err)
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resp,
	})
}
