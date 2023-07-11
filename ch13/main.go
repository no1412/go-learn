package main

import "fmt"

/*
*
参数传递：值、引用及指针之间的区别
*/
func main() {
	p := person{name: "张三", age: 18}
	fmt.Printf("main函数：p的内存地址为%p\n", &p)
	modifyPerson(p)
	fmt.Println("person name:", p.name, ",age:", p.age)
}

type person struct {
	name string
	age  int
}

func modifyPerson(p person) {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n", &p)
	p.name = "李四"
	p.age = 20
}
