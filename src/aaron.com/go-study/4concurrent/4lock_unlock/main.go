package main

// 互斥锁

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	x      = 0
	ticket = 100
	wg     sync.WaitGroup
	lock   sync.Mutex
)

func main() {
	//fmt.Println("-------------------sample1-------------------")
	//wg.Add(3)
	//go sum()
	//go sum()
	//go sum()
	//wg.Wait()
	//fmt.Println("sum：", x)

	fmt.Println("-------------------sample2-------------------")
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")
	wg.Wait()
	fmt.Println("程序结束了")
}

func sum() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func saleTickets(name string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		lock.Lock() //上锁
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			lock.Unlock() //条件不满足，也要解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		lock.Unlock() //解锁
	}
}
