package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
并发模式：Go 语言中即学即用的高效并发模式
*/
func main() {
	//selectTimeout()
	//hotpot()
	pipeline2()
}

// select timeout 模式
func selectTimeout() {
	result := make(chan string)
	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
}

// 工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

// 工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

// 工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

// 流水线
func pipeline() {
	fittings := buy(10)
	phones := build(fittings)
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}
}

// 扇入函数（组件），把多个chanel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	//把一个channel中的数据发送到out中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	//扇入，需要启动多个goroutine用于处于多个channel中的数据
	for _, cs := range ins {
		go p(cs)
	}
	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 流水线2 打包组改为多个
func pipeline2() {
	fittings := buy(1000)

	phones1 := build(fittings)
	phones2 := build(fittings)
	phones3 := build(fittings)

	// 扇入处理
	mergePhones := merge(phones1, phones2, phones3)

	packs := pack(mergePhones)

	for p := range packs {
		fmt.Println(p)
	}
}
