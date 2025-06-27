package router

import "github.com/gin-gonic/gin"

func Router()  {
 	r:= gin.Default()
	r.POST("task/create", func(ctx *gin.Context) {

	})
	//r.get  статус задачи . дата создания задачи. время выполнения задачи 
}