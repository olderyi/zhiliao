package main

import (

	"seckill_micro_zhiliao/handler"
	pb "seckill_micro_zhiliao/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "seckill_micro_zhiliao"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterSeckillmicrozhiliaoHandler(srv.Server(), new(handler.Seckillmicrozhiliao)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
