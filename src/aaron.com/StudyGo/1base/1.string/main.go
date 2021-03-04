package main

import (
	"fmt"
	"strings"
)

// %T：查看类型
// %d：十进制数
// %b：二进制数
// %o：八进制数
// %x：六进制数
// %c：字符
// %s：字符串
// %p：指针
// %v：值
// %f：浮点数
// %t：布尔值
// %q：用双引号包括字符串值
func main() {
	fmt.Println("1、获取字符串长度")
	fmt.Println("当前字符串长度是：", len("abcd"))

	fmt.Println("2、字符串拼接")
	fmt.Println("aa" + "bb")
	fmt.Println(fmt.Sprintf("%v%v", "cc", "dd"))

	fmt.Println("3、字符串切割、遍历")
	a := "a.b.c.d.e.f"
	s := strings.Split(a, ".")
	fmt.Println("切割后结果：", s)
	fmt.Println("3.1、遍历")
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("方法一：index:%v,value:%v \n", i, s[i])
	//}
	for i, v := range s {
		fmt.Printf("方法二：index:%v,value:%v \n", i, v)
	}

	fmt.Println("4、切片拼接成字符串")
	fmt.Println(strings.Join(s, ","))

	fmt.Println("5、字符串判断")
	fmt.Println(strings.Contains(a, "a"))

	fmt.Println("6、字符串前缀、后缀判断")
	fmt.Println(strings.HasPrefix(a, "a"))
	fmt.Println(strings.HasSuffix(a, "n"))

	fmt.Println("7、获取字符在字符串中的下标")
	fmt.Println(strings.Index(a, "b"))

	fmt.Println("8、修改字符串")
	// Go语言的字符有以下两种：
	// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。例：var b := 'x'
	// rune类型，代表一个 UTF-8字符。例：var b := '中'
	s1 := "hello"
	byteS1 := []byte(s1) // 强制类型转换
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "红客"
	runeS2 := []rune(s2)
	runeS2[0] = '黑'
	fmt.Println(string(runeS2))

	fmt.Println("9、练习题：打印字符串中的汉字")
	s3 := "abc深圳市"
	for _, v := range s3 {
		if v > 'z' {
			fmt.Println(string(v))
		}
	}
}
