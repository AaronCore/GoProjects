package main

import "fmt"

// 空接口
// 声明：interface{}
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
func main() {
	fmt.Println("1、空接口")
	// 定义空接口
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	// 空接口作为map值
	fmt.Println("1.1、空接口使用")
	var userInfo = make(map[string]interface{}, 100)
	userInfo["name"] = "aaron"
	userInfo["age"] = 27
	userInfo["married"] = false
	fmt.Println(userInfo)
}
