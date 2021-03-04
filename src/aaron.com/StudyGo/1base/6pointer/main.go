package main

import "fmt"

// 指针 pointer，&：取内存地址；*：根据内存地址取值
// new 给基本类型（int、string、bool...）申请内存，返回的是对应类型的指针，很少使用
// make 给slice、map、chan申请内存，函数返回的是这三个类型的本身
func main() {
	fmt.Println("1、指针使用")
	a := "abc"
	a1 := &a  // 获取内存地址
	a2 := *a1 // 根据内存地址获取值
	fmt.Println("内存地址：", a1)
	fmt.Printf("%T \n", a1) // *string：string类型的指针
	fmt.Println("根据内存地址获取值", a2)

	fmt.Println("2、控指针")
	// 声明格式：func new(Type) *Type
	// 说明：
	// 1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
	// 2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	b := new(int)
	b1 := new(bool)
	fmt.Printf("%T\n", b)  // *int
	fmt.Printf("%T\n", b1) // *bool
	fmt.Println(*b)        // 0
	fmt.Println(*b1)
	// 声明格式：func make(t Type, size ...IntegerType) Type
	// 1.二者都是用来做内存分配的。
	// 2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// 3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	var c map[string]int
	c = make(map[string]int, 10)
	c["测试"] = 100
	fmt.Println(c)

	fmt.Println("3、练习")
	var num int
	fmt.Println(&num)
	var ptr *int
	ptr = &num
	*ptr = 55
	fmt.Println(num)
}
