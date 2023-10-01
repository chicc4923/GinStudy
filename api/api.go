package api

import (
	"Gin_study/model"
	"Gin_study/table"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTodo API层的新增事项
func CreateTodo(c *gin.Context) {
	var data table.Todo
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
		return
	}
	if err := model.CreateTodo(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "服务端出现错误！",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": data,
		})
	}
}

func GetTodoList(c *gin.Context) {
	list := make([]table.Todo, 0)
	err := model.GetTodo(list)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "服务端出现错误！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": list,
		})
	}
}
