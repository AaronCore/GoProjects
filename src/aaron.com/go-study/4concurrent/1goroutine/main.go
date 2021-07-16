package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		WaitGroup：同步等待组
			Add()，设置等待组中要执行的子 goroutine的数量
			Wait()，让主goroutine出于等待
			Done()，让等待组中的counter计数器的值减1，同Add(-1)
	*/
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go printNo(i)
	}
	wg.Wait() // 等待所有登记的 goroutine 都结束
}

func printNo(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("No:", i)
}
