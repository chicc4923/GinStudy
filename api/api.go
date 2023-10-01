package api

import (
	"Gin_study/model"
	"Gin_study/table"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoView struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(c *gin.Context) {
	var data table.TodoList
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
		return
	}
	if err := model.GetAll(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
}
