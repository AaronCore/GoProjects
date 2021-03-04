package main

import (
	"fmt"
	"os"
)

type study struct {
	id   int64
	name string
}

type studyMgr struct {
	allStudy map[int64]study
}

// 查看所有
func (s studyMgr) showStudys() {
	for _, v := range s.allStudy {
		fmt.Printf("学号：%d 学生：%v \n", v.id, v.name)
	}
}

// 添加
func (s studyMgr) addStudy() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	stu := study{
		id:   id,
		name: name,
	}
	s.allStudy[stu.id] = stu
	fmt.Println("添加成功")
}

// 编辑
func (s studyMgr) editStudy() {
	var id int64
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)

	value, ok := s.allStudy[id]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("查询结果：学号：%d 学生：%v \n", value.id, value.name)
	fmt.Println("请输入需要修改的名称：")
	var newName string
	fmt.Scanln(&newName)
	value.name = newName
	s.allStudy[id] = value
	fmt.Println("修改成功")
}

// 删除
func (s studyMgr) delStudy() {
	var id int64
	fmt.Println("请输入删除的学生学号：")
	fmt.Scanln(&id)

	_, ok := s.allStudy[id]
	if !ok {
		fmt.Println("查无此人")
		return
	}

	delete(s.allStudy, id)
	fmt.Println("删除成功")
}

// 声明全局变量
var (
	smr studyMgr
)

// 结构体练习
func main() {
	smr = studyMgr{
		allStudy: make(map[int64]study, 30),
	}
	for {
		showMenu()
		fmt.Println("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)
		switch choice {
		case 1:
			smr.showStudys()
		case 2:
			smr.addStudy()
		case 3:
			smr.editStudy()
		case 4:
			smr.delStudy()
		case 5:
			os.Exit(1)
		default:
			fmt.Printf("输入错误...")
		}
	}
}
func showMenu() {
	fmt.Println("----- webcome... -----")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}
