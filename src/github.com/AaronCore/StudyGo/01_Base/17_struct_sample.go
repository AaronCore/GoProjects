package main

import "fmt"

// 	  结构体
//    type 类型名 struct {
//        字段名 字段类型
//        字段名 字段类型
//        …
//    }
//    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
//    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
//    3.字段类型：表示结构体字段的具体类型。
func main() {
	fmt.Println("----- 1、结构体 ------")
	p1 := person{
		name:  "aaron",
		city:  "深圳",
		age:   27,
		bobby: []string{"听歌", "打球"},
	}
	fmt.Printf("%T\n", p1)
	fmt.Println(p1)

	fmt.Println("----- 2、匿名结构体 ------")
	var user struct {
		name string
		age  int
	}
	user.name = "aaa"
	user.age = 18
	fmt.Println(user)

	fmt.Println("----- 3、指针类型结构体 ------")
	//  什么时候应该使用指针类型接收者?
	//    1.需要修改接收者中的值
	//    2.接收者是拷贝代价比较大的大对象
	//    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	// var p2 = new(person)
	p2 := &person{
		name:  "p2",
		city:  "广州",
		age:   20,
		bobby: []string{"跑步", "游泳"},
	}
	fmt.Printf("%T \n", p2)
	fmt.Println(p2)

	fmt.Println("----- 4、构造函数 ------")
	b1 := newStudent("aaa", 20)
	fmt.Println(b1)

	fmt.Println("----- 5、方法和接收者 ------")

	//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	//        函数体
	//}

	//    1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
	//    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
	//    3.方法名、参数列表、返回参数：具体格式与函数定义相同。

	// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	c1 := newStudent("张三", 23)
	c1.Dream()

}

type person struct {
	name  string
	city  string
	age   int
	bobby []string
}

type student struct {
	name string
	age  int
}

func (p student) Dream() {
	fmt.Println("dream", p.name)
}

// 构造函数，约定以new开头
func newPerson(name string, city string, age int, bobby []string) *person {
	return &person{
		name:  name,
		city:  "深圳",
		age:   age,
		bobby: bobby,
	}
}
func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func study() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
