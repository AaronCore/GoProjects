package main

import "fmt"

func main() {
	fmt.Println("1、for")
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				goto over
			}
			fmt.Printf("%v-%c \n", i, j)
		}
	}

over:
	fmt.Println("结束了")

	fmt.Println("2、range")
	s := "abcdef"
	for i, v := range s {
		fmt.Printf("下标：%d,值：%v \n", i, string(v))
	}
}
