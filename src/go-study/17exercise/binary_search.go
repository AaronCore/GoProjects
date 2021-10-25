package main

import (
	"fmt"
	"math/rand"
	"time"
)

// binary search(二分查找)
func main() {
	list := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		list = append(list, i+1)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		v := rand.Intn(1_000_000-1) + 1
		fmt.Printf("针对 %d 进行二分查找：\n", v)
		idx := binarySearch(list, v)
		fmt.Printf("%d 的索引位置是：[%d]\n", v, idx)
		fmt.Println("---------------------------------")
	}
}

// 二分查找：数组必须是有序
func binarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	step := 0
	for {
		step = step + 1
		if low <= high {
			mid := (low + high) / 2
			guess := list[mid]
			if guess == target {
				fmt.Printf("一共查找了 %d 次\n", step)
				return mid
			}
			if guess > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}
