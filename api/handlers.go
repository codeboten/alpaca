package main

import "github.com/gin-gonic/gin"

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func functionCreate(c *gin.Context) {
	var f serverlessFunction
	c.Bind(&f)
	if len(f.Path) == 0 || len(f.Container) == 0 {
		c.JSON(422, gin.H{
			"message": "invalid request",
		})
		return
	}
	err := setValue(f.Path, f.Container)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"message":  "created",
		"function": f,
	})
}

func functionDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleted",
	})
}

func functionGet(c *gin.Context) {
	id := c.Param("id")
	val, err := getValue(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	f := serverlessFunction{id, val}
	c.JSON(200, gin.H{
		"function": f,
	})
}

func functionUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func functionsList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
