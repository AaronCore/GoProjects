package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组
func main() {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "user index"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "user login"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "user login POST"})
		})
	}

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop index"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop cart"})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop checkout"})
		})
		// 嵌套路由组
		test := shopGroup.Group("test")
		test.GET("/aaa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop test aaa"})
		})
	}

	r.Run(":650")
}
