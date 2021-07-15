package main

import (
	"errors"
	"fmt"
	"math"
	"net"
)

func main() {
	fmt.Println(errors.New("自定义错误"))
	fmt.Println(fmt.Errorf("错误码:%d\n", 100))

	addr, err := net.LookupHost("ccc.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
	fmt.Println("------------自定义错误------------")
	radius := 3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok {
			fmt.Println("当前错误半径：", err.radius)
		}
		return
	}
	fmt.Println("面积：", area)
}

// 1.定义结构体，表示错误类型
type areaError struct {
	msg    string
	radius float64
}

// 2.实现error接口，就是实现Error()方法
func (a *areaError) Error() string {
	return fmt.Sprintf("error：半径：%.2f,%s", a.radius, a.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法", radius}
	}
	return math.Pi * radius * radius, nil
}
