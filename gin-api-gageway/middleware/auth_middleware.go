package middleware

import (
	"gomicro/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//验证token是否传了
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || tokenString[0:7] != "Bearer " {
			//TODO回传错误401
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			//终止请求
			context.Abort()
			return
		}
		token, claims, err := common.ParseToken(tokenString[7:])
		if err != nil || !token.Valid {
			//TODO回传错误401
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			//终止请求
			context.Abort()
			return
		}
		//user存入上下文
		context.Set("user", claims.UserName)
		context.Next()
	}
}
