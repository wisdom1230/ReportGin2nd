package middlewares

import (
	"ReportGin2nd/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": 3003,
				"msg":    "token为空",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": 3004,
				"msg":    "token格式错误",
			})
			ctx.Abort()
			return
		}

		claims, err := util.ParseToken(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": 3005,
				"msg":    "无效token",
			})
			ctx.Abort()
			return
		}
		if err := util.ValidateUser(*claims); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status": 3006,
				"msg":    "用户名或者密码错误",
			})
		}

		ctx.Set("genome", claims.Login.User)
		ctx.Next()
	}
}
