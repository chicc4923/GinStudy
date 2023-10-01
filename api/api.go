package api

import (
	"Gin_study/model"
	"Gin_study/table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var data table.Todo
	err := c.BindJSON(&data)
	fmt.Println(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数输入错误！",
		})
		return
	}
	fmt.Printf("%#v\n", data)
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
