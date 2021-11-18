package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// gin获取form表单数据
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	// 页面跳转、获取表单数据
	r.GET("/login", loginView)
	r.POST("/login", getForm)

	r.Run(":650")
}
