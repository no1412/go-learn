package main

import "testing"

/*
*
1. 含有单元测试代码的 go 文件必须以 _test.go 结尾，Go 语言测试工具只认符合这个规则的文件。
2. 单元测试文件名 _test.go 前面的部分最好是被测试的函数所在的 go 文件的文件名，比如以上示例中单元测试文件叫 main_test.go，因为测试的 Fibonacci 函数在 main.go 文件里。
3. 单元测试的函数名必须以 Test 开头，是可导出的、公开的函数。
4. 测试函数的签名必须接收一个指向 testing.T 类型的指针，并且不能返回任何值。
5. 函数名最好是 Test + 要测试的函数名，比如例子中是 TestFibonacci，表示测试的是 Fibonacci 这个函数。

单元测试
go test -v ./xxgo_folder

单元测试覆盖率
go test -v --coverprofile=ch18.cover ./xxgo_folder

生成HTML 格式的单元测试覆盖率报告
go tool cover -html=ch18.cover -o=xxx.html
*/
func TestFibonacci(t *testing.T) {
	//预先定义的一组斐波那契数列作为测试用例
	fsMap := map[int]int{}
	fsMap[-1] = 0
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}
	}
}

/*
*
基准测试
go test -bench=.
*/
func BenchmarkFibonacci(b *testing.B) {
	n := 10

	b.ReportAllocs() //开启内存统计

	b.ResetTimer() //重置计时器

	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

/*
*
并发基准测试
*/
func BenchmarkFibonacciRunParallel(b *testing.B) {
	//b.ReportAllocs() //开启内存统计
	n := 10
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FibonacciRefine(n)
		}
	})
}
