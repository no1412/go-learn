package main

import (
	"fmt"
	"os"
)

/*
*
性能优化：Go 语言如何进行代码检查和优化？
*/
func main() {
	//checkTest()
	//newStringEscape()
}

const name = "阿萨德"

func checkTest() {
	os.Mkdir("tmp", 0666)
}

// 逃逸分析 go build -gcflags="-m -l" ./ch19/main.go
func newStringEscape() *string {
	s := new(string)
	*s = "一个字符串"
	return s
}

// 逃逸分析 go build -gcflags="-m -l" ./ch19/main.go
func newStringNotEscape() string {
	s := new(string)
	*s = "一个字符串"

	fmt.Println("啊啊啊啊")
	return *s
}
