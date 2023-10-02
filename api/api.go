package api

import (
	"Gin_study/model"
	"Gin_study/table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// GetTodoList 获取事项列表
func GetTodoList(c *gin.Context) {
	list := make([]*table.Todo, 0)
	err, _ := model.GetTodo(&list)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "服务端出现错误！",
		})
	} else {
		c.JSON(http.StatusOK, list)
	}
}

// DeleteTodo API 层根据 ID 删除事项
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
	}
	fmt.Println("id", id)
	err = model.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "删除失败！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success!",
		})
	}
}

// UpdateTodo 更新事项状态 TODO: 这里似乎有问题，数据库方面并没有更新
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
	}
	var data table.Todo
	data, err = model.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
		return
	}
	err = model.UpdateTodosByID(data, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "更新失败！",
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
