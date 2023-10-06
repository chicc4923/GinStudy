package api

import (
	"Gin_study/model"
	"Gin_study/table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
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

/*
正向代理：
代理会隐藏客户端的真实信息（IP、端口），以自己的身份代替客户端在互联网上发起请求，
并将结果转发给客户端。代理可以保护客户端，帮助客户端访问自己无法访问的网络，客户端需要将特定请求或全部请求主动配置为请求代理服务器。
反向代理：
反向代理会隐藏服务端的真实信息（IP、端口），把自己作为服务端暴露在互联网中，通过把请求转发给真实服务器处理，拿到结果再返回，来对外提供服务。
反向代理保护了服务端，隔离了有效环境，并进行多个服务器的负载均衡。
服务端需要将自己配置到反向代理中，然后将反向代理暴露在公网。

*/

func Proxy(c *gin.Context) {
	host := "127.0.0.1:8080"
	ur := c.Param("path")
	fmt.Println(ur)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = host
			req.URL.Path = "/v1" + ur
		},
	}
	fmt.Println(">>>>>>", c.Request.URL.Path)
	c.Request.URL.Path = "/v1" + ur
	fmt.Println(">>>>>>2", c.Request.URL.Path)
	proxy.ServeHTTP(c.Writer, c.Request)

}
