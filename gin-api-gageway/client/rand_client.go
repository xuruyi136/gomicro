package client

import (
	"gomicro/proto/rand"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var (
	randClient rand.RandService
)

func GetRandClient() rand.RandService {
	if randClient == nil {
		consulRegistry := consul.NewRegistry(
			registry.Addrs("121.196.195.242:8500"),
		)
		service := micro.NewService(
			micro.Registry(consulRegistry))
		randClient = rand.NewRandService("go.micro.ocean.rand", service.Client())
	}
	return randClient
}
