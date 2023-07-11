package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
*
Context：你必须掌握的多线程并发控制神器
*/
func main() {
	//watchDogTest1()
	//watchDogTest2()
	//watchDogTest3()
	getUserTest()
}

func watchDogTest() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool) //用来停止监控狗
	go func() {
		defer wg.Done()
		watchDog(stopCh, "【监控狗1】")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stopCh <- true              //发停止指令
	wg.Wait()
}
func watchDog(stopCh chan bool, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

// 使用context取消协程
func watchDogTest2() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog2(ctx, "【监控狗1】")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}

// 使用context取消多个协程
func watchDogTest3() {
	var wg sync.WaitGroup
	wg.Add(3)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog2(ctx, "【监控狗1】")
	}()
	go func() {
		defer wg.Done()
		watchDog2(ctx, "【监控狗2】")
	}()
	go func() {
		defer wg.Done()
		watchDog2(ctx, "【监控狗3】")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}
func watchDog2(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

func getUserTest() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	wg.Add(1) //记得这里要改为4，原来是3，因为要多启动一个协程

	//省略其他无关代码
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()
	wg.Wait()
}
func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户ID为：", userId)
			time.Sleep(1 * time.Second)
		}
	}
}
