package main

import (
	"fmt"
)

// 匿名结构体、嵌套结构体
func main() {
	fmt.Println("1、匿名结构体")
	p1 := card{
		"aaron",
		20,
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)

	fmt.Println("2、嵌套结构体")
	p2 := product{
		id:          10,
		productName: "产品1",
		productType: productType{
			id:              100,
			productTypeName: "产品类型1",
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.productName, p2.productType.productTypeName)
}

// card
// 匿名结构体
type card struct {
	string // 匿名字段
	int
}

// product
//嵌套结构体
type product struct {
	id          int64
	productName string
	productType productType
}

// productType
type productType struct {
	id              int64
	productTypeName string
}
