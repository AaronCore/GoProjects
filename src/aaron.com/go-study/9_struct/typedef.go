package main

import "fmt"

type aInt int   // 自定义类型
type bInt = int // 类型别名,只会在代码中存在，编译完成时并不会存在

// 类型别名和自定义类型
func main() {

	f1()
	f2()
	a := aInt(10)
	a.hello()
}

func f1() {
	var x aInt
	x = 100
	fmt.Printf("type1：%T %d \n", x, x)
}

// 给自定义类型加方法
// 只能给当前包下的自定义类型加方法
func (a aInt) hello() {
	fmt.Println("Hello, 我是一个int...")
}

func f2() {
	var x bInt
	x = 100
	fmt.Printf("type2：%T %d \n", x, x)
}
