package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4/registry"
	"product_micro_zhiliao/handler"
	pb "product_micro_zhiliao/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "product_micro_zhiliao"
	version = "latest"
)

func main() {
	//consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// router

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Server(grpc.NewServer()),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	if err := pb.RegisterProductmicrozhiliaoHandler(srv.Server(), new(handler.Productmicrozhiliao)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
