package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Gin 框架允许开发者在处理请求的工程中，加入用户自己的钩子函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
// 中间件必须是一个 gin.HandlerFunc 类型，比如下面的代码就是一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	c.Next() // 调用后续处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost :%v\n", cost)
	fmt.Println("m1 out ...")
}

// 此函数也可以视为中间件
func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
func middleware() {
	e := gin.Default()
	e.Use(m1)
	e.GET("/index", m1, index)
	e.Run(":8090")
}
