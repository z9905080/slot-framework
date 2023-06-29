package adapter

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	microService "slot-framework/pkg/micro_service"
	protoGameLogic "slot-framework/pkg/protobuf/proto_gen/game_logic"
)

type GrpcServer struct {
	micro.Service
}

// NewMicroServer
func NewMicroServer(reg registry.Registry, gameLogicHandler protoGameLogic.GameLogicServiceHandler) *GrpcServer {
	service := microService.NewService(microService.GameLogic, reg)
	if err := protoGameLogic.RegisterGameLogicServiceHandler(service.Server(), gameLogicHandler); err != nil {
		panic(err)
	}

	return &GrpcServer{Service: service}
}
