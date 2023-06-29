package micro_service

import (
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
)

type MicroServiceName string

func (n MicroServiceName) String() string {
	return string(n)
}

const (
	GameLogic = MicroServiceName("game_logic")
	Player    = MicroServiceName("player")
)

// NewService
func NewService(serviceName MicroServiceName, reg registry.Registry) micro.Service {
	myClient := client.NewClient(
		client.PoolSize(500),
		client.RequestTimeout(300*time.Second),
		client.Retry(client.RetryOnError),
	)

	opts := []micro.Option{
		micro.Name(serviceName.String()), // The service name to register in the registry_provider
		micro.Client(myClient),
	}

	if reg != nil {
		opts = append(opts, micro.Registry(reg))
	}

	//destOpts := make([]micro.Option, len(opts))
	//copy(destOpts, opts)
	//spew.Dump(destOpts)

	service := micro.NewService(opts...)
	//service.Init()
	return service
}
