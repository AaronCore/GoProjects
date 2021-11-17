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

// queryString
func queryString(c *gin.Context) {
	//query := c.Query("query")
	query := c.DefaultQuery("query", "")
	//query, ok := c.GetQuery("query")
	//if !ok {
	//	fmt.Println("get query failed.")
	//}
	c.JSON(http.StatusOK, gin.H{"query": query})
}

// loginView
func loginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func getForm(c *gin.Context) {
	//获取form表单数据
	// 方式1
	userName := c.PostForm("userName")
	passWord := c.PostForm("passWord")
	// 方式2
	//userName := c.DefaultPostForm("userName", "")
	//passWord := c.DefaultPostForm("passWord", "")
	// 方式3
	//userName, _ := c.GetPostForm("userName")
	//passWord, _ := c.GetPostForm("passWord")

	c.HTML(http.StatusOK, "index.html", gin.H{"userName": userName, "passWord": passWord})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	// 返回Josn数据
	r.GET("/json", getJson)
	// 获取参数
	r.GET("/web", queryString)
	// 页面跳转、获取表单数据
	r.GET("/login", loginView)
	r.POST("/login", getForm)

	r.Run(":650")
}
