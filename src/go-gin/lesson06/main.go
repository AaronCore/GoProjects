package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

//文件上传
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		//dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)
		c.SaveUploadedFile(f, dst)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

		// 多文件上传
		//form, _ := c.MultipartForm()
		//files := form.File["file"]
		//for index, file := range files {
		//	dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
		//	// 上传文件到指定的目录
		//	c.SaveUploadedFile(file, dst)
		//}
		//c.JSON(http.StatusOK, gin.H{
		//	"message": fmt.Sprintf("%d files uploaded!", len(files)),
		//})
	})

	r.Run(":655")
}
