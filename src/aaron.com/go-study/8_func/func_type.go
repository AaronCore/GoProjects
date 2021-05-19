package main

import "fmt"

func main() {
	// 函数类型

	a := ff1
	fmt.Printf("ff1：%T \n", a)

	b := ff2
	fmt.Printf("ff2：%T \n", b)

	ff3(ff2)

	c := ff5(ff2)
	fmt.Printf("ff5：%T \n", c)
}

func ff1() {
	fmt.Println("vvvv")
}

func ff2() int {
	return 50
}

// 函数可以作为参数类型
func ff3(x func() int) {
	fmt.Println("ff3：", x())
}

func ff4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
func ff5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}
