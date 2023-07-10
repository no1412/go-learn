package main

import (
	"fmt"
	"unicode/utf8"
)

/*
集合类型：如何正确使用 array、slice 和 map
*/
func main() {

	// 数组声明
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)

	array1 := []string{1: "b", 3: "d"}
	fmt.Println(array1)

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}

	fmt.Println()

	// 数组循环
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	fmt.Println()

	for _, v := range array {
		fmt.Printf("对应值:%s\n", v)
	}

	fmt.Println()

	// 基于数组生成切片  slice:=array[start:end]
	arrays := [5]string{"a", "b", "c", "d", "e"}
	slice := arrays[2:5]
	fmt.Println(slice)

	// 切片修改
	slice[1] = "f"
	fmt.Println(arrays)

	fmt.Println()

	// 切片声明
	slice1 := make([]string, 4, 8)
	fmt.Println(slice1)
	slice2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice2)

	fmt.Println()

	// append
	//追加一个元素
	//slice3 := append(slice1, "f")
	//多加多个元素
	slice4 := append(slice1, "z", "g", "l", "x")
	//追加另一个切片
	//slice5 := append(slice1, slice...)
	//fmt.Println(slice3)
	fmt.Println(slice4)
	//fmt.Println(slice5)

	fmt.Println()

	// Map 声明初始化
	//nameAgeMap := make(map[string]int)
	//nameAgeMap["飞雪无情"] = 20
	nameAgeMap := map[string]int{"飞雪无情": 20}

	//添加键值对或者更新对应 Key 的 Value

	nameAgeMap["飞雪无情"] = 30

	//获取指定 Key 对应的 Value

	age, ok := nameAgeMap["null"]
	fmt.Println(age)
	if !ok {
		fmt.Println("value not exist", ok)
	}

	//
	//测试 for range
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	fmt.Println()

	// String 和 []byte
	s := "Hello飞雪无情"
	bs := []byte(s)

	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, i2 := range s {
		fmt.Println(i, i2)
	}
}
