package main

import (
	"gomicro/proto/rand"
	"gomicro/rand-service/handler"
	"gomicro/rand-service/middleware"

	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//注册
	consulRegistry := consul.NewRegistry(
		registry.Addrs("121.196.195.242:8500"),
	)
	middleware.InitSentinel()
	service := micro.NewService(
		micro.Name("go.micro.ocean.rand"),
		micro.Registry(consulRegistry),
		//包装器
		micro.WrapHandler(middleware.LimitingWrapper()),
	)
	//生命周期过程
	service.Init()

	//TODO 暴露接口到对应服务
	_ = rand.RegisterRandHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
