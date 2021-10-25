package main

import "fmt"

// 只能为当前包内命名类型定义方法。
// 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
// 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
// 不支持方法重载，receiver 只是参数签名的组成部分。
// 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

// Go方法定义格式：
// func (接收者变量 接收者类型) 方法名称(参数列表)(返回值列表){
//		函数体
// }
// 参数和返回值可以省略
func main() {
	// 值类型调用方法
	//u1 := user{"golang", "golang@golang.com"}
	//u1.Notify()
	//// 指针类型调用方法
	//u2 := user{"go", "go@go.com"}
	//u3 := &u2
	//u3.Notify()

	// 值类型调用方法
	u1 := user{"golang", "golang@golang.com"}
	u1.notify()

	// 指针类型调用方法
	// 什么时候应该使用指针类型接收者?
	// 1.需要修改接收者中的值
	// 2.接收者是拷贝代价比较大的大对象
	// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	u2 := user{"go", "go@go.com"}
	u3 := &u2
	u3.notify1()
}

// 结构体
// 标识符：变量名、函数名、类型名、方法名
// Go语言中如果标识符首字母大写，则表示对外部可见（暴露的、公有的）
// 例：user内部可见，User外部可见
type user struct {
	name  string
	email string
}

// 方法
// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名称手写小字母表示

// 值接收者：传拷贝进去
func (u user) notify() {
	fmt.Printf("%v : %v \n", u.name, u.email)
}

// 指针接收者:传内存地址进去
func (u *user) notify1() {
	fmt.Printf("%v : %v \n", u.name, u.email)
}

// 无参数、无返回值
func (u user) method0() {
}

// 单参数、无返回值
func (u user) method1(i int) {
}

// 多参数、无返回值
func (u user) method2(x, y int) {
}

// 无参数、单返回值
func (u user) method3() (i int) {
	return
}

// 多参数、多返回值
func (u user) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (u *user) method5() {
}

// 单参数、无返回值
func (u *user) method6(i int) {
}

// 多参数、无返回值
func (u *user) method7(x, y int) {
}

// 无参数、单返回值
func (u *user) method8() (i int) {
	return
}

// 多参数、多返回值
func (u *user) method9(x, y int) (z int, err error) {
	return
}
