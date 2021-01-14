package main

import "fmt"

func main() {
	// 遍历

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			fmt.Println("跳出循环")
			break
		}
		if i == 3 {
			fmt.Println("跳过循环")
			continue
		}
	}

	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	var str = "abc深圳市"
	for _, v := range str {
		fmt.Printf("value：%c \n", v)
	}

	flag := false
	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				flag = true
				break
			}
			fmt.Printf("%v-%c \n", i, j)
		}
		if flag {
			break
		}
	}

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
}
