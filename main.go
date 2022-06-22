package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	token, err := util.GenerateToken(util.Login{User: "user", Password: "password"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	router := gin.Default()
	router.Use(middlewares.CrossDomain())	
	router.GET("/greeting", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"greeting": "Hello World!",
		})
	})
	
	router.POST("/kpi/login", apis.LoginHandler)
	v1 := router.Group("/kpi").Use(middlewares.AuthenticateJWT())
	{
		v1.GET("/", apis.IndexHandler)
		v1.GET("/index", apis.IndexHandler)
		v1.GET("/logout", apis.LogoutHandler)
		v1.GET("/display/:department", apis.DisplayHandler)
		v1.POST("/submit/:department", apis.SubmitHandler)
	}
	router.Run(":8000")
}

