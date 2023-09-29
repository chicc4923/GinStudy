package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func router() {
	e := gin.Default()
	e.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	e.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	e.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	e.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "https://www.baidu.com",
		})
	})

	// 路由组
	// 视频的首页和详情页
	videoG := e.Group("/video")
	//e.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "/video/index"})
	//})
	{
		videoG.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/video/index"})
		})
	}

	e.Run()
}
