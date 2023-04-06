package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/server/grpc"
	"user_micro_zhiliao/handler"
	pb "user_micro_zhiliao/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "user_micro_zhiliao"
	version = "latest"
)

func main() {
	// consul
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Server(grpc.NewServer()),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register handler
	if err := pb.RegisterUsermicrozhiliaoHandler(srv.Server(), new(handler.Usermicrozhiliao)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
