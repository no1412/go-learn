package main

import (
	"fmt"
	"sync"
	"time"
)

/**
同步原语：sync 包让你对并发控制得心应手
*/

// 共享的资源
var sum = 0
var rwmutex sync.RWMutex

func main() {
	//run()

	//doOnce()

	race()
}

// sync.WaitGroup
func run() {
	var wg sync.WaitGroup

	// 1000 +10 协程
	wg.Add(1010)

	//开启1000个协程让sum+10
	for i := 0; i < 1000; i++ {
		go func() {
			// 释放计数器
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			// 释放计数器
			defer wg.Done()
			go fmt.Println("read sum value", readSum())
		}()
	}

	// 阻塞，等待计数器归零
	wg.Wait()
	fmt.Println("sum value", readSum())
}

func add(i int) {
	rwmutex.Lock()
	sum += i
	rwmutex.Unlock()
}

func readSum() int {
	// 只获取读锁
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	b := sum
	return b
}

// sync.Once
func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// 10个人赛跑，1个裁判发号施令
func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")

			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()

		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}
