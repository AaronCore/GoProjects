package main

import "fmt"

// 冒泡排序
func main() {
	a := [...]int{5, 9, 6, 8, 12, 42}
	fmt.Println(a)
	num1 := len(a)
	for i := 0; i < num1; i++ {
		for j := i + 1; j < num1; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
