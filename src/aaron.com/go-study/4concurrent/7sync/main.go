package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Once
// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
	m        sync.Map
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(x int) {
			key := strconv.Itoa(x)
			m.Store(key, x)
			v, _ := m.Load(key)
			fmt.Printf("key：%s,value：%d\n", key, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
