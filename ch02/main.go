package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
数据类型：你必须掌握的数据类型有哪些
*/
func main() {
	// var 变量名 类型 = 表达式
	var i int = 10
	fmt.Print(i)

	// 浮点数
	var f32 float32 = 2.2
	var f64 float64 = 10.3456

	fmt.Println("f32: ", f32, "f64: ", f64)

	// 布尔型
	var bf bool = false
	var bt bool = true
	fmt.Println("bf:", bf, "bt:", bt)

	// 字符串
	var s1 string = "hello"
	var s2 string = "world"
	fmt.Println("s1+s2:", s1+s2)

	// 零值
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	// 变量简短声明
	si := 10
	sbf := false
	ss1 := "Hello"
	fmt.Println(si, sbf, ss1)

	// 指针
	pi := &i
	fmt.Println(*pi)

	// 赋值
	i = 20
	fmt.Println("i的新值是", i)

	// 常量的定义
	const name = "飞雪无情"

	// iota
	const (
		one   = iota + 1
		two   = 2
		three = 3
		four  = 4
	)
	fmt.Println(one, two, three, four)

	// 字符串和数字互转
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	// Strings 包
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1, "H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1, "o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
}
