package main

import "fmt"

/*
*
指针详解：在什么情况下应该使用指针？
*/
func main() {
	point()
}

func point() {
	name := "飞雪无情"
	nameP := &name //取地址
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)
}
