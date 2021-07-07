package main

import (
	"fmt"
	"math"
)

// 打印九九乘法表
func sample1() {
	fmt.Println("九九乘法表")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}

// 水仙花数：三位数:[100,999]
// 每个位上的数字的立方和，刚好等于该数字本身
// 比如：153 1*1*1 + 5*5*5 3*3*3 = 1+125+27 = 153
func sample2() {
	fmt.Println("水仙花数")
	for i := 100; i < 999; i++ {
		x := i / 100     // 百位
		y := i / 10 % 10 // 十位
		z := i % 10      // 个位
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
}

// 打印 2-100 之间的素质
// 素数：只能被1和自身本身整除
func sample3() {
	fmt.Println("素数打印")
	for i := 2; i < 100; i++ {
		flag := true
		// 判断到根号i就可以，不需要到i的前一个
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}

// 1.三个语句都可以配合标签(label)使用
// 2.标签名区分大小写，定以后若不使用会造成编译错误
// 3.continue、break配合标签(label)可用于多层循环跳出
func main() {
	sample1()
	sample2()
	sample3()
}

// 条件语句if
func f1() {
	var a = 100
	var b = 200
	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}

// 条件语句switch
func f2() {
	var x interface{}
	//写法一：
	switch i := x.(type) { // 带初始化语句
	case nil:
		fmt.Printf(" x 的类型 :%T\r\n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	//写法二
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	//写法三
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	//写法三
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
	//写法四
	var n = 0
	switch { //省略条件表达式，可当 if...else if...else
	case n > 0 && n < 10:
		fmt.Println("i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("i > 10 and i < 20")
	default:
		fmt.Println("def")
	}
}

// 条件语句select
// 每个case都必须是一个通信
// 所有channel表达式都会被求值
// 所有被发送的表达式都会被求值
// 如果任意某个通信可以进行，它就执行；其他被忽略。
// 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
// 否则：
// 如果有default子句，则执行该语句。
// 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
func f3() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := <-c3: // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

// 条件语句for
func f4() {
	var b = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	fmt.Println("使用循环嵌套来输出2到100间的素数：")
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

// 条件语句range
func f5() {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}

	a := [3]int{0, 1, 2}
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}
	fmt.Println(a) // 输出 [100, 101, 102]。

	k := []int{1, 2, 3, 4, 5}
	for i, v := range k { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			k = k[:3]  // 对 slice 的修改，不会影响 range。
			k[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}
}

// goto使用
// goto是调整执行位置，与continue、break配合标签(label)的结果并不相同
func f6() {
	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				goto over
			}
			fmt.Printf("%v-%c \n", i, j)
		}
	}
over:
	fmt.Println("结束了")
}
