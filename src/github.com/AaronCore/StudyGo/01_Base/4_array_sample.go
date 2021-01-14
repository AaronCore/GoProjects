package main

import "fmt"

func main() {
	// 数组

	// 1.数组初始化方式1

	// 数组长度是数组类型的一部分
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	var a1 [2]bool //声明
	a1 = [2]bool{true, false}
	fmt.Printf("%T \n", a1)

	// 2.数组初始化方式2
	// [...]根据初始值自动推断数组的长度
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a2)

	// 3.数组初始化方式3
	a3 := [5]int{1, 2}
	a4 := [5]int{0: 1, 2: 2} //根据索引初始化
	fmt.Println(a3, a4)

	// 数组遍历
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}
	for _, v := range a2 {
		//fmt.Printf("index:%d,value:%v \n", i, v)
		fmt.Printf("value：%v", v)
	}

	// 4.数组初始化方式4
	var a5 [3][2]int
	a5 = [3][2]int{
		{10, 11},
		{20, 21},
		[2]int{25, 26},
	}
	a6 := [3][2]int{{1, 2}, {2, 3}}
	fmt.Println(a5, a6)

	for _, i1 := range a5 {
		for _, i2 := range i1 {
			fmt.Println(i2)
		}
	}

	// 练习
	a7 := [...]int{1, 3, 5, 7, 8}
	var num int
	for _, v := range a7 {
		num += v
	}
	fmt.Printf("总和：%d \n", num)

	for i := 0; i < len(a7); i++ {
		for j := i + 1; j < len(a7); j++ {
			if a7[i]+a7[j] == 8 {
				fmt.Printf("(%d %d) \n", i, j)
			}
		}
	}
}
