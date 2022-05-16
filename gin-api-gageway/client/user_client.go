package client

import (
	"gomicro/proto/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

var (
	userClient user.UserService
)

func GetUserClient() user.UserService {
	if userClient == nil {
		consulRegistry := consul.NewRegistry(
			registry.Addrs("121.196.195.242:8500"),
		)
		service := micro.NewService(
			micro.Registry(consulRegistry))
		userClient = user.NewUserService("go.micro.ocean.user", service.Client())
	}
	return userClient
}
