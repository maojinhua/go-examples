package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context) {

	id := ctx.DefaultQuery("id", "0")
	ctx.JSON(
		http.StatusOK, gin.H{"id": id},
	)

}

type Course struct {
	Name    string `form:"name" binding:"required"`
	Teacher string `form:"teacher" binding:"required"`
	Price   string `form:"price" binding:"number"`
}

func AddCourse(ctx *gin.Context) {
	req := &Course{}
	// ShouldBind 发生错误的时候可以让我们自己处理错误
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	// 如果发生错误 Bind 直接返回，不会让我们自己处理错误
	// err := ctx.BindJSON(req)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }
	// id := ctx.DefaultQuery("id", "0")
	// ctx.JSON(
	// 	http.StatusOK, gin.H{"req": req},
	// )
	resp := &AddCourseResponse{
		Name: req.Name,
		Teacher: req.Teacher,
		Duration: "Duration",
	}
	ctx.ProtoBuf(http.StatusOK,resp)
}
