package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
gin 包使用
go get -u github.com/gin-gonic/gin
go mod tidy
*/
func main() {
	fmt.Println("先导入fmt包，才能使用")

	r := gin.Default()

	r.Run()
}
