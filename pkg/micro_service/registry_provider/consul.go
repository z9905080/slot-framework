package registry_provider

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

func NewConsulRegistry() registry.Registry {
	return consul.NewRegistry(
		registry.Addrs([]string{"127.0.0.1:8500"}...),
	)
}
