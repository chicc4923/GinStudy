package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解析表单参数
func formParse() {
	e := gin.Default()
	e.LoadHTMLFiles("./template/login.html", "./template/index.html")
	e.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 第一种获取 form 表单提交的数据
	// 其余两种参考 querystring
	// login post
	e.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"password": password,
		})
	})

	e.Run()
}
