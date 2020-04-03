package main

import "github.com/gin-gonic/gin"

func main() {

	// set mode

	// gin.SetMode(gin.ReleaseMode)

	// default route
	r := gin.Default() // return default router

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":    1,
			"name":  "张三",
			"class": "清华",
		})
	})

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
	r.Run(":20000")

}
