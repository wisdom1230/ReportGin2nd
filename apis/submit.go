package apis

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SubmitHandler(ctx *gin.Context) {
	var container ContainerForm2
	var db *sql.DB
	insert := "insert inot kpis(`quarter`, `index`, `department`, `performance`, `certification`, `creation`) values (?, ?, ?, ?, ?, ?);"
	dsn := "root:12345678@tcp(127.0.0.1:3306)/erp?charset=utf8&parseTime=True"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	ctx.ShouldBind(&container)

	for _, item := range container.Kpi {
		fmt.Printf("entry:\t%v\n", item)
		var destFileName string
		if item.Certification == "NoneCertification" {
			destFileName = ""
		} else {
			sourceFile, err := filepath.Abs(path.Join("certification", item.Department, item.Certification))
			if err != nil {
				panic(err.Error())
			}
			err = os.Rename(sourceFile, destFileName)
			if err != nil {
				panic(err.Error())
			}
		}
		_, err := db.Exec(insert, item.Quarter, item.Index, item.Department, item.Performance, item.Certification, item.Creation)
		if err != nil {
			panic(err.Error())
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 2021,
		"msg":    "success",
	})
}
