package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginRateLimiter/breaker"
	"github.com/maojinhua/ginRateLimiter/limiter"
	"github.com/maojinhua/ginRateLimiter/middleware"
)

func main(){
	r := gin.Default()

	r.GET("/ping",middleware.Limiter(limiter.NewLimiter(3*time.Second,4,1)),func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})

	b:=breaker.NewBreaker(4,4,2,15*time.Second)
	r.GET("/ping1",func(c *gin.Context) {
		err := b.Exec(func() error {
			value,_:=c.GetQuery("value")
			if value == "a"{
				return errors.New("value 为 a 返回错误")
			}
			return nil
		})
		if err!=nil{
			fmt.Println("err ",err)
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	r.Run()
}