package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("1、runtime.Gosched()")
	gosched()

	fmt.Println("2、runtime.Goexit()")
	goexit()

	fmt.Println("3、runtime.GOMAXPROCS()")
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func goexit() {
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

func gosched() {
	go func(x string) {
		for i := 0; i < 3; i++ {
			fmt.Println(x)
		}
	}("world...")
	for i := 0; i < 3; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello...")
	}
}
