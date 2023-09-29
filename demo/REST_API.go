package demo

import "github.com/gin-gonic/gin"

//	func sayHello(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "hello",
//		})
//	}
func restApi() {
	r := gin.Default()
	//r.GET("/hello", sayHello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})
	r.Run()
}
