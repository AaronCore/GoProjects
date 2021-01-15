package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 练习1：判断字符串中汉字符串

	fmt.Println("1、判断字符串中的汉字")
	s := "s广东G省tZ深Y圳z市福田m区abC"
	var count int
	str := make([]string, 20)
	for _, value := range s {
		// 判断方式1
		//if value > 'z' {
		//	count++
		//	str = append(str, string(value))
		//	fmt.Println(string(value))
		//}
		// 判断方式2
		if unicode.Is(unicode.Han, value) {
			count++
			str = append(str, string(value))
			//fmt.Println(string(value))
		}
	}
	fmt.Println("拼接字符串中的汉字：", strings.Join(str, ""))
	fmt.Println("字符串中汉字的个数：", count)

	// 练习2：字符串切割
	fmt.Println("2、统计字母出现的次数")
	a := "a.b.c.d.a.r.g.d.b.h.r.a"
	a1 := strings.Split(a, ".")
	mapSlice := make(map[string]int, 20)
	for _, value := range a1 {
		_, ok := mapSlice[value]
		if !ok {
			mapSlice[value] = 1
		} else {
			mapSlice[value] += 1
		}
	}
	fmt.Println(mapSlice)

	// 练习2：回文判断，字符串从左往右读跟从右往左读都是一样的
	fmt.Println("2、回文判断，字符串从左往右读跟从右往左读都是一样的")
	ss := "A上海自来水来自海上A"
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	fmt.Println("[]rune：", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
