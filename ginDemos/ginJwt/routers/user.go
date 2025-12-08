package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginJwt/middleware"
	"github.com/maojinhua/ginJwt/web"
)

func InitUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.Auth())
	v1.Group("/users").GET("/", web.GetUser).POST("/", web.AddUser)
}
