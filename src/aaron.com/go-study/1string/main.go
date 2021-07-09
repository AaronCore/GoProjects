package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
	%T：查看类型
	%d：十进制数
	%b：二进制数
	%o：八进制数
	%x：六进制数
	%c：字符
	%s：字符串
	%p：指针
	%v：值
	%f：浮点数
	%t：布尔值
	%q：用双引号包括字符串值
*/
func main() {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("world")
	fmt.Println(buffer.String())

	var builder strings.Builder
	builder.WriteString("ABC")
	builder.WriteString("DEF")
	fmt.Println(builder.String())
}

// 求长度
func f1() {
	fmt.Println("len：", len("abcd"))
}

// 拼接字符串
func f2() {
	fmt.Println(fmt.Sprintf("%v%v", "中国", "广东"))
}

// 判断是否包含
func f3() {
	fmt.Println(strings.Contains("bca", "a"))
}

// 字符串前缀、后缀判断
func f4() {
	var str = "ahn"
	fmt.Println(strings.HasPrefix(str, "a"))
	fmt.Println(strings.HasSuffix(str, "n"))
}

// 获取字符在字符串中的下标
func f5() {
	fmt.Println(strings.Index("a.b.c.d.e.f", "b"))
}

// 切片拼接成字符串
func f6() {
	var str = "a.b.c.d.e.f"
	s := strings.Split(str, ".")
	fmt.Println(strings.Join(s, ","))
}

// 遍历
func f7() {
	var str = "a.b.c.d.e.f"
	s := strings.Split(str, ".")
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("index:%v,value:%v \n", i, s[i])
	//}
	for i, v := range s {
		fmt.Printf("index:%v,value:%v \n", i, v)
	}
}

// 字符串修改
func f8() {
	/*
		Go语言的字符有以下两种：
			uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。例：var b := 'x'
			rune类型，代表一个 UTF-8字符。例：var b := '中'
	*/
	s1 := "hello"
	byteS1 := []byte(s1) // 强制类型转换
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "红客"
	runeS2 := []rune(s2)
	runeS2[0] = '黑'
	fmt.Println(string(runeS2))
}

// 练习题：打印字符串中的汉字
func f9() {
	s3 := "abc深圳市"
	for _, v := range s3 {
		if v > 'z' {
			fmt.Println(string(v))
		}
	}
}
