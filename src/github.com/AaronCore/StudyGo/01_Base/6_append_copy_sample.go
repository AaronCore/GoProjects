package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切片扩容

	s1 := []string{"深圳", "北京", "上海"}
	s1 = append(s1, "广州", "太原")
	ss := []string{"成都", "重庆"}
	s1 = append(s1, ss...) // ...表示拆开追加
	fmt.Printf("1	s1=%v	len=%d	cap=%d \n", s1, len(s1), cap(s1))

	s2 := make([]string, 5)
	copy(s2, ss)
	fmt.Printf("2	s2=%v	len=%d	cap=%d \n", s2, len(s2), cap(s2))

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 = append(s3[:2], s3[3:]...)
	fmt.Printf("3	s3=%v	len=%d	cap=%d \n", s3, len(s3), cap(s3))

	s4 := []int{8, 9, 7, 3, 4, 2, 1}
	// 顺序
	sort.Ints(s4[:])
	fmt.Printf("4 顺序	s4=%v \n", s4)
	// 逆序
	sort.Sort(sort.Reverse(sort.IntSlice(s4)))
	fmt.Printf("4 逆序	s4=%v \n", s4)

	s5 := []string{"abc", "12", "kk", "Jordan", "Ko", "DD"}
	// 顺序
	sort.Strings(s5)
	fmt.Printf("5 顺序	s5=%v \n", s5)
	// 逆序
	sort.Sort(sort.Reverse(sort.StringSlice(s5)))
	fmt.Printf("5 逆序	s5=%v \n", s5)



}
