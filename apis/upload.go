package apis

import (
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	department := ctx.Param(file.Filename)
	fileName, err := filepath.Abs(path.Join(".", "cert", department, file.Filename))
	if err != nil {
		panic(err)
	}
	ctx.SaveUploadedFile(file, fileName)
	ctx.JSON(http.StatusOK, gin.H{
		"status":   2000,
		"msg":      "success",
		"filename": file.Filename,
	})
}
