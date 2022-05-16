package main

import (
	"gomicro/gin-api-gageway/client"
	"gomicro/gin-api-gageway/handler"

	"gomicro/gin-api-gageway/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	//解决回调函数
	randclient := client.GetRandClient()
	userclient := client.GetUserClient()
	apiHandler := handler.GetAPIHandler(randclient, userclient)
	// v1.GET("/v1", apiHandler.Rand)
	user := v1.Group("/user")
	user.POST("/registry", apiHandler.RegistryUser)
	user.POST("/login", apiHandler.LoginUser)
	user.GET("/rand", middleware.AuthMiddleware(), apiHandler.Rand)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
