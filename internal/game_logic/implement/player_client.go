package implement

import (
	"github.com/asim/go-micro/v3/registry"
	microService "slot-framework/pkg/micro_service"
	protoGamePlayer "slot-framework/pkg/protobuf/proto_gen/player"
)

func NewPlayerService(reg registry.Registry) protoGamePlayer.PlayerService {
	service := microService.NewService(microService.Player, reg)
	gameLogicSrv := protoGamePlayer.NewPlayerService(service.Name(), service.Client())
	return gameLogicSrv
}
