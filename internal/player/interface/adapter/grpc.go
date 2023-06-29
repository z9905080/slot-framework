package adapter

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	microService "slot-framework/pkg/micro_service"
	protoGamePlayer "slot-framework/pkg/protobuf/proto_gen/player"
)

type GrpcServer struct {
	micro.Service
}

// NewMicroServer
func NewMicroServer(reg registry.Registry, gameLogicHandler protoGamePlayer.PlayerServiceHandler) *GrpcServer {
	service := microService.NewService(microService.Player, reg)
	if err := protoGamePlayer.RegisterPlayerServiceHandler(service.Server(), gameLogicHandler); err != nil {
		panic(err)
	}
	return &GrpcServer{Service: service}
}
