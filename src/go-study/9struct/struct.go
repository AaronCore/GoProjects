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
	fmt.Println("p1：", p1)

	var p2 person
	p2.name = "abc"
	fmt.Println("p2：", p2)

	p3 := person{}
	p3.name = "efg"
	fmt.Println("p3：", p3)

	fmt.Println("----- 2、匿名结构体 ------")
	var user struct {
		name string
		age  int
	}
	user.name = "aaa"
	user.age = 18
	fmt.Println(user)

	fmt.Println("----- 3、指针类型结构体 ------")
	// var p4 = new(person)
	p4 := &person{
		name:  "p2",
		city:  "广州",
		age:   20,
		bobby: []string{"跑步", "游泳"},
	}
	fmt.Printf("%T \n", p4)
	fmt.Println("p4：", p4)
}

// 定义 person 结构体
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
type person struct {
	name  string
	city  string
	age   int
	bobby []string
}

// 面试题
func structSample() {
	m := make(map[string]*student)
	students := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range students {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
