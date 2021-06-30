package main

import (
	"fmt"
	"sort"
)

// 切片 Slice
// 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
// 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。
// 切片扩容策略：
// 首先判断，如果新申请容量(cap)大于2倍的旧容量(old.cap)，最终容量(newcap)就是新申请的容量(cap)。
// 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即(newcap=doublecap)
// 否则判断，如果旧切片长度大于等于1024，则最终容量(newcap)从旧容量(old.cap)开始循环增加原来的1/4，
//   即(newcap=old.cap,for {newcap += newcap/4})直到最终容量(newcap)大于等于新申请的容量(cap)，即(newcap >= cap)
// 如果最终容量(cap)计算值溢出，则最终容量(cap)就是新申请容量(cap)。
func main() {
	fmt.Println("1、切片声明")
	var s1 []int
	var s2 []string
	fmt.Println(s1 == nil, s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"福田区", "南山区", "宝安区", "龙岗区"}
	fmt.Printf("s1 len：%d，cap：%d \n", len(s1), cap(s1))
	fmt.Printf("s2 len：%d，cap：%d \n", len(s2), cap(s2))

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

	fmt.Println("2、切片扩容")
	ss := []string{"成都", "重庆"}
	s2 = append(s2, ss...) // ...表示拆开追加
	fmt.Printf("1	s2=%v	len=%d	cap=%d \n", s1, len(s2), cap(s2))

	a1 := make([]string, 5)
	copy(a1, ss)
	fmt.Printf("2	a1=%v	len=%d	cap=%d \n", a1, len(a1), cap(a1))

	b1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b1 = append(b1[:2], b1[3:]...)
	fmt.Printf("3	b1=%v	len=%d	cap=%d \n", b1, len(b1), cap(b1))

	b2 := []int{8, 9, 7, 3, 4, 2, 1}
	// 顺序
	sort.Ints(b2[:])
	fmt.Printf("4 顺序	b2=%v \n", b2)
	// 逆序
	sort.Sort(sort.Reverse(sort.IntSlice(b2)))
	fmt.Printf("4 逆序	b2=%v \n", b2)

	b5 := []string{"abc", "12", "kk", "Jordan", "Ko", "DD"}
	// 顺序
	sort.Strings(b5)
	fmt.Printf("5 顺序	b5=%v \n", b5)
	// 逆序
	sort.Sort(sort.Reverse(sort.StringSlice(b5)))
	fmt.Printf("5 逆序	s5=%v \n", b5)
}
