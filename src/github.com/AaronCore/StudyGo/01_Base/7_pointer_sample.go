package main

import (
	"fmt"
)

func main() {
	// 指针，&：取内存地址；*：根据内存地址取值
	// new 给基本类型（int、string、bool...）申请内存，返回的是对应类型的指针，很少使用
	// make 给slice、map、chan申请内存，函数返回的是这三个类型的本身

	str := "abc"
	a1 := &str // 获取内存地址
	a2 := *a1  // 根据内存地址获取值
	fmt.Println("内存地址：", a1)
	fmt.Printf("%T \n", a1) // *string：string类型的指针
	fmt.Println("根据内存地址获取值", a2)

	var p *string
	fmt.Println(p)
	fmt.Printf("%s \n", p)
	if p != nil {
		fmt.Println("not null")
	} else {
		fmt.Println("null")
	}

	//  func new(Type) *Type
	//  1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
	//  2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	var b = new(int)
	*b = 100
	fmt.Println(&b)
	fmt.Println(*b)

	// 小练习
	//程序定义一个int变量num的地址并打印
	//将num的地址赋给指针ptr，并通过ptr去修改num的值
	fmt.Println("----- 小练习 -----")
	var num int
	fmt.Println(&num)
	var ptr *int
	ptr = &num
	*ptr = 55
	fmt.Println(num)
}
