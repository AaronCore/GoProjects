package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	c  chan int
	wg sync.WaitGroup
)

/*
	chan声明定义
		声明通道：
			var 通道名 chan 数据类型
		创建通道：如果通道为nil(就是不存在)，就需要先创建通道
			通道名 = make(chan 数据类型)
	双向chan：
	chan发送、接收，chan在<-左为发送，在<-右接收
		chan T -->
			chan <- date，写，写出数据
				例：a <- data
			data <- chan，读, 读出数据
				例：data := <- a
				例：v, ok := <- a
	单向chan：
		1.chan <- T，只写
			//该函数接收，只写的通道
			func fun1(ch chan <- int){
				// 函数内部，对于ch只能写数据，不能读数据
				ch <- 100
				fmt.Println("fun1函数结束。。")
			}
		2.<- chan T，只读

	单、双向chan声明及使用：
		1.双向chan
			ch1 := make(chan string)
				ch1 <- 100
				data :=<-ch1
		2.单向，只写
			ch2 := make(chan <- int) // 单向，只写，不能读
				ch2 <- 1000
				data := <- ch2 // error：<-ch2 //invalid operation: <-ch2 (receive from send-only type chan<- int)
			// 只写chan函数
			func fun1(ch chan <- int){
				// 函数内部，对于ch只能写数据，不能读数据
				ch <- 100
				fmt.Println("fun1函数结束。。")
			}
		3.单向，只读
			ch3 := make(<- chan int) //单向，只读，不能写
				ch3 <- 100 // error：ch3 <- 100 //invalid operation: ch3 <- 100 (send to receive-only type <-chan int)
			// 只读chan函数
			func fun2(ch <-chan int){
				//函数内部，对于ch只能读数据，不能写数据
				data := <- ch
				fmt.Println("fun2函数，从ch中读取的数据是：",data)
			}
	chan注意点：
		1.用于goroutine，传递消息的。
		2.通道，每个都有相关联的数据类型, nil chan，不能使用，类似于nil map，不能直接存储键值对。
		3.使用通道传递数据：<- chan <- data,发送数据到通道。向通道中写数据 data <- chan,从通道中获取数据。从通道中读数据。
		4.阻塞： 发送数据：chan <- data,阻塞的，直到另一条goroutine，读取数据来解除阻塞 读取数据：data <- chan,也是阻塞的。直到另一条goroutine，写出数据解除阻塞。
		5.本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。
		6.对一个关闭的通道再发送值就会导致panic。
		7.对一个关闭的通道进行接收会一直获取值直到通道为空。
		8.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		9.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
*/

func main() {
	fmt.Println("1、channel声明及使用")
	wg.Add(1)
	c = make(chan int) // 无缓冲的通道
	// c = make(chan int, 10) // 有缓冲的通道
	go rev(c)
	c <- 66 // 发送值
	fmt.Println("发送成功...")
	wg.Wait()
	//close(c) // 关闭通道

	fmt.Println("2、channel练习---通道循环取值")
	chan1()

	fmt.Println("2、channel练习---单向通道")
	chan2()
}

func rev(x chan int) {
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

// 单向通道
func chan2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1) // 只写
	go squarer(ch1, ch2)
	printer(ch2) // 只读
}

// 1.ch1 chan<- int是一个只能发送的通道，可以发送但是不能接收
func counter(ch1 chan<- int) {
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

// ch2 chan<- int是一个只能接收的通道，可以接收但是不能发送
func printer(ch2 <-chan int) {
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 1.ch1 <-chan int是一个只能接收的通道，可以接收但是不能发送。
// 2.ch2 chan<- int是一个只能发送的通道，可以发送但是不能接收；
func squarer(ch1 <-chan int, ch2 chan<- int) {
	for i := range ch1 {
		ch2 <- i * i
	}
	close(ch2)
}

// goroutine里处理错误
func sample() {
	e := make(chan error)
	wgDone := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
	}()
	go func() {
		err := returnError()
		if err != nil {
			e <- err
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-e:
		close(e)
		fmt.Println(err)
	}

	time.Sleep(time.Second)
}

func returnError() error {
	return errors.New("报错了")
}
