package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的 path(URI)参数
func uriParse() {
	e := gin.Default()
	// http://127.0.0.1:8080/user/max/12
	// ep :http://127.0.0.1:8080/blog/2023/10
	e.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	e.Run()
}
