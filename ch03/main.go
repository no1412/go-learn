package main

import "fmt"

/*
控制结构：if、for、switch 逻辑语句的那些事儿
*/
func main() {

	// if 条件语句

	if i := 7; i > 10 {
		fmt.Println("i>10")
	} else if i > 5 {
		fmt.Println("i>5 and i <=10")
	} else {
		fmt.Println("i<5")
	}

	// switch 选择语句
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}

	// for 循环语句
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 500 {
			break
		}
	}
	fmt.Println("the sum is", sum)
}
