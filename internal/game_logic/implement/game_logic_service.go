package implement

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	microService "slot-framework/pkg/micro_service"
)

func NewGameLogicServiceServer(registry registry.Registry) micro.Service {
	service := microService.NewService(microService.GameLogic, registry)
	return service
}
