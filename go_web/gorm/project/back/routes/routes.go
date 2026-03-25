package routes

import (
	"go_web/gorm/project/back/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 新增笔记
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有笔记
		v1Group.GET("/todo", controller.GetTodoList)

		// 更新笔记状态
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除笔记
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
