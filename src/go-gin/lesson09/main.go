package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "jack") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		fmt.Printf("耗时：%d\n", cost)
	}
}

// 中间件
func main() {
	r := gin.Default()
	// 注册一个全局中间件
	r.Use(StatCost())

	r.GET("index", func(c *gin.Context) {
		name := c.MustGet("name") // 从上下文取值
		c.JSON(http.StatusOK, gin.H{"message": name})
	})

	// 为某个路由单独注册
	r.GET("index1", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name")
		c.JSON(http.StatusOK, gin.H{"message": name})
	})

	// 为路由组注册中间件
	// 方式1
	//shopGroup := r.Group("/shop", StatCost())
	//{
	//	shopGroup.GET("/index", func(c *gin.Context) {
	//		name := c.MustGet("name")
	//		c.JSON(http.StatusOK, gin.H{"message": name})
	//	})
	//}
	// 方式2
	shopGroup := r.Group("/shop")
	shopGroup.Use(StatCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			name := c.MustGet("name")
			c.JSON(http.StatusOK, gin.H{"message": name})
		})
	}

	r.Run(":650")
}
