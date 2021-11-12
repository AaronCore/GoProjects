package main

import (
	"github.com/gin-gonic/gin"
)

func satHello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello gin"})
}

func main() {
	r := gin.Default()

	r.GET("/hello", satHello)

	r.Run()
}
