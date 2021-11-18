package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	UserName string `form:"username" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//参数绑定
func main() {
	r := gin.Default()

	// 绑定JSON的示例 ({"username": "Aton1", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"UserName": login.UserName,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (username=Bas&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"UserName": login.UserName,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?username=abc&password=123456)
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			fmt.Println(login)
			c.JSON(http.StatusOK, gin.H{
				"UserName": login.UserName,
				"Password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.Run(":650")
}
