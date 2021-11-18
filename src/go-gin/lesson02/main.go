package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// gin获取queryString
func main() {
	r := gin.Default()

	// 获取参数
	r.GET("/web", queryString)

	r.Run(":650")
}
