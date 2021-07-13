package main

import "fmt"

//    接口是一个或多个方法签名的集合。
//    任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
//    就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
//    这称为Structural Typing。
//    所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
//    当然，该类型还可以有其他方法。
//    接口只有方法声明，没有实现，没有数据字段。
//    接口可以匿名嵌入其他接口，或嵌入到结构中。
//    对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
//    只有当接口存储的类型和对象都为nil时，接口才等于nil。
//    接口调用不会做receiver的自动转换。
//    接口同样支持匿名字段方法。
//    接口也可实现类似OOP中的多态。
//    空接口可以作为任何类型数据的容器。
//    一个类型可实现多个接口。
//    接口命名习惯以 er 结尾。

// 	  声明方式：
//    type 接口类型名 interface{
//        方法名1( 参数列表1 ) 返回值列表1
//        方法名2( 参数列表2 ) 返回值列表2
//        …
//    }
//    其中
//    1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
// 	  2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
//    3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

// 使用值接收者实现接口和使用指针接收者实现接口的区别？
// 使用值接收者实现接口：结构体类型和结构体指针类型的变量都可以接收
// 使用指针接收者实现接口：只能存结构体指针类型的变量
func main() {
	fmt.Println("1、接口使用")
	var x sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪

	fmt.Println("2、空接口断言")
	// 格式：x.(T)
	// x：表示类型为interface{}的变量
	// T：表示断言x可能是的类型。
	var m interface{}
	m = "aaron"
	v, ok := m.(string)
	if !ok {
		fmt.Println("类型断言失败")
	} else {
		fmt.Println(v)
	}
}

// sayer 接口
type sayer interface {
	say()
}

// 定义结构体
type (
	cat struct{}
	dog struct{}
)

// 接口实现
func (c cat) say() {
	fmt.Println("喵喵喵...")
}
func (d dog) say() {
	fmt.Println("汪汪汪...")
}
func speak(s sayer) {
	s.say()
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupported type!")
	}
}
