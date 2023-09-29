package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnJson() {
	e := gin.Default()
	e.GET("/json", func(c *gin.Context) {
		// 1. 使用 map 序列化 json
		data := map[string]interface{}{
			"name":    "max",
			"message": "hello",
			"age":     19,
		}
		c.JSON(http.StatusOK, data)
	})
	// 2. 使用 struct 序列化 json
	type msg struct {
		Name    string // 不可导出字段无法序列化 如果首字母一定要小写，可以使用 tag 指定字段名
		Message string `json:"message"`
		Age     int
	}
	e.GET("/newJson", func(c *gin.Context) {
		data := msg{
			"max",
			"hello",
			19,
		}
		c.JSON(http.StatusOK, data)
	})
	e.Run()
}
