package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func redirect() {
	e := gin.Default()

	e.GET("/redirect", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	// 跳转到百度
	e.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	// 路由跳转
	e.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" // 把请求的 URL 修改
		e.HandleContext(c)        // 继续后续处理
	})
	e.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	e.Run()
}
