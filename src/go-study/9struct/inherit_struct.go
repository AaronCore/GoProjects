package main

import "fmt"

// animal 动物
// 结构体字段的可见性：结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
type animal struct {
	name string
}

type dog struct {
	feet    int8
	*animal //通过嵌套匿名结构体实现继承
}

func (a *animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (d *dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// 结构体继承
func main() {
	d1 := &dog{
		feet: 4,
		animal: &animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
