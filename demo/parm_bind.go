package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func parmBind() {
	e := gin.Default()

	e.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := user{
		//	Username: username,
		//	Password: password,
		//}
		u := user{}
		err := c.ShouldBind(&u) // 函数传参是值传递，所以如果不加取址符，修改的是 u 的副本，而不是 u 本身
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	e.Run()
}
