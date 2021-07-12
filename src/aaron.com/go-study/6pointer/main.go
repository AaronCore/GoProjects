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

	fmt.Println("3、数组指针")
	// 数组指针：var arr1 *[3]int
	arr1 := [4]int{1, 2, 3, 4}
	var p1 *[4]int
	p1 = &arr1
	fmt.Println("p1：", p1)
	p1[0] = 11
	fmt.Println("arr1：", arr1)

	fmt.Println("4、指针数组")
	// 指针数组：var arr2 [3]*int
	num1 := 1
	num2 := 2
	num3 := 3
	num4 := 4
	p2 := [4]int{num1, num2, num3, num4}
	fmt.Println("p2：", p2)
	p3 := [4]*int{&num1, &num2, &num3, &num4}
	fmt.Println("p3：", p3)
	p2[0] = 10
	fmt.Println("p2：", p2)
	*p3[0] = 100
	fmt.Println("p3：", p3)
	fmt.Println("num1：", num1)

	fmt.Println("5、指针的指针")
	m := 99
	var ptr *int
	var pptr **int
	ptr = &m    // 指针ptr地址
	pptr = &ptr // 指向ptr地址
	fmt.Printf("变量m=%d\n", m)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	fmt.Println("6、指针函数")
	f1 := f1()
	fmt.Printf("f1的类型：%T，地址：%p，数值：%v\n", f1, &f1, f1)
	f2 := f2()
	fmt.Printf("f2的类型：%T，地址：%p，数值：%v\n", f2, &f2, f2)
	fmt.Printf("f2指针中存储的数组地址：%p\n", f2)

	fmt.Println("7、练习")
	var mm int
	fmt.Println(&mm)
	var mmPtr *int
	mmPtr = &mm
	*mmPtr = 55
	fmt.Println(mm)
}

func f1() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func f2() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("f2()中arr的地址：%p\n", &arr)
	return &arr
}
