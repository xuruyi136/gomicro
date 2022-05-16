package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = "ZXZldmVAQEBqb2kzdTMyNDIzNDI="

type DceanClaims struct {
	UserName string `json:"username"`
	Uid      uint   `json:"id"`
	jwt.StandardClaims
}

func GenToken(userName string, uid uint) string {
	claims := DceanClaims{
		userName,
		uid,
		jwt.StandardClaims{
			//过期时间
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			//生成时期
			IssuedAt: time.Now().Unix(),
			//生成者
			Issuer: "Service",
		},
	}
	//token实例化
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//签名加密
	tokenstring, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return tokenstring
}

func ParseToken(tokenString string) (*jwt.Token, DceanClaims, error) {
	claims := DceanClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return token, claims, err
}
