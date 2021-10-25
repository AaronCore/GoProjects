package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Id     int    `json:id` //结构体标签：格式声明：`key1:"value1" key2:"value2"`
	Name   string `json:name`
	Gender string `json:gender`
}

type Company struct {
	Title     string      `json:title`
	Employees []*Employee `json:employees`
}

// 结构体与json
func main() {
	c := &Company{
		Title:     "100",
		Employees: make([]*Employee, 0, 50),
	}
	for i := 1; i <= 10; i++ {
		emp := &Employee{
			Id:     i,
			Name:   fmt.Sprintf("%02d", i),
			Gender: "男",
		}
		c.Employees = append(c.Employees, emp)
	}
	fmt.Println("1.JSON序列化：结构体-->JSON格式的字符串")
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	fmt.Println("2.JSON反序列化：JSON格式的字符串-->结构体")
	str := `{"title":"100","employees":[{"id":1,"name":"01","gender":"男"},{"id":2,"name":"02","gender":"男"}]}`
	c1 := &Company{} // 传指针是为了json.Unmarshal内部修改值
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
