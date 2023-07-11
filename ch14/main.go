package main

import "fmt"

/*
*
内存分配：new 还是 make？什么情况下该用谁
*/
func main() {
	// 值类型
	//valueType()

	// 指针类型
	pointTypeThrowError()
}

func valueType() {
	var s string
	s = "张三"
	fmt.Println(s)
}

// 指针类型，没有初始化，直接赋值会报错
func pointTypeThrowError() {
	var s *string
	*s = "张三"
	fmt.Println(*s)

}
