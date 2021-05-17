package main

import "fmt"

// 数组
func main() {
	fmt.Println("1、数组声明方式")
	//方式1
	a1 := [3]int{1, 2, 3}
	//方式2，[...]根据初始值自动推断数组的长度
	a2 := [...]int{1, 2, 3, 4, 5}
	//方式3,根据索引初始化
	a3 := [5]int{0: 1, 2: 2}
	fmt.Println("a1：", a1)
	fmt.Println("a2：", a2)
	fmt.Println("a3：", a3)

	fmt.Println("2、数组遍历")
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}
	for _, v := range a2 {
		fmt.Printf("value：%v \n", v)
	}

	fmt.Println("3、多维数组声明")
	var a5 [3][2]int
	a5 = [3][2]int{
		{10, 11},
		{20, 21},
		[2]int{25, 26},
	}
	a6 := [3][2]int{{1, 2}, {2, 3}}
	fmt.Println(a5, a6)

	fmt.Println("4、练习一")
	a7 := [...]int{1, 3, 5, 7, 8}
	var num int
	for _, v := range a7 {
		num += v
	}
	fmt.Printf("总和：%d \n", num)

	fmt.Println("4、练习二")
	// 找出两个元素之和等于8的下标
	for i := 0; i < len(a7); i++ {
		for j := i + 1; j < len(a7); j++ {
			if a7[i]+a7[j] == 8 {
				fmt.Printf("(%d %d) \n", i, j)
			}
		}
	}
}
