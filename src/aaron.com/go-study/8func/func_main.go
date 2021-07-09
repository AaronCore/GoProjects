package main

import "fmt"

/*
	go函数特点：
		1.无需声明原型。
		2.支持不定 变参。
		3.支持多返回值。
		4.支持命名返回参数。
		5.支持匿名函数和闭包。
		6.函数也是一种类型，一个函数可以赋值给变量。
		7.不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
		8.不支持 重载 (overload)
		9.不支持 默认参数 (default parameter)。

	参数传递：传递的都是数据副本。修改数据，对于原始的数据没有影响
		值类型的数据，默认都是值传递：基本数据类型int、float、bool、string、array、struct...
	引用传递：传递的是数据的地址。导致多个变量指向同一个内存地址
		引用类型的数据，默认都是引用传递：指针、slice、map、chan...
*/
var m = 10 // 全局变量
func main() {
	fmt.Println("----- 1.函数测试 -----")
	a := f4(1, 2)
	fmt.Println(a)
	a1, a2 := f5()
	fmt.Println(a1, a2)

	// Go语言中函数传递的都是值（复制、粘贴）
	fmt.Println("----- 2.函数传递测试 -----")
	x := [3]int{1, 2, 3}
	fmt.Println("x：", x)
	f7(x)
	fmt.Println("x：", x)
	y := x
	y[0] = 55
	fmt.Println("y：", y)

	fmt.Println("----- 3.变量作用域 -----")
	f9()

	fmt.Println("----- 4.匿名函数 -----")
	// 匿名函数，函数内部没有办法什么带名称的函数
	var f10 = func(x, y int) {
		fmt.Println(x + y)
	}
	f10(1, 2)
	// 如果只是调用一次的函数，可以简写成立即执行的函数
	func() {
		fmt.Println("立即执行的函数...")
	}()
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
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

// 有参有返回值函数,命名返回值
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

func f7(x [3]int) {
	x[0] = 100
}

// 可变参数函数
func f8(x ...int) {
	fmt.Println(x)
}

// 全局变量函数
func f9() {
	fmt.Println("m：", m)
}

// 递归函数

// 数字阶乘
func factorial(x int) int {
	if x <= 1 {
		return 1
	}
	return x * factorial(x-1)
}

// 斐波那契数列
func fibonacci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
