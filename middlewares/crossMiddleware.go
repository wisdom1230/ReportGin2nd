package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrossDomain() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Accept, Authorization, Origin, X-Requested-With, Content-Type, Content-Length, X-CSRF-Token, Token, Session")
			ctx.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Content-Length")
			ctx.Header("Access-Control-Allow-Credentials", "false")
			ctx.Set("Content-type", "application/json")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
