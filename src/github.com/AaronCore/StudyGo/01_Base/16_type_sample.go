package main

import "fmt"

type aInt int   // 自定义类型
type bInt = int // 类型别名,只会在代码中存在，编译完成时并不会存在
func main() {
	// 类型别名和自定义类型

	type1()
	type2()
}

func type1() {
	var x aInt
	x = 100
	fmt.Printf("type1：%T %d \n", x, x)
}

func type2() {
	var x bInt
	x = 100
	fmt.Printf("type2：%T %d \n", x, x)
}
