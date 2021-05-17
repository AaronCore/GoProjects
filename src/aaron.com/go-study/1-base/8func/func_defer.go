package main

import (
	"fmt"
)

func main() {
	// 函数不能重载，不能嵌套，不支持默认参数
	// defer延时调用，return前调用，多个defer语句，按照先进后出的方式执行
	// defer用途：关闭文件句柄，锁资源释放，数据库连接资源释放

	fmt.Println("1	----- func -----")

	fu1("abc")

	fmt.Println("2	----- defer -----")

	// Go语言中函数的return不是原子操作，在底层是分为两步执行的
	// 第一步：返回值赋值
	// defer
	// 第二步：真正的RET返回
	// 函数中如果存在defer,那么defer的执行时机是在第一步和第二步之间

	//fmt.Println("start...")
	//defer fmt.Println("111")
	//defer fmt.Println("222")
	//defer fu1("aaa")
	//defer fu1("bbb")
	//fmt.Println("end...")

	fmt.Println("de1：", de1())
	fmt.Println("de2：", de2())
	fmt.Println("de3：", de3())
	fmt.Println("de4：", de4())
	fmt.Println("de5：", de5())

	fmt.Println("3	----- defer应用 -----")
	a := 1
	b := 2
	defer de6("1", a, de6("10", a, b))
	a = 0
	defer de6("2", a, de6("20", a, b))
	b = 1
}

func fu1(name string) {
	fmt.Println("你好：", name)
}

func de1() int {
	x := 5
	defer func() {
		x++ // 此处修改的是x值，不是返回的值
	}()
	return x
}

func de2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=5
}

func de3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x值
	}()
	return x // 返回值=y=x=5
}

func de4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中副本值
	}(x)
	return 5 // 返回值=x=5
}

func de5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}

func de6(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
