package main

import (
	"fmt"
	"sync"
)

var (
	c  chan int
	wg sync.WaitGroup
)

// 1.对一个关闭的通道再发送值就会导致panic。
// 2.对一个关闭的通道进行接收会一直获取值直到通道为空。
// 3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
// 4.关闭一个已经关闭的通道会导致panic。
func main() {
	fmt.Println("1、channel声明及使用")
	wg.Add(1)
	c = make(chan int) // 无缓冲的通道
	// c = make(chan int, 10) // 有缓冲的通道
	go recv(c)
	c <- 66 // 发送值
	fmt.Println("发送成功...")
	wg.Wait()
	//close(c) // 关闭通道

	fmt.Println("1、channel练习---通道循环取值")
	chan1()
}

func recv(x chan int) {
	defer wg.Done()
	ret := <-c // 接收值
	fmt.Println("接收成功...", ret)
}

func chan1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~50的数发送到ch1中
	go func() {
		for i := 0; i < 50; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			// 通道关闭后再取值ok=false
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
