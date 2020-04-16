package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "accept,x-requested-with,Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		ctx.Next()
	}
}
