package main

import (
	"fmt"
)

func main() {
	// 流程判断

	age := 18
	if age > 18 {
		fmt.Println("成年")
	} else if age >= 35 {
		fmt.Println("中年人")
	} else {
		fmt.Println("未成年")
	}

	n := 2
	switch n {
	case 1:
		fmt.Println(1)
		fallthrough // 执行满足条件的下一个case
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(2)
	}

	switch a := 3; a {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}
}
