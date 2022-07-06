package apis

import (
	"localKPI/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var form util.Login
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 2007,
			"msg":    err.Error(),
		})
		return
	}
	if !(form.User == "user" && form.Password == "password") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": 2009,
			"msg":    "用户名或者密码错误",
		})
		return
	}
	var login util.Login
	login.User = "user"
	login.Password = "password"
	token, err := util.GenerateToken(login)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 2008,
			"msg":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 2005,
		"msg":    "success",
		"token":  token,
	})
}
