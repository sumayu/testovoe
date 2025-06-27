package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sumayu/testovoe/src/internal/task"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/task/create", func(c *gin.Context) {
		response := task.Create()
		c.JSON(200, response)
	})

	r.GET("/task/info", func(c *gin.Context) {
		response := task.Get()
		if response == nil {
			c.JSON(404, gin.H{"error": "task not found"})
			return
		}
		c.JSON(200, response)
	})

	r.DELETE("/task/delete", func(c *gin.Context) {
		response := task.Delete()
		if response == nil {
			c.JSON(404, gin.H{"error": "no task to delete"})
			return
		}
		c.JSON(200, response)
	})

	return r
}