package main

import (
	"errors"
	"fmt"
	"strconv"
)

/**
错误处理：如何通过 error、deferred、panic 等处理错误
*/

// error工厂函数
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能为负数")
	} else {
		return a + b, nil
	}
}

// 自定义error
type commonError struct {
	errorCode int    // 错误码
	errorMsg  string // 错误信息
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

func add1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{errorCode: 1, errorMsg: "a或者b不能为负数"}
	} else {
		return a + b, nil
	}
}

// 错误嵌套
type MyError struct {
	err error
	msg string
}

func (me *MyError) Error() string {
	return me.err.Error() + me.msg
}

func main() {
	i, err := strconv.Atoi("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	_, err1 := add(-1, -2)
	fmt.Println(err1)
	_, err2 := add1(-1, -2)
	fmt.Println(err2)

	fmt.Println("---------------")

	sum, err4 := add1(-1, -2)
	if cm, ok := err4.(*commonError); ok {
		fmt.Println(sum, cm.errorCode)
	}

	fmt.Println("---------------")

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	// 判断error是否匹配
	fmt.Println(errors.Is(w, e))
	fmt.Println(errors.As(w, &e))
}
