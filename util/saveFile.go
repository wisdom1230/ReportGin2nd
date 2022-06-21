package util

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var Prefix string = "./certificate"

func SavePath(prefix, department, quarter string) string {
	target := filepath.Join(prefix, department, quarter)
	_, err := os.Stat(target)
	if err != nil {
		os.MkdirAll(filepath.Join(prefix, department, quarter), os.ModePerm)
	}
	fileName := fmt.Sprint(time.Now().Format("20060102150405")) + "-" + quarter + ".zip"
	return filepath.Join(target, fileName)
}
