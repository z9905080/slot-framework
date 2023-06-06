package implement

import (
	"github.com/asim/go-micro/v3/registry"
	microService "slot-framework/pkg/micro_service"
	protoGameLogic "slot-framework/pkg/protobuf/proto_gen/game_logic"
)

func NewGameLogicService(registry registry.Registry) protoGameLogic.GameLogicService {
	service := microService.NewService(microService.GameLogic, registry)
	gameLogicSrv := protoGameLogic.NewGameLogicService(service.Name(), service.Client())
	return gameLogicSrv
}
