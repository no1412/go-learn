package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
*
SliceHeader：slice 如何高效处理数据？
*/
func main() {
	//sliceTest()
	//sliceAutoResizeTest()
	byteToString()
}

// slice test
func sliceTest() {
	ss := []string{"飞雪无情", "张三"}
	ss = append(ss, "李四", "王五")
	fmt.Println(ss)
}

// slice auto resize
func sliceAutoResizeTest() {
	ss := []string{"飞雪无情", "张三"}
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	ss = append(ss, "李四", "王五")
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	fmt.Println(ss)
}

func byteToString() {
	s := "飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}
