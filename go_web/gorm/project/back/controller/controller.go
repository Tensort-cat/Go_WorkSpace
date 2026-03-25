package controller

import (
	"go_web/gorm/project/back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "出错"})
		return
	}

	if models.CreateTodo(&todo) != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
		return
	}
	ctx.JSON(http.StatusOK, "添加成功")
}

func GetTodoList(ctx *gin.Context) {
	todos, err := models.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
		return
	}
	ctx.JSON(http.StatusOK, &todos)
}

func UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "出错"})
		return
	}
	err := models.UpdateTodo(id, &todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
		return
	}
	ctx.JSON(http.StatusOK, "修改成功")
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
		return
	}
	ctx.JSON(http.StatusOK, "删除成功")
}
