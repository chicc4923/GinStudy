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
	e.GET("/v2/*path", api.Proxy)
	TodoG := e.Group("/v1")
	{
		TodoG.GET("/todo:id")
		TodoG.GET("/todo", api.GetTodoList)
		TodoG.POST("/todo", api.CreateTodo)
		TodoG.PUT("/todo/:id", api.UpdateTodo)
		TodoG.DELETE("/todo/:id", api.DeleteTodo)
	}
	e.Run()
	return e
}
