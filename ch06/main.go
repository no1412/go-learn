package main

import (
	"fmt"
)

/**
struct 和 interface：结构体与接口都实现了哪些功能
*/

// 结构体
type person struct {
	name string
	age  int
	addr address
}

type address struct {
	province string
	city     string
}

// 接口
type stringor interface {
	toString() string
}

func (p person) toString() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func (addr address) toString() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s stringor) {
	fmt.Println(s.toString())
}

func main() {
	p := person{"张三", 12, address{"河南省", "郑州市"}}
	fmt.Println(p)

	printString(p)
	printString(p.addr)
}
