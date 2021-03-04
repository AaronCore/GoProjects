package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

// 读
func read() {
	defer wg.Done()
	//lock.Lock()
	rwLock.RLock()
	//fmt.Println("r：", x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
}

// 写
func write() {
	defer wg.Done()
	//lock.Lock()
	rwLock.Lock()
	x += 1
	//fmt.Println("w：", x)
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()
	rwLock.Unlock()
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	//time.Sleep(time.Second * 1)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}
