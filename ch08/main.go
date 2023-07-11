package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
并发基础：Goroutines 和 Channels 的声明与使用
*/
func main() {
	//ch := make(chan string)
	//
	//go func() {
	//    go fmt.Println("飞雪无情")
	//    ch <- "goroutine 完成"
	//    ch <- "goroutine 完成1"
	//}()
	//
	//fmt.Println("我是 main goroutine")
	//
	//v := <-ch
	//
	//fmt.Println("接收到的chan中的值为：", v)
	//fmt.Println("接收到的chan中的值为：", <-ch)

	fmt.Println("---------------------------")

	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}
func downloadFile(chanName string) string {
	//模拟下载文件,可以自己随机time.Sleep点时间试试
	duration := time.Duration(rand.Int() % 10000000000)
	fmt.Println(chanName, "wait", duration)
	time.Sleep(duration)
	return chanName + ":filePath"
}
