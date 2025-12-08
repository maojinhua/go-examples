package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginJwt/util"
)

var token = "1234567"

// jwt 权限校验
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("access_token")
		fmt.Println("access_token ", accessToken)
		jwtPayload := util.JwtPayload{}
		err := util.Verify(accessToken,&jwtPayload)
		if err != nil {
			log.Println("jwt 认证失败 ",err)
			ctx.JSON(http.StatusForbidden,gin.H{
				"error":"身份认证失败",
			})
			ctx.Abort()
		}
		// if accessToken!=token {
		// 	ctx.JSON(http.StatusForbidden,gin.H{
		// 		"message":"token 校验失败",
		// 	})
		// 	ctx.Abort()
		// }
		fmt.Println("jwtPayload ",jwtPayload)
		ctx.Set("auth_info",jwtPayload)
		ctx.Next()
	}
}
