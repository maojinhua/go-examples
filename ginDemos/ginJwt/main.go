package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginJwt/middleware"
	"github.com/maojinhua/ginJwt/routers"
)

func main() {
	LoggerMiddleware(LoggerMiddleware(initDB))("sss")
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors1(), middleware.Cors(), CheckAuth("abcde"), CheckAuth2("abcde"))
	routers.InitRouters(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Group("/")
	r.Run()
}

// 闭包形式的中间件的好处：闭包形式的中间件可以传递参数
func CheckAuth(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("CheckAuth 开始", param)
		ctx.Next()
		fmt.Println("CheckAuth 结束")
	}
}

// 闭包形式的中间件可以传递参数
func CheckAuth2(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("CheckAuth222 开始", param)
		ctx.Next()
		fmt.Println("CheckAuth222 结束")
	}
}

// 非闭包函数不能传递数据
// func CheckAuth(ctx *gin.Context){
// 	fmt.Println("call check Auth func")
// 	ctx.Next()
// }

func initDB(connStr string) {
	fmt.Println("初始化数据库", connStr)
}

// 演示中间件原理：使用函数包裹，一个洋葱模型，可以多层嵌套
func LoggerMiddleware(in func(connStr string)) func(connStr string) {
	return func(connStr string) {
		log.Println("打印日志")
		in(connStr)
		log.Println("日志打印结束")
	}
}
