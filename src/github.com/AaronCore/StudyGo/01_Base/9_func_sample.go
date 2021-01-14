package main

import "fmt"

func main() {
	a := f4(1, 2)
	fmt.Println(a)
	a1, a2 := f5()
	fmt.Println(a1, a2)
}

// 无参无返回值函数
func f1() {
	fmt.Println("abc")
}

// 无参有返回值函数
func f2() int {
	return 2
}

// 有参无返回值函数
func f3(x int, y int) {
	m := x + y
	fmt.Println(m)
}

// 有参有返回值函数
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多返回值函数
func f5() (int, string) {
	return 1, "A"
}

// 同类型参数简写函数
func f6(a, b string) string {
	return a + b
}
