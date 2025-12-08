package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	id := ctx.DefaultQuery("id","0")
	ctx.JSON(
		http.StatusOK, gin.H{"id": id},
	)
}

type User struct {
	Name    string `form:"name" binding:"required"`
	// e164 表示手机号码格式
	Phone string `form:"phone" binding:"required,e164"`
}

func AddUser(ctx *gin.Context) {
	req := &Course{}
	// ShouldBind 发生错误的时候可以让我们自己处理错误
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(
		http.StatusOK, gin.H{"req": req},
	)
}