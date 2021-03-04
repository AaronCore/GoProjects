package main

import "fmt"

type dog1 struct{}

func (d *dog1) move() {
	fmt.Println("狗会动")
}

type mover interface {
	move()
}

// 指针接收者接口
// 使用值接收者实现接口和使用指针接收者实现接口的区别？
// 使用值接收者实现接口：结构体类型和结构体指针类型的变量都可以接收
// 使用指针接收者实现接口：只能存结构体指针类型的变量
func main() {
	var x mover
	//var wangcai = dog1{} // 旺财是dog类型
	//x = wangcai          // x可以接收dog类型
	var fugui = &dog1{} // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}
