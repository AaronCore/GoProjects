package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//http重定向
func main() {
	r := gin.Default()

	// http重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// 路由重定向
	r.GET("test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	r.GET("test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.Run(":650")
}
