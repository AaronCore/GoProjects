package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置逻辑CPU数量
}

func main() {
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS) // 获取操作系统
	fmt.Println("逻辑CPU数量:", runtime.NumCPU())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine()) // 获取当前存在的Go程数

	fmt.Println("1、runtime.Gosched()")
	goSched()

	fmt.Println("2、runtime.Goexit()")
	goExit()

	fmt.Println("3、runtime.GOMAXPROCS()")
	go f1()
	go f2()
	time.Sleep(time.Second)
}

func f1() {
	for i := 1; i < 10; i++ {
		fmt.Println("f1():", i)
	}
}

func f2() {
	for i := 1; i < 10; i++ {
		fmt.Println("f2():", i)
	}
}

func goExit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
}

func goSched() {
	go func(x string) {
		for i := 0; i < 5; i++ {
			fmt.Println(x)
		}
	}("world...")
	for i := 0; i < 5; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello...")
	}
}
