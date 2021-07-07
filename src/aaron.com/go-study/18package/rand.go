package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数
func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子数据
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(1000))
	}

	// 获取10-50区间随机数
	num1 := rand.Intn(50) + 10
	fmt.Println(num1)
}
