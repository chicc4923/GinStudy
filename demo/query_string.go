package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func queryString() {
	e := gin.Default()
	// GET 请求 url? 后是 querystring
	// key=value 格式，多个 key-value 之间用 & 连接
	// eq: /web?query=max&age=ccc
	e.GET("/web", func(c *gin.Context) {
		//获取浏览器请求的 query string 参数
		name := c.Query("query") // 通过 query 获取请求中携带的 querystring 参数 http://127.0.0.1:8080/web?query=max
		age := c.Query("age")    // url 可以通过 & 符号查询两个参数，比如：http://127.0.0.1:8080/web?query=max&age=19 todo: 为什么返回的 json 里 age 在前，name 在后
		//name := c.DefaultQuery("query", "ccc") // 如果没有 query，就返回指定的默认值：http://127.0.0.1:8080/web?aaa=123
		//name, ok := c.GetQuery("query") // 如果没有 query，返回 false
		//if !ok {
		//	name = "ccc"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	e.Run()
}
