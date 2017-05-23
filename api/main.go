package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/functions")
	{
		v1.POST("/", functionCreate)
		v1.GET("/", functionsList)
		v1.GET("/:id", functionGet)
		v1.PUT("/:id", functionUpdate)
		v1.DELETE("/:id", functionDelete)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
