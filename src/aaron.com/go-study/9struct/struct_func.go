package main

import "fmt"

// 	  方法定义格式：
//    func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//        函数体
//    }
//    说明：
//    1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。
//      例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
//    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
//    3.方法名、参数列表、返回参数：具体格式与函数定义相同。
// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func main() {
	fmt.Println("----- 1、构造函数 ------")
	b1 := newStudent("aaa", 20)
	fmt.Println(b1)

	fmt.Println("----- 2、方法和接收者 ------")
	c1 := newStudent("张三", 23)
	c1.dream()

	fmt.Println("----- 3、指针类型的接收者 ------")
	p1 := newStudent("测试", 25)
	fmt.Println(p1.age) // 25
	p1.setAge(30)
	fmt.Println(p1.age) // 30

	fmt.Println("----- 4、值类型的接收者 ------")
	p2 := newStudent("测试", 25)
	p2.dream()
	fmt.Println(p2.age) // 25
	p1.setAge1(30)      // (*p1).SetAge2(30)
	fmt.Println(p2.age) // 25
}

type student struct {
	name string
	age  int8
}

// 构造函数：约定以new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针，节省内存开销
func newStudent(name string, age int8) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func (s student) dream() {
	fmt.Printf("%s的梦想是学好Go语言...\n", s.name)
}

// 使用指针接收者
func (s *student) setAge(newAge int8) {
	s.age = newAge
}

// 使用值接收者
func (s student) setAge1(newAge int8) {
	s.age = newAge
}
