package router

import (
	"Gin_study/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouters 初始化路由
func SetupRouters() *gin.Engine {

	e := gin.Default()
	// 解析静态目录
	e.LoadHTMLGlob("./dist/index.html")
	// 获取静态文件 css js
	e.Static("./static", "./dist/static")
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	TodoG := e.Group("/v1")
	{
		TodoG.GET("/todolist")
		TodoG.POST("/todolist", api.CreateTodo)
		TodoG.PUT("/todolist")
		TodoG.DELETE("/todolist")
	}
	e.Run()
	return e
}
