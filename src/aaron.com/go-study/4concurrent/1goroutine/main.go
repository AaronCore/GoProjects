package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	f1()
}

func f1() {
	for i := 1; i <= 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的 goroutine 都结束
}

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("No.", i)
}
