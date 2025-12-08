package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginJwt/web"
)

func InitCourse(r *gin.Engine){
	v1:=r.Group("/v1")
	v1.Group("/course").GET("/",web.GetCourse).POST("/",web.AddCourse)
}