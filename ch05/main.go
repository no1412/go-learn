package main

import (
	"errors"
	"fmt"
)

/*
*
函数和方法：Go 语言中的函数和方法到底有什么不同
*/
func sum(a int, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a,b 不能小于0")
	}
	return a + b, nil
}

func sum1(params ...int) (sum int, err error) {
	for _, param := range params {
		sum += param
	}
	return sum, err
}

// 匿名函数和闭包
func increase() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 不同于函数的方法
type Age int

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(20)
}

func main() {
	result, err := sum(1, -2)
	fmt.Println(result, err)

	result1, err1 := sum1(1, 2, 3, 4)
	fmt.Println(result1, err1)

	inc := increase()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println()

	age := Age(25)
	age.String()
	age.Modify()
	age.String()

}
