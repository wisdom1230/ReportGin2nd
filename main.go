package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/greeting", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"greeting": "Hello World!",
		})
	})
	router.Run(":8000")
}
