package main

import (
	"fmt"
	"strings"
)

func main() {
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

	path := "D:\\Program Files (x86)\\Haozip"
	fmt.Println(path)

	// 字符串拼接
	fmt.Println("----- 1、字符串拼接 ------")
	name := "aaron"
	age := 27
	address := "深圳市 福田区"
	str := fmt.Sprintf("%v %v %v", name, age, address)
	fmt.Println(str)

	// 字符串切割、遍历
	fmt.Println("----- 2、字符串切割、遍历 ------")
	str1 := "a,b,c"
	ret := strings.Split(str1, ",")
	fmt.Println(ret)

	for i := 0; i < len(ret); i++ {
		fmt.Printf("方法一	index:%v,value:%v \n", i, ret[i])
	}
	for i, v := range ret {
		fmt.Printf("方法二	index:%v,value:%v \n", i, v)
	}

	// 数组拼接
	fmt.Println("----- 3、数组拼接 ------")
	fmt.Println(strings.Join(ret, "+"))

	// 字符串判断
	fmt.Println("----- 4、字符串判断 ------")
	fmt.Println(strings.Contains(name, "a"))

	// 字符串前后缀判断
	fmt.Println("----- 5、字符串前后缀判断 ------")
	fmt.Println(strings.HasPrefix(name, "a"))
	fmt.Println(strings.HasSuffix(name, "n"))

	// 获取字符串下标
	fmt.Println("----- 6、获取字符串下标 ------")
	fmt.Println(strings.Index(name, "o"))

	// 字符串修改
	fmt.Println("----- 7、字符串修改 ------")

	str2 := "白花菜"
	str3 := []rune(str2) // 将字符串强制转换rune切片
	str3[0] = '红'
	fmt.Println(str3)
	fmt.Println(string(str3)) // 将rune切片转成字符

	str4 := "abc"
	str5 := []byte(str4)
	str5[0] = 'A'
	fmt.Println(str5)
	fmt.Println(string(str5))

	str6 := "abc深圳市"
	for _, v := range str6 {
		if v > 'z' {
			fmt.Println(string(v))
		}
	}

	// 类型转换
	fmt.Println("----- 8、字符串修改 ------")
	a := 100
	a1 := float64(a)
	fmt.Println(a1)
}
