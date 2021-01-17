package main

import "fmt"

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

func f7(x [3]int) {
	x[0] = 100
}

// 可变参数函数
func f8(x ...int) int {
	fmt.Println(x)
	return 10
}
