package apis

import (
	"fmt"
	"ReportGin2nd/util"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoginHandler(ctx *gin.Context) {
	var form util.Login
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": 2004,
			"msg":    err.Error(),
		})
		fmt.Println(form)
		return
	}

	if !(form.User == "user" && form.Password == "password") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": 2006,
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
			"status": 2007,
			"msg":    err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 2000,
		"msg":    "success",
		"token":  token,
	})
}

func LogoutHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 2018,
		"msg":    "logout",
		"token":  "",
	})
}

