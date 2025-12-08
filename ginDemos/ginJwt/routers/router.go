package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginJwt/web"
)

func InitRouters(r *gin.Engine) {
	r.GET("/login",web.Login)
	InitCourse(r)
	InitUser(r)
	// // 路由分组：
	// // 	1.根据模块划分接口
	// // 	2.根据版本划分接口
	// course := r.Group("/course")
	// {
	// 	course.GET("/", getCourse)
	// }

	// r.Group("/user").GET("/", getUser)

}
