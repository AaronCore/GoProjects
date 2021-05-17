package main

import "fmt"

// 流程控制
func main() {
	fmt.Println("1、if else")
	a := 20
	if a > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	fmt.Println("2、switch")
	var grade string = "B"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
		fallthrough // 强制执行后面的case代码
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}

	fmt.Println("2、select")
	//    每个case都必须是一个通信
	//    所有channel表达式都会被求值
	//    所有被发送的表达式都会被求值
	//    如果任意某个通信可以进行，它就执行；其他被忽略。
	//    如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	//    否则：
	//    如果有default子句，则执行该语句。
	//    如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
