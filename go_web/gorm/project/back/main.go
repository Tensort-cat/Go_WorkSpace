package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `json:"title"`
	Status    bool           `json:"status"`
}

func main() {
	r := gin.Default()
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 增加笔记
	r.POST("/v1/todo", func(ctx *gin.Context) {
		var todo Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "出错"})
			return
		}
		res := db.Create(&todo)
		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
			return
		}
		ctx.JSON(http.StatusOK, "添加成功")
	})

	// 查看笔记
	r.GET("/v1/todo", func(ctx *gin.Context) {
		todos := []Todo{}
		res := db.Find(&todos)
		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
			return
		}
		ctx.JSON(http.StatusOK, &todos)
	})

	// 更新待办事项状态
	r.PUT("/v1/todo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var todo Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "出错"})
			return
		}
		res := db.Model(&Todo{}).Where("id = ?", id).Update("status", todo.Status)
		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
			return
		}
		ctx.JSON(http.StatusOK, "修改成功")
	})

	// 删除笔记
	r.DELETE("/v1/todo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		res := db.Delete(&Todo{}, id)
		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器出错"})
			return
		}
		ctx.JSON(http.StatusOK, "删除成功")
	})

	r.Run()
}
