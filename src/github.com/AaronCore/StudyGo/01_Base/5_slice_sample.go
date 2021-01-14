package main

import "fmt"

func main() {
	// 切片 slice,引用类型
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
	// 切片唯一合法的比较操作是和nil比较。
	// 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例
	/*
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	*/

	var s1 []int
	var s2 []string

	fmt.Println(s1 == nil, s2 == nil)

	//if s1 == nil {
	//	fmt.Println("空")
	//} else {
	//	fmt.Println("非空")
	//}

	s1 = []int{1, 2, 3}
	s2 = []string{"福田区", "南山区", "宝安区", "龙岗区"}

	fmt.Println(s1 == nil, s2 == nil, s1, s2)
	fmt.Printf("s1 len：%d，cap：%d \n", len(s1), cap(s1))
	fmt.Printf("s2 len：%d，cap：%d \n", len(s2), cap(s2))

	// 由数组得到切片
	s3 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19} // len 10,cap 10
	s4 := s3[0:7]                                     // => [0,4]
	s5 := s3[:3]                                      // => [0,3]
	s6 := s3[2:]                                      // => [2,len(s3)-1]
	s7 := s3[:]                                       // => [0,len(s3)-1]
	fmt.Printf("s4：%v \n", s4)
	fmt.Printf("s5：%v \n", s5)
	fmt.Printf("s6：%v \n", s6)
	fmt.Printf("s7：%v \n", s7)

	// 切片容量是底层数组从切片的第一个元素到最后的一个元素的数量
	// 切片的容量是指底层数组的容量
	fmt.Printf("s4 len：%d，cap：%d \n", len(s4), cap(s4))
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("s6 len：%d，cap：%d \n", len(s6), cap(s6))

	// 切片再切片
	s8 := s6[3:]
	fmt.Printf("s8 len：%d，cap：%d \n", len(s8), cap(s8))

	// 切片是引用类型，都指向了底层的一个数组
	fmt.Printf("s6：%v \n", s6)
	s3[9] = 190 // 修改了底层数组的值
	fmt.Println("s6：", s6)
	fmt.Println("s8：", s8)

	// 使用make函数创建切片

	// var slice []type = make([]type, len,cap) 不写cap的话，cap的默认等于len
	s10 := make([]int, 5, 10)
	fmt.Printf("s10=%v	len=%d	cap=%d \n", s1, len(s10), cap(s10))
}
