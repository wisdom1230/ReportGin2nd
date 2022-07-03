package apis

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayHandler(ctx *gin.Context) {
	var query conform
	var result ReturnDisplay
	var returns []ReturnDisplay
	var quarters = []string{"2022Q1", "2022Q2", "2022Q3", "2022Q4"}
	var db *sql.DB
	selects := "select * from `kpis` where `quarter`=? and `department`=? and `index`=? order by `creation` desc, `id` desc limit 1;"
	department := ctx.Param("department")
	dsn := "root:123346@tcp(127.0.0.1)/erp?charset=utf8&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	for _, quarter := range quarters{
		for _, index_ := range Order[department]
		{
			err := db.QueryRow(selects, quarter, department, index_).Scan(&query.Quarter, &query.Department, &query.Creation, &query.Index, &query.Id, &query.Performance, &query.Certification)
			if err != nil {
				if err.Error() == "sql: no rows in result set" {
					result.Date = quarter
					result.File = ""
					result.Index = index_
					result.Number = "0"
					returns = append(returns, result)
				}else{
					panic(err)
				}
			}else{
				result.Date = quarter.Quarter
				result.File = query.Certification
				result.Index = query.Index
				result.Number = query.Performance
				result.Unit = Indexes[department][index_]
				returns = append(returns, result)
			}
		}
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"status":2200,
			"msg":"success",
			"data":returns,
		}
	)
}
