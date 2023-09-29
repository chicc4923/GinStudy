package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func uploadFile() {
	e := gin.Default()
	e.LoadHTMLFiles("./template/index.html")
	e.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	e.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		// 将读取到的文件保存到服务端本地
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			dest := path.Join("./", file.Filename)
			_ = c.SaveUploadedFile(file, dest)
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		}

	})

	e.Run()
}
