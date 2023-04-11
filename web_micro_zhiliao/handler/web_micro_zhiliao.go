package handler

import (
	"context"
	"go-micro.dev/v4/logger"

	pb "web_micro_zhiliao/proto"
)

type Webmicrozhiliao struct{}

func (e *Webmicrozhiliao) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Webmicrozhiliao.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
