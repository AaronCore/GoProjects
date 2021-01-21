package main

import "fmt"

// 空接口
// 声明：interface{}
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
func main() {

	fmt.Println("1、空接口")

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

	fmt.Println("2、空接口断言")
	// 格式：x.(T)
	// x：表示类型为interface{}的变量
	// T：表示断言x可能是的类型。
	var m interface{}
	m = "aaron"
	v, ok := m.(string)
	if !ok {
		fmt.Println("类型断言失败")
	} else {
		fmt.Println(v)
	}

}
