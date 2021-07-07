package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

// 普通版加函数
func add() {
	defer wg.Done()
	x++
}

// 互斥锁版加函数
func mutexAdd() {
	defer wg.Done()
	lock.Lock()
	x++
	lock.Unlock()
}

// 原子操作版加函数
func atomicAdd() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	start := time.Now()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add() 		//普通版add函数 不是并发安全的
		//go mutexAdd() 	//加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() //原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()

	end := time.Now()
	fmt.Println("x：", x)
	fmt.Println(end.Sub(start))
}
