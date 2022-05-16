package main

import (
	"gomicro/proto/user"
	"gomicro/user-service/handler"

	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//注册
	consulRegistry := consul.NewRegistry(
		registry.Addrs("121.196.195.242:8500"),
	)
	//服务名称
	service := micro.NewService(
		micro.Name("go.micro.ocean.user"),
		micro.Registry(consulRegistry),
	)
	//生命周期过程
	service.Init()

	//TODO 暴露接口到对应服务

	_ = user.RegisterUserHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
