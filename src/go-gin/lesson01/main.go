package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSON
func getJson(c *gin.Context) {
	// 方式1：gin.H
	//c.JSON(http.StatusOK, gin.H{"message": "hello gin"})

	// 方式2：使用map
	//data := map[string]interface{}{
	//	"name": "abc",
	//	"age":  20,
	//}
	//c.JSON(http.StatusOK, data)

	// 方式3：结构体
	type user struct {
		Name    string
		Age     int
		Address string
	}

	data := user{
		Name:    "aaron",
		Age:     25,
		Address: "深圳",
	}
	c.JSON(http.StatusOK, data)
}

// gin返回json数据
func main() {
	r := gin.Default()

	// 返回Josn数据
	r.GET("/json", getJson)

	r.Run(":650")
}
