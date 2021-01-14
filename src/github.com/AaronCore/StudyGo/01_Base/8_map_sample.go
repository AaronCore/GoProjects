package main

import "fmt"

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

	// map[KeyType]ValueType
	// KeyType:表示键的类型。
	// ValueType:表示键对应的值的类型。

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	// make(map[KeyType]ValueType, [cap])

	m1 := make(map[string]int, 5) // map必须初始化才能使用

	m1["A"] = 10
	m1["B"] = 20

	fmt.Println("m1", m1)
	fmt.Println("A：", m1["A"])

	// map声明并赋值
	m2 := map[string]string{
		"A": "AA",
		"B": "BB",
	}
	fmt.Println("m2", m2)

	//
}
