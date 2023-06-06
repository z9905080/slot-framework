package adapter

import (
	"github.com/asim/go-micro/v3"
	protoGameLogic "slot-framework/pkg/protobuf/proto_gen/game_logic"
)

type GrpcServer struct {
	micro.Service
}

// NewMicroServer
func NewMicroServer(service micro.Service, gameLogicHandler protoGameLogic.GameLogicServiceHandler) *GrpcServer {
	if err := protoGameLogic.RegisterGameLogicServiceHandler(service.Server(), gameLogicHandler); err != nil {
		panic(err)
	}

	return &GrpcServer{Service: service}
}
