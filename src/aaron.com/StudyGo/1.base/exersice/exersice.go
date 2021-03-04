package main

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	f1()
	f2()
	f3()
	f4()
}

// 练习1：判断字符串中汉字符串
func f1() {
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
}

// 练习2：字符串切割
func f2() {
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

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

func f3() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
func dispatchCoin() int {
	for _, name := range users {
		num := 0
		for _, value := range name {
			v := strings.ToLower(string(value))
			switch v {
			case "e":
				num += 1
			case "i":
				num += 2
			case "o":
				num += 3
			case "u":
				num += 4
			}
		}
		distribution[name] = num
		coins -= num
	}
	fmt.Println(distribution)
	return coins
}

// 冒泡排序
func f4() {
	fmt.Println("冒泡排序")
	a := [...]int{5, 9, 6, 8, 12, 42}
	fmt.Println(a)
	num1 := len(a)
	for i := 0; i < num1; i++ {
		for j := i + 1; j < num1; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
