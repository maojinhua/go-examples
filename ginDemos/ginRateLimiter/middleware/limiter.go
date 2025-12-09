package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maojinhua/ginRateLimiter/limiter"
)

func Limiter(l *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.Allow() {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "可用令牌数不足",
			})
			c.Abort()
		}
		c.Next()
	}

}
