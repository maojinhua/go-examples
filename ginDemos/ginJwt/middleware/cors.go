package middleware

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors1() gin.HandlerFunc{
	return  func(ctx *gin.Context) {
		fmt.Println("call corse")
	}

}

func Cors()gin.HandlerFunc{
	return cors.New(
		cors.Config{
			// 同意所有的域名
			AllowAllOrigins: true,
			AllowHeaders: []string{
				"Origin","Content-Length","Content-Type",
			},
			AllowMethods: []string{
				"GET","POST","DELETE","PUT","HEAD","OPTIONS",
			},
		},
	)
}