package handler

import (
	"context"
	"gomicro/proto/rand"
	"gomicro/proto/user"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	randClient rand.RandService
	userClient user.UserService
}

func GetAPIHandler(randservice rand.RandService, userService user.UserService) *APIHandler {
	return &APIHandler{
		randClient: randservice,
		userClient: userService,
	}
}

//获取随机数
func (s *APIHandler) Rand(ctx *gin.Context) {
	request := rand.RandRequest{}
	_ = ctx.ShouldBindQuery(&request)
	response, err := s.randClient.GetRand(context.Background(), &request)
	if err != nil {
		log.Println(err)
	}
	//上下文读取username
	username, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"result":   response.GetMax(),
	})

}

//注册
func (s *APIHandler) RegistryUser(ctx *gin.Context) {
	request := user.UserRegistryRequest{}

	_ = ctx.ShouldBindJSON(&request)
	response, err := s.userClient.UserRegistry(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": response.GetMessage(),
		"token":   response.GetToken(),
	})
}

//登录
func (s *APIHandler) LoginUser(ctx *gin.Context) {
	request := user.UserLoginRequest{}
	_ = ctx.ShouldBindJSON(&request)
	response, _ := s.userClient.UserLogin(context.Background(), &request)
	ctx.JSON(http.StatusOK, gin.H{
		"message": response.GetMessage(),
		"token":   response.GetToken(),
	})
}
