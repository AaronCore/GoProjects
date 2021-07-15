package main

import "fmt"

func main() {
	m1 := Mouse{
		name: "双飞燕",
	}
	fmt.Println(m1.name)
	f1 := FlashDisk{
		name: "金士顿",
	}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)

	var usb Usb
	usb = m1
	usb.start()
	usb.end()

	m1.edit()
	f1.delete()

	var peo People = &Student{}
	think := "hi"
	fmt.Println(peo.speak(think))
}

// Usb 接口
type Usb interface {
	start()
	end()
}

// Keyboard 嵌套接口
type Keyboard interface {
	Usb
	typing()
}

// People 接口
type People interface {
	speak(string) string
}

// Mouse 结构体
type Mouse struct {
	name string
}

// FlashDisk 结构体
type FlashDisk struct {
	name string
}

// Student 结构体
type Student struct{}

// Mouse 实现接口
func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作...")
}
func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，安全退出。")
}
func (m Mouse) edit() {
	fmt.Println(m.name, "编辑文件。")
}

// FlashDisk 实现接口
func (f FlashDisk) start() {
	fmt.Println(f.name, "准备就绪，可以开始存储数据...")
}
func (f FlashDisk) end() {
	fmt.Println(f.name, "结束工作，安全弹出。")
}
func (f FlashDisk) delete() {
	fmt.Println(f.name, "删除数据。")
}

// 测试方法
func testInterface(usb Usb) {
	usb.start()
	usb.end()
}

func (s *Student) speak(think string) (talk string) {
	if think == "hi" {
		talk = "Hello"
	} else {
		talk = "go go"
	}
	return
}
