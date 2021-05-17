package main

import "fmt"

// 闭包,闭包=函数+引用环境
// 底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，现在内部查找，找不到往外层找
func main() {
	fmt.Println("1	-----  -----")
	//fun1(fun3(100, 200))
	//fun4(fun2, 10, 20)()
	ret := fun4(fun2, 1, 2)
	fun1(ret)

	fmt.Println("2	----- 闭包 -----")
	a2 := adder()(100)
	fmt.Println("a2：", a2)

	fmt.Println("3	----- 示例 -----")
	a3, a4 := cale(100)
	fmt.Println(a3(1), a4(2))
	fmt.Println(a3(3), a4(4))
}

func fun1(f func()) {
	fmt.Println("this is fu1")
	f()
}

func fun2(x, y int) {
	fmt.Println("this is fu2")
	fmt.Println(x + y)
}

func fun3(x, y int) func() {
	return func() {
		fmt.Println(x + y)
	}
}

func fun4(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func cale(x int) (func(int) int, func(int) int) {
	add := func(y int) int {
		x += y
		return x
	}
	sub := func(y int) int {
		x -= y
		return x
	}
	return add, sub
}
